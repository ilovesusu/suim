// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/ilovesusu/suim/app/logic/service/internal/biz"
	"github.com/ilovesusu/suim/app/logic/service/internal/conf"
	"github.com/ilovesusu/suim/app/logic/service/internal/data"
	"github.com/ilovesusu/suim/app/logic/service/internal/server"
	"github.com/ilovesusu/suim/app/logic/service/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData, logger)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, logger)
	greeterService := service.NewGreeterService(greeterUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, greeterService, logger)
	grpcServer := server.NewGRPCServer(confServer, greeterService, logger)
	registrar := server.NewEtcdServer(confServer, logger)
	app := newApp(logger, httpServer, grpcServer, registrar)
	return app, func() {
		cleanup()
	}, nil
}
