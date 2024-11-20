package dto

type PaginatedResponse struct {
	Data    interface{} `json:"data"`
	Page    int         `json:"page"`
	PerPage int         `json:"perPage"`
	Total   int         `json:"total"`
}
