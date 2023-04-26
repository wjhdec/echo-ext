package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/wjhdec/echo-ext/pkg/elog"
	"io"
)

var _ echo.Logger = (*echoLogger)(nil)

func newEchoLogger(log elog.Logger) *echoLogger {
	return &echoLogger{log}
}

type echoLogger struct {
	elog.Logger
}

func (e echoLogger) Output() io.Writer {
	return e.Logger.Output()
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
	e.Logger.Debug(i...)
}

func (e echoLogger) Printf(format string, args ...interface{}) {
	e.Logger.Debugf(format, args...)
}

func (e echoLogger) Printj(j log.JSON) {
	e.Logger.Debug(j)
}

func (e echoLogger) Debugj(j log.JSON) {
	e.Logger.Debug(j)
}

func (e echoLogger) Infoj(j log.JSON) {
	e.Logger.Info(j)
}

func (e echoLogger) Warnj(j log.JSON) {
	e.Logger.Warn(j)
}

func (e echoLogger) Errorj(j log.JSON) {
	e.Logger.Error(j)
}

func (e echoLogger) Fatalj(j log.JSON) {
	e.Logger.Fatal(j)
}

func (e echoLogger) Panicj(j log.JSON) {
	e.Logger.Panic(j)
}
