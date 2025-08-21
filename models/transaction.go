package models

type Transaction struct {
	ID          uint    `json:"id" gorm:"primaryKey"`
	UserID      uint    `json:"user_id"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"` // "income" or "expense"
	Category    string  `json:"category"` // "groceries" or "rent"
	Description string  `json:"description"`
	Date        string  `json:"date"`
}
