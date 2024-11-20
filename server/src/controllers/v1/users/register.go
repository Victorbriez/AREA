package users

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"net/mail"
	"server/src/config"
	"server/src/models"
	"server/src/models/dto"
	"time"
)

// Register godoc
// @Summary Register a new account
// @Tags users
// @Accept json
// @Param input body dto.Register true "Registration Data"
// @Produce json
// @Success 201 {object} dto.OauthCallbackResponse "User created"
// @Failure 400 {object} dto.RequestError
// @Failure 409 {object} dto.RequestError
// @Failure 500 {object} dto.RequestError
// @Router /v1/users/register [post]
func Register(c *gin.Context) {
	var requestData dto.Register
	var user models.User

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, dto.RequestError{Error: "Invalid Parameters"})
		return
	}

	request := config.DB.Where(&models.User{Email: requestData.Email}).First(&models.User{})
	if request.Error != nil && !errors.Is(request.Error, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "Internal Server Error"})
		return
	}

	if request.RowsAffected > 0 {
		c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "User Already Exists"})
		return
	}

	if !CheckEmail(requestData.Email) {
		c.JSON(http.StatusBadRequest, dto.RequestError{Error: "Invalid Email"})
		return
	}

	user.Email = requestData.Email
	user.Password = HashPassword(requestData.Password)
	user.Username = requestData.Name

	request = config.DB.Create(&user)
	if request.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "Internal Server Error"})
		return
	}

	session := uuid.New().String()
	config.Redis.Set(context.Background(), session, user.ID, 24*time.Hour)
	c.JSON(http.StatusOK, dto.OauthCallbackResponse{Token: session})
}

func CheckEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	return true
}
