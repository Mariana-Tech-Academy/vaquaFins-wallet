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

	//1) Ask repo: “Give me account 123 (the from-account), but lock it so no one else changes it right now.
	//var existingAccount models.Account
	fromAccount, err := t.Repo.FindAccount(AccountNum)
	if err != nil{
		return err
	}
	toAccount, err := t.Repo.FindAccount(RecipientID)

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
	 //2) Check if account.UserID == userID (ownership).
	 //  so that when Im calling this function in handler it will bring both the account number and the user ID
	
     //5) Ask repo: “Also load account 456 (the to-account).”
     //6) Subtract from one, add to the other.
     //7) Tell repo: “Save new balances + record transfer row.”

	


	}
 

