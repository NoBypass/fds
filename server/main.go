package main

import (
	"fmt"
	"net/http"
	"server/src/api/handlers/logger"
	"server/src/api/resolvers"
	"server/src/graph/generated"
	"server/src/middleware"
	"server/src/utils"
)

func main() {
	logger.Log("Starting server", logger.INFO)

	_, ctx, err := utils.ConnectDB()
	if err != nil {
		logger.Error(err)
		return
	} else {
		logger.Log("Connected to database", logger.SUCCESS)
	}

	ws, err := middleware.Auth(ctx, resolvers.WebSocketHandler)
	if err != nil {
		logger.Error(err)
		return
	}
	graphql, err := middleware.Auth(ctx, resolvers.GraphQLHandler)
	if err != nil {
		logger.Error(err)
		return
	}

	http.Handle("/ws", ws)
	http.Handle("/graphql", graphql)

	generated.InitSchema()
	logger.Log("Server started & graphql initialized", logger.SUCCESS)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
