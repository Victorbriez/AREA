package providerSlug

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/utils"
	"net/http"
	"server/src/config"
	"server/src/middleware"
	"server/src/models"
	"server/src/models/dto"
	"strconv"
	"strings"
)

// URL godoc
// @Security Session Token
// @Summary Get OAuth url
// @Tags oauth
// @Param provider path  string true "Provider Slug"
// @Param type query string true "Link Type" Enums(link, register, login)
// @Param scopes query []string false "Scopes" collectionFormat()
// @Produce json
// @Success 200 {object} dto.OauthConnectionResponse
// @Failure 400 {object} dto.RequestError
// @Failure 401 {object} dto.RequestError
// @Failure 403 {object} dto.RequestError
// @Failure 404 {object} dto.RequestError
// @Failure 500 {object} dto.RequestError
// @Router /v1/oauth/{provider}/url [get]
func URL(c *gin.Context) {
	providerSlug := c.Param("provider")
	urlType := c.Query("type")
	urlScopes := strings.Split(c.Query("scopes"), ",")
	var provider models.Provider
	var newScopes []models.Scope

	switch urlType {
	case "link", "register", "login":
		break
	default:
		c.JSON(http.StatusBadRequest, dto.RequestError{Error: "Missing parameter 'type' (link, register or login)."})
		return
	}

	request := config.DB.Preload("Scopes").Where(models.Provider{Slug: providerSlug}).First(&provider)
	if request.Error != nil {
		c.JSON(http.StatusNotFound, dto.Error("Failed find the provider."))
		return
	}

	if urlType == "link" {
		_, userID := middleware.GetAuthozirationInfo(c)
		if c.IsAborted() {
			return
		}

		urlType = "link-" + strconv.Itoa(int(userID))
	}

	for _, scope := range provider.Scopes {
		if utils.Contains(urlScopes, scope.Scope) {
			newScopes = append(newScopes, scope)
		} else if scope.Required {
			newScopes = append(newScopes, scope)
		}
	}
	urlType += "?"

	provider.Scopes = newScopes

	for index, scope := range provider.Scopes {
		urlType += scope.Scope
		if index < len(provider.Scopes)-1 {
			urlType += "&"
		}
	}

	c.JSON(http.StatusOK, provider.GenerateOAuthURL(urlType))
}
