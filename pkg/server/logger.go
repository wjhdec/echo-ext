package server

import (
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
)

var _ echo.Logger = (*echoLogger)(nil)

func newEchoLogger(log *zap.Logger, out io.Writer) *echoLogger {
	return &echoLogger{log.Sugar(), out}
}

type echoLogger struct {
	*zap.SugaredLogger
	out io.Writer
}

func (e echoLogger) Output() io.Writer {
	return e.out
}

func (e echoLogger) SetOutput(io.Writer) {
}

func (e echoLogger) Prefix() string {
	return ""
}

func (e echoLogger) SetPrefix(string) {
}

func (e echoLogger) Level() log.Lvl {
	return log.INFO
}

func (e echoLogger) SetLevel(log.Lvl) {
}

func (e echoLogger) SetHeader(string) {
}

func (e echoLogger) Print(i ...interface{}) {
	e.Debug(i...)
}

func (e echoLogger) Printf(format string, args ...interface{}) {
	e.Debugf(format, args...)
}

func (e echoLogger) Printj(j log.JSON) {
	e.Desugar().Debug("", zap.Any("", j))
}

func (e echoLogger) Debugj(j log.JSON) {
	e.Desugar().Debug("", zap.Any("", j))
}

func (e echoLogger) Infoj(j log.JSON) {
	e.Desugar().Info("", zap.Any("", j))
}

func (e echoLogger) Warnj(j log.JSON) {
	e.Desugar().Warn("", zap.Any("", j))
}

func (e echoLogger) Errorj(j log.JSON) {
	e.Desugar().Error("", zap.Any("", j))
}

func (e echoLogger) Fatalj(j log.JSON) {
	e.Desugar().Fatal("", zap.Any("", j))
}

func (e echoLogger) Panicj(j log.JSON) {
	e.Desugar().Panic("", zap.Any("", j))
}
