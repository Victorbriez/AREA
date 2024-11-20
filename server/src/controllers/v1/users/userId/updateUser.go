package userId

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/config"
	"server/src/controllers/v1/users"
	"server/src/models"
	"server/src/models/dto"
	"strconv"
)

// UpdateUser godoc
// @Security Session Token
// @Summary Updates a user
// @Tags users
// @Accept json
// @Param userId path string true "User ID (integer or 'me')"
// @Param input body dto.UserUpdate true "Updated Data"
// @Produce json
// @Success 200 {object} dto.ShortUser
// @Failure 400 {object} dto.RequestError
// @Failure 403 {object} dto.RequestError
// @Failure 409 {object} dto.RequestError
// @Failure 500 {object} dto.RequestError
// @Router /v1/users/{userId} [put]
func UpdateUser(c *gin.Context) {
	var newData dto.UserUpdate
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

	if err := c.ShouldBindJSON(&newData); err != nil {
		c.JSON(http.StatusBadRequest, dto.RequestError{Error: "Invalid Parameters"})
		return
	}

	if newData.Email != "" {
		request = config.DB.Where(&models.User{Email: newData.Email}).First(&models.User{})
		if request.RowsAffected > 0 {
			c.JSON(http.StatusConflict, dto.RequestError{Error: "User email is already taken"})
			return
		}
		if !users.CheckEmail(newData.Email) {
			c.JSON(http.StatusBadRequest, dto.RequestError{Error: "The provided email address is not valid"})
			return
		}
		user.Email = newData.Email
	}

	if newData.Name != "" {
		user.Username = newData.Name
	}

	if newData.Password != "" {
		user.Password = users.HashPassword(newData.Password)
	}

	if newData.Admin != nil {
		if !value.(models.User).Admin {
			c.JSON(http.StatusUnauthorized, dto.RequestError{Error: "You must be admin to update permissions"})
			return
		}
		user.Admin = *newData.Admin
	}

	request = config.DB.Save(&user)
	if request.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "Failed to save new data"})
		return
	}

	c.JSON(http.StatusOK, dto.ShortUser{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	})
}
