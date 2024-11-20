package scopeId

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/config"
	"server/src/models"
	"server/src/models/dto"
	"strconv"
)

// DeleteScope godoc
// @Security Session Token
// @Summary Delete a scope
// @Tags scopes
// @Param provider path string true "Provider Slug"
// @Param scopeId path int true "Scope ID"
// @Produce json
// @Success 204 "Scope Deleted"
// @Failure 400 {object} dto.RequestError
// @Failure 401 {object} dto.RequestError
// @Failure 403 {object} dto.RequestError
// @Failure 404 {object} dto.RequestError "Provider or Scope not found"
// @Failure 500 {object} dto.RequestError
// @Router /v1/providers/{provider}/scopes/{scopeId} [delete]
func DeleteScope(c *gin.Context) {
	providerSlug := c.Param("provider")
	scopeId, _ := strconv.Atoi(c.Param("scopeId"))
	var provider models.Provider
	var scope models.Scope

	if config.DB.Where(models.Provider{Slug: providerSlug}).First(&provider).Error != nil {
		c.JSON(http.StatusNotFound, dto.RequestError{Error: "Failed find the provider."})
		return
	}

	if scopeId == 0 || config.DB.Where(models.Scope{ID: scopeId, ProviderID: provider.ID}).First(&scope).Error != nil {
		c.JSON(http.StatusNotFound, dto.RequestError{Error: "Failed find the scope."})
		return
	}

	if config.DB.Delete(&scope).Error != nil {
		c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "Failed delete the scope."})
	}

	c.Status(http.StatusNoContent)
}
