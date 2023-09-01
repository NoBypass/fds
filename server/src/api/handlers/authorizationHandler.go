package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"server/src/graph/generated"
	"server/src/utils"
)

func NewClaims(role string) (*CustomClaims, error) {
	possible := []string{"admin", "bot", "subscriber", "user"}
	for _, p := range possible {
		if p == role {
			return &CustomClaims{
				Role: role,
			}, nil
		}
	}
	return nil, fmt.Errorf("invalid role")
}

type CustomClaims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}

func (c *CustomClaims) Sign(subject *generated.Account) *CustomClaims {
	jsonSubject, err := json.Marshal(subject)
	if err != nil {
		return c
	}

	c.Subject = string(jsonSubject)
	c.StandardClaims.Issuer = "Root CA"
	c.StandardClaims.IssuedAt = utils.GetNowInMs()

	return c
}

func (c *CustomClaims) Generate() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	tokenString, err := token.SignedString([]byte("your-secret-key")) // TODO use env variable
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJWT(tokenString string) (*CustomClaims, error) {
	c := &CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, c, func(token *jwt.Token) (interface{}, error) {
		return []byte("your-secret-key"), nil // TODO use env variable
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return c, nil
}
