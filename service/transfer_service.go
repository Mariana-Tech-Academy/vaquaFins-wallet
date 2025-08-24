package service

import (
	"errors"
	"fmt"
	"vaqua/models"
	"vaqua/repository"

	"gorm.io/gorm"
)

type TransferService struct {
	Repo  repository.TransferRepository
	Trepo repository.TransactionRepository
}

func (s *TransferService) TransferMoney(transfer *models.Transfer) error {
	// sender must match user_id + account_num(sender must own the account)
	fromAccount, err := s.Repo.FindAccountByUser(transfer.UserID, transfer.FromAccountNum)
	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("sender account not found")
		}
		return err

	}
	if fromAccount == nil {
		return fmt.Errorf("sender account lookup returned nil")
	}

	//}

	//if err != nil {
		//return err
	// recipient: look up by account_num only
	toAccount, err := s.Repo.FindRecipientAccount(transfer.RecipientAccountNumber)
	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("sender account not found")
		}
		return err

	}
	if fromAccount == nil {
		return fmt.Errorf("sender account lookup returned nil")
	}


	if fromAccount.AccountBalance < transfer.Amount {
		return errors.New("not enough funds")
	}

	fromAccount.AccountBalance -= transfer.Amount
	toAccount.AccountBalance += transfer.Amount

	err = s.Repo.UpdateBalance(fromAccount)
	if err != nil {
		return err
	}
	err = s.Repo.UpdateBalance(toAccount)
	if err != nil {
		return err
	}

	// record the transfer itself
	err = s.Repo.CreateTransfer(transfer)
	if err != nil {
		return err
	}

    // record the transaction
	transaction := models.Transaction{
		UserID:                 transfer.UserID,
		Type:                   "transfer",
		Amount:                 transfer.Amount,
		RecipientAccountNumber: transfer.RecipientAccountNumber,
		Description:            "Transfer to " + transfer.RecipientAccountNumber,
	}

    // Create transaction record
	err = s.Trepo.CreateTransaction(&transaction)
	if err != nil {
		return err
	}

	return nil

}
