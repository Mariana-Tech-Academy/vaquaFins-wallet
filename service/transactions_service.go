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
/*package service

import (
	"errors"
	"vaqua/models"
	"vaqua/repository"
)

type TransactionService struct {
	Repo repository.TransactionRepository
}

func (s *TransactionService) CreateTransaction(user *models.Transaction) error {
	// //call the createTransaction
	err := s.Repo.CreateTransaction(user)
	if err != nil {
		return err
	}
	//check if transaction was already created.
	_, err = s.Repo.GetTransactionByID(user.UserID)
	if err != nil {
		return errors.New("transaction already created")
	}
	return nil

}

func (s *TransactionService) GetTransactions(transaction *models.Transaction) error {
	TransactionExist, err := s.Repo.GetTransactionByID(transaction.ID)
	if err != nil {
		return err
	}
	if TransactionExist != nil {
		return err
	}
	return nil
}
*/