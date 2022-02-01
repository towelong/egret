// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/towelong/egret"
	"github.com/towelong/egret/example/internal/pkg/config"
	"github.com/towelong/egret/example/internal/repo"
	"github.com/towelong/egret/example/internal/server"
	"github.com/towelong/egret/example/internal/service"
	"github.com/towelong/egret/example/internal/usecase"
	"go.uber.org/zap"
)

func initApp(c *config.Config, l *zap.Logger) *egret.App {
	panic(
		wire.Build(
			repo.ProvideSet,
			usecase.ProvideSet,
			service.ProvideSet,
			server.ProvideSet,
			newApp,
		),
	)
}
