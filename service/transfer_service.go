package service

import (
	"fmt"
	
	"vaqua/repository"
)

type TransferService struct {
	Repo repository.TransferRepository
}


//Holds business rules: ownership, balance, amount.

 func( t *TransferService) TransferMoney(AccountNum, RecipientID uint ,Amount int64 ) error{

	//1) Ask repo: “Give me account 123 (the from-account)
	fromAccount, err := t.Repo.FindAccount(AccountNum)
	if err != nil{
		return err
	}
	toAccount, err := t.Repo.FindAccount(RecipientID)
//Check if amount > 0.

	if fromAccount.AccountBalance < int(Amount){
	return fmt.Errorf("not enough funds")
	}

	fromAccount.AccountBalance -= int(Amount)
	toAccount.AccountBalance += int(Amount)

	err= t.Repo.UpdateBalance(fromAccount)
	if err!= nil{
		return err
	}

	err = t.Repo.UpdateBalance(toAccount)
	if err!= nil{
		return err
	}

	return nil
	

	


	}
 

