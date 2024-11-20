package users

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"server/src/config"
	"server/src/models"
	"server/src/models/dto"
	"time"
)

// Login godoc
// @Summary Login into a user
// @Tags users
// @Accept json
// @Param input body dto.Login true "Login Data"
// @Produce json
// @Success 201 {object} dto.OauthCallbackResponse "Token"
// @Failure 400 {object} dto.RequestError
// @Failure 409 {object} dto.RequestError
// @Failure 500 {object} dto.RequestError
// @Router /v1/users/login [post]
func Login(c *gin.Context) {
	var requestData dto.Login
	var user models.User

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, dto.RequestError{Error: "Invalid Parameters"})
		return
	}

	request := config.DB.Where(&models.User{Email: requestData.Email}).First(&user)
	if request.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "Internal Server Error"})
		return
	}

	hashedPassword := HashPassword(requestData.Password)

	if request.RowsAffected == 0 || user.Password != hashedPassword {
		c.JSON(http.StatusUnauthorized, dto.RequestError{Error: "Username or password incorrect"})
		return
	}

	session := uuid.New().String()
	config.Redis.Set(context.Background(), session, user.ID, 24*time.Hour)
	c.JSON(http.StatusOK, dto.OauthCallbackResponse{Token: session})

}

func HashPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
