package models

import "time"

type Log struct {
	ID      int       `gorm:"primaryKey;autoIncrement"`
	Event   string    `gorm:"not null"`
	LogTime time.Time `gorm:"not null"`
	UserID  *int      `gorm:"index"`
	User    *User     `gorm:"foreignKey:UserID"`
}
