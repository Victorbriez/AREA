package providerSlug

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/config"
	"server/src/models"
	"server/src/models/dto"
	"strconv"
)

// GetScopes godoc
// @Security Session Token
// @Summary Get user accepted scopes by provider
// @Tags users
// @Param userId path string true "User ID (integer or 'me')"
// @Param provider path string true "Provider slug"
// @Produce json
// @Success 200 {object} dto.ProviderDTO
// @Failure 400 {object} dto.RequestError
// @Failure 401 {object} dto.RequestError
// @Failure 403 {object} dto.RequestError
// @Failure 404 {object} dto.RequestError
// @Failure 500 {object} dto.RequestError
// @Router /v1/users/{userId}/providers/{provider}/scopes [get]
func GetScopes(c *gin.Context) {
	var provider models.Provider
	var scopes []models.UserScope
	var finalScopes []dto.SimpleScope
	providerSlug := c.Param("provider")
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

	request := config.DB.Preload("Scopes").Where(&models.Provider{Slug: providerSlug}).First(&provider)
	if request.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "Internal server error"})
		return
	}

	if provider.Slug != providerSlug {
		c.JSON(http.StatusNotFound, dto.RequestError{Error: "Provider not found"})
		return
	}

	request = config.DB.Preload("Scope").Where(&models.UserScope{ProviderID: provider.ID, UserID: id}).Find(&scopes)
	if request.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "Internal server error"})
		return
	}

	finalScopes = make([]dto.SimpleScope, len(scopes))
	for i, scope := range scopes {
		finalScopes[i] = scope.GetScope()
	}

	c.JSON(http.StatusOK, finalScopes)
}
