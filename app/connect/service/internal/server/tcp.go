package server

import "github.com/ilovesusu/suim/app/connect/service/internal/server/tcp"

func NewTCPServer()  *tcp.TcpServer{
	srv := tcp.NewServer()
	return srv
}
