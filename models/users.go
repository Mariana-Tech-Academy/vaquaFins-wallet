package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	AccountNum uint `json:"accountnum" gorm:"unique;not null"`
	AccountBalance int `json:"account_balance"`
	Name string `json:"name" gorm:"unique;not null"`
	Email string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Role string `json:"role" gorm:"not null"`
	Transactions []Transaction `json:"transactions" gorm:"foreignKey:UserID"`
}