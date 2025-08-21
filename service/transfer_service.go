package service

import (
	"errors"
	"vaqua/database"
	"vaqua/models"
	"vaqua/repository"
)

func TransferMoney(req *models.TransferRequest) error {
	if req.Amount <= 0 {
		return errors.New("transfer amount must be greater than zero")
	}
	fromUser, err := repository.GetUserByID(req.FromUserID)
	if err != nil {
		return errors.New("sender not found")
	}
	toUser, err := repository.GetUserByID(req.ToUserID)
	if err != nil {
		return errors.New("receiver not found")
	}
	if fromUser.Balance < req.Amount {
		return errors.New("insufficient balance")
	}
	tx := database.DB.Begin()
	fromUser.Balance -= req.Amount
	toUser.Balance += req.Amount
	if err := repository.UpdateUserBalance(tx, fromUser); err != nil {
		tx.Rollback()
		return err
	}
	if err := repository.UpdateUserBalance(tx, toUser); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
