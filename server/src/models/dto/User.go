package dto

type UserDTO struct {
	Email    string        `json:"email"`
	Username string        `json:"username"`
	Admin    bool          `json:"admin"`
	Accounts []ProviderDTO `json:"accounts"`
}

type ShortUser struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}
