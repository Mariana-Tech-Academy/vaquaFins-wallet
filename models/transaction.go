package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID                 uint    `json:"user_id" gorm:"not null"`
	Type                   string  `json:"type" gorm:"not null"`
	Amount                 float64 `json:"amount" gorm:"not null"`
	RecipientAccountNumber string    `json:"recipient_account_number"`
	Description            string  `json:"description"`
}

/*package models

import (
	"gorm.io/gorm"
)



type Transaction struct {
	gorm.Model
	UserID      uint    `json:"userId" gorm:"not null"`//this is what should connect them
	Type        string  `json:"type" gorm:"not null"`
	Amount      float64 `json:"amount" gorm:"not null"`
	RecipientID  string   `json:"recipient_id" gorm:"not null"`
	BalanceAfter int64
	Reference    string `json:"reference" gorm:"size:255"`
}*/



//every transaction needs a userID and RecipientID

/*type Transactions struct {
	gorm.Model
	TranferID    uint  `json:"TransferID" gorm:"index;not null"`
	AccountID    uint  `json:"AccountID" gorm:"index;not null"`
	Amount       int64 `json:"amount" gorm:"not null"`
	BalanceAfter int64
	Reference    string `json:"reference" gorm:"size:255"`
}*/