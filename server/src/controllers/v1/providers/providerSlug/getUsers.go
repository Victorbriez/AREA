package providerSlug

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/config"
	"server/src/models"
	"server/src/models/dto"
)

func GetUsers(c *gin.Context) {
	var userProviders []models.UserProvider
	var shortUsers []dto.ShortUser
	var provider models.Provider

	request := config.DB.Where(&models.Provider{Slug: c.Param("provider")}).Find(&provider)
	if request.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.Error("Internal Server Error"))
		return
	}

	if provider.ID == 0 {
		c.JSON(404, dto.Error("Provider not found"))
		return
	}

	request = config.DB.Where(&models.UserProvider{ProviderID: provider.ID}).Preload("User").Find(&userProviders)
	if request.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.Error("Internal Server Error"))
		return
	}

	for _, userProvider := range userProviders {
		shortUsers = append(shortUsers, userProvider.GetShortUser())
	}

	c.JSON(200, shortUsers)
}