package log

import "log"

// TODO: 需要封装zap的logger
func Info(format string, v ...interface{})  {
	log.Printf(format, v)
}

func Error(format string, v ...interface{})  {
	log.Printf(format, v)
}

func Fatal(format string, v ...interface{})  {
	log.Fatalf(format, v)
}