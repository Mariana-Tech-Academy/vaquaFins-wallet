package service

import (
	"errors"
	"vaqua/middleware"
	"vaqua/models"
	"vaqua/repository"
	"vaqua/utils"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) CreateUser(user *models.User) error {
	//check if user exist already
	userExists, err := s.Repo.GetUserByEmail(user.Email)

	if err != nil {
		return err //DB error
	}

	if userExists != nil {
		return errors.New("email already in use")
	}
	// hash the password
	hashpass, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashpass

	// call the create method
	err = s.Repo.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) LogInUser(request models.User) (string, error) {
	user, err := s.Repo.GetUserByEmail(request.Email)
	if err != nil {
		return "", err
	}
	err = utils.ComparePassword(user.Password, request.Password)
	if err != nil {
		return "", err
	}
	token, err := middleware.GenerateJWT(user.ID) //taken out user.Role, we don't need it right? (P)
	if err != nil {
		return "", err
	}
	return token, nil
}
