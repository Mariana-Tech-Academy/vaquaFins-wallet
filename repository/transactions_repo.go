package repository

import "vaqua/models"

type TransactionRepository interface {
	CreateTransaction(transaction *models.Transaction) error
	GetTransactionByID(id uint) (models.Transaction, error)

}

type TransactionRepo struct{}