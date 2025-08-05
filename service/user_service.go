package service

import (
	"vaqua/middleware"
	"errors"
	"vaqua/models"
	"vaqua/repository"
	"vaqua/utils"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) RegisterUser(user *models.User) error {
    //check if user exist already
    _, err := s.Repo.GetUserByEmail(user.Email)
    if err == nil {
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
// check if user is in db
// compare password
// generate token
func (s *UserService) LoginUser(request models.User) (string, error) {
    user, err := s.Repo.GetUserByEmail(request.Email)
    if err != nil {
        return "", err
    }
    err = utils.ComparePassword(user.Password, request.Password)
    if err != nil {
        return "", err
    }
    token, err := middleware.GenerateJWT(user.ID, user.Role)
    if err != nil {
        return "", err
    }
    return token, nil
}