package actionId

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/config"
	"server/src/models"
	"server/src/models/dto"
	"strconv"
)

// DeleteAction godoc
// @Security Session Token
// @Summary Delete an action
// @Tags actions
// @Param actionId path int true "Action ID"
// @Produce json
// @Success 204 "Action Deleted"
// @Failure 400 {object} dto.RequestError
// @Failure 401 {object} dto.RequestError
// @Failure 403 {object} dto.RequestError
// @Failure 404 {object} dto.RequestError "Action not found"
// @Failure 500 {object} dto.RequestError
// @Router /v1/action/{actionId} [delete]
func DeleteAction(c *gin.Context) {
	actionId := c.Param("actionId")
	atoi, err := strconv.Atoi(actionId)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.RequestError{Error: "Invalid action ID."})
		return
	}

	var action models.Action

	if config.DB.Where(&models.Action{ID: atoi}).First(&action).Error != nil {
		c.JSON(http.StatusNotFound, dto.RequestError{Error: "Failed find the action."})
		return
	}

	if config.DB.Where(&models.Action{ID: atoi}).Delete(&action).Error != nil {
		c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "Failed delete the action."})
	}

	c.Status(http.StatusNoContent)
}
