package providerSlug

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tidwall/gjson"
	"gorm.io/gorm"
	"gorm.io/gorm/utils"
	"io"
	"net/http"
	"server/src/config"
	"server/src/models"
	"server/src/models/dto"
	"strconv"
	"strings"
	"time"
)

// Callback godoc
// @Security Session Token
// @Summary Get OAuth url
// @Tags oauth
// @Param provider path  string true "Provider Slug"
// @Param code query string true "Provider code"
// @Param state query string true "Provider state"
// @Produce json
// @Success 200 {object} dto.OauthCallbackResponse
// @Failure 400 {object} dto.RequestError
// @Failure 401 {object} dto.RequestError
// @Failure 403 {object} dto.RequestError
// @Failure 404 {object} dto.RequestError
// @Failure 500 {object} dto.RequestError
// @Router /v1/oauth/{provider}/callback [get]
func Callback(c *gin.Context) {
	providerSlug := c.Param("provider")
	var provider models.Provider
	var validScopes []models.UserScope
	var acceptedScopes []models.UserScope

	state := c.Query("state")
	callbackAction, err := config.Redis.Get(context.Background(), state).Result()

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error("State is invalid or expired."))
		return
	}
	scopeList := strings.Split(callbackAction, "?")[1]
	urlScopes := strings.Split(scopeList, "&")
	callbackAction = strings.Split(callbackAction, "?")[0]

	result := config.DB.Preload("Scopes").Where(models.Provider{Slug: providerSlug}).First(&provider)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.Error("Failed find the provider."))
		return
	}

	code := c.Query("code")

	token, err := provider.Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error("Failed to exchange token"))
		return
	}

	client := provider.Client(c, token)
	rq, err := http.NewRequest("GET", provider.UserInfoEndpoint, nil)

	rq.Header.Set("Client-Id", provider.ClientID)
	resp, err := client.Do(rq)
	if err != nil || resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, dto.Error("Failed to get user info from provider"))
		return
	}
	defer func() { _ = resp.Body.Close() }()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Error("Failed to read response"))
		return
	}

	var userInfo dto.UserInfoDTO
	if value := gjson.GetBytes(data, provider.UserNameField).String(); value != "" {
		userInfo.Name = value
	} else {
		c.JSON(http.StatusInternalServerError, dto.Error("Failed to get user name"))
		return
	}

	if provider.UserEmailField != "" {
		if value := gjson.GetBytes(data, provider.UserEmailField).String(); value != "" {
			userInfo.Email = value
		} else {
			c.JSON(http.StatusInternalServerError, dto.Error("Failed to get user email"))
			return
		}
	} else {
		userInfo.Email = ""
	}

	if (strings.HasPrefix(callbackAction, "register") || strings.HasPrefix(callbackAction, "login")) && userInfo.Email == "" {
		c.JSON(http.StatusUnauthorized, dto.RequestError{Error: "You cannot use this provider to login / register"})
		return
	}

	if value := gjson.GetBytes(data, provider.UserIDField).String(); value != "" {
		userInfo.ID = value
	} else {
		c.JSON(http.StatusInternalServerError, dto.Error("Failed to get user id"))
		return
	}

	user := models.User{
		Email:    userInfo.Email,
		Username: userInfo.Name,
	}
	tx := config.DB.Begin()

	var sqlToken = models.Token{
		AccessToken:   token.AccessToken,
		RefreshToken:  token.RefreshToken,
		Expiry:        token.Expiry,
		UserProviders: provider.Users,
	}

	result = config.DB.Where(&sqlToken).Save(&sqlToken)
	if result.Error != nil {
		tx.Rollback()
		return
	}

	if strings.HasPrefix(callbackAction, "register") {
		CallbackRegister(&user, &provider, c, tx, sqlToken.ID, &userInfo)
	} else if strings.HasPrefix(callbackAction, "link-") {
		id, _ := strconv.ParseInt(strings.Split(callbackAction, "-")[1], 10, 64)
		CallbackLink(&user, &provider, c, int(id), tx, sqlToken.ID, &userInfo)
	} else {
		request := config.DB.Where(models.UserProvider{ProviderID: provider.ID, ExternalAccountID: userInfo.ID}).First(&models.UserProvider{})
		if request.RowsAffected != 1 {
			c.JSON(http.StatusInternalServerError, dto.Error("Failed to get user info"))
			return
		}
		request = config.DB.Where(&models.User{Email: userInfo.Email}).First(&user)
		if request.RowsAffected != 1 {
			c.JSON(http.StatusInternalServerError, dto.Error("Failed to get user info"))
			return
		}
		CallbackLogin(&user, &provider, c, &userInfo)
	}

	for _, scope := range provider.Scopes {
		if utils.Contains(urlScopes, scope.Scope) {
			validScopes = append(validScopes, models.UserScope{
				ScopeID:    scope.ID,
				ProviderID: provider.ID,
				UserID:     user.ID,
			})
		}
	}

	request := tx.Where(&models.UserScope{UserID: user.ID, ProviderID: provider.ID}).Find(&acceptedScopes)
	if request.Error != nil {
		tx.Rollback()
		return
	}

	var validScopesCopy []models.UserScope
	for _, validScope := range validScopes {
		found := false
		for _, acceptedScope := range acceptedScopes {
			if validScope.ScopeID == acceptedScope.ScopeID {
				found = true
			}
		}
		if !found {
			validScopesCopy = append(validScopesCopy, validScope)
		}
	}

	validScopes = validScopesCopy

	if len(validScopes) > 0 {
		request = tx.Save(&validScopes)
		if request.Error != nil {
			tx.Rollback()
			return
		}
	}

	if c.IsAborted() {
		tx.Rollback()
		return
	}

	tx.Commit()
}

