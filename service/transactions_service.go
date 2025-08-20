package service

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
	if err == nil {
		return errors.New("transaction already created.")
	}
	return nil

	// //something needs to be rectified.
	//i don't think i need a password for getTransactionsByID
	// hashPass, err := utils.HashPassword(transaction.Password)
	// if err != nil {
	// 	return err
	// }
	// Transaction.Password = hashPass
}

func (s *TransactionService) GetTransactions(transaction *models.Transaction) error {
	TransactionExist, err := s.Repo.GetTransactionByID(transaction.UserID)
	if err != nil {
		return err
	}
	if TransactionExist != nil {
		return err
	}
	return nil
}
