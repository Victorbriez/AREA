package models

import "time"

type User struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Email     string `gorm:"unique"`
	Password  string
	Username  string         `gorm:"unique"`
	CreatedAt time.Time      `gorm:"default:now()"`
	Admin     bool           `gorm:"default:false"`
	Scopes    []UserScope    `gorm:"foreignKey:UserID"`
	Providers []UserProvider `gorm:"foreignKey:UserID"`
	Flows     []Flow         `gorm:"foreignKey:UserID"`
	Logs      []Log          `gorm:"foreignKey:UserID"`
}
