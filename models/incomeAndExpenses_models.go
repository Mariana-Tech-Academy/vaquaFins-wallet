package models

import (

	"gorm.io/gorm"
)

type IncomeAndExpenses struct {
    gorm.Model
    AccountID   uint    `gorm:"index;not null"`
    Type        string  `gorm:"not null;index"`
    Amount      float64 `gorm:"not null"`
    CounterpartyID *uint `gorm:"index"`
    Description string
}