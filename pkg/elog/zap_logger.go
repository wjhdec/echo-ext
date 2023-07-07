package elog

import (
	"fmt"
	"io"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/wjhdec/echo-ext/pkg/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(cfg config.Config) (*zap.Logger, io.Writer) {
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
	return zap.New(zCore, zap.AddStacktrace(zapcore.WarnLevel)), writer
}

func NewConsoleLogger() *zap.Logger {
	logger, err := zap.NewDevelopmentConfig().Build()
	if err != nil {
		fmt.Printf("build log error: %+v \n", err)
	}
	return logger
}
