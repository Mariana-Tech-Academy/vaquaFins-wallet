package models

import "gorm.io/gorm"

type Account struct{
	gorm.Model
	UserId uint `json:"user_id" gorm:"index;not null"`
	AccountNum     string          `json:"accountnum" gorm:"size:9;uniqueIndex"`
	AccountBalance int64         `json:"account_balance" gorm:"not null;default:0"`
}