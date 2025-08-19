package service

import (
	"errors"
	"vaqua/models"
	"vaqua/repository"
)

type TransferService struct {
	Repo repository.TransferRepository
}

// Holds business rules: ownership, balance, amount.
func (s *TransferService) TransferMoney(transfer *models.Transfer) error {

	//1) Ask repo: “Give me account 123 (the from-account)
	fromAccount, err := s.Repo.FindAccount(transfer.AccountNum)
	if err != nil {
		return err
	}
	toAccount, err := s.Repo.FindAccount(transfer.RecipientAccountNumber)
	if err != nil {
		return err
	}
	//Check if amount > 0.
	if fromAccount.AccountBalance < (transfer.Amount) {
		return errors.New("not enough funds")
	}
	fromAccount.AccountBalance -= (transfer.Amount)
	toAccount.AccountBalance += (transfer.Amount)

	err = s.Repo.UpdateBalance(fromAccount)
	if err != nil {
		return err
	}

	err = s.Repo.UpdateBalance(toAccount)
	if err != nil {
		return err
	}

	return nil

}
