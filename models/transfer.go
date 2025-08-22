package models

import "gorm.io/gorm"


type Transfer struct {
    gorm.Model // gives you ID, CreatedAt, UpdatedAt, DeletedAt
    UserID                 uint   `json:"user_id" gorm:"not null"`
    FromAccountNum         string   `json:"from_account_num" gorm:"not null"`
    AccountBalance         float64   `json:"account_balance" gorm:"not null"`
    RecipientAccountNumber string   `json:"recipient_account_number" gorm:"not null"`
    Amount                 float64   `json:"amount" gorm:"not null"`
    Description            string `json:"description" gorm:"not null"`
}

/*type Transfer struct {
	gorm.Model
	FromAccountID          uint   `json:"From_Account_ID" gorm:"index;not null"`
	RecipientAccountNumber uint   `json:"recipient_account_number" gorm:"index;not null"`
	Amount                 int64  `json:"amount" gorm:"not null"`
	Reference              string `json:"reference" gorm:"size:255"`
}*/
