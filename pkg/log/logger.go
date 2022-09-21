package log

import (
	"io"
	"log"
	"os"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func Config(verbose bool, debug bool) {
	log.SetOutput(NewLogCapture(logger))
	logger.Formatter = &logrus.TextFormatter{
		DisableQuote: true,
	}
	if verbose {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.InfoLevel)
	}
	logger.SetReportCaller(debug)
	if verbose || debug {
		logger.Debugf("logger configured: verbose=%v, debug=%v", verbose, debug)
	}
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	logFile := helper.Join(home, ".apigear/logs/app.log")
	ljack := newLogFileRotator(logFile)

	if verbose { // log to console and file
		logger.SetOutput(io.MultiWriter(os.Stderr, ljack))
	} else { // log to reporter and file
		logger.SetOutput(ljack)
		logger.AddHook(NewReportHook())
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
