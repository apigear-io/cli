package log

import (
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func SetVerbose(verbose bool) {
	if verbose {
		logger.SetReportCaller(true)
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}
}

var Debug = logger.Debug
var Debugf = logger.Debugf
var Info = logger.Info
var Infof = logger.Infof
var Warn = logger.Warn
var Warnf = logger.Warnf
var Error = logger.Error
var Errorf = logger.Errorf
var Panic = logger.Panic
var Panicf = logger.Panicf
var Fatal = logger.Fatal
var Fatalf = logger.Fatalf
