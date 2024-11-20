package flowId

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/config"
	"server/src/models"
	"server/src/models/dto"
	"strconv"
	"time"
)

// UpdateFlow godoc
// @Security Session Token
// @Summary Updates a flow
// @Tags flows
// @Accept json
// @Param input body dto.SimpleFlowPost true "Flow data"
// @Param flowId path int true "Flow ID"
// @Produce json
// @Success 201 {object} dto.SimpleFlow
// @Failure 400 {object} dto.RequestError
// @Failure 401 {object} dto.RequestError
// @Failure 403 {object} dto.RequestError
// @Failure 404 {object} dto.RequestError
// @Failure 500 {object} dto.RequestError
// @Router /v1/flow/{flowId} [put]
func UpdateFlow(c *gin.Context) {
	var newData dto.SimpleFlowPost
	var flow models.Flow
	flowId, _ := strconv.Atoi(c.Param("flowId"))

	if err := c.ShouldBindJSON(&newData); err != nil {
		c.JSON(http.StatusBadRequest, dto.RequestError{Error: "Invalid Parameters"})
		return
	}

	request := config.DB.Where(&models.Flow{ID: flowId}).First(&flow)
	if request.Error != nil {
		c.JSON(http.StatusNotFound, dto.RequestError{Error: "Flow not found"})
		return
	}

	if newData.Name != "" {
		flow.Name = newData.Name
	}

	if newData.Active != nil {
		flow.Active = *newData.Active
	}

	if newData.FirstStep != 0 {
		var step models.FlowStep
		request = config.DB.Where(&models.FlowStep{ID: newData.FirstStep}).First(&step)
		if request.Error != nil {
			c.JSON(http.StatusBadRequest, dto.RequestError{Error: "Step not found"})
			return
		}
		flow.FirstStep = step.ID
	}

	if newData.NextRunAt != 0 {
		flow.NextRunAt = time.Unix(newData.NextRunAt, 0)
	}

	if err := config.DB.Save(&flow).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "Save flow failed"})
		return
	}

	c.JSON(http.StatusOK, flow.GetSimpleFlow())
}
