package db

import (
	"context"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/wjhdec/echo-ext/pkg/elog"
)

type logAdapter struct {
	logger elog.Logger
}

// NewAdaptor set zap logger as backend as an example on how it process log from sqldblogger.Log().
func NewAdaptor(logger elog.Logger) sqldblogger.Logger {
	return &logAdapter{logger: logger}
}

// Log implement sqldblogger.Logger and log it as is.
// To use context.Context values, please copy this file and adjust to your needs.
func (la *logAdapter) Log(_ context.Context, level sqldblogger.Level, msg string, data map[string]interface{}) {
	fields := make([]any, 0, len(data)*2)
	for k, v := range data {
		fields = append(fields, k, v)
	}

	switch level {
	case sqldblogger.LevelError:
		la.logger.Errorw(msg, fields...)
	case sqldblogger.LevelInfo:
		la.logger.Infow(msg, fields...)
	case sqldblogger.LevelDebug:
		la.logger.Debugw(msg, fields...)
	default:
		// trace will use zap debug
		la.logger.Debugw(msg, fields...)
	}
}
