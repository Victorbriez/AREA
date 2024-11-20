package providers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/src/config"
	"server/src/models"
	"server/src/models/dto"
	"server/src/utils"
	"strconv"
)

// GetAllProviders godoc
// @Security Session Token
// @Summary Retrieve all providers
// @Tags providers
// @Param page query int false "Page" default(1)
// @Param pageSize query int false "Page size" default(10)
// @Produce json
// @Success 200 {object} dto.PaginatedResponse{data=dto.ProviderDTO} "OK"
// @Failure 400 {object} dto.RequestError
// @Failure 401 {object} dto.RequestError
// @Failure 403 {object} dto.RequestError
// @Failure 404 {object} dto.RequestError
// @Failure 500 {object} dto.RequestError
// @Router /v1/providers [get]
func GetAllProviders(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	var providers []models.Provider

	result := config.DB.Find(&providers)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.RequestError{Error: "Failed find providers."})
	}

	var dtoProviders []dto.ProviderDTO

	for _, provider := range providers {
		dtoProviders = append(dtoProviders, provider.GetSimpleProvider())
	}

	c.JSON(http.StatusOK, utils.Paginate(dtoProviders, page, pageSize, len(dtoProviders)))
}
