package repository

import (
	"vaqua/database"
	"vaqua/models"
)

func GetTransactionByID(id uint) (*models.Transaction, error) {
	var transaction models.Transaction
	if err := database.DB.First(&transaction, id).Error; err != nil {
		return nil, err
	}
	return &transaction, nil
}
func CreateTransaction(t *models.Transaction) (*models.Transaction, error) {
	if err := database.DB.Create(t).Error; err != nil {
		return nil, err
	}
	return t, nil
}

func GetTransactionsByCategory(userID uint, category string) ([]models.Transaction, error) {

	var transactions []models.Transaction
	err := database.DB.Where("user_id = ? AND category = ?", userID,
		category).Find(&transactions).Error
	return transactions, err
}
func GetIncomeExpenseSummary(userID uint) (float64, float64, error) {
	var incomeSum float64
	var expenseSum float64
	if err := database.DB.Model(&models.Transaction{}).Where("user_id = ? AND type = ?", userID,
		"income").Select("SUM(amount)").Scan(&incomeSum).Error; err != nil {
		return 0, 0, err
	}
	if err := database.DB.Model(&models.Transaction{}).Where("user_id = ? AND type = ?", userID,
		"expense").Select("SUM(amount)").Scan(&expenseSum).Error; err != nil {
		return 0, 0, err
	}
	return incomeSum, expenseSum, nil
}

func GetTransactionsByType(userID uint, txType string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := database.DB.Where("user_id = ? AND type = ?", userID, txType).Find(&transactions).Error
	return transactions, err
}

func GetTransactionsByUser(userID uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := database.DB.Where("user_id = ?", userID).Find(&transactions).Error
	return transactions, err
}