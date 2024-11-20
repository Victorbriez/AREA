package models

import "server/src/models/dto"

type UserScope struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	ScopeID    int
	ProviderID int
	UserID     int
	Scope      Scope    `gorm:"foreignKey:ScopeID"`
	User       User     `gorm:"foreignKey:UserID"`
	Provider   Provider `gorm:"foreignKey:ProviderID"`
}

func (up *UserScope) GetShortUser() dto.ShortUser {
	return dto.ShortUser{
		ID:       up.ID,
		Username: up.User.Username,
		Email:    up.User.Email,
	}
}

func (up *UserScope) GetScope() dto.SimpleScope {
	return dto.SimpleScope{
		ID:       up.ScopeID,
		Scope:    up.Scope.Scope,
		Required: up.Scope.Required,
	}
}
