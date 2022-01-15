package logger

import (
	"github.com/google/wire"
	"github.com/towelong/egret/pkg/log"
	"go.uber.org/zap"
)

var ProvideSet = wire.NewSet(New)

// New 基于外部pkg中logger的配置
func New() *zap.Logger {
	return log.NewLogger()
}
