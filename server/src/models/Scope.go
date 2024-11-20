package models

import (
	"server/src/models/dto"
)

type Scope struct {
	ID         int         `gorm:"primaryKey;autoIncrement"`
	Scope      string      `gorm:"not null"`
	ProviderID int         `gorm:"not null"`
	Required   bool        `gorm:"default:false"`
	UserScopes []UserScope `gorm:"foreignKey:ScopeID"`
	Actions    []Action    `gorm:"foreignKey:ScopeID"`
}

func (s *Scope) GetSimpleScope() dto.SimpleScope {
	return dto.SimpleScope{
		ID:       s.ID,
		Scope:    s.Scope,
		Required: s.Required,
	}
}

func (s *Scope) GetScopeDetails() dto.ScopeDetails {
	return dto.ScopeDetails{
		ID:          s.ID,
		Scope:       s.Scope,
		Required:    s.Required,
		UsersCount:  -1,
		ActionCount: -1,
	}
}
