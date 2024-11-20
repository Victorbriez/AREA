package dto

type ActionDescription struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Service struct {
	Name      string              `json:"name"`
	Actions   []ActionDescription `json:"actions"`
	Reactions []ActionDescription `json:"reactions"`
}

type About struct {
	Client struct {
		Host string `json:"host"`
	} `json:"client"`
	Server struct {
		CurrentTime int64     `json:"current_time"`
		Services    []Service `json:"services"`
	} `json:"server"`
}
