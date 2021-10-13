package server

import (
	etcd "github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/ilovesusu/suim/app/connect/service/internal/conf"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func NewEtcdServer(c *conf.Server, logger log.Logger) registry.Registrar {
	//service register
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{c.Etcd.Addr},
	})

	if err != nil {
		panic(err)
	}
	reg := etcd.New(cli)
	//etcd.Namespace("suim"),
	//etcd.RegisterTTL(12),
	return reg
}
