package models

import "server/src/models/dto"

type UserProvider struct {
	ID                  int      `gorm:"primaryKey;autoIncrement"`
	ExternalAccountID   string   `gorm:"not null"`
	ExternalAccountName string   `gorm:"not null"`
	ProviderID          int      `gorm:"not null"`
	UserID              int      `gorm:"not null"`
	TokenID             int      `gorm:"not null"`
	Provider            Provider `gorm:"foreignKey:ProviderID"`
	User                User     `gorm:"foreignKey:UserID"`
	Token               Token    `gorm:"foreignKey:TokenID"`
}

func (up *UserProvider) GetShortUser() dto.ShortUser {
	return dto.ShortUser{
		ID:       up.ID,
		Username: up.User.Username,
		Email:    up.User.Email,
	}
}

func (up *UserProvider) GetProvider() dto.ProviderDTO {
	return dto.ProviderDTO{
		ID:           up.ProviderID,
		ProviderName: up.Provider.Name,
		ProviderSlug: up.Provider.Slug,
	}
}
