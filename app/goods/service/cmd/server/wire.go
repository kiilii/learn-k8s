//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"learn-k8s/app/goods/service/internal/biz"
	"learn-k8s/app/goods/service/internal/conf"
	"learn-k8s/app/goods/service/internal/data"
	"learn-k8s/app/goods/service/internal/server"
	"learn-k8s/app/goods/service/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Naming, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(
		wire.Build(
			server.ProviderSet,
			data.ProviderSet,
			biz.ProviderSet,
			service.ProviderSet,
			newApp,
		),
	)

}
