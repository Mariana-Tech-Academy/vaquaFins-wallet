package repository

import (
	"errors"
	"vaqua/db"
	"vaqua/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type AccountRepository interface{
Create(account *models.Account) error
Save( account *models.Account) error
FindByNumberAndUserForUpdate(number string, UserId uint) (*models.Account, error)
FindAccountByNumber(number string) (*models.Account, error)
}

type AccountRepo struct{

}
//it creates an account  table  in DB
func (r *AccountRepo) Create(account *models.Account) error{
	err := db.DB.Create(account).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *AccountRepo) Save( account *models.Account) error{
	err := db.DB.Save(account).Error
	if err != nil {
		return err
	}
	return nil
}

//This is to find the Sender (ownership +lock)- this is to debit the sender safely for transfer later on
func (r *AccountRepo) FindByNumberAndUserForUpdate(number string, UserId uint) (*models.Account, error){
	var account models.Account

	err := db.DB.Clauses(clause.Locking{Strength: "UPDATE"}).
	Where("accountnum = ? AND user_id=?", number, UserId).First(&account).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // no user found but an error
	}

	if err != nil {
		return &models.Account{}, err
	}

	return &account, nil // user found
}


//This is to find the recipients accountnumber when making a tarnsfer in the future
func (r *AccountRepo) FindAccountByNumber(number string) (*models.Account, error){
	var account models.Account
	err := db.DB.Where("accountnum = ?", number).First(&account).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // no user found but an error
	}

	if err != nil {
		return &models.Account{}, err
	}

	return &account, nil // user found
}





