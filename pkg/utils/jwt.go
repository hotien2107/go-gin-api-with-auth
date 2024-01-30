package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secretkey"

// Function generate new Token
func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour).Unix(), // expired time is 2 hours
	})

	accessToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
