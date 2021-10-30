package conn

import (
	"github.com/alberliu/gn"
	"github.com/ilovesusu/suim/app/connect/service/internal/service"
)

type Handler struct {
	S *service.ShopAdmin
}

func (h *Handler) OnConnect(c *gn.Conn) {
	// 初始化连接数据
	conn := &Conn{
		CoonType: CoonTypeTCP,
		TCP:      c,
		S:        h.S,
	}
	c.SetData(conn)
	//todo log的实现
	//logger.Logger.Debug("connect:", zap.Int32("fd", c.GetFd()), zap.String("addr", c.GetAddr()))
}

func (*Handler) OnMessage(c *gn.Conn, bytes []byte) {
	conn := c.GetData().(*Conn)
	conn.HandleMessage(bytes)
}

func (*Handler) OnClose(c *gn.Conn, err error) {
	conn := c.GetData().(*Conn)
	//logger.Logger.Debug("close", zap.String("addr", c.GetAddr()), zap.Int64("user_id", connect.UserId),
	//	zap.Int64("device_id", connect.DeviceId), zap.Error(err))

	DeleteConn(conn.DeviceId)

	if conn.UserId != 0 {
		//todo 调用客户端离线rpc方法
		//_, _ = rpc.LogicIntClient.Offline(context.TODO(), &pb.OfflineReq{
		//	UserId:     connect.UserId,
		//	DeviceId:   connect.DeviceId,
		//	ClientAddr: c.GetAddr(),
		//})
	}
}
