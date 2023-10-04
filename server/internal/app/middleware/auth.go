package middleware

import (
	"context"
	"github.com/rs/cors"
	"net/http"
	"server/internal/pkg/auth"
)

const (
	SUBSCRIBER = "subscriber"
	ADMIN      = "admin"
	USER       = "user"
	BOT        = "bot"
)

var c = cors.New(cors.Options{
	AllowedOrigins:   []string{"http://localhost:5173"},
	AllowedMethods:   []string{"GET", "POST"},
	AllowedHeaders:   []string{"Authorization", "Content-Type"},
	AllowCredentials: true,
})

func Auth(ctx context.Context, run func(ctx context.Context)) http.Handler {
	return c.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx = context.WithValue(ctx, "req", r)
		ctx = context.WithValue(ctx, "res", w)

		claims, err := auth.ParseJWT(ctx, r.Header.Get("Authorization"))
		if err == nil {
			ctx = context.WithValue(ctx, "claims", claims)
		}

		// RateLimiterMiddleware(ctx)
		run(ctx)
	}))
}
