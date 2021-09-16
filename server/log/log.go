package log

import "os"

var logger = initLogger()
var (
	Info   = logger.Info
	Warn   = logger.Warn
	Error  = logger.Error
	DPanic = logger.DPanic
	Panic  = logger.Panic
	Fatal  = logger.Fatal
	Debug  = logger.Debug
)

func initLogger() *Logger {
	return New(os.Stderr, InfoLevel)
}

func Get() *Logger {
	return logger
}

func Sync() error {
	if logger != nil {
		return logger.Sync()
	}
	return nil
}

func ResetDefault(logger *Logger) {
	logger = logger
	Info = logger.Info
	Warn = logger.Warn
	Error = logger.Error
	DPanic = logger.DPanic
	Panic = logger.Panic
	Fatal = logger.Fatal
	Debug = logger.Debug
}
