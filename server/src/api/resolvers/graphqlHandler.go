package resolvers

import (
	"context"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"net/http"
	"server/src/api/handlers/logger"
	"server/src/graph/generated"
	"server/src/utils"
)

func GraphQLHandler(ctx context.Context) http.Handler {
	h := handler.New(&handler.Config{
		Schema:   &generated.RootSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// If-Else statement to use GraphiQL along with the GraphQL handler
		if r.Method == "POST" {
			ctx = context.WithValue(ctx, "request", r)
			ctx = context.WithValue(ctx, "response", w)
			claims, err := utils.ParseJWT(r.Header.Get("Authorization"))
			if err == nil {
				ctx = context.WithValue(ctx, "claims", claims)
			}

			var requestBody struct {
				Query         string `json:"query"`
				OperationName string `json:"operationName"`
			}
			if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
				http.Error(w, "Invalid request body", http.StatusBadRequest)
				return
			}

			result := graphql.Do(graphql.Params{
				Schema:        generated.RootSchema,
				RequestString: requestBody.Query,
				Context:       ctx,
			})
			if len(result.Errors) != 0 {
				http.Error(w, result.Errors[0].Message, http.StatusInternalServerError)
				return
			}

			err = json.NewEncoder(w).Encode(result)
			if err != nil {
				http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
				return
			}
			logger.Log("GraphQL query executed", logger.SUCCESS, requestBody.OperationName)
		} else {
			h.ServeHTTP(w, r)
		}
	})
}
