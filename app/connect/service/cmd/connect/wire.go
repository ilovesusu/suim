//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/ilovesusu/suim/app/connect/service/internal/biz"
	"github.com/ilovesusu/suim/app/connect/service/internal/conf"
	"github.com/ilovesusu/suim/app/connect/service/internal/data"
	"github.com/ilovesusu/suim/app/connect/service/internal/server"
	"github.com/ilovesusu/suim/app/connect/service/internal/service"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(data.ProviderSet, biz.ProviderSet, service.ProviderSet, server.ProviderSet, newApp))
}
