package services

import (
	"errors"
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

func (s *AuthService) SignUp(u *models.User) error {
	// check is email is valid
	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return errors.New("email is invalid")
	}

	// check password is valid
	// password has not contain space
	if utils.IsContainSpace(u.Password) {
		return errors.New("password must not contain spaces")
	}
	// password has contain number
	if !utils.IsContainNumber(u.Password) {
		return errors.New("password must contain number")
	}
	// password has contain capital letter
	if !utils.IsContainCapitalLetter(u.Password) {
		return errors.New("password must contain capital letter")
	}

	// hash password before store user
	hashPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	err = s.repo.SignUp(u.Email, hashPass)

	if err != nil {
		return err
	}

	return nil
}

/*
Login function
Input: email, password
Output: accessToken, refreshToken, error
*/
func (s *AuthService) Login(u *models.User) (string, string, error) {
	// get password from email
	hashPass, err := s.repo.GetHashPassByEmail(u.Email)
	if err != nil {
		return "", "", err
	}

	// check password
	isValidPassword := utils.ComparePassword(u.Password, hashPass)
	if !isValidPassword {
		return "", "", errors.New("password is invalid")
	}

	userId, err := s.repo.GetUserIdByEmail(u.Email)
	if err != nil {
		return "", "", err
	}

	//login success -> generate JWT token
	accessTokenString, err := utils.GenerateToken(u.Email, userId, false)
	if err != nil {
		return "", "", err
	}

	refreshTokenString, err := utils.GenerateToken(u.Email, userId, true)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func (s *AuthService) GenNewAccessToken(refreshToken string) (string, error) {
	// validate refresh token
	if refreshToken == "" {

		return "", errors.New("missing refresh token")
	}
	email, userId, err := utils.VerifyToken(refreshToken)
	if err != nil {
		return "", errors.New("refresh token is invalid: " + err.Error())
	}

	// generate new access token
	accessToken, err := utils.GenerateToken(email, userId, false)
	if err != nil {
		return "", errors.New("generate access token failed: " + err.Error())
	}

	return accessToken, nil
}
