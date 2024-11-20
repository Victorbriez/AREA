package scopes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/config"
	"server/src/models"
	"server/src/models/dto"
)

// AddScope godoc
// @Security Session Token
// @Summary Add a scope
// @Tags scopes
// @Param provider path string true "Provider Slug"
// @Accept json
// @Param input body dto.ScopePost true "Scope data"
// @Produce json
// @Success 201 {object} dto.SimpleScope
// @Failure 400 {object} dto.RequestError
// @Failure 401 {object} dto.RequestError
// @Failure 403 {object} dto.RequestError
// @Failure 404 {object} dto.RequestError "Provider not found"
// @Failure 500 {object} dto.RequestError
// @Router /v1/providers/{provider}/scopes [post]
func AddScope(c *gin.Context) {
	var requestData dto.ScopePost
	provider := models.Provider{}
	providerSlug := c.Param("provider")

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, dto.RequestError{Error: "Invalid Parameters"})
		return
	}

	if config.DB.Where(models.Provider{Slug: providerSlug}).First(&provider).Error != nil {
		c.JSON(http.StatusNotFound, dto.RequestError{Error: "Failed find the provider."})
		return
	}

	var scope models.Scope
	if config.DB.Where(models.Scope{Scope: requestData.Scope, ProviderID: provider.ID}).First(&scope).RowsAffected == 1 {
		c.JSON(http.StatusConflict, dto.RequestError{Error: "The scope already exists."})
		return
	}

	newScope := models.Scope{
		Scope:      requestData.Scope,
		ProviderID: provider.ID,
		Required:   *requestData.Required,
	}

	if err := config.DB.Create(&newScope).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "Failed save the scope."})
		return
	}

	c.JSON(http.StatusCreated, newScope.GetSimpleScope())
}
