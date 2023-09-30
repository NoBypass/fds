package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"server/src/api/handlers/logger"
)

func Respond[T any](ctx context.Context, data *T) {
	if ctx.Value("conn") != nil {
		conn := ctx.Value("conn").(*websocket.Conn)
		err := conn.WriteJSON(data)
		if err != nil {
			Error(ctx, err)
		}
		return
	}

	res := ctx.Value("res").(http.ResponseWriter)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	body, err := json.Marshal(*data)
	if err != nil {
		Error(ctx, err)
	}
	_, err = res.Write(body)
	if err != nil {
		Error(ctx, err)
	}
}

func Error(ctx context.Context, err error) {
	res := ctx.Value("res").(http.ResponseWriter)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusInternalServerError)

	_, err = res.Write([]byte(err.Error()))
	if err != nil {
		logger.Error(err, "Error while writing error response")
	}
}

func HttpError(ctx context.Context, status int, message string) error {
	res := ctx.Value("res").(http.ResponseWriter)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(status)

	_, err := res.Write([]byte(message))
	if err != nil {
		Error(ctx, err)
	}
	return fmt.Errorf(message)
}
