package models

import (
	"context"
	"golang.org/x/oauth2"
	"math/rand"
	"net/http"
	"server/src/config"
	"server/src/models/dto"
	"time"
)

type Provider struct {
	ID                 int            `gorm:"primaryKey;autoIncrement"`
	Name               string         `gorm:"not null"`
	Slug               string         `gorm:"unique;not null"`
	ClientID           string         `gorm:"not null"`
	ClientSecret       string         `gorm:"not null"`
	RedirectURL        string         `gorm:"not null"`
	AuthEndpoint       string         `gorm:"not null"`
	TokenEndpoint      string         `gorm:"not null"`
	DeviceCodeEndpoint string         `gorm:"not null"`
	UserInfoEndpoint   string         `gorm:"not null"`
	UserIDField        string         `gorm:"not null"`
	UserEmailField     string         `gorm:"not null"`
	UserNameField      string         `gorm:"not null"`
	Scopes             []Scope        `gorm:"foreignKey:ProviderID"`
	Users              []UserProvider `gorm:"foreignKey:ProviderID"`
}

func (provider Provider) Config() *oauth2.Config {
	var scopes []string
	for _, scope := range provider.Scopes {
		scopes = append(scopes, scope.Scope)
	}

	var Endpoint = oauth2.Endpoint{
		AuthURL:       provider.AuthEndpoint,
		TokenURL:      provider.TokenEndpoint,
		DeviceAuthURL: provider.DeviceCodeEndpoint,
		AuthStyle:     oauth2.AuthStyleAutoDetect,
	}

	var Config = oauth2.Config{
		ClientID:     provider.ClientID,
		ClientSecret: provider.ClientSecret,
		Endpoint:     Endpoint,
		//TODO: Use a global config table in DB to get hostname of the instance in order to craft URL ?
		RedirectURL: provider.RedirectURL,
		Scopes:      scopes,
	}
	return &Config
}

func (provider Provider) GenerateOAuthURL(callbackAction string) dto.OauthConnectionResponse {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 64)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	state := string(b)
	config.Redis.Set(context.Background(), state, callbackAction, time.Minute*5)
	return dto.OauthConnectionResponse{OauthUrl: provider.Config().AuthCodeURL(state), State: state}
}

func (provider Provider) Client(c context.Context, token *oauth2.Token) *http.Client {
	return provider.Config().Client(c, token)
}

func (provider Provider) Exchange(c context.Context, code string, opts ...oauth2.AuthCodeOption) (*oauth2.Token, error) {
	return provider.Config().Exchange(c, code, opts...)
}

func (provider Provider) GetSimpleProvider() dto.ProviderDTO {
	return dto.ProviderDTO{
		ID:           provider.ID,
		ProviderName: provider.Name,
		ProviderSlug: provider.Slug,
	}
}

func (provider Provider) GetProviderDetails() dto.ProviderDetails {
	return dto.ProviderDetails{
		ProviderName:       provider.Name,
		ProviderSlug:       provider.Slug,
		ClientID:           provider.ClientID,
		RedirectURL:        provider.RedirectURL,
		AuthEndpoint:       provider.AuthEndpoint,
		TokenEndpoint:      provider.TokenEndpoint,
		DeviceCodeEndpoint: provider.DeviceCodeEndpoint,
		UserInfoEndpoint:   provider.UserInfoEndpoint,
	}
}
