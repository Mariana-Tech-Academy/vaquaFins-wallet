package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID         uint          `json:"userid" gorm:"uniqueIndex;autoIncrement;not null"`
	AccountNum     uint          `json:"accountnum" gorm:"not null"`
	AccountBalance int           `json:"account_balance"`
	Name           string        `json:"name" gorm:"not null"`
	Email          string        `json:"email" gorm:"unique;not null"`
	Password       string        `json:"password" gorm:"not null"`
	Transactions   []Transaction `json:"transactions" gorm:"foreignKey:UserID"`
	Role           string        `json:"role" gorm:"not null"`
}
