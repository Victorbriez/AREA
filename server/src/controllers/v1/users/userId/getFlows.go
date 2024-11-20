package userId

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/config"
	"server/src/models"
	"server/src/models/dto"
	"strconv"
)

// GetFlows godoc
// @Security Session Token
// @Summary Get user details
// @Tags users
// @Param userId path string true "User ID (integer or 'me')"
// @Produce json
// @Success 200 {object} dto.SimpleFlow
// @Failure 400 {object} dto.RequestError
// @Failure 401 {object} dto.RequestError
// @Failure 403 {object} dto.RequestError
// @Failure 404 {object} dto.RequestError
// @Failure 500 {object} dto.RequestError
// @Router /v1/users/{userId}/flows [get]
func GetFlows(c *gin.Context) {
	var userFlows []models.Flow
	var finalFlows []dto.SimpleFlow
	routeId := c.Param("userId")
	value, exists := c.Get("user")
	id := -1

	if routeId == "me" {
		if !exists {
			c.JSON(http.StatusUnauthorized, dto.RequestError{Error: "You must be logged in to access this resource"})
			return
		}
		id = value.(models.User).ID
	} else {
		tempId, err := strconv.Atoi(routeId)
		if err != nil {
			c.JSON(http.StatusNotFound, dto.RequestError{Error: "User not found"})
			return
		}
		id = tempId
	}

	request := config.DB.Where(&models.Flow{UserID: id}).Find(&userFlows)
	if request.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "Internal server error"})
		return
	}

	finalFlows = make([]dto.SimpleFlow, len(userFlows))
	for i, userFlow := range userFlows {
		finalFlows[i] = userFlow.GetSimpleFlow()
	}

	c.JSON(http.StatusOK, finalFlows)
}
