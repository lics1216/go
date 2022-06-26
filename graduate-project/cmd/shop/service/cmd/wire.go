// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"eshop/cmd/shop/service/internal/biz"
	"eshop/cmd/shop/service/internal/conf"
	"eshop/cmd/shop/service/internal/data"
	"eshop/cmd/shop/service/internal/server"
	"eshop/cmd/shop/service/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, *conf.Auth, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
