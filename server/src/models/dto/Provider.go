package dto

type ProviderDTO struct {
	ID           int    `json:"id"`
	ProviderName string `json:"provider_name"`
	ProviderSlug string `json:"provider_slug"`
}

type ProviderDetails struct {
	ProviderName       string `json:"provider_name"`
	ProviderSlug       string `json:"provider_slug"`
	ClientID           string `json:"client_id"`
	RedirectURL        string `json:"redirect_url"`
	AuthEndpoint       string `json:"auth_endpoint"`
	TokenEndpoint      string `json:"token_endpoint"`
	DeviceCodeEndpoint string `json:"device_code_endpoint"`
	UserInfoEndpoint   string `json:"user_info_endpoint"`
}

type ProviderPost struct {
	ProviderName       string `json:"provider_name" binding:"required"`
	ProviderSlug       string `json:"provider_slug" binding:"required"`
	ClientID           string `json:"client_id" binding:"required"`
	ClientSecret       string `json:"client_secret" binding:"required"`
	RedirectURL        string `json:"redirect_url" binding:"required"`
	AuthEndpoint       string `json:"auth_endpoint" binding:"required"`
	TokenEndpoint      string `json:"token_endpoint" binding:"required"`
	DeviceCodeEndpoint string `json:"device_code_endpoint" binding:"required"`
	UserInfoEndpoint   string `json:"user_info_endpoint" binding:"required"`
	UserIDField        string `json:"user_id_field" binding:"required"`
	UserEmailField     string `json:"user_email_field"`
	UserNameField      string `json:"user_name_field" binding:"required"`
}

type UserInfoDTO struct {
	ID    string `json:"sub"`
	Email string `json:"email"`
	Name  string `json:"name"`
}
