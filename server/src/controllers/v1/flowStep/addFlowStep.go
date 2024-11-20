package flowStep

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/config"
	"server/src/models"
	"server/src/models/dto"
)

// AddFlowStep godoc
// @Security Session Token
// @Summary Add a new flow step
// @Tags flow steps
// @Accept json
// @Param input body dto.FlowStepPost true "Flow data"
// @Produce json
// @Success 201 {object} dto.FlowStep
// @Failure 400 {object} dto.RequestError
// @Failure 401 {object} dto.RequestError
// @Failure 403 {object} dto.RequestError
// @Failure 404 {object} dto.RequestError
// @Failure 500 {object} dto.RequestError
// @Router /v1/flowsteps [post]
func AddFlowStep(c *gin.Context) {
	var requestData dto.FlowStepPost

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, dto.RequestError{Error: "Invalid Parameters"})
		return
	}

	newAction := models.FlowStep{
		FlowID:       requestData.FlowID,
		PreviousStep: requestData.PreviousStep,
		NextStep:     requestData.NextStep,
		ActionID:     requestData.ActionID,
	}

	if err := config.DB.Create(&newAction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "Failed save the flow step."})
		return
	}

	c.JSON(http.StatusCreated, newAction.GetSimpleStep())
}
