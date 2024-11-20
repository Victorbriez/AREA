package flow

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/config"
	"server/src/models"
	"server/src/models/dto"
	"time"
)

// AddFlow godoc
// @Security Session Token
// @Summary Add a new flow
// @Tags flows
// @Accept json
// @Param input body dto.SimpleFlowPost true "Flow data"
// @Produce json
// @Success 201 {object} dto.SimpleFlow
// @Failure 400 {object} dto.RequestError
// @Failure 401 {object} dto.RequestError
// @Failure 403 {object} dto.RequestError
// @Failure 404 {object} dto.RequestError
// @Failure 500 {object} dto.RequestError
// @Router /v1/flow [post]
func AddFlow(c *gin.Context) {
	var requestData dto.SimpleFlowPost
	value, _ := c.Get("user")

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, dto.RequestError{Error: "Invalid Parameters"})
		return
	}

	newAction := models.Flow{
		Name:      requestData.Name,
		UserID:    value.(models.User).ID,
		FirstStep: requestData.FirstStep,
		RunEvery:  requestData.RunEvery,
		NextRunAt: time.Unix(requestData.NextRunAt, 0),
		Active:    *requestData.Active,
	}

	if err := config.DB.Create(&newAction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "Failed save the action."})
		return
	}

	c.JSON(http.StatusCreated, newAction.GetSimpleFlow())
}
