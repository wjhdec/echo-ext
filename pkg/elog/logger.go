package elog

import "io"

// Logger 日志
type Logger interface {
	Output() io.Writer
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Panic(args ...any)
	Fatal(args ...any)

	Debugf(template string, args ...any)
	Infof(template string, args ...any)
	Warnf(template string, args ...any)
	Errorf(template string, args ...any)
	Panicf(template string, args ...any)
	Fatalf(template string, args ...any)

	Debugw(msg string, kv ...any)
	Infow(msg string, kv ...any)
	Warnw(msg string, kv ...any)
	Errorw(msg string, kv ...any)
	Panicw(msg string, kv ...any)
	Fatalw(msg string, kv ...any)

	// Origin 原始logger，可能某些情况会用到
	Origin() any
}
