package actionId

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/config"
	"server/src/models"
	"server/src/models/dto"
	"strconv"
)

// AddActionField godoc
// @Security Session Token
// @Summary Add a new field for an action
// @Tags actions
// @Accept json
// @Param input body dto.ActionFieldPost true "Action field data"
// @Param actionId path int true "Action ID"
// @Produce json
// @Success 201 {object} dto.ActionField
// @Failure 400 {object} dto.RequestError
// @Failure 401 {object} dto.RequestError
// @Failure 403 {object} dto.RequestError
// @Failure 404 {object} dto.RequestError
// @Failure 500 {object} dto.RequestError
// @Router /v1/action/{actionId}/fields [post]
func AddActionField(c *gin.Context) {
	actionId := c.Param("actionId")
	atoi, err := strconv.Atoi(actionId)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.RequestError{Error: "Invalid action ID."})
		return
	}

	var requestData dto.ActionField

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, dto.RequestError{Error: "Invalid Parameters"})
		return
	}

	var action models.ActionField
	if config.DB.Where(models.ActionField{Name: requestData.Name}).First(&action).RowsAffected == 1 {
		c.JSON(http.StatusConflict, dto.RequestError{Error: "This action field already exists."})
		return
	}

	tmp := uint(atoi)

	newActionField := models.ActionField{
		Name:     requestData.Name,
		IsInput:  requestData.IsInput,
		JsonPath: requestData.JsonPath,
		ActionID: &tmp,
	}

	if err := config.DB.Create(&newActionField).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "Failed save the action field."})
		return
	}

	c.JSON(http.StatusCreated, newActionField.GetSimpleField())
}
