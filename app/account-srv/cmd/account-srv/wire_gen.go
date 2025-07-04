// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"shopx/app/account-srv/internal/biz"
	"shopx/app/account-srv/internal/conf"
	"shopx/app/account-srv/internal/data"
	"shopx/app/account-srv/internal/server"
	"shopx/app/account-srv/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	userUsecase := biz.NewUserUsecase(userRepo, logger)
	frontLoginService := service.NewFrontLoginService(userUsecase)
	grpcServer := server.NewGRPCServer(confServer, frontLoginService, logger)
	httpServer := server.NewHTTPServer(confServer, frontLoginService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
