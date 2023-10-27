package logext

import (
	"fmt"
	"log/slog"
)

// Error log error with stack trace
func Error(err error) {
	ErrorWithMsg(err.Error(), err)
}

// ErrorWithMsg log error with message
func ErrorWithMsg(msg string, err error) {
	slog.Error(msg, slog.String("stacktrace", fmt.Sprintf("%+v", err)))
}
