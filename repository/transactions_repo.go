package repository

import (
	"vaqua/db"
	"vaqua/models"
)

type TransactionRepository interface {
	GetTransactionByID(id uint) (*models.Transaction, error)
	GetTransactionsByUserID(userID uint, transactions *[]models.Transaction) error
}
type TransactionRepo struct {
}

func (r *TransactionRepo) GetTransactionByID(id uint) (*models.Transaction, error) {

	var transaction models.Transaction
	err := db.DB.Where("id = ?", id).First(&transaction).Error
	if err != nil {
		return &transaction, err
	}
	return &transaction, nil
}
func (r *TransactionRepo) GetTransactionsByUserID(userID uint, transactions *[]models.Transaction) error {
	err := db.DB.Where("user_id = ?", userID).Find(transactions).Error
	if err != nil {
		return err
	}
	return nil
}


/*package repository

import (
	"vaqua/db"
	"vaqua/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	GetTransactionByID(id uint) (*models.Transaction, error)
	CreateTransaction(transaction *models.Transaction) error
}
type TransactionRepo struct {
	DB *gorm.DB
}

func (r *TransactionRepo) GetTransactionByID(id uint) (*models.Transaction, error) {

	var transaction models.Transaction

	err := db.DB.Where("id = ?", id).First(&transaction).Error
	if err != nil {
		return &transaction, err
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

}*/
