package js

import (
	"github.com/apigear-io/cli/pkg/log"
	"github.com/dop251/goja_nodejs/console"
)

type LogPrinter struct {
}

var _ console.Printer = (*LogPrinter)(nil)

func (lp *LogPrinter) Log(s string) {
	log.Info().Msg(s)
}

func (lp *LogPrinter) Warn(s string) {
	log.Warn().Msg(s)
}

func (lp *LogPrinter) Error(s string) {
	log.Error().Msg(s)
}
