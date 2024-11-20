package dto

type ActionField struct {
	ID       uint   `json:"id"`
	IsInput  bool   `json:"is_input"`
	Name     string `json:"name"`
	JsonPath string `json:"json_path"`
}

type ActionFieldPost struct {
	Name     string `json:"name"`
	IsInput  bool   `json:"is_input"`
	JsonPath string `json:"json_path"`
}
