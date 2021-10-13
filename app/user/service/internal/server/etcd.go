package server

import (
	etcd "github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/ilovesusu/suim/app/user/service/internal/conf"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func NewETCDServer(c *conf.Server, logger log.Logger) registry.Registrar {
	//service register
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{c.Etcd.Addr},
		DialTimeout: time.Duration(c.Etcd.Timeout) * time.Second,
	})

	if err != nil {
		_ = logger.Log(log.LevelError, "ETCD", err)
		panic(err)
	}
	reg := etcd.New(cli)
	_ = logger.Log(log.LevelInfo, "ETCD", "etcd init success!")
	return reg
}
