package main

import (
	"fmt"
	"github.com/rs/cors"
	"net/http"
	"server/src/api/handlers/logger"
	"server/src/api/resolvers"
	"server/src/graph/generated"
	"server/src/utils"
)

func main() {
	logger.Log("Starting server", logger.INFO)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	_, ctx, err := utils.ConnectDB()
	if err != nil {
		logger.Error(err)
		return
	} else {
		logger.Log("Connected to database", logger.SUCCESS)
	}

	http.Handle("/ws", resolvers.WebSocketHandler(ctx))
	http.Handle("/graphql", c.Handler(resolvers.GraphQLHandler(ctx)))

	generated.InitSchema()
	logger.Log("Server started & graphql initialized", logger.SUCCESS)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
