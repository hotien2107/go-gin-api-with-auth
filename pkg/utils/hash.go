package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(pass string, hasPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hasPass), []byte(pass))
	return err == nil
}
