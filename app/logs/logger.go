package logs

import (
	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

func Info(message string, fields ...zap.Field) {
	Logger.Info(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	Logger.Error(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	Logger.Debug(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	Logger.Warn(message, fields...)
}