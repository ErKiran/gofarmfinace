package models

import (
	"time"

	"github.com/google/uuid"
)

type Crop struct {
	Model     `json:"-"`
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"Id"`
	Name      string    `json:"name"`
	PlantedAt time.Time `json:"plantedAt"`
	FarmSize  float64   `json:"farmSize"`
	FarmID    uuid.UUID `json:"farmID"`
	UserID    uuid.UUID `json:"userId"`
}
