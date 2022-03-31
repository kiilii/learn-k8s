// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"goods/internal/biz"
	"goods/internal/conf"
	"goods/internal/data"
	"goods/internal/server"
	"goods/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	goodsRepo := data.NewGoodsRepo(dataData, logger)
	goodsUsecase := biz.NewGoodsUsecase(goodsRepo, logger)
	goodsService := service.NewGoodsService(goodsUsecase)
	httpServer := server.NewHTTPServer(confServer, goodsService, logger)
	grpcServer := server.NewGRPCServer(confServer, goodsService, logger)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
