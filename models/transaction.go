package models

import (
	
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID      uint    `json:"user_id" gorm:"not null"`
	Type        string  `json:"type" gorm:"not null;index"`
	Amount      float64 `json:"amount" gorm:"not null"`
	RecipientID uint    `json:"recipient_id" gorm:"not null"`
}






//every transaction needs a userID and RecipientID


