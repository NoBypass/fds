package middleware

import (
	"context"
	"net/http"
	"server/src/auth"
)

const (
	SUBSCRIBER = "subscriber"
	ADMIN      = "admin"
	USER       = "user"
	BOT        = "bot"
)

func Auth(ctx context.Context, run func(ctx context.Context)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx = context.WithValue(ctx, "req", r)
		ctx = context.WithValue(ctx, "res", w)

		claims, err := auth.ParseJWT(ctx, r.Header.Get("Authorization"))
		if err == nil {
			ctx = context.WithValue(ctx, "claims", claims)
		}

		RateLimiterMiddleware(ctx)
		run(ctx)
	})
}