func CallbackRegister(user *models.User, provider *models.Provider, c *gin.Context, tx *gorm.DB, tokenId int, userInfo *dto.UserInfoDTO) {
	account := models.UserProvider{
		ExternalAccountID: userInfo.ID,
		UserID:            user.ID,
		ProviderID:        provider.ID,
	}

	result := tx.Where(account).Find(&account)
	if result.RowsAffected != 0 {
		c.JSON(http.StatusForbidden, dto.Error("Account already exist."))
		c.Abort()
		return
	}

	result = tx.Create(user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.Error("Account already exist."))
		c.Abort()
		return
	}

	account = models.UserProvider{
		UserID:              user.ID,
		TokenID:             tokenId,
		ExternalAccountID:   userInfo.ID,
		ExternalAccountName: userInfo.Name,
		ProviderID:          provider.ID,
	}

	result = tx.Create(&account)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.Error("Failed to find provider for account."))
		c.Abort()
		return
	}
	session := uuid.New().String()
	config.Redis.Set(context.Background(), session, user.ID, 24*time.Hour)
	c.JSON(http.StatusCreated, dto.OauthCallbackResponse{Token: session})
}

func CallbackLogin(user *models.User, provider *models.Provider, c *gin.Context, userInfo *dto.UserInfoDTO) {
	account := models.UserProvider{
		ExternalAccountID: userInfo.ID,
		UserID:            user.ID,
		ProviderID:        provider.ID,
	}

	result := config.DB.Where(account).First(&account)
	if result.Error != nil {
		c.JSON(http.StatusForbidden, dto.Error("Account is not registered."))
		c.Abort()
		return
	}

	session := uuid.New().String()
	config.Redis.Set(context.Background(), session, account.UserID, 24*time.Hour)
	c.JSON(http.StatusOK, dto.OauthCallbackResponse{Token: session})
}

func CallbackLink(user *models.User, provider *models.Provider, c *gin.Context, userID int, tx *gorm.DB, tokenId int, userInfo *dto.UserInfoDTO) {
	account := models.UserProvider{
		ExternalAccountID:   userInfo.ID,
		ExternalAccountName: userInfo.Name,
		UserID:              userID,
		TokenID:             tokenId,
		ProviderID:          provider.ID,
	}

	result := tx.Where(models.User{ID: userID}).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusForbidden, dto.Error("Account is not registered."))
		c.Abort()
		return
	}

	result = tx.Where(account).Create(&account)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.Error("This provider is already linked to this account."))
		c.Abort()
		return
	}

	session := uuid.New().String()
	config.Redis.Set(context.Background(), session, user.ID, 24*time.Hour)
	c.JSON(http.StatusOK, dto.OauthCallbackResponse{Token: session})
}
