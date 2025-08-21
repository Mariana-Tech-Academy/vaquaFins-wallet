package repository

import (
	"vaqua/db"
	"vaqua/models"
)

type TransferRepository interface {
	FindAccountByUser(userID uint, FromAccountNum string) (*models.Transfer, error)
	FindRecipientAccount(RecipientAccountNumber string) (*models.Transfer, error)
	CreateTransfer(t *models.Transfer) error
	UpdateBalance(acc *models.Transfer) error
	 
}

type TransferRepo struct {
}

func (r *TransferRepo) FindAccountByUser(userID uint, FromAccountNum string) (*models.Transfer, error) {
    var acc models.Transfer
    result := db.DB.Where("from_account_num = ? AND user_id = ?", FromAccountNum, userID).First(&acc)
    if result.Error != nil {
        return nil, result.Error
    }
    return &acc, nil
}

func (r *TransferRepo) FindRecipientAccount(RecipientAccountNumber string) (*models.Transfer, error) {
    var acc models.Transfer
    result := db.DB.Where("recipient_account_number = ?", RecipientAccountNumber).First(&acc)
    if result.Error != nil {
        return nil, result.Error
    }
    return &acc, nil
}


func (r *TransferRepo) CreateTransfer(t *models.Transfer) error {
	return db.DB.Create(t).Error
}

func (r *TransferRepo) UpdateBalance(acc *models.Transfer) error {
    return db.DB.
        Model(&models.Transfer{}).
        Where("account_num = ?", acc.FromAccountNum).
        Update("account_balance", acc.AccountBalance).Error
}

	
	
	
