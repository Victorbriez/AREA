package dto

type SimpleScope struct {
	ID       int    `json:"id"`
	Scope    string `json:"scope"`
	Required bool   `json:"required"`
}

type ScopePost struct {
	Scope    string `json:"scope" binding:"required"`
	Required *bool  `json:"required" binding:"required"`
}

type ScopeDetails struct {
	ID          int    `json:"id"`
	Scope       string `json:"scope"`
	Required    bool   `json:"required"`
	UsersCount  int    `json:"users_count"`
	ActionCount int    `json:"action_count"`
}
