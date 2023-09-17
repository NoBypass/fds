package middleware

import (
	"context"
	"net/http"
	"server/src/api/handlers"
	"server/src/utils"
)

func Auth(ctx context.Context, next func(context.Context) http.Handler) (http.Handler, error) {
	claims, err := utils.ParseJWTbyCxt(ctx)
	if err != nil {
		return nil, handlers.NewHttpError(ctx, http.StatusUnauthorized, "invalid token")
	}
	ctx = context.WithValue(ctx, "claims", claims)
	return next(ctx), nil
}
