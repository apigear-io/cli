package log

import (
	"github.com/apigear-io/hub/log"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func Config(verbose bool, debug bool) {
	if verbose {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}
	logger.SetReportCaller(debug)
	if verbose || debug {
		log.Infof("logger configured: verbose=%v, debug=%v", verbose, debug)
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
