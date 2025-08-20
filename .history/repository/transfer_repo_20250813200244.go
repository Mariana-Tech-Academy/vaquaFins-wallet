// package repository

// import (
// 	"errors"
// 	"vaqua/db"
// 	"vaqua/models"

// 	"gorm.io/gorm"
// )

// type TransferRepository interface {
// 	CreateTransfer(transfer *models.Transfer) error
// 	DeleteTransfer(transfer *models.Transfer) error
// 	//is getting the transaction not the same as getting the transfer transaction.
// 	//GetTransferByID(id uint) (models.Transfer, error)
// }

// type TransferRepo struct{}

// func (r *TransferRepo) CreateTransfer(transfer *models.Transfer) error {
// 	err := db.DB.Create(transfer).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func (r *TransferRepo) GetTransferByID(id uint) (*models.Transfer, error) {
// 	var transfer models.Transfer
// 	err := db.DB.Where("id = ?", id).Find(&transfer).Error
// 	if errors.Is(err, gorm.ErrRecordNotFound) {
// 		return nil, nil
// 	}
// 	if err != nil {
// 		return &models.Transfer{}, err
// 	}
// 	return &transfer, nil
// }
// func (r *TransferRepo) DeleteTransfer(transfer *models.Transfer) error {
// 	err := db.DB.Delete(transfer).Error
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
