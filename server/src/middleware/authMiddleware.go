package middleware

import (
	"context"
	"net/http"
)

func Auth(ctx context.Context, next func(context.Context) http.Handler) (http.Handler, error) {
	// TODO
	return next(ctx), nil
}
