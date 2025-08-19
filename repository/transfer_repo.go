package repository

import (
	"vaqua/db"
	"vaqua/models"
)

type TransferRepository interface {
	FindAccount(AccountNum uint) (*models.Transfer, error)
	UpdateBalance(acc *models.Transfer) error
	CreateTransfer(Amount uint) (*models.Transfer, error)
}

type TransferRepo struct {
}

//Find account numberid

func (r *TransferRepo) FindAccount(AccountNum uint) (*models.Transfer, error) {

	var acc models.Transfer
	result := db.DB.Where("accountnum = ? and User_id = ?", AccountNum).First(&acc)
	if result.Error != nil {
		return nil, result.Error
	}
	return &acc, nil
}
func (r *TransferRepo) CreateTransfer(Amount uint) (*models.Transfer, error) {
	newTransfer := &models.Transfer{
		Amount: Amount,
	}
	err := db.DB.Create(newTransfer).Error
	if err != nil {
		return nil, err
	}
	return newTransfer, nil
}

func (r *TransferRepo) UpdateBalance(acc *models.Transfer) error {
	return db.DB.Save(acc).Error
}
