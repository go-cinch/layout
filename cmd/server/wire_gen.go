// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-cinch/layout/internal/biz"
	"github.com/go-cinch/layout/internal/conf"
	"github.com/go-cinch/layout/internal/data"
	"github.com/go-cinch/layout/internal/server"
	"github.com/go-cinch/layout/internal/service"
	"github.com/go-kratos/kratos/v2"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(c *conf.Bootstrap) (*kratos.App, func(), error) {
	universalClient, err := data.NewRedis(c)
	if err != nil {
		return nil, nil, err
	}
	db, err := data.NewDB(c)
	if err != nil {
		return nil, nil, err
	}
	tracerProvider, err := data.NewTracer(c)
	if err != nil {
		return nil, nil, err
	}
	dataData, cleanup := data.NewData(universalClient, db, tracerProvider)
	greeterRepo := data.NewGreeterRepo(dataData)
	transaction := data.NewTransaction(dataData)
	greeterUseCase := biz.NewGreeterUseCase(greeterRepo, transaction)
	hellowordService := service.NewHellowordService(greeterUseCase)
	grpcServer := server.NewGRPCServer(c, hellowordService)
	httpServer := server.NewHTTPServer(c, hellowordService)
	app := newApp(grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
