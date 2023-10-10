package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"server/internal/app/middleware"
	"server/internal/app/resolvers"
	"server/internal/pkg/generated"
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

	r := mux.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Auth)
	r.Use(middleware.RateLimiter)

	r.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	r.Handle("/graphql", handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolvers.Resolver{}})))

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
