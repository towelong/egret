package main

import (
	"github.com/gin-gonic/gin"
	"github.com/towelong/egret"
)

func newApp(srv *gin.Engine) *egret.App {
	return egret.New(
		egret.Name("shop"),
		egret.Version("1.0.0"),
		egret.Addr("8081"),
		egret.Server(srv),
	)
}

func main() {
	app := initApp()
	if err := app.Run(); err != nil {
		panic(err)
	}
}
