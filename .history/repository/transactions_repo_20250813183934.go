package repository

import (
	"vaqua/db"
	"vaqua/models"
)

//import "vaqua/models"

type TransactionRepository interface {
	CreateTransaction(transaction *models.Transaction) error
	GetTransactionByID(id uint) (*models.Transaction, error)
	DeleteTransaction(transaction *models.Transaction) error
}

type TransactionRepo struct {
}

func (r *TransactionRepo) GetTransactionByID(UserID uint) (*models.Transaction, error) {
	var transaction models.Transaction
	err := db.DB.Where("UserID = ?", UserID).First(&transaction).Error
	if err != nil {
		return &models.Transaction{}, err
	}
	return &transaction, nil
}

// create the createTransaction method
func (r *TransactionRepo) CreateTransaction(transaction *models.Transaction) error {
	err := db.DB.Create(transaction).Error
	if err != nil {
		return err
	}
	return nil

}
func (r *TransactionRepo) DeleteTransaction(transaction *models.Transaction) error {
	err := db.DB.Delete(transaction).Error
	if err != nil {
		return err
	}
	return nil
}
