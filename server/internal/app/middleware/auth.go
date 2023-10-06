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

func Auth(next http.Handler) http.Handler {
	return c.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, err := auth.ParseJWT(r.Header.Get("Authorization"))
		if err == nil {
			r = r.WithContext(context.WithValue(r.Context(), "claims", claims))
		}

		next.ServeHTTP(w, r)
	}))
}
