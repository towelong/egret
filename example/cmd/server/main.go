package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/towelong/egret"
	"github.com/towelong/egret/example/internal/pkg/config"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "conf", "../../configs/config.yaml", "config path, eg: -conf config.yaml")
}

func newApp(c *config.Config, srv *gin.Engine) *egret.App {
	return egret.New(
		egret.Name(c.APP.Name),
		egret.Version(c.APP.Version),
		egret.Addr(c.APP.Addr),
		egret.Server(srv),
	)
}

func main() {
	flag.Parse()
	c, err := config.New(configPath)
	if err != nil {
		panic(err)
	}
	app := initApp(c)
	if err := app.Run(); err != nil {
		panic(err)
	}
}
