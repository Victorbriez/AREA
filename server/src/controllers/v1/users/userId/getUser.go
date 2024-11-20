package userId

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/config"
	"server/src/models"
	"server/src/models/dto"
	"strconv"
)

// GetUser godoc
// @Security Session Token
// @Summary Get user details
// @Tags users
// @Param userId path string true "User ID (integer or 'me')"
// @Produce json
// @Success 201 {object} dto.ProviderDTO
// @Failure 400 {object} dto.RequestError
// @Failure 401 {object} dto.RequestError
// @Failure 403 {object} dto.RequestError
// @Failure 404 {object} dto.RequestError
// @Failure 500 {object} dto.RequestError
// @Router /v1/users/{userId} [get]
func GetUser(c *gin.Context) {
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
			c.JSON(http.StatusNotFound, dto.RequestError{Error: "The user doesn't exists"})
			return
		}
		id = tempId
	}

	var user models.User
	request := config.DB.Where(&models.User{ID: id}).First(&user)
	if request.Error != nil {
		c.JSON(http.StatusNotFound, dto.RequestError{Error: "User not found"})
		return
	}

	c.JSON(http.StatusOK, dto.ShortUser{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	})
}
