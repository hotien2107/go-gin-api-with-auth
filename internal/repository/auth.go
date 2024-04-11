package repository

import (
	"errors"
	"fmt"

	"gin-rest-api.com/basic/internal/db/postgres"
)

type AuthRepository struct {
	*postgres.PsqlDB
}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{
		postgres.NewPsqlDB(),
	}
}

func (r *AuthRepository) SignUp(email string, password string) error {
	// query string
	query := `
		INSERT INTO users(email, password) VALUES($1, $2)
	`

	fmt.Println("r.DB", r.DB)

	_, err := r.DB.Exec(query, email, password)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthRepository) GetHashPassByEmail(email string) (string, error) {
	// query string
	query := `
		SELECT password FROM users WHERE email = $1
	`

	row := r.DB.QueryRow(query, email)

	var hashPass string

	err := row.Scan(&hashPass)
	if err != nil {
		return "", errors.New("GET PASSWORD FROM EMAIL FAILED: " + err.Error())
	}

	return hashPass, nil
}

func (r *AuthRepository) GetUserIdByEmail(email string) (int64, error) {
	// query string
	query := `
		SELECT id FROM users WHERE email = $1
	`

	row := r.DB.QueryRow(query, email)

	var userId int64

	err := row.Scan(&userId)
	if err != nil {
		return 0, errors.New("GET USER ID FROM EMAIL FAILED: " + err.Error())
	}

	return userId, nil
}
