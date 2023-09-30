package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"server/src/graph/generated/models"
	"server/src/utils"
)

func (c *CustomClaims) Sign(subject *models.Account) *CustomClaims {
	jsonSubject, err := json.Marshal(subject)
	if err != nil {
		return c
	}

	c.Subject = string(jsonSubject)
	c.StandardClaims.Issuer = "Root CA"
	c.StandardClaims.IssuedAt = utils.GetNowInMs()

	return c
}

func (c *CustomClaims) Generate(ctx context.Context) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	tokenString, err := token.SignedString([]byte(ctx.Value("env").(utils.ENV).JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJWT(ctx context.Context, tokenString string) (*CustomClaims, error) {
	c := &CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, c, func(token *jwt.Token) (interface{}, error) {
		return []byte(ctx.Value("env").(utils.ENV).JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return c, nil
}
