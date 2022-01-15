package server

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/towelong/egret/example/api/v1"
	"github.com/towelong/egret/example/internal/service"
	"github.com/towelong/egret/middleware"
	_ "github.com/towelong/egret/pkg/validate"
	"go.uber.org/zap"
)

func NewHttpServer(s *service.ShopInterface, logger *zap.Logger) *gin.Engine {
	srv := gin.New()
	srv.Use(middleware.Error)
	srv.Use(middleware.CORS)
	srv.Use(middleware.New(logger).Log)
	srv.Use(middleware.Recovery(logger))
	v1.RegisterShopServiceHTTPServer(srv, s)
	return srv
}
