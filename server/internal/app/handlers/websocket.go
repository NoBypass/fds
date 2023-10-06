package handlers

import (
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

func WebSocketHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "could not upgrade connection to websocket: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer conn.Close()

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
					http.Error(w, "could not write message to websocket: "+err.Error(), http.StatusInternalServerError)
					break
				}
				continue
			}

			handler := GraphQLHandler()
			handler.ServeHTTP(w, r)
		}
	})
}
