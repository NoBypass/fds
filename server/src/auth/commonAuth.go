package auth

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"server/src/api/handlers"
)

type CustomClaims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}

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

func Allow(ctx context.Context, roles []string) error {
	claims := ctx.Value("claims").(*CustomClaims)
	for _, role := range roles {
		if claims.Role == role {
			return nil
		}
	}
	return handlers.HttpError(ctx, http.StatusUnauthorized, fmt.Sprintf("you don't have permission to access this. role: %s", claims.Role))
}