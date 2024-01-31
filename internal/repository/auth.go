package repository

import (
	"errors"

	"gin-rest-api.com/basic/internal/db"
)

type AuthRepository struct{}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

func (*AuthRepository) SignUp(email string, password string) error {
	// query string
	query := `
		INSERT INTO users(email, password) VALUES($1, $2)
	`

	_, err := db.DB.Exec(query, email, password)
	if err != nil {
		return err
	}

	return nil
}

func (*AuthRepository) GetHashPassByEmail(email string) (string, error) {
	// query string
	query := `
		SELECT password FROM users WHERE email = $1
	`

	row := db.DB.QueryRow(query, email)

	var hashPass string

	err := row.Scan(&hashPass)
	if err != nil {
		return "", errors.New("GET PASSWORD FROM EMAIL FAILED: " + err.Error())
	}

	return hashPass, nil
}

func (*AuthRepository) GetUserIdByEmail(email string) (int64, error) {
	// query string
	query := `
		SELECT id FROM users WHERE email = $1
	`

	row := db.DB.QueryRow(query, email)

	var userId int64

	err := row.Scan(&userId)
	if err != nil {
		return 0, errors.New("GET USER ID FROM EMAIL FAILED: " + err.Error())
	}

	return userId, nil
}
