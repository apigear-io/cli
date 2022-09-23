package git

import zlog "github.com/rs/zerolog/log"

// import logger "github.com/apigear-io/cli/pkg/log"

// var log = logger.TopicLogger("git")

var log = zlog.With().Str("topic", "git").Logger()
