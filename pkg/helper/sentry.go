package helper

import (
	"time"

	"github.com/getsentry/sentry-go"
)

var CLI_DSN = "https://a93da0b87a654895a96ab3a1e4101792@o170438.ingest.sentry.io/4503898727907328"

func SentryInit(dsn string) error {
	return sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		TracesSampleRate: 1.0,
	})
}

func SentryFlush() {
	sentry.Flush(2 * time.Second)
}

func SentryCaptureError(err error) {
	sentry.CaptureException(err)
}

func SentryCaptureMessage(message string) {
	sentry.CaptureMessage(message)
}
