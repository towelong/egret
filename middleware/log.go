package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/towelong/egret/pkg/utils"
	"go.uber.org/zap"
	"io/ioutil"
	"math"
	"strings"
	"time"
)

type Log struct {
	log *zap.Logger
}

func New(log *zap.Logger) *Log {
	return &Log{log: log}
}

func (l *Log) Log(ctx *gin.Context) {
	start := time.Now()
	data, err := ctx.GetRawData()
	if err != nil {
		l.log.Error("[middleware Log: ]" + err.Error())
		ctx.Abort()
		return
	}
	query := ctx.Copy().Request.URL.RawQuery
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data)) // 重新赋值
	contentType := ctx.Copy().ContentType()
	ctx.Next()
	costs := time.Since(start)
	latency := int(math.Ceil(float64(costs.Nanoseconds()) / 1000000.0))
	newData := strings.TrimSpace(string(data))
	body := strings.ReplaceAll(newData, "\n", "")
	body = strings.ReplaceAll(body, " ", "")
	if body == "" || strings.HasPrefix(contentType, "multipart") {
		body = "{}"
	}
	var token string
	if token = ctx.Copy().Request.Header.Get("Authorization"); token == "" {
		token = ""
	}
	msg := fmt.Sprintf(`[%s] -> [%s] from: %s costs: %vms User-Agent: [%s] token: [%s] 
data: { 
	params: %s, 
	body: %s
}`,
		ctx.Request.Method,
		ctx.Request.RequestURI,
		ctx.ClientIP(),
		latency,
		ctx.Request.UserAgent(),
		token,
		utils.FormatQuery(query),
		body,
	)
	l.log.Info(msg)
}
