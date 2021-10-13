package websocket

import (
	"context"
	"fmt"
	"github.com/alberliu/gn"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/ilovesusu/suim/app/connect/service/internal/server/conn"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 65536,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocketServer struct {
	server *http.Server
}

func NewServer() *WebSocketServer {
	//gn.SetLogger(logger.Sugar)

	srv := WebSocketServer{}
	r := mux.NewRouter()
	r.HandleFunc("/ws", wsHandler)
	srv.server = &http.Server{Addr: "127.0.0.1:8887", Handler: r}
	return &srv
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	wsConn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		//logger.Sugar.Error(err)
		return
	}

	conn := &conn.Conn{
		CoonType: conn.ConnTypeWS,
		WS:       wsConn,
	}
	DoConn(conn)
}

// DoConn 处理连接
func DoConn(conn *conn.Conn) {
	defer RecoverPanic()

	for {
		err := conn.WS.SetReadDeadline(time.Now().Add(12 * time.Minute))
		if err != nil {
			HandleReadErr(conn, err)
			return
		}
		_, data, err := conn.WS.ReadMessage()
		if err != nil {
			HandleReadErr(conn, err)
			return
		}

		conn.HandleMessage(data)
	}
}

// HandleReadErr 读取conn错误
func HandleReadErr(conn *conn.Conn, err error) {
	//logger.Logger.Debug("read tcp error：", zap.Int64("user_id", conn.UserId),
	//	zap.Int64("device_id", conn.DeviceId), zap.Error(err))
	str := err.Error()
	// 服务器主动关闭连接
	if strings.HasSuffix(str, "use of closed network connection") {
		return
	}

	conn.Close()
	// 客户端主动关闭连接或者异常程序退出
	if err == io.EOF {
		return
	}
	// SetReadDeadline 之后，超时返回的错误
	if strings.HasSuffix(str, "i/o timeout") {
		return
	}
}

// RecoverPanic 恢复panic
func RecoverPanic() {
	err := recover()
	if err != nil {
		//logger.Logger.DPanic("panic", zap.Any("panic", err), zap.String("stack", GetStackInfo()))
	}
}

// GetStackInfo 获取Panic堆栈信息
func GetStackInfo() string {
	buf := make([]byte, 4096)
	n := runtime.Stack(buf, false)
	return fmt.Sprintf("%s", buf[:n])
}

func (s *WebSocketServer) Start(ctx context.Context) error {
	err := s.server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *WebSocketServer) Stop(ctx context.Context) error {
	err := s.server.Close()
	if err != nil {
		return err
	}
	return nil
}

func (s *WebSocketServer) Endpoint() (*url.URL, error) {
	// todo 注册中心实现方法
	return url.Parse("127.0.0.1")
}

type TcpHandler struct{}

func (*TcpHandler) OnConnect(c *gn.Conn) {
	// 初始化连接数据
	conn := &conn.Conn{
		CoonType: conn.CoonTypeTCP,
		TCP:      c,
	}
	c.SetData(conn)
	//todo log的实现
	//logger.Logger.Debug("connect:", zap.Int32("fd", c.GetFd()), zap.String("addr", c.GetAddr()))
}
