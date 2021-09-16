package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
)

type Level = zapcore.Level

const (
	InfoLevel   Level = zap.InfoLevel   // 0, default level
	WarnLevel   Level = zap.WarnLevel   // 1
	ErrorLevel  Level = zap.ErrorLevel  // 2
	DPanicLevel Level = zap.DPanicLevel // 3, used in development log
	// PanicLevel logs a message, then panics
	PanicLevel Level = zap.PanicLevel // 4
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel Level = zap.FatalLevel // 5
	DebugLevel Level = zap.DebugLevel // -1
)

type Logger struct {
	logger *zap.SugaredLogger
	level  Level
}

func New(writer io.Writer, level Level) *Logger {
	if writer == nil {
		panic("writer is nil")
	}
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		zapcore.AddSync(writer),
		level,
	)
	logger := &Logger{
		logger: zap.New(core).Sugar(),
		level:  level,
	}
	return logger
}

func (l *Logger) Debug(msg string, fields ...interface{}) {
	l.logger.Debugf(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...interface{}) {
	l.logger.Infof(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...interface{}) {
	l.logger.Warnf(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...interface{}) {
	l.logger.Errorf(msg, fields...)
}

func (l *Logger) DPanic(msg string, fields ...interface{}) {
	l.logger.DPanicf(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...interface{}) {
	l.logger.Panicf(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...interface{}) {
	l.logger.Fatalf(msg, fields...)
}

func (l *Logger) Sync() error {
	return l.logger.Sync()
}
