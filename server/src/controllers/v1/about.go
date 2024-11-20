package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/config"
	"server/src/models"
	"server/src/models/dto"
	"time"
)

func About(c *gin.Context) {
	var providers []models.Provider
	var actions []models.Action
	var about dto.About

	result := config.DB.Find(&providers)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "Failed find providers."})
		return
	}
	about.Client.Host = c.ClientIP()
	about.Server.CurrentTime = time.Now().Unix()
	about.Server.Services = []dto.Service{}

	for _, provider := range providers {
		service := dto.Service{
			Name:      provider.Name,
			Actions:   []dto.ActionDescription{},
			Reactions: []dto.ActionDescription{},
		}

		result = config.DB.Joins("JOIN scopes ON scopes.id = actions.scope_id").Where("provider_id = ?", provider.ID).Find(&actions).Order("scopes.name ASC, actions.name ASC")
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "Failed find actions for provider " + provider.Name + "."})
			return
		}
		for _, action := range actions {
			if action.Type == models.ActionEnum {
				service.Actions = append(service.Actions, dto.ActionDescription{
					Name:        action.Name,
					Description: action.Description,
				})
			} else if action.Type == models.TriggerEnum {
				service.Reactions = append(service.Reactions, dto.ActionDescription{
					Name:        action.Name,
					Description: action.Description,
				})
			} else {
				c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "Failed find actions for provider " + provider.Name + "."})
				return
			}
		}

		about.Server.Services = append(about.Server.Services, service)
	}
	c.JSON(http.StatusOK, about)
}
