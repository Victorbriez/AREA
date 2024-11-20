package scopes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/config"
	"server/src/models"
	"server/src/models/dto"
	"server/src/utils"
	"strconv"
)

// GetAllScopes godoc
// @Security Session Token
// @Summary Retrieve all scopes related to a provider
// @Tags scopes
// @Param page query int false "Page" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Param provider path  string true "Provider Slug"
// @Produce json
// @Success 200 {object} dto.PaginatedResponse{data=dto.SimpleScope} "OK"
// @Failure 400 {object} dto.RequestError
// @Failure 401 {object} dto.RequestError
// @Failure 403 {object} dto.RequestError
// @Failure 404 {object} dto.RequestError
// @Failure 500 {object} dto.RequestError
// @Router /v1/providers/{provider}/scopes [get]
func GetAllScopes(c *gin.Context) {
	providerSlug := c.Param("provider")
	provider := models.Provider{}
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	result := config.DB.Where(models.Provider{Slug: providerSlug}).First(&provider)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "Failed find the provider."})
		return
	}

	var scopes []models.Scope

	result = config.DB.Where("provider_id = ?", provider.ID).Find(&scopes)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "Failed find the scopes of the provider."})
	}

	var dtoScopes []dto.SimpleScope

	for _, scope := range scopes {
		dtoScopes = append(dtoScopes, scope.GetSimpleScope())
	}

	c.JSON(http.StatusOK, utils.Paginate(dtoScopes, page, pageSize, len(dtoScopes)))
}
