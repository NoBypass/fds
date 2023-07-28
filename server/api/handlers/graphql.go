package handlers

import (
	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
	dbutils "server/db/utils"
)

func GraphQLHandler(schema *graphql.Schema) http.Handler {
	h := handler.New(&handler.Config{
		Schema:   schema,
		Pretty:   true,
		GraphiQL: true,
	})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If-Else statement to use GraphiQL along with the GraphQL handler
		if r.Method == "POST" {
			_, ctx, _ := dbutils.ConnectDB()

			var requestBody struct {
				Query string `json:"query"`
			}
			if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}

			result := graphql.Do(graphql.Params{
				Schema:        *schema,
				RequestString: requestBody.Query,
				Context:       ctx,
			})
			if len(result.Errors) > 0 {
				log.Printf("wrong result, unexpected errors: %v", result.Errors)
				return
			}

			err := json.NewEncoder(w).Encode(result)
			if err != nil {
				return
			}
		} else {
			h.ServeHTTP(w, r)
		}
	})
}
