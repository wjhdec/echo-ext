package elog

import (
	"echoext/pkg/config"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
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

type baseLogger struct {
	zap.SugaredLogger
	writer io.Writer
}

func (l *baseLogger) Output() io.Writer {
	return l.writer
}

func Default() Logger {
	zlogOnce.Do(func() {
		jack := new(lumberjack.Logger)
		cfg := config.New()
		if err := cfg.Unmarshal("logger", jack); err != nil {
			panic(err)
		}
		writer := zapcore.AddSync(jack)
		encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())

		level, err := zapcore.ParseLevel(cfg.StrValue("logger.level"))
		if err != nil {
			panic(err)
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
