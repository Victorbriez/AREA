package actionId

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/config"
	"server/src/models"
	"server/src/models/dto"
	"strconv"
)

// GetActionFields godoc
// @Security Session Token
// @Summary Get all fields of an action
// @Tags actions
// @Param actionId path int true "Action ID"
// @Produce json
// @Success 200 {object} dto.PaginatedResponse{data=dto.ActionField} "OK"
// @Failure 400 {object} dto.RequestError
// @Failure 401 {object} dto.RequestError
// @Failure 403 {object} dto.RequestError
// @Failure 404 {object} dto.RequestError "Action not found"
// @Failure 500 {object} dto.RequestError
// @Router /v1/action/{actionId}/fields [get]
func GetActionFields(c *gin.Context) {
	actionId := c.Param("actionId")
	atoi, err := strconv.Atoi(actionId)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.RequestError{Error: "Invalid action ID."})
		return
	}

	var action models.Action

	if config.DB.Where(&models.Action{ID: atoi}).Preload("Fields").First(&action).Error != nil {
		c.JSON(http.StatusNotFound, dto.RequestError{Error: "Failed find the action."})
		return
	}

	var actionFields []dto.ActionField

	for _, field := range action.Fields {
		actionFields = append(actionFields, field.GetSimpleField())
	}

	c.JSON(http.StatusOK, actionFields)
}
