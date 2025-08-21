package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID                 uint    `json:"user_id" gorm:"not null"`
	Type                   string  `json:"type" gorm:"not null"`
	Amount                 float64 `json:"amount" gorm:"not null"`
	RecipientAccountNumber uint    `json:"recipient_account_number"`
	Description            string  `json:"description"`
}
