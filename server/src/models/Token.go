package models

import "time"

type Token struct {
	ID            int            `gorm:"primaryKey;autoIncrement"`
	AccessToken   string         `gorm:"not null"`
	RefreshToken  string         `gorm:"not null"`
	Expiry        time.Time      `gorm:"not null"`
	UserProviders []UserProvider `gorm:"foreignKey:TokenID"`
}
