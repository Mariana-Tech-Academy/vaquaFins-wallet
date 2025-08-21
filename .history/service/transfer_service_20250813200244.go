// package service

// import (
// 	"errors"
// 	"vaqua/models"
// 	"vaqua/repository"
// )

// type TransferService struct {
// 	Repo repository.TransferRepository
// }

// func (s *TransferService) CreateTransfer(transfer *models.Transfer) error {
// 	_, err := s.Repo.GetTransferByID(transfer.UserID)
// 	if err == nil {
// 		return errors.New("Do you want to transfer these funds again?")
// 	}
// }
