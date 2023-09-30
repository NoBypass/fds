package main

import (
	"context"
	"fmt"
	"net/http"
	"server/src/api/handlers/logger"
	"server/src/api/resolvers"
	"server/src/graph/generated"
	"server/src/middleware"
	"server/src/repository/db"
	"server/src/utils"
)

// var c = cors.New(cors.Options{
// 	AllowedOrigins:   []string{"http://localhost:5173"},
// 	AllowedMethods:   []string{"GET", "POST"},
// 	AllowedHeaders:   []string{"Authorization", "Content-Type"},
// 	AllowCredentials: true,
// })

func main() {
	logger.Log("Starting server", logger.INFO)

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

	http.Handle("/ws", middleware.Auth(ctx, resolvers.WebSocketHandler))
	http.Handle("/graphql", middleware.Auth(ctx, resolvers.GraphQLHandler))

	generated.InitSchema()
	logger.Log("Server started & graphql initialized", logger.SUCCESS)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
