package service

import (
"vaqua/repository"
)

type TransactionService struct {
	Repo repository.TransactionRepository
}