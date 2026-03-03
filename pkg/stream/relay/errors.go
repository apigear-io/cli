package relay

import (
	"github.com/apigear-io/cli/pkg/stream/relay/internal/client"
	"github.com/apigear-io/cli/pkg/stream/relay/internal/core"
)

// Connection errors (from core)
var (
	ErrConnectionNotFound  = core.ErrConnectionNotFound
	ErrConnectionClosed    = core.ErrConnectionClosed
	ErrPoolClosed          = core.ErrPoolClosed
	ErrDuplicateConnection = core.ErrDuplicateConnection
)

// Client errors (from client)
var (
	ErrClientNotFound      = client.ErrClientNotFound
	ErrClientAlreadyExists = client.ErrClientAlreadyExists
	ErrNotConnected        = client.ErrNotConnected
	ErrAlreadyStarted      = client.ErrAlreadyStarted
)
