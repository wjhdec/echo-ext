package elog

import (
	"io"

	"go.uber.org/zap"
)

var (
	std    = NewConsoleLogger()
	writer io.Writer
)

// OverrideGlobalLogger 覆盖当前全局logger
func OverrideGlobalLogger(logger *zap.Logger, out io.Writer) {
	std = logger
	writer = out
}

func Out() io.Writer {
	return writer
}

// GlobalLogger 全局logger
func GlobalLogger() *zap.Logger {
	return std
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...any) {
	std.Sugar().Debug(args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...any) {
	std.Sugar().Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...any) {
	std.Sugar().Warn(args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...any) {
	if len(args) == 1 {
		if err, ok := args[0].(error); ok {
			std.Error("", zap.Error(err))
		} else {
			std.Sugar().Error(args...)
		}
	} else {
		std.Sugar().Error(args...)
	}
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...any) {
	std.Sugar().Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...any) {
	std.Sugar().Fatal(args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...any) {
	std.Sugar().Debugf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...any) {
	std.Sugar().Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...any) {
	std.Sugar().Warnf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...any) {
	std.Sugar().Errorf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(format string, args ...any) {
	std.Sugar().Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalf(format string, args ...any) {
	std.Sugar().Fatalf(format, args...)
}
