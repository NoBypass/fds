package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"server/internal/app/handlers"
	"server/internal/app/middleware"
	"server/internal/pkg/db"
	"server/internal/pkg/generated"
	"server/internal/pkg/misc"
	"time"
)

func main() {
	fmt.Println("Starting FDS server")
	generated.InitSchema()

	r := mux.NewRouter()
	env := misc.FetchEnv()
	ctx := context.Background()
	ctx = context.WithValue(ctx, "env", env)

	driver, cache := db.Connect(ctx)
	ctx = context.WithValue(ctx, "driver", driver)
	ctx = context.WithValue(ctx, "cache", cache)
	fmt.Println("Connected to database")

	// TODO use middleware to handle auth, rate limiting, etc.
	r.Handle("/ws", middleware.Auth(ctx, handlers.WebSocketHandler))
	r.Handle("/graphql", middleware.Auth(ctx, handlers.GraphQLHandler))

	fmt.Println("Server started & graphql initialized")
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
