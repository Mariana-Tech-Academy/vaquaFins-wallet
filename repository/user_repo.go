package repository

import (
	"time"
	"vaqua/models"

	"errors"
	"vaqua/db"

	"gorm.io/gorm"
)

//import "vaqua/models"

type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
	//EditUser(user *models.User) error
	BlacklistToken(jti string, expiresAt time.Time) error
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

// BlacklistToken adds a JWT to a blacklist table
func (r *UserRepo) BlacklistToken(jti string, expiresAt time.Time) error {
	blacklistedToken := &models.BlacklistedToken{
		JTI:       jti,
		ExpiresAt: expiresAt,
	}
	err := db.DB.Create(blacklistedToken).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) IsTokenBlacklisted(jti string) (bool, error) {
	var token models.BlacklistedToken
	result := db.DB.Where("jti = ?", jti).First(&token)

	if result.RowsAffected > 0 {
		return true, nil
	}

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return false, result.Error
	}

	return false, nil
}

