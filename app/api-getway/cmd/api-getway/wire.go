//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"api-getway/internal/biz"
	"api-getway/internal/client"
	"api-getway/internal/conf"
	"api-getway/internal/server"
	"api-getway/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Naming, log.Logger) (*kratos.App, func(), error) {
	panic(
		wire.Build(
			server.ProviderSet,
			biz.ProviderSet,
			service.ProviderSet,
			client.ProviderSet,
			newApp,
		),
	)
}
