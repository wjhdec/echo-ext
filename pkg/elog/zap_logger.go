package elog

import (
	"github.com/wjhdec/echo-ext/pkg/config"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	blog     *baseLogger
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

// Default 默认日志，如果读取内容出错，则使用默认的zap设置
func Default() Logger {
	zlogOnce.Do(func() {
		jack := new(lumberjack.Logger)
		cfg, err := config.New()
		if err != nil {
			zap.S().Error(err)
			jack = nil
		} else if err := cfg.UnmarshalByKey("logger", jack); err != nil {
			zap.S().Error(err)
			jack = nil
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
		if cfg != nil {
			level, err = zapcore.ParseLevel(cfg.StrValueByKey("logger.level"))
			if err != nil {
				zap.S().Errorf("log read config error, use warn default: %+v", err.Error())
			}
		}
		zCore := zapcore.NewCore(encoder, writer, level)
		l := zap.New(zCore, zap.AddStacktrace(zapcore.WarnLevel))
		zap.ReplaceGlobals(l)
		blog = &baseLogger{
			SugaredLogger: *l.Sugar(),
			writer:        writer,
		}
	})
	return blog

}
