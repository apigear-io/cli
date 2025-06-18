package sim

import (
	"github.com/dop251/goja_nodejs/console"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

type LogPrinter struct {
	logger *zerolog.Logger
}

func NewLogPrinter(logger *zerolog.Logger) *LogPrinter {
	if logger == nil {
		logger = &zlog.Logger
	}
	return &LogPrinter{logger: logger}
}

var _ console.Printer = (*LogPrinter)(nil)

func (lp *LogPrinter) Log(s string) {
	lp.logger.Info().Msg(s)
}

func (lp *LogPrinter) Warn(s string) {
	lp.logger.Warn().Msg(s)
}

func (lp *LogPrinter) Error(s string) {
	lp.logger.Error().Msg(s)
}
