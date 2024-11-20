package flow

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"server/src/config"
	"server/src/models"
	"strconv"
	"strings"
	"text/template"
	"time"
)

func RunFlow() {
	for {
		var flows []models.Flow
		time.Sleep(100 * time.Millisecond)
		if err := config.DB.Where("active = ? AND next_run_at <= ?", true, time.Now()).Find(&flows).Error; err != nil {
			fmt.Println("Error retrieving flows:", err)
			return
		}
		for _, flow := range flows {
			flow := flow
			go func() { runIndividualFlow(&flow) }()
		}
	}
}

func runIndividualFlow(flow *models.Flow) {
	flowVariable := map[string]interface{}{}
	flowVariable["flowName"] = flow.Name
	flowVariable["flowLastRun"] = flow.NextRunAt.Unix() - int64(flow.RunEvery)

	result := config.DB.Model(&flow).Where("active = ?", true).Update("NextRunAt", time.Now().Add(time.Duration(flow.RunEvery)*time.Second))
	if result.Error != nil {
		fmt.Println("Error updating flow:", result.Error)
	}

	var done = false
	var safeGuard = 0
	var step models.FlowStep
	config.DB.Where(models.FlowStep{ID: flow.FirstStep}).Preload("Next").Preload("Action").Preload("Action.Fields").Preload("Action.Scope").First(&step)
	for !done {
		safeGuard++
		err := exectureAction(flow, &step.Action, &flowVariable)
		if err != nil {
			fmt.Println("Error executing action:", err)
			return
		}

		if step.NextStep == nil {
			done = true
			continue
		}
		if safeGuard > 250 {
			fmt.Println("Endless flow:", flow.ID)
			logError(flow, "Endless flow.")
			return
		}

		step = models.FlowStep{ID: step.Next.ID}
		config.DB.Preload("Next").Preload("Action").Preload("Action.Fields").Preload("Action.Scope").First(&step)
	}
	log := config.DB.Create(&models.FlowRun{
		FlowID:     flow.ID,
		ExecutedAt: time.Now(),
		Logs:       "Flow finish successfully",
		Successful: true,
	})
	if err := log.Error; err != nil {
		fmt.Println("Error during logging:", err)
	}
}

func exectureAction(flow *models.Flow, action *models.Action, flowVariable *map[string]interface{}) error {
	if action == nil {
		logError(flow, "Action is incorrectly formed.")
		return errors.New("action is nil")
	}

	var provider models.Provider

	if err := config.DB.Where(&models.Provider{ID: action.Scope.ProviderID}).First(&provider).Error; err != nil {
		logError(flow, "Can't retrieve the provider for this action.")
		return errors.New("can't retrieve provider")
	}

	var userProvider models.UserProvider

	if err := config.DB.Where(&models.UserProvider{UserID: flow.UserID, ProviderID: provider.ID}).First(&userProvider).Error; err != nil {
		logError(flow, "Can't retrieve the user's provider info.")
		return errors.New("can't retrieve userProvider")
	}

	var token models.Token

	if err := config.DB.Where(&models.Token{ID: userProvider.TokenID}).First(&token).Error; err != nil {
		logError(flow, "Can't retrieve token.")
		return errors.New("can't retrieve token")
	}

	tmplURL, err := template.New("url").Parse(action.URL)
	if err != nil {
		logError(flow, "Can't parse URL.")
		return errors.New("can't parse URL")
	}

	var bufURL bytes.Buffer

	err = tmplURL.Execute(&bufURL, flowVariable)
	if err != nil {
		logError(flow, "Can't parse URL.")
		return errors.New("can't parse URL")
	}

	var request *http.Request

	if action.Method == "GET" {
		request, err = http.NewRequest(string(action.Method), bufURL.String(), nil)
	} else {

		tmplBody, err := template.New("body").Parse(action.Body)
		if err != nil {
			logError(flow, "Can't parse Body.")
			return errors.New("can't parse Body")
		}

		var bufBody bytes.Buffer

		err = tmplBody.Execute(&bufBody, flowVariable)
		if err != nil {
			logError(flow, "Can't parse Body.")
			return errors.New("can't parse Body")
		}

		request, err = http.NewRequest(string(action.Method), bufURL.String(), bytes.NewReader(bufBody.Bytes()))
	}

	if err != nil {
		logError(flow, "Can't make request to service")
		return errors.New("can't make request")
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))
	request.Header.Set("Client-Id", provider.ClientID)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		logError(flow, "Can't make request to service: "+action.Name)
		return errors.New("fail to call remove service: " + action.Name)
	}
	defer response.Body.Close()

	jsonByte, err := io.ReadAll(response.Body)
	if err != nil {
		logError(flow, "Invalid service response")
		return errors.New("invalid service response")
	}

	if !(response.StatusCode == 200 || response.StatusCode == 204) {
		fmt.Println(string(jsonByte))
		logError(flow, "Service: '"+action.Name+"' respond with an invalid status code: "+strconv.Itoa(response.StatusCode))
		return errors.New("service: '" + action.Name + "' respond with an invalid status code")
	}

	for _, field := range action.Fields {
		if field.IsInput {
			continue
		}
		jsonValue := gjson.GetBytes(jsonByte, field.JsonPath).String()
		safeValue := strings.ReplaceAll(jsonValue, `"`, "'")
		(*flowVariable)[field.Name] = safeValue
	}

	return nil
}

func logError(flow *models.Flow, message string) {
	fmt.Println(message)
	log := config.DB.Create(&models.FlowRun{
		FlowID:     flow.ID,
		ExecutedAt: time.Now(),
		Logs:       message,
		Successful: false,
	})
	if err := log.Error; err != nil {
		fmt.Println("Error during logging:", err)
	}
}
