package v1

import (
	"context"
	"github.com/gin-gonic/gin"
)

type ShopServiceHTTPServer interface {
	GetUserById(context.Context, *GetUserReq) (*GetUserResp, error)
}

type ShopService struct {
	server ShopServiceHTTPServer
	router gin.IRouter
}

func RegisterShopServiceHTTPServer(r gin.IRouter, srv ShopServiceHTTPServer) {
	s := ShopService{
		server: srv,
		router: r,
	}
	s.RegisterService()
}

func (s *ShopService) RegisterService() {
	v1 := s.router.Group("/v1")
	{
		v1.GET("/user/:user_id", s.GetUserById)
	}
}
