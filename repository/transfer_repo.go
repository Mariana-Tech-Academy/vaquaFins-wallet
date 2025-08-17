package repository

import (
	"vaqua/models"

	"gorm.io/gorm"
)

//import "vaqua/models"

type TransferRepository interface {
	
	FindAccount(AccountNum uint) (*models.Account, error)
	UpdateBalance( acc *models.Account) error
	//GetAccountBalance( AccountNum uint, UserID uint) (AccountBalance uint, err error)
}

type TransferRepo struct{
	db *gorm.DB
}

//Find account numberid

func(r *TransferRepo) FindAccount( AccountNum uint, UserID uint) (*models.Account, error){
	var acc models.Account
	result := r.db.Where("accountnum = ? and User_id = ?", AccountNum, UserID).First(&acc)
	if result.Error!= nil{
		return nil, result.Error
	}
	return &acc, nil

}


func (r*TransferRepo) UpdateBalance( acc *models.Account) error{
	return r.db.Save(acc).Error
}
/*func(r *TransferRepo) GetAccountBalance( AccountNum uint, UserID uint) (AccountBalance uint, err error){
var bal models.Account
balance:= db.Db.Where("account_balance=? and User_id=?",AccountBalance, UserID).First(&bal)
if balance.Error!= nil{
		return 0, balance.Error
	}
	return 0, nil
	}*/
