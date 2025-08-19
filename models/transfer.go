package models

type Transfer struct { // a transfer is always from a senders point of view
	UserID                 uint   `json:"user_id" gorm:"not null"`
	AccountNum             uint   `json:"account_number" gorm:"not null"`
	AccountBalance         uint   `json:"account_balance" gorm:"not null"`
	RecipientAccountNumber uint   `json:"recipient_account_number" gorm:"not null"`
	Amount                 uint   `json:"amount" gorm:"not null"`
	Description            string `json:"description" gorm:"not null"`
}

//AccountNum, RecipientID uint ,Amount
