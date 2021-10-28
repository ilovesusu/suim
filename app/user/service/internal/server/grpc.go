package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/ilovesusu/suim/api/user/service/v1/friend"
	"github.com/ilovesusu/suim/api/user/service/v1/user"
	"github.com/ilovesusu/suim/app/user/service/internal/conf"
	"github.com/ilovesusu/suim/app/user/service/internal/service"
)

// NewGRPCServer 创建有一个grpc服务
func NewGRPCServer(c *conf.Server, us *service.UserService, fs *service.FriendService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	user.RegisterUserServer(srv, us)
	friend.RegisterFriendServer(srv, fs)
	_ = logger.Log(log.LevelInfo, "GRPC", "grpc init success!")
	return srv
}
