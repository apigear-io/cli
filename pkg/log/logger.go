package log

import "go.uber.org/zap"

var logger, _ = zap.NewDevelopment()
var sugar = logger.Sugar()

var Debug = sugar.Debug
var Debugf = sugar.Debugf
var Info = sugar.Info
var Warn = sugar.Warn
var Warnf = sugar.Warnf
var Error = sugar.Error
var Errorf = sugar.Errorf
var DPanic = sugar.DPanic
var Panic = sugar.Panic
var Fatal = sugar.Fatal
var Fatalf = sugar.Fatalf
var Infof = sugar.Infof
