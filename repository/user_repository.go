package repository

import (
"vaqua/database"
"vaqua/models"
"gorm.io/gorm"
)

func GetUserByID(id uint) (*models.User, error) {

	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func UpdateUserBalance(tx *gorm.DB, user *models.User) error {
return tx.Save(user).Error
}
