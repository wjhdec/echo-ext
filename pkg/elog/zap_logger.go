package elog

import (
	"io"
	"os"
	"sync"

	"github.com/wjhdec/echo-ext/pkg/config"
	"gopkg.in/natefinch/lumberjack.v2"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	blog     Logger
	zlogOnce sync.Once
)

type LoggerConfig struct {
	lumberjack.Logger
	Level string
}

var _ Logger = &baseLogger{}

type baseLogger struct {
	zap.SugaredLogger
	writer io.Writer
}

func (l *baseLogger) Output() io.Writer {
	return l.writer
}

func NewLogger(cfg config.Config) Logger {
	jack := new(lumberjack.Logger)
	if cfg != nil {
		if err := cfg.UnmarshalByKey("logger", jack); err != nil {
			zap.S().Error(err)
		}
	} else {
		zap.S().Warn("config not found, use default logger")
	}
	var writer zapcore.WriteSyncer
	var encoder zapcore.Encoder
	if jack != nil && jack.Filename != "" {
		writer = zapcore.AddSync(jack)
		encoder = zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	} else {
		writer = zapcore.AddSync(os.Stdout)
		encoder = zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	}
	level := zapcore.DebugLevel
	var err error
	if cfg != nil {
		level, err = zapcore.ParseLevel(cfg.StrValueByKey("logger.level"))
		if err != nil {
			zap.S().Errorf("log read config error, use warn default: %+v", err.Error())
		}
	}
	zCore := zapcore.NewCore(encoder, writer, level)
	l := zap.New(zCore, zap.AddStacktrace(zapcore.WarnLevel))
	zap.ReplaceGlobals(l)
	return &baseLogger{
		SugaredLogger: *l.Sugar(),
		writer:        writer,
	}
}

// Default 默认日志，如果读取内容出错，则使用默认的zap设置
func Default() Logger {
	zlogOnce.Do(func() {
		cfg, err := config.New()
		if err != nil {
			zap.S().Error(err)
		}
		blog = NewLogger(cfg)
	})
	return blog

}
