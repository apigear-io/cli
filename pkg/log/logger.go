package log

import "go.uber.org/zap"

var logger, _ = zap.NewDevelopment()

var Info = logger.Info
var Debug = logger.Debug
var Warn = logger.Warn
var Error = logger.Error
var Fatal = logger.Fatal
var Panic = logger.Panic
var With = logger.With
var DPanic = logger.DPanic

func GetLogger() *zap.Logger {
	return logger
}
