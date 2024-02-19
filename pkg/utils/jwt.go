package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secretkey"

// Function generate new Token
func GenerateToken(email string, userId int64, isRefreshToken bool) (string, error) {
	var expiredTime int64
	if isRefreshToken {
		expiredTime = time.Now().Add(time.Hour * 8760).Unix() // 1 year
	} else {
		expiredTime = time.Now().Add(time.Hour).Unix() // 1 hour
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    expiredTime,
	})

	accessToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

// VerifyToken is a function that checks the validity of a JWT token
// return userId & errors
func VerifyToken(token string) (int64, error) {
	// Initialize a new JWT parser with the given token
	parser, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Verify if the token's signing method is HMAC
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("UNEXPECTED SIGNING METHOD")
		}
		// Return the secret key as the verification key
		return []byte(secretKey), nil
	})

	if err != nil {
		// If there was an error parsing the token, return an error
		return 0, errors.New("Cannot parse token: " + err.Error())
	}

	// Check if the token is valid
	tokenIsValid := parser.Valid
	if !tokenIsValid {
		return 0, errors.New("TOKEN IS INVALID")
	}

	// Type assert the token's claims to jwt.MapClaims
	claims, ok := parser.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("INVALID TOKEN CLAIMS")
	}

	// Check if the token has expired
	if _, ok = claims["exp"]; ok {
		unixTime := claims["exp"].(float64)
		currentTime := time.Now().Unix()
		if currentTime > int64(unixTime) {
			return 0, errors.New("TOKEN HAS EXPIRED")
		}
	}

	userId, ok := claims["userId"].(float64)
	if !ok {
		return 0, errors.New("CANNOT GET USER ID IN TOKEN")
	}

	return int64(userId), nil
}
