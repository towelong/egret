package logger

import (
	"github.com/towelong/egret/pkg/log"
	"go.uber.org/zap"
)

// NewLogger 基于外部pkg中logger的配置
func NewLogger() *zap.Logger {
	return log.NewLogger()
}
