package models

import "server/src/models/dto"

type FlowStep struct {
	ID           int       `gorm:"primaryKey;autoIncrement"`
	FlowID       int       `gorm:"not null"`
	PreviousStep *int      `gorm:"index"`
	NextStep     *int      `gorm:"index"`
	ActionID     int       `gorm:"not null"`
	Flow         Flow      `gorm:"foreignKey:FlowID"`
	Previous     *FlowStep `gorm:"foreignKey:PreviousStep"`
	Next         *FlowStep `gorm:"foreignKey:NextStep"`
	Action       Action    `gorm:"foreignKey:ActionID"`
}

func (FS FlowStep) GetSimpleStep() dto.FlowStep {
	return dto.FlowStep{
		ID:           FS.ID,
		FlowID:       FS.FlowID,
		PreviousStep: FS.PreviousStep,
		NextStep:     FS.NextStep,
		ActionID:     FS.ActionID,
	}
}
