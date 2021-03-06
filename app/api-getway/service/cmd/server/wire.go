//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"learn-k8s/app/api-getway/service/internal/biz"
	"learn-k8s/app/api-getway/service/internal/conf"
	"learn-k8s/app/api-getway/service/internal/data"
	"learn-k8s/app/api-getway/service/internal/server"
	"learn-k8s/app/api-getway/service/internal/service"
	"go.opentelemetry.io/otel/sdk/trace"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Naming, log.Logger, *trace.TracerProvider) (*kratos.App, func(), error) {
	panic(
		wire.Build(
			server.ProviderSet,
			biz.ProviderSet,
			service.ProviderSet,
			data.ProviderSet,
			newApp,
		),
	)
}
