package log

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func Config(verbose bool, debug bool) {
	log.SetOutput(NewLogCapture(logger))
	if verbose {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}
	logger.SetReportCaller(debug)
	if verbose || debug {
		logger.Debugf("logger configured: verbose=%v, debug=%v", verbose, debug)
	}
	logger.AddHook(NewReportHook())
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	logFile := filepath.Join(home, ".apigear/logs/app.log")
	ljack := newLogFileRotator(logFile)
	if debug {
		logger.SetOutput(io.MultiWriter(os.Stderr, ljack))
	} else {
		logger.SetOutput(ljack)
	}
}

var Debug = logger.Debug
var Debugf = logger.Debugf
var Debugln = logger.Debugln
var Info = logger.Info
var Infof = logger.Infof
var Infoln = logger.Infoln
var Warn = logger.Warn
var Warnf = logger.Warnf
var Warnln = logger.Warnln
var Error = logger.Error
var Errorf = logger.Errorf
var Errorln = logger.Errorln
var Panic = logger.Panic
var Panicf = logger.Panicf
var Panicln = logger.Panicln
var Fatal = logger.Fatal
var Fatalf = logger.Fatalf
var Fatalln = logger.Fatalln

func TopicLogger(topic string) *logrus.Entry {
	return logger.WithField("topic", topic)
}

func Logger() *logrus.Logger {
	return logger
}
