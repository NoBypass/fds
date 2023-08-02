package utils

import (
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func GenerateJWT(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("your-secret-key")) // TODO use env variable
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
