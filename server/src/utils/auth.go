package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"server/src/api/handlers"
	"server/src/graph/generated/models"
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

func (c *CustomClaims) Sign(subject *models.Account) *CustomClaims {
	jsonSubject, err := json.Marshal(subject)
	if err != nil {
		return c
	}

	c.Subject = string(jsonSubject)
	c.StandardClaims.Issuer = "Root CA"
	c.StandardClaims.IssuedAt = GetNowInMs()

	return c
}

func (c *CustomClaims) Generate(ctx context.Context) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	tokenString, err := token.SignedString([]byte(ctx.Value("env").(ENV).JWTSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJWT(ctx context.Context, tokenString string) (*CustomClaims, error) {
	c := &CustomClaims{}
	token, err := jwt.ParseWithClaims(tokenString, c, func(token *jwt.Token) (interface{}, error) {
		return []byte(ctx.Value("env").(ENV).JWTSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return c, nil
}

func Error(ctx context.Context, err error) error {
	return handlers.NewHttpError(ctx, http.StatusUnauthorized, err.Error())
}

func Allow(ctx context.Context, roles []string) error {
	claims := ctx.Value("claims").(*CustomClaims)
	for _, role := range roles {
		if claims.Role == role {
			return nil
		}
	}
	return Error(ctx, fmt.Errorf("you don't have permission to access this. role: %s", claims.Role))
}
