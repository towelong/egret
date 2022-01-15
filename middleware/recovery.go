package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/towelong/egret/errors"
	"go.uber.org/zap"
	"net"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
)

func Recovery(log *zap.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}
				httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
				if brokenPipe {
					msg := fmt.Sprintf("error -> [%s] request -> [%s] stack -> [%s]", err, ctx.Request.URL, string(debug.Stack()))
					log.Error(msg)
					// If the connection is dead, we can't write a status to it.
					_ = ctx.Error(err.(error))
					ctx.Abort()
					return
				}
				msg := fmt.Sprintf("[panic] error -> [%s] request -> [%s] stack -> [%s]", err, string(httpRequest), string(debug.Stack()))
				log.Error(msg)
				_ = ctx.Error(errors.Unknown)
			}
		}()
		ctx.Next()
	}
}
