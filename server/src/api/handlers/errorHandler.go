package handlers

import (
	"context"
	"errors"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"net/http"
)

func NewHttpError(ctx context.Context, status int, message string) error {
	w := ctx.Value("response").(http.ResponseWriter)
	http.Error(w, message, status)
	return errors.New(message)
}

func CheckIfFound(ctx context.Context, result *neo4j.EagerResult, message string) {
	if result.Records == nil || len(result.Records) == 0 {
		err := NewHttpError(ctx, http.StatusNotFound, message)
		if err != nil {
			return
		}
	}
}
