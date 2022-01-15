package egret

import (
	"context"
	"github.com/gin-gonic/gin"
	"os"
	"time"
)

type Option func(o *options)

type options struct {
	name    string
	version string
	timeout time.Duration
	ctx     context.Context
	addr    string
	sigs    []os.Signal
	srv     *gin.Engine
}

func Name(name string) Option {
	return func(o *options) {
		o.name = name
	}
}

func Version(version string) Option {
	return func(o *options) {
		o.version = version
	}
}

func Timeout(timeout time.Duration) Option {
	return func(o *options) {
		o.timeout = timeout
	}
}

func Context(ctx context.Context) Option {
	return func(o *options) {
		o.ctx = ctx
	}
}

func Addr(addr string) Option {
	return func(o *options) {
		o.addr = addr
	}
}

func Signal(sigs ...os.Signal) Option {
	return func(o *options) { o.sigs = sigs }
}

func Server(srv *gin.Engine) Option {
	return func(o *options) {
		o.srv = srv
	}
}
