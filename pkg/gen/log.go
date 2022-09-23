package gen

import zlog "github.com/rs/zerolog/log"

var log = zlog.With().Str("topic", "gen").Logger()
