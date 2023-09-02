package main

import (
	"fmt"
	"github.com/rs/cors"
	"net/http"
	"server/src/api/resolvers"
	"server/src/graph/generated"
)

func main() {
	fmt.Println("Starting server...")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // Set the allowed origins here
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	http.Handle("/graphql", c.Handler(resolvers.GraphQLHandler(&generated.RootSchema)))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	// dbutils.CloseDB(driver, ctx)
}
