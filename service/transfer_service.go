package service

import (
	"errors"
	"vaqua/models"
	"vaqua/repository"
    
	
)

type TransferService struct {
	Repo repository.TransferRepository
   
}

func (s *TransferService) TransferMoney(transfer *models.Transfer) error {
    // sender must match user_id + account_num(sender must own the account)
    fromAccount, err := s.Repo.FindAccountByUser(transfer.UserID, transfer.FromAccountNum)
    if err != nil {
        return err
    }

    // recipient: look up by account_num only
    toAccount, err := s.Repo.FindRecipientAccount(transfer.RecipientAccountNumber)
    if err != nil {
        return err
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
	if  err != nil {
        return err
    }

    // record the transfer itself
    err = s.Repo.CreateTransfer(transfer); 
	if err != nil {
        return err
    }

    return nil

}