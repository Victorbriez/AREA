package action

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/config"
	"server/src/models"
	"server/src/models/dto"
)

// AddAction godoc
// @Security Session Token
// @Summary Add a new action
// @Tags actions
// @Accept json
// @Param input body dto.ActionPost true "Action data"
// @Produce json
// @Success 201 {object} dto.Action
// @Failure 400 {object} dto.RequestError
// @Failure 401 {object} dto.RequestError
// @Failure 403 {object} dto.RequestError
// @Failure 404 {object} dto.RequestError
// @Failure 500 {object} dto.RequestError
// @Router /v1/action [post]
func AddAction(c *gin.Context) {
	var requestData dto.ActionPost

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, dto.RequestError{Error: "Invalid Parameters"})
		return
	}

	var action models.Action
	config.DB.Where(models.Action{Name: requestData.Name}).First(&action)

	newAction := models.Action{
		Name:        requestData.Name,
		Description: requestData.Description,
		Type:        models.ActionType(requestData.Type),
		Method:      models.ActionMethod(requestData.Method),
		URL:         requestData.URL,
		Body:        requestData.Body,
		ScopeID:     requestData.ScopeId,
	}

	if err := config.DB.Create(&newAction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "Failed save the action."})
		return
	}

	c.JSON(http.StatusCreated, newAction.GetSimpleAction())
}
