package elog

import (
	"fmt"
	"github.com/wjhdec/echo-ext/pkg/config"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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
			fmt.Printf("read logger config error : %+v \n", err)
		}
	} else {
		fmt.Println("config not found, use default logger instead")
	}
	var writer zapcore.WriteSyncer
	var encoder zapcore.Encoder
	if jack != nil && jack.Filename != "" {
		writer = zapcore.AddSync(jack)
		peCfg := zap.NewProductionEncoderConfig()
		peCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoder = zapcore.NewJSONEncoder(peCfg)
	} else {
		fmt.Println("can not find log file config, use stdout instead")
		writer = zapcore.AddSync(os.Stdout)
		devCfg := zap.NewDevelopmentEncoderConfig()
		devCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoder = zapcore.NewConsoleEncoder(devCfg)
	}
	level := zapcore.DebugLevel
	var err error
	if cfg != nil {
		level, err = zapcore.ParseLevel(cfg.StrValueByKey("logger.level"))
		if err != nil {
			fmt.Printf("log read config error, use warn default: %+v \n", err.Error())
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
