package providers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/config"
	"server/src/models"
	"server/src/models/dto"
)

// AddProvider godoc
// @Security Session Token
// @Summary Add a new provider
// @Tags providers
// @Accept json
// @Param input body dto.ProviderPost true "Provider data"
// @Produce json
// @Success 201 {object} dto.ProviderDTO
// @Failure 400 {object} dto.RequestError
// @Failure 401 {object} dto.RequestError
// @Failure 403 {object} dto.RequestError
// @Failure 404 {object} dto.RequestError
// @Failure 500 {object} dto.RequestError
// @Router /v1/providers [post]
func AddProvider(c *gin.Context) {
	var requestData dto.ProviderPost

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, dto.RequestError{Error: "Invalid Parameters"})
		return
	}

	var provider models.Provider
	if config.DB.Where(models.Provider{Slug: requestData.ProviderSlug}).First(&provider).RowsAffected == 1 {
		c.JSON(http.StatusConflict, dto.RequestError{Error: "This provider already exists."})
		return
	}

	newProvider := models.Provider{
		Name:               requestData.ProviderName,
		Slug:               requestData.ProviderSlug,
		ClientID:           requestData.ClientID,
		ClientSecret:       requestData.ClientSecret,
		RedirectURL:        requestData.RedirectURL,
		AuthEndpoint:       requestData.AuthEndpoint,
		TokenEndpoint:      requestData.TokenEndpoint,
		DeviceCodeEndpoint: requestData.DeviceCodeEndpoint,
		UserInfoEndpoint:   requestData.UserInfoEndpoint,
		UserIDField:        requestData.UserIDField,
		UserEmailField:     requestData.UserEmailField,
		UserNameField:      requestData.UserNameField,
	}

	if err := config.DB.Create(&newProvider).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "Failed save the provider."})
		return
	}

	c.JSON(http.StatusCreated, newProvider.GetSimpleProvider())
}
