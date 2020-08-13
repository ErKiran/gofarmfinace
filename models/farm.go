package models

import uuid "github.com/satori/go.uuid"

type Farm struct {
	Model      `json:"-"`
	ID         uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Name       string    `gorm:"not null;" json:"name"`
	Location   string    `json:"location"`
	FarmSize   float64   `json:"farmSize"`
	FarmTypeID uuid.UUID `json:"farmTypeID"`
	UserID     uuid.UUID `json:"userId"`
}
