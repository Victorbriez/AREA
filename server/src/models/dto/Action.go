package dto

type Action struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Method      string `json:"method"`
}

type ActionPost struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Method      string `json:"method"`
	URL         string `json:"url"`
	Body        string `json:"body"`
	ScopeId     *int   `json:"scope_id"`
}
