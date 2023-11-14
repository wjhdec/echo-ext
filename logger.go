package echoext

import (
	"fmt"
	"log/slog"
	"os"
)

func logError(err error) {
	logErrorWithMsg(err.Error(), err)
}

func logErrorWithMsg(msg string, err error) {
	slog.Error(msg, slog.String("stacktrace", fmt.Sprintf("%+v", err)))
}

func logFatalWithMsg(msg string, err error) {
	logErrorWithMsg(msg, err)
	os.Exit(1)
}
