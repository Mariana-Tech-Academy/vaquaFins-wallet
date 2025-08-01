package repository

import "vaqua/models"

type TransferRepository interface {
	CreateTransfer(transfer *models.Transfer) error
	GetTransferByID(id uint) (models.Transfer, error)
}

type TransferRepo struct{}
