package models

import ("gorm.io/gorm"


)

type Transfer struct {
	gorm.Model
	AccountNum uint `json:"accountnum" gorm:"unique;not null"`
	UserID uint `json:"User_id"`
	AccountBalance int `json:"account_balance"`
	Amount int64 `json:"amount" gorm:"not null"`
	RecipientID uint `json:"recipient_id" gorm:"not null"`
	Name string `json:"name" gorm:"unique;not null"`
	Email string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Role string `json:"role" gorm:"not null"`
	
}
