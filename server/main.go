package main

import (
	"fmt"
	"net/http"
	"server/api/handlers"
	"server/api/schemas"
)

func main() {
	fmt.Println("Starting server...")

	http.Handle("/graphql", handlers.GraphQLHandler(&schemas.RootSchema))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

	// dbutils.CloseDB(driver, ctx)
}
