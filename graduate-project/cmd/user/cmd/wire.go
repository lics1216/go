// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"eshop/cmd/user/internal/biz"
	"eshop/cmd/user/internal/conf"
	"eshop/cmd/user/internal/data"
	"eshop/cmd/user/internal/data/cache"
	"eshop/cmd/user/internal/data/db"
	"eshop/cmd/user/internal/server"
	"eshop/cmd/user/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Registry, *conf.Data, log.Logger, *tracesdk.TracerProvider) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, cache.ProviderSet, db.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
