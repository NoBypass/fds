package handlers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketHandler(ctx context.Context) {
	req := ctx.Value("request").(*http.Request)
	res := ctx.Value("response").(http.ResponseWriter)

	conn, err := upgrader.Upgrade(res, req, nil)
	if err != nil {
		http.Error(res, "could not upgrade connection to websocket: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer conn.Close()
	ctx = context.WithValue(ctx, "conn", conn)

	for {
		_, msgBytes, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var query struct {
			OperationName string `json:"operationName"`
			Query         string `json:"query"`
			Variables     any    `json:"variables"`
		}
		err = json.Unmarshal(msgBytes, &query)
		if err != nil {
			err := conn.WriteMessage(websocket.TextMessage, []byte("invalid request"))
			if err != nil {
				res := ctx.Value("res").(http.ResponseWriter)
				res.WriteHeader(http.StatusInternalServerError)
				res.Write([]byte("invalid request"))
				break
			}
			continue
		}

		GraphQLHandler(ctx)
	}
}
