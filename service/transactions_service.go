package service

import (
	"vaqua/models"
	"vaqua/repository"
)

type TransactionService struct {
	Repo repository.TransactionRepository
}

func (s *TransactionService) GetTransactions(transaction *models.Transaction) (*models.Transaction,error ){
	TransactionExist, err := s.Repo.GetTransactionByID(uint(transaction.ID))
	if err != nil {
		return nil,err
	}
	return TransactionExist, nil
}
func (s *TransactionService)GetTransactionsByUserID(userID uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := s.Repo.GetTransactionsByUserID(userID, &transactions)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
