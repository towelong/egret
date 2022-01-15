package egret

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type AppInfo interface {
	Name() string
	Version() string
}

type App struct {
	opts   options
	ctx    context.Context
	cancel func()
	server *http.Server
}

func (a *App) Name() string {
	return a.opts.name
}

func (a *App) Version() string {
	return a.opts.version
}

func New(opts ...Option) *App {
	o := options{
		ctx:     context.Background(),
		timeout: 10 * time.Second,
		sigs:    []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
	}
	for _, opt := range opts {
		opt(&o)
	}
	ctx, cancel := context.WithCancel(o.ctx)
	addr := fmt.Sprintf(":%s", o.addr)
	srv := &http.Server{
		Addr:    addr,
		Handler: o.srv,
	}
	return &App{
		ctx:    ctx,
		cancel: cancel,
		opts:   o,
		server: srv,
	}
}

func (a *App) Run() error {
	eg, ctx := errgroup.WithContext(a.ctx)
	eg.Go(func() error {
		return a.server.ListenAndServe()
	})
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, a.opts.sigs...)
	eg.Go(func() error {
		select {
		case <-ctx.Done():
			fmt.Println("Timeout Done")
			return ctx.Err()
		case <-quit:
			fmt.Println("Server Stop")
			err := a.Stop()
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err := eg.Wait(); err != nil && errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (a *App) Stop() error {
	ctx, cancelFunc := context.WithTimeout(a.ctx, a.opts.timeout)
	defer cancelFunc()
	return a.server.Shutdown(ctx)
}
