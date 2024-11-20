package providerSlug

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/config"
	"server/src/models"
	"server/src/models/dto"
)

// GetProviderDetails godoc
// @Security Session Token
// @Summary Retrieve provider details
// @Tags providers
// @Param provider path  string true "Provider Slug"
// @Produce json
// @Success 200 {object} dto.ProviderDetails "OK"
// @Failure 400 {object} dto.RequestError
// @Failure 401 {object} dto.RequestError
// @Failure 403 {object} dto.RequestError
// @Failure 404 {object} dto.RequestError "Provider not found"
// @Failure 500 {object} dto.RequestError
// @Router /v1/providers/{provider} [get]
func GetProviderDetails(c *gin.Context) {
	providerSlug := c.Param("provider")
	var provider models.Provider

	if config.DB.Where(models.Provider{Slug: providerSlug}).First(&provider).Error != nil {
		c.JSON(http.StatusNotFound, dto.RequestError{Error: "Failed find the provider."})
		return
	}

	c.JSON(http.StatusOK, provider.GetProviderDetails())
}
