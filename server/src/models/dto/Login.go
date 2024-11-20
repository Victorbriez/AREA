package dto

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Register struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type UserUpdate struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Admin    *bool  `json:"admin"`
}
