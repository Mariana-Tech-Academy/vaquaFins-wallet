package service

import (
	"errors"
	"fmt"
	"time"
	"vaqua/middleware"
	"vaqua/models"
	"vaqua/repository"
	"vaqua/utils"

	"github.com/golang-jwt/jwt/v5"
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

// check if user is in db
// compare password
// generate token
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

func (s *UserService) LogoutUser(tokenString string) error {
	token, err := middleware.VerifyJWT(tokenString) 
	if err != nil {
		return fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("unable to extract claims")
	}

	jti, ok := claims["jti"].(string)
	if !ok {
		return fmt.Errorf("jti claim not found in token")
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return fmt.Errorf("exp claim not found in token")
	}

	expirationTime := time.Unix(int64(exp), 0)

	
	return s.Repo.BlacklistToken(jti, expirationTime)
}
