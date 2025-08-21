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

	//check if transaction was already created.
	_, err := s.Repo.GetTransactionByID(user.UserID)
	if err == nil {
		return errors.New("transaction already created.")
	}
	// //something needs to be rectified....
	// hashPass, err := utils.HashPassword(transaction.Password)
	// if err != nil {
	// 	return err
	// }
	// Transaction.Password = hashPass

	// //call the createTransaction
	// err = s.Repo.CreateTransaction(user)
	// if err != nil {
	// 	return err
	// }
	// // return nil
	return nil
}

func(s *TransactionService) GetTransactions(transaction *models.Transaction) error{
	_, err := s.Repo.GetTransactions(transaction.UserID)
	if err = nil {
		return err
	}
	return &transaction, nil
}