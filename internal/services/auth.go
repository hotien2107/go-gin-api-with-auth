package services

import (
	"errors"
	"log"
	"net/mail"

	"gin-rest-api.com/basic/internal/models"
	"gin-rest-api.com/basic/internal/repository"
	"gin-rest-api.com/basic/pkg/utils"
)

type AuthService struct {
	repo *repository.AuthRepository
}

func NewAuthService() *AuthService {
	// Initialize and return a new AuthService instance
	return &AuthService{
		repo: repository.NewAuthRepository(),
	}
}

func (s *AuthService) SignUp(u *models.User) (int64, error) {
	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return 0, errors.New("EMAIL IS INVALID")
	}

	// hash password before store user
	hashPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return 0, err
	}

	log.Println(s)

	eventId, err := s.repo.SignUp(u.Email, hashPass)

	if err != nil {
		return 0, err
	}

	return eventId, nil
}

func (s *AuthService) Login(u *models.User) error {
	// get password from email
	hashPass, err := s.repo.GetHashPassByEmail(u.Email)
	if err != nil {
		return err
	}

	// check password
	isValidPassword := utils.ComparePassword(u.Password, hashPass)
	if !isValidPassword {
		return errors.New("PASSWORD IS INVALID")
	}

	return nil
}
