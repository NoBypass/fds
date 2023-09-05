package resolvers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/graphql-go/graphql"
	"net/http"
	"server/src/api/handlers"
	"server/src/api/handlers/logger"
	"server/src/utils"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketHandler(schema *graphql.Schema, ctx context.Context) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "could not upgrade connection to websocket: "+err.Error(), http.StatusBadRequest)
			return
		}
		defer conn.Close()

		ctx = context.WithValue(ctx, "request", r)
		ctx = context.WithValue(ctx, "response", w)

		claims, err := handlers.ParseJWT(r.Header.Get("Authorization"))
		if err == nil {
			ctx = context.WithValue(ctx, "claims", claims)
		}

		requestLimit := 3
		if claims != nil {
			switch claims.Role {
			case "admin":
				requestLimit = -1
			case "bot":
				requestLimit = -1
			case "subscriber":
				requestLimit = -1
			case "user":
				requestLimit = 500
			}
		}

		wsId := utils.GenerateUUID(claims)
		if claims != nil {
			logger.Log(fmt.Sprintf("websocket connection opened with '%s' with role '%s', ID:%s", claims.StandardClaims.Subject, claims.Role, wsId), logger.SUCCESS)
		} else {
			logger.Log("anonymous websocket connection opened, ID:"+wsId, logger.WARN)
		}

		for {
			_, msgBytes, err := conn.ReadMessage()
			if err != nil {
				logger.Log("client disconnected, ID:"+wsId, logger.INFO)
				break
			}

			requestLimit--
			if requestLimit == 0 {
				logger.Log("disconnecting websocket connection due to rate limit, ID:"+wsId, logger.INFO)
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
					logger.Error(fmt.Errorf("disconnecting; could not send response: %w, ID:%s", err, wsId))
					break
				}
				continue
			}

			result := graphql.Do(graphql.Params{
				Schema:        *schema,
				RequestString: query.Query,
				Context:       ctx,
			})
			if len(result.Errors) != 0 {
				fmt.Println(result.Errors[0].Message)
				err := conn.WriteMessage(websocket.TextMessage, []byte(result.Errors[0].Message))
				if err != nil {
					logger.Error(fmt.Errorf("disconnecting; could not send response: %w, ID:%s", err, wsId))
					break
				}
				continue
			}

			var response struct {
				OperationName string      `json:"operationName"`
				Data          interface{} `json:"data"`
				Errors        interface{} `json:"errors"`
			}
			response.OperationName = query.OperationName
			response.Data = result.Data
			response.Errors = result.Errors

			jsonResponse, err := json.Marshal(response)
			if err != nil {
				if err != nil {
					logger.Error(fmt.Errorf("disconnecting; could not send response: %w, ID:%s", err, wsId))
					break
				}
				continue
			}

			err = conn.WriteMessage(websocket.TextMessage, jsonResponse)
			if err != nil {
				logger.Error(fmt.Errorf("disconnecting; could not send response: %w, ID:%s", err, wsId))
				break
			}

			err = json.NewEncoder(w).Encode(result)
			if err != nil {
				err := conn.WriteMessage(websocket.TextMessage, []byte("could not send response"))
				if err != nil {
					logger.Error(fmt.Errorf("disconnecting; could not send response: %w, ID:%s", err, wsId))
					break
				}
				break
			}
		}
	})
}
