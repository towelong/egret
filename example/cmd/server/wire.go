// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/towelong/egret"
	"github.com/towelong/egret/example/internal/pkg/config"
	"github.com/towelong/egret/example/internal/pkg/logger"
	"github.com/towelong/egret/example/internal/repo"
	"github.com/towelong/egret/example/internal/server"
	"github.com/towelong/egret/example/internal/service"
	"github.com/towelong/egret/example/internal/usecase"
)

func initApp(c *config.Config) *egret.App {
	panic(
		wire.Build(
			repo.ProvideSet,
			usecase.ProvideSet,
			service.ProvideSet,
			server.ProvideSet,
			logger.ProvideSet,
			newApp,
		),
	)
}
