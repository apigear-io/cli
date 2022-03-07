package logger

import "go.uber.org/zap"

var logger, _ = zap.NewDevelopment()

func Get() *zap.SugaredLogger {
	return logger.Sugar()
}
