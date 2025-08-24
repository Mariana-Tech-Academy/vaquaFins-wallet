package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	UserId         uint   `json:"user_id" gorm:"index;not null"`
	AccountNum     string `json:"account_num" gorm:"size:9;uniqueIndex"`
	AccountBalance  float64 `json:"account_balance" gorm:"not null;default:200"`
}
