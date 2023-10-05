package main

import (
	"context"
	"github.com/fatih/color"
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
	color.New(color.FgHiMagenta, color.Bold).Println(`
8888888  88888     88888      88888                                        
88       88  88   88   88    88   88   8888                   8888         
88888    88   88   888        888     88  88  88 88  88  88  88  88  88 88 
88       88   88     888        888   888888  888 8  88  88  888888  888 8 
88       88  88   88   88    88   88  88      88      8888   88      88    
88       88888     88888      88888    88888  88       88     88888  88`)
	color.New(color.FgHiWhite).Println("\nLogger output:")

	generated.InitSchema()

	r := mux.NewRouter()
	env := misc.FetchEnv()
	ctx := context.Background()
	ctx = context.WithValue(ctx, "env", env)

	driver, cache := db.Connect(ctx)
	ctx = context.WithValue(ctx, "driver", driver)
	ctx = context.WithValue(ctx, "cache", cache)

	r.Use(middleware.Logger)

	// TODO use middleware to handle auth, rate limiting, etc.
	r.Handle("/ws", middleware.Auth(ctx, handlers.WebSocketHandler))
	r.Handle("/graphql", middleware.Auth(ctx, handlers.GraphQLHandler))

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
