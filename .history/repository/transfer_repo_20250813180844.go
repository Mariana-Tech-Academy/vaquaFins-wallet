package repository

import (
	"errors"
	"vaqua/db"
	"vaqua/models"

	"gorm.io/gorm"
)

//import "vaqua/models"

type TransferRepository interface {
	//CreateTransfer(transfer *models.Transfer) error
	//GetTransferByID(id uint) (models.Transfer, error)
}

type TransferRepo struct{}

func (r *TransferRepo) CreateTransfer(transfer *models.Transaction) error {
	err := db.DB.Create(transfer).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *TransferRepo) GetTransferByID(id uint) (transfer *models.Transaction) {
	err := db.DB.Where("id = ?", id).Find(&transfer).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if err != nil {
		return &models.Transfer{}, err
	}
	return &transfer, nil
}
