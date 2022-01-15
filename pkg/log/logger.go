package log

import (
	"github.com/natefinch/lumberjack"
	"github.com/towelong/egret/pkg/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
)

type Options func(l *Logger)

type Logger struct {
	// 日期格式
	timeLayout string
	// 日志存放目录
	logPath string
	// 文件大小限制,单位MB
	maxSize int
	// 日志文件保留天数
	maxAge int
	// 最大保留日志文件数量
	maxBackups int
}

func (l *Logger) TimeLayout(layout string) Options {
	return func(l *Logger) {
		l.timeLayout = layout
	}
}

func (l *Logger) LogPath(path string) Options {
	return func(l *Logger) {
		// 容错处理，判断传进来的目录是否有 "/"
		if strings.HasSuffix(path, "/") {
			l.logPath = path
		} else {
			l.logPath = path + "/"
		}
	}
}

func (l *Logger) MaxSize(maxSize int) Options {
	return func(l *Logger) {
		l.maxSize = maxSize
	}
}

func (l *Logger) MaxAge(maxAge int) Options {
	return func(l *Logger) {
		l.maxAge = maxAge
	}
}

func (l *Logger) MaxBackups(maxBackups int) Options {
	return func(l *Logger) {
		l.maxBackups = maxBackups
	}
}

func NewLogger(opts ...Options) *zap.Logger {
	l := &Logger{
		timeLayout: utils.FullTimeHorizontalLayout,
		logPath:    "./logs/",
		maxSize:    2,
		maxAge:     30,
		maxBackups: 100,
	}
	for _, opt := range opts {
		opt(l)
	}
	var coreArr []zapcore.Core
	// 获取编码器
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(l.timeLayout) //指定时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder         // 按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder               // 显示完整文件路径
	encoder := zapcore.NewConsoleEncoder(encoderConfig)                  // NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	// 日志级别
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { // error级别
		return lev >= zap.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool { // info和debug级别,debug级别是最低的
		return lev < zap.ErrorLevel && lev >= zap.DebugLevel
	})
	t := time.Now().Format(utils.YearMonthDayTimeHorizontalLayout)
	// info文件writeSyncer
	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   l.logPath + t + "_info" + ".log", // 日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    l.maxSize,                        // 文件大小限制,单位MB
		MaxBackups: l.maxBackups,                     // 最大保留日志文件数量
		MaxAge:     l.maxAge,                         // 日志文件保留天数
		Compress:   false,                            // 是否压缩处理
	})
	infoFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)), lowPriority) //第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	// error文件writeSyncer
	errorFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   l.logPath + t + "_error" + ".log", // 日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    l.maxSize,                         // 文件大小限制,单位MB
		MaxBackups: l.maxBackups,                      // 最大保留日志文件数量
		MaxAge:     l.maxAge,                          // 日志文件保留天数
		Compress:   false,                             // 是否压缩处理
	})
	errorFileCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(errorFileWriteSyncer, zapcore.AddSync(os.Stdout)), highPriority) // 第三个及之后的参数为写入文件的日志级别,ErrorLevel模式只记录error级别的日志
	coreArr = append(coreArr, infoFileCore)
	coreArr = append(coreArr, errorFileCore)
	// zap.AddCaller()为显示文件名和行号，可省略
	log := zap.New(zapcore.NewTee(coreArr...), zap.AddStacktrace(zapcore.WarnLevel))
	return log
}
