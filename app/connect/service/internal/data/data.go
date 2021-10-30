package data

import (
	"context"
	etcd "github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	v1 "github.com/ilovesusu/suim/api/logic/service/v1"
	"github.com/ilovesusu/suim/app/connect/service/internal/conf"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewLogicClient, NewDiscovery, NewRegistrar)

// Data .
type Data struct {
	lc v1.LogicClient
}

// NewData 数据库连接 .
func NewData(c *conf.Data, lc v1.LogicClient, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{lc: lc}, cleanup, nil
}

func NewDiscovery(conf *conf.Server) registry.Discovery {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{conf.Etcd.Addr},
	})

	if err != nil {
		panic(err)
	}
	reg := etcd.New(cli)
	return reg
}

func NewRegistrar(conf *conf.Server) registry.Registrar {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{conf.Etcd.Addr},
	})

	if err != nil {
		panic(err)
	}
	reg := etcd.New(cli)
	return reg
}

func NewLogicClient(r registry.Discovery) v1.LogicClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///Logic/cwy-system"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	c := v1.NewLogicClient(conn)
	return c
}
