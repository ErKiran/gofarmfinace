package models

import uuid "github.com/satori/go.uuid"

type FarmType struct {
	Model       `json:"-"`
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name        string    `gorm:"not null;" json:"name"`
	Description string    `gorm:"not null;" json:"description"`
}
