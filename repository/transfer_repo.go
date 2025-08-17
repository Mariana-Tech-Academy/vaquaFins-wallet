package repository

import (
	"vaqua/models"

	"gorm.io/gorm"
)

//import "vaqua/models"

type TransferRepository interface {
	
	FindAccount(AccountNum uint) (*models.Transfer, error)
	UpdateBalance( acc *models.Transfer) error
	
}

type TransferRepo struct{
	db *gorm.DB
}

//Find account numberid

func(r *TransferRepo) FindAccount( AccountNum uint, UserID uint) (*models.Transfer, error){
	var acc models.Transfer
	result := r.db.Where("accountnum = ? and User_id = ?", AccountNum, UserID).First(&acc)
	if result.Error!= nil{
		return nil, result.Error
	}
	return &acc, nil

}


func (r*TransferRepo) UpdateBalance( acc *models.Transfer) error{
	return r.db.Save(acc).Error
}

