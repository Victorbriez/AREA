package models

import "time"

type FlowRun struct {
	ID         int       `gorm:"primaryKey;autoIncrement"`
	FlowID     int       `gorm:"not null"`
	ExecutedAt time.Time `gorm:"default:now()"`
	Logs       string
	Successful bool
	Flow       Flow `gorm:"foreignKey:FlowID"`
}
