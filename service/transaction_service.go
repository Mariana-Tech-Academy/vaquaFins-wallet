package service

import (
	"vaqua/models"
	"vaqua/repository"
)

func GetTransactionByID(id uint) (*models.Transaction, error) {
	return repository.GetTransactionByID(id)
}
func CreateTransaction(t *models.Transaction) (*models.Transaction, error) {
	return repository.CreateTransaction(t)
}
func GetTransactionsByCategory(userID uint, category string) ([]models.Transaction, error) {
	return repository.GetTransactionsByCategory(userID, category)
}
func GetIncomeExpenseSummary(userID uint) (float64, float64, error) {
	return repository.GetIncomeExpenseSummary(userID)
}
func GetTransactionsByType(userID uint, txType string) ([]models.Transaction, error) {
	return repository.GetTransactionsByType(userID, txType)
}
func GetTransactionsByUser(userID uint) ([]models.Transaction, error) {
	return repository.GetTransactionsByUser(userID)
}

