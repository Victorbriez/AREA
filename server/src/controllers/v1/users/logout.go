package users

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/config"
	"server/src/models/dto"
)

// Logout godoc
// @Security Session Token
// @Summary Logout current user
// @Tags users
// @Produce json
// @Success 204 "Disconnected"
// @Failure 400 {object} dto.RequestError
// @Failure 401 {object} dto.RequestError
// @Failure 403 {object} dto.RequestError
// @Failure 500 {object} dto.RequestError
// @Router /v1/users/logout [post]
func Logout(c *gin.Context) {
	_, exists := c.Get("user")

	if !exists {
		c.JSON(http.StatusUnauthorized, dto.RequestError{Error: "You need to be logged in"})
		return
	}

	token, _ := c.Get("token")

	config.Redis.Del(context.Background(), token.(string))
	c.Status(http.StatusNoContent)
}
