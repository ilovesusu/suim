package server

import (
	"github.com/ilovesusu/suim/app/connect/service/internal/server/websocket"
	"github.com/ilovesusu/suim/app/connect/service/internal/service"
)

func NewWebSocketServer(s *service.ShopAdmin) *websocket.WebSocketServer {
	srv := websocket.NewServer(s)
	return srv
}
