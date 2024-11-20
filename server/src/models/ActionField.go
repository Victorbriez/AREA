package models

import "server/src/models/dto"

type ActionField struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	IsInput  bool   `gorm:"not null"`
	Name     string `gorm:"not null"`
	JsonPath string `gorm:"not null"`
	ActionID *uint  `gorm:"index"`
}

func (af ActionField) GetSimpleField() dto.ActionField {
	return dto.ActionField{
		ID:       af.ID,
		IsInput:  af.IsInput,
		Name:     af.Name,
		JsonPath: af.JsonPath,
	}
}
