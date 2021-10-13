package tcp

import (
	"context"
	"fmt"
	"github.com/alberliu/gn"
	"github.com/ilovesusu/suim/app/connect/service/internal/server/conn"
	"net/url"
	"time"
)

type TcpServer struct {
	server *gn.Server
}

func NewServer() *TcpServer {
	//gn.SetLogger(logger.Sugar)

	srv := TcpServer{}

	var err error
	// todo address conf 中获取
	srv.server, err = gn.NewServer("127.0.0.1:8888", &conn.Handler{},
		gn.NewHeaderLenDecoder(2),
		gn.WithReadBufferLen(256),
		gn.WithTimeout(5*time.Minute, 11*time.Minute),
		gn.WithAcceptGNum(10),
		gn.WithIOGNum(100))
	if err != nil {
		//logger.Sugar.Error(err)
		panic(err)
	}
	return &srv
}

func (s *TcpServer) Start(ctx context.Context) error {
	s.server.Run()
	return nil
}

func (s *TcpServer) Stop(ctx context.Context) error {
	// todo stop is shit
	s.server.Stop()
	time.Sleep(3 * time.Second)
	fmt.Println(123)
	return nil
}

func (s *TcpServer) Endpoint() (*url.URL, error) {
	// todo 注册中心实现方法
	return url.Parse("127.0.0.1")
}
