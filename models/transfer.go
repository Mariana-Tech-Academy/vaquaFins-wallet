package models

import ("gorm.io/gorm"


)

type Transfer struct {
	gorm.Model
	AccountNum uint `json:"accountnum" gorm:"unique;not null"`
	UserID uint `json:"user_id" gorm:"not null"`
	Type string `json:"type" gorm:"unique;not null"`
	AccountBalance int `json:"account_balance"`
	Amount int64 `json:"amount" gorm:"not null"`
	RecipientID uint `json:"recipient_id" gorm:"not null"`
}