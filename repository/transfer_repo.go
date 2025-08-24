package repository

import (
	"vaqua/db"
	"vaqua/models"
)

type TransferRepository interface {
	FindAccountByUser(userID uint, FromAccountNum string) (*models.Account, error)
	FindRecipientAccount(RecipientAccountNumber string) (*models.Account, error)
	CreateTransfer(t *models.Transfer) error
	UpdateBalance(acc *models.Account) error 
}

type TransferRepo struct {
}

// /
func (r *TransferRepo) FindAccountByUser(userID uint, FromAccountNum string) (*models.Account, error) { //id was userId
	var acc models.Account
	result := db.DB.Where("account_num = ? AND user_id = ?", FromAccountNum, userID).First(&acc)
	if result.Error != nil {
		return nil, result.Error
	}
	return &acc, nil
}

func (r *TransferRepo) FindRecipientAccount(RecipientAccountNumber string) (*models.Account, error) {
	var acc models.Account
	result := db.DB.Where("account_num = ?", RecipientAccountNumber).First(&acc)
	if result.Error != nil {
		return nil, result.Error
	}
	return &acc, nil
}

func (r *TransferRepo) CreateTransfer(t *models.Transfer) error {
	return db.DB.Create(t).Error
}

func (r *TransferRepo) UpdateBalance(acc *models.Account) error {
	return db.DB.
		Model(&models.Account{}).
		Where("account_num = ?", acc.AccountNum).
		Update("account_balance", acc.AccountBalance).Error
}

/*func (r *TransferRepo) FindAccountByUser(userID uint, FromAccountNum string) (*models.Transfer, error) {
    var acc models.Transfer
    result := db.DB.Where("from_account_num = ? AND user_id = ?", FromAccountNum, userID).First(&acc)
    if result.Error != nil {
        return nil, result.Error
    }
    return &acc, nil
}*/
