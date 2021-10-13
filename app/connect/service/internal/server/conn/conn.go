package conn

import (
	"container/list"
	"sync"
	"time"

	"github.com/alberliu/gn"
	"github.com/gorilla/websocket"
)

const (
	CoonTypeTCP int8 = 1 // tcp连接
	ConnTypeWS  int8 = 2 // websocket连接
)

var encoder = gn.NewHeaderLenEncoder(2, 1024)

type Conn struct {
	CoonType int8            // 连接类型
	TCP      *gn.Conn        // tcp连接
	WSMutex  sync.Mutex      // WS写锁
	WS       *websocket.Conn // websocket连接
	UserId   int64           // 用户ID
	DeviceId int64           // 设备ID
	RoomId   int64           // 订阅的房间ID
	Element  *list.Element   // 链表节点
}

// Write 写入数据
func (c *Conn) Write(bytes []byte) error {
	if c.CoonType == CoonTypeTCP {
		return encoder.EncodeToWriter(c.TCP, bytes)
	} else if c.CoonType == ConnTypeWS {
		return c.WriteToWS(bytes)
	}
	//logger.Logger.Error("unknown conn type", zap.Any("conn", c))
	return nil
}

// WriteToWS 消息写入WebSocket
func (c *Conn) WriteToWS(bytes []byte) error {
	c.WSMutex.Lock()
	defer c.WSMutex.Unlock()

	err := c.WS.SetWriteDeadline(time.Now().Add(10 * time.Millisecond))
	if err != nil {
		return err
	}
	return c.WS.WriteMessage(websocket.BinaryMessage, bytes)
}

// Close 关闭
func (c *Conn) Close() error {
	// 取消设备和连接的对应关系
	if c.DeviceId != 0 {
		DeleteConn(c.DeviceId)
	}

	// todo 取消订阅，需要异步出去，防止重复加锁造成死锁
	//go func() {
	//	SubscribedRoom(c, 0)
	//}()

	if c.DeviceId != 0 {
		//todo 调用客户端离线rpc方法
		//_, _ = rpc.LogicIntClient.Offline(context.TODO(), &pb.OfflineReq{
		//	UserId:     c.UserId,
		//	DeviceId:   c.DeviceId,
		//	ClientAddr: c.GetAddr(),
		//})
	}

	if c.CoonType == CoonTypeTCP {
		return c.TCP.Close()
	} else if c.CoonType == ConnTypeWS {
		return c.WS.Close()
	}
	return nil
}

func (c *Conn) GetAddr() string {
	if c.CoonType == CoonTypeTCP {
		return c.TCP.GetAddr()
	} else if c.CoonType == ConnTypeWS {
		return c.WS.RemoteAddr().String()
	}
	return ""
}

// HandleMessage todo 消息处理
func (c *Conn) HandleMessage(bytes []byte) {
	c.Write([]byte("woaini"))
	//var input = new(pb.Input)
	//err := proto.Unmarshal(bytes, input)
	//if err != nil {
	//	logger.Logger.Error("unmarshal error", zap.Error(err))
	//	return
	//}
	//logger.Logger.Debug("HandleMessage", zap.Any("input", input))
	//
	//// 对未登录的用户进行拦截
	//if input.Type != pb.PackageType_PT_SIGN_IN && c.UserId == 0 {
	//	// 应该告诉用户没有登录
	//	return
	//}
	//
	//switch input.Type {
	//case pb.PackageType_PT_SIGN_IN:
	//	c.SignIn(input)
	//case pb.PackageType_PT_SYNC:
	//	c.Sync(input)
	//case pb.PackageType_PT_HEARTBEAT:
	//	c.Heartbeat(input)
	//case pb.PackageType_PT_MESSAGE:
	//	c.MessageACK(input)
	//case pb.PackageType_PT_SUBSCRIBE_ROOM:
	//	c.SubscribedRoom(input)
	//default:
	//	logger.Logger.Error("handler switch other")
	//}
}

// Send todo 下发消息
//func (c *Conn) Send(pt pb.PackageType, requestId int64, message proto.Message, err error) {
//	var output = pb.Output{
//		Type:      pt,
//		RequestId: requestId,
//	}
//
//	if err != nil {
//		status, _ := status.FromError(err)
//		output.Code = int32(status.Code())
//		output.Message = status.Message()
//	}
//
//	if message != nil {
//		msgBytes, err := proto.Marshal(message)
//		if err != nil {
//			logger.Sugar.Error(err)
//			return
//		}
//		output.Data = msgBytes
//	}
//
//	outputBytes, err := proto.Marshal(&output)
//	if err != nil {
//		logger.Sugar.Error(err)
//		return
//	}
//
//	err = c.Write(outputBytes)
//	if err != nil {
//		logger.Sugar.Error(err)
//		c.Close()
//		return
//	}
//}

// Heartbeat todo 心跳
func (c *Conn) Heartbeat() {
	//func (c *Conn) Heartbeat(input *pb.Input) {
	//c.Send(pb.PackageType_PT_HEARTBEAT, input.RequestId, nil, nil)
	//
	//logger.Sugar.Infow("heartbeat", "device_id", c.DeviceId, "user_id", c.UserId)
}

// MessageACK todo 消息收到回执(可能会合并到通知)
//func (c *Conn) MessageACK(input *pb.Input) {
//	var messageACK pb.MessageACK
//	err := proto.Unmarshal(input.Data, &messageACK)
//	if err != nil {
//		logger.Sugar.Error(err)
//		return
//	}
//
//	_, _ = rpc.LogicIntClient.MessageACK(grpclib.ContextWithRequestId(context.TODO(), input.RequestId), &pb.MessageACKReq{
//		UserId:      c.UserId,
//		DeviceId:    c.DeviceId,
//		DeviceAck:   messageACK.DeviceAck,
//		ReceiveTime: messageACK.ReceiveTime,
//	})
//}
