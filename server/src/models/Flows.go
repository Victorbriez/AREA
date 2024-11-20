package models

import "time"
import "server/src/models/dto"

type Flow struct {
	ID        int        `gorm:"primaryKey;autoIncrement"`
	Name      string     `gorm:"not null"`
	UserID    int        `gorm:"not null"`
	FirstStep int        `gorm:"not null"`
	RunEvery  int        `gorm:"not null"`
	NextRunAt time.Time  `gorm:"type:timestamp;not null"`
	Active    bool       `gorm:"default:false"`
	User      User       `gorm:"foreignKey:UserID"`
	Steps     []FlowStep `gorm:"foreignKey:FlowID"`
}

func (f *Flow) GetSimpleFlow() dto.SimpleFlow {
	return dto.SimpleFlow{
		ID:     f.ID,
		Name:   f.Name,
		Active: f.Active,
	}
}
