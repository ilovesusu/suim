// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/ilovesusu/suim/app/connect/service/internal/conf"
	"github.com/ilovesusu/suim/app/connect/service/internal/server"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, data *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	tcpServer := server.NewTCPServer()
	webSocketServer := server.NewWebSocketServer()
	app := newApp(logger, tcpServer, webSocketServer)
	return app, func() {
	}, nil
}