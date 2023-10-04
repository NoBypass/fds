package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"server/src/api/handlers/logger"
	"server/src/api/resolvers"
	"server/src/graph/generated"
	"server/src/middleware"
	"server/src/repository/db"
	"server/src/utils"
	"time"
)

func main() {
	logger.Log("Starting server", logger.INFO)

	r := mux.NewRouter()
	env := utils.FetchEnv()
	ctx := context.Background()
	ctx = context.WithValue(ctx, "env", env)

	driver, cache, err := db.Connect(ctx)
	ctx = context.WithValue(ctx, "driver", driver)
	ctx = context.WithValue(ctx, "cache", cache)
	if err != nil {
		logger.Error(err)
		return
	} else {
		logger.Log("Connected to database", logger.SUCCESS)
	}

	// TODO use middleware to handle auth, rate limiting, etc.
	r.Handle("/ws", middleware.Auth(ctx, resolvers.WebSocketHandler))
	r.Handle("/graphql", middleware.Auth(ctx, resolvers.GraphQLHandler))

	generated.InitSchema()

	logger.Log("Server started & graphql initialized", logger.SUCCESS)
	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
