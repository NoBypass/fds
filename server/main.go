package main

import (
	"fmt"
	"github.com/rs/cors"
	"net/http"
	"server/src/api/resolvers"
	"server/src/graph/generated"
	"server/src/utils"
)

func main() {
	fmt.Println("Starting server...")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	_, ctx, _ := utils.ConnectDB()
	http.Handle("/ws", resolvers.WebSocketHandler(&generated.RootSchema, ctx))
	http.Handle("/graphql", c.Handler(resolvers.GraphQLHandler(&generated.RootSchema, ctx)))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	// dbutils.CloseDB(driver, ctx)
}
