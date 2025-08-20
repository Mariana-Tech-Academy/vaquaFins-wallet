package repository

import (
	"vaqua/models"

	"errors"
	"vaqua/db"

	"gorm.io/gorm"
)

//import "vaqua/models"

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	
}

type UserRepo struct{}

func (r *UserRepo) CreateUser(user *models.User) error { //added pointer receiver to fix error in main.go
	err := db.DB.Create(user).Error

	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := db.DB.Where("email = ?", email).First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // no user found but an error
	}

	if err != nil {
		return &models.User{}, err
	}

	return &user, nil // user found
}
