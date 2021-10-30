package server

import (
	"github.com/ilovesusu/suim/app/connect/service/internal/server/tcp"
	"github.com/ilovesusu/suim/app/connect/service/internal/service"
)

func NewTCPServer(s *service.ShopAdmin) *tcp.TcpServer {
	srv := tcp.NewServer(s)
	return srv
}
