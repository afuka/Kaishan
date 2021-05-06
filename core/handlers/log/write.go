package log

import (
	"go.uber.org/zap/zapcore"
)

func Info(msg string, fields ...zapcore.Field) {
	logger.Info(msg, fields...)
}

func Debug(msg string, fields ...zapcore.Field) {
	logger.Debug(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field)  {
	logger.Error(msg, fields...)
}