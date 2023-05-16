package elog

var (
	std = NewConsoleLogger()
)

// OverrideGlobalLogger 覆盖当前全局logger
func OverrideGlobalLogger(logger Logger) {
	std = logger
}

// GlobalLogger 全局logger
func GlobalLogger() Logger {
	return std
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	std.Debug(args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	std.Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	std.Warn(args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	std.Error(args...)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	std.Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(format string, args ...interface{}) {
	std.Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}

func Debugw(msg string, kv ...any) {
	std.Debugw(msg, kv...)
}
func Infow(msg string, kv ...any) {
	std.Infow(msg, kv...)
}
func Warnw(msg string, kv ...any) {
	std.Warnw(msg, kv...)
}
func Errorw(msg string, kv ...any) {
	std.Errorw(msg, kv...)
}
func Panicw(msg string, kv ...any) {
	std.Panicw(msg, kv...)
}
func Fatalw(msg string, kv ...any) {
	std.Fatalw(msg, kv...)
}
