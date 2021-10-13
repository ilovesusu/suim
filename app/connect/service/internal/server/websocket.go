package server

import (
	"github.com/ilovesusu/suim/app/connect/service/internal/server/websocket"
)

func NewWebSocketServer() *websocket.WebSocketServer {
	srv := websocket.NewServer()
	return srv
}
