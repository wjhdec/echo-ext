package logext

import (
	"fmt"
	"log/slog"
)

// LogError log error with stack trace
func LogError(err error) {
	LogErrorMsg(err.Error(), err)
}

// LogErrorMsg log error with message
func LogErrorMsg(msg string, err error) {
	slog.Error(msg, slog.String("stacktrace", fmt.Sprintf("%+v", err)))
}
