package repository

import "vaqua/models"

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	EditUser(user *models.User) error
}

type UserRepo struct{}