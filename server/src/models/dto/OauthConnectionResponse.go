package dto

type OauthConnectionResponse struct {
	OauthUrl string `json:"oauth_url"`
	State    string `json:"state"`
}
