package models

import "server/src/models/dto"

type ActionType string

const (
	TriggerEnum ActionType = "TRIGGER"
	ActionEnum  ActionType = "ACTION"
)

type ActionMethod string

const (
	Get     ActionMethod = "GET"
	Post    ActionMethod = "POST"
	Put     ActionMethod = "PUT"
	Delete  ActionMethod = "DELETE"
	Options ActionMethod = "OPTION"
)

type Action struct {
	ID          int          `gorm:"primaryKey;autoIncrement"`
	Name        string       `gorm:"not null"`
	Description string       `gorm:"not null;type:text"`
	Type        ActionType   `gorm:"type:action_type;not null"`
	Method      ActionMethod `gorm:"type:action_method;not null"`
	URL         string       `gorm:"not null"`
	Body        string
	Fields      []ActionField `gorm:"foreignKey:ActionID"`
	ScopeID     *int          `gorm:"index"`
	Scope       *Scope        `gorm:"foreignKey:ScopeID"`
}

func (action Action) GetSimpleAction() dto.Action {
	return dto.Action{
		ID:          action.ID,
		Name:        action.Name,
		Description: action.Description,
		Type:        string(action.Type),
		Method:      string(action.Method),
	}
}
