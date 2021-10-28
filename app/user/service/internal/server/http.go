package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/ilovesusu/suim/api/user/service/v1/user"
	"github.com/ilovesusu/suim/app/user/service/internal/conf"
	"github.com/ilovesusu/suim/app/user/service/internal/service"
)

// NewHTTPServer 创建http服务
func NewHTTPServer(c *conf.Server, us *service.UserService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	user.RegisterUserHTTPServer(srv, us)
	_ = logger.Log(log.LevelInfo, "HTTP", "http init success!")
	return srv
}
