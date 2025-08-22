package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string    `json:"name" gorm:"not null"`
	Email    string    `json:"email" gorm:"uniqueIndex;not null"`
	Password string    `json:"password" gorm:"not null"`
	Account  *Account `gorm:"foreignKey:UserId;references:ID" json:”account”`
	Transactions   []Transaction `json:"transactions" gorm:"foreignKey:UserID"`
}

/*type User struct {
	gorm.Model
	AccountNum     uint          `json:"accountnum" gorm:"uniqueIndex"`
	AccountBalance int           `json:"account_balance"`
	Name           string        `json:"name" gorm:"not null"`
	Email          string        `json:"email" gorm:"unique;not null"`
	Password       string        `json:"password" gorm:"not null"`
	Transactions   []Transaction `json:"transactions" gorm:"foreignKey:UserID"`
	Role           string        `json:"role" gorm:"not null"`
}*/
