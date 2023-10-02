package resolvers

import (
	"context"
	"encoding/json"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"net/http"
	"server/src/api/handlers"
	"server/src/graph/generated"
)

func GraphQLHandler(ctx context.Context) {
	req := ctx.Value("req").(*http.Request)
	res := ctx.Value("res").(http.ResponseWriter)
	h := handler.New(&handler.Config{
		Schema:   &generated.RootSchema,
		Pretty:   true,
		GraphiQL: true,
	})

	// If-Else statement to use GraphiQL along with the GraphQL handler
	if req.Method == "POST" {
		var requestBody struct {
			Query         string                 `json:"query"`
			OperationName string                 `json:"operationName"`
			Variables     map[string]interface{} `json:"variables"`
		}
		if err := json.NewDecoder(req.Body).Decode(&requestBody); err != nil {
			handlers.HttpError(ctx, http.StatusBadRequest, "Invalid request body")
			return
		}

		result := graphql.Do(graphql.Params{
			Schema:         generated.RootSchema,
			RequestString:  requestBody.Query,
			Context:        ctx,
			VariableValues: requestBody.Variables,
		})
		if len(result.Errors) != 0 {
			handlers.Error(ctx, result.Errors[0])
			return
		}

		handlers.Respond(ctx, &result)
	} else {
		h.ServeHTTP(res, req)
	}
}
