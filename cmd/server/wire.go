//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-cinch/layout/internal/biz"
	"github.com/go-cinch/layout/internal/conf"
	"github.com/go-cinch/layout/internal/data"
	"github.com/go-cinch/layout/internal/pkg/idempotent"
	"github.com/go-cinch/layout/internal/pkg/task"
	"github.com/go-cinch/layout/internal/server"
	"github.com/go-cinch/layout/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(c *conf.Bootstrap) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, task.ProviderSet, idempotent.ProviderSet, service.ProviderSet, newApp))
}
