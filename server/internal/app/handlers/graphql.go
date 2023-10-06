package handlers

import (
	"context"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"net/http"
	"server/internal/pkg/generated"
	"time"
)

type GraphQLBody struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

func GraphQLHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()

		req := ctx.Value("req").(*http.Request)
		res := ctx.Value("res").(http.ResponseWriter)
		c := make(chan bool, 1)

		go func() {
			h := handler.New(&handler.Config{
				Schema:     &generated.RootSchema,
				Pretty:     true,
				GraphiQL:   false,
				Playground: true,
			})

			// If-Else statement to use GraphiQL along with the GraphQL handler
			if req.Method == "POST" {
				requestBody := req.Context().Value("requestBody").(GraphQLBody)

				result := graphql.Do(graphql.Params{
					Schema:         generated.RootSchema,
					RequestString:  requestBody.Query,
					Context:        ctx,
					VariableValues: requestBody.Variables,
				})

				_ = json.NewEncoder(res).Encode(result)
			} else {
				h.ServeHTTP(res, req)
			}
			c <- true
		}()

		select {
		case <-c:
			return
		case <-ctx.Done():
			http.Error(res, "Request took too long to execute.", http.StatusRequestTimeout)
			return
		}
	})
}
