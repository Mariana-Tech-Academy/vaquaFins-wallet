package models

import "gorm.io/gorm"

type Transfer struct {
    gorm.Model // gives you ID, CreatedAt, UpdatedAt, DeletedAt
    UserID                 uint   `json:"user_id" gorm:"not null"`
    AccountNum             uint   `json:"account_num" gorm:"not null"`
    AccountBalance         uint   `json:"account_balance" gorm:"not null"`
    RecipientAccountNumber uint   `json:"recipient_account_number" gorm:"not null"`
    Amount                 uint   `json:"amount" gorm:"not null"`
    Description            string `json:"description" gorm:"not null"`
}


