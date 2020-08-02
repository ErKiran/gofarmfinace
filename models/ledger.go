package models

import "github.com/google/uuid"

type Ledger struct{
	Model `json:"-"`
	ID uuid.UUID `gorm:"type:uuid;primary_key" json:"Id"`
	TransactionName string `json:"transactionName"`
	Amount float64 `json:"amount"`
	UserID uuid.UUID `json:"userId"`
	CropID uuid.UUID `json:"cropId"`
	TransactionType string `gorm:"type:"transaction_type" json:"transactionType"`
	IsCashedTransaction bool `json:"isCashedTransaction"`
}
