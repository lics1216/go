// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"eshop/cmd/order/internal/biz"
	"eshop/cmd/order/internal/conf"
	"eshop/cmd/order/internal/data"
	"eshop/cmd/order/internal/server"
	"eshop/cmd/order/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
