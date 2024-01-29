package repository

import (
	"errors"

	"gin-rest-api.com/basic/internal/db"
)

type AuthRepository struct{}

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{}
}

func (*AuthRepository) SignUp(email string, password string) (int64, error) {
	// query string
	query := `
		INSERT INTO users(email, password)
		VALUES (?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(email, password)
	if err != nil {
		return 0, err
	}

	eventId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return eventId, nil
}

func (*AuthRepository) GetHashPassByEmail(email string) (string, error) {
	// query string
	query := `
		SELECT password FROM users WHERE email = ?
	`

	row := db.DB.QueryRow(query, email)

	var hashPass string

	err := row.Scan(&hashPass)
	if err != nil {
		return "", errors.New("EMAIL IS NOT REGISTERED")
	}

	return hashPass, nil
}
