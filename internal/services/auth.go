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
	// check is email is valid
	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return 0, errors.New("EMAIL IS INVALID")
	}

	// check password is valid
	// password has not contain space
	if utils.IsContainSpace(u.Password) {
		return 0, errors.New("PASSWORD HAS SPACE")
	}
	// password has contain number
	if !utils.IsContainNumber(u.Password) {
		return 0, errors.New("PASSWORD MUST CONTAIN NUMBER")
	}
	// password has contain capital letter
	if !utils.IsContainCapitalLetter(u.Password) {
		return 0, errors.New("PASSWORD MUST CONTAIN CAPITAL LETTER")
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

func (s *AuthService) Login(u *models.User) (string, error) {
	// get password from email
	hashPass, err := s.repo.GetHashPassByEmail(u.Email)
	if err != nil {
		return "", err
	}

	// check password
	isValidPassword := utils.ComparePassword(u.Password, hashPass)
	if !isValidPassword {
		return "", errors.New("PASSWORD IS INVALID")
	}

	userId, err := s.repo.GetUserIdByEmail(u.Email)
	if err != nil {
		return "", err
	}

	//login success -> generate JWT token
	tokenString, err := utils.GenerateToken(u.Email, userId)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
