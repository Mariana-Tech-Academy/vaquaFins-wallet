package repository

import (
	"vaqua/db"
	"vaqua/models"
)

type TransferRepository interface {
	FindAccountByUser(UserID uint,AccountNum uint) (*models.Transfer, error)
	FindRecipientAccount(RecipientAccountNumber uint) (*models.Transfer, error)
	CreateTransfer(t *models.Transfer) error
	UpdateBalance(acc *models.Transfer) error
	 
}

type TransferRepo struct {
}

func (r *TransferRepo) FindAccountByUser(userID uint, accountNum uint) (*models.Transfer, error) {
    var acc models.Transfer
    result := db.DB.Where("account_num = ? AND user_id = ?", accountNum, userID).First(&acc)
    if result.Error != nil {
        return nil, result.Error
    }
    return &acc, nil
}

func (r *TransferRepo) FindRecipientAccount(RecipientAccountNumber uint) (*models.Transfer, error) {
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
        Where("account_num = ?", acc.AccountNum).
        Update("account_balance", acc.AccountBalance).Error
}

	
	
	
