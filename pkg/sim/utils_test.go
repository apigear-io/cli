package sim

import (
	"testing"
	"time"

	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
)

func setupServer(t *testing.T) (*nats.Conn, func()) {
	opts := &server.Options{
		ServerName: "apigear_server",
		DontListen: true,
	}
	server, err := server.NewServer(opts)
	assert.NoError(t, err)
	assert.NotNil(t, server)
	server.Start()
	if !server.ReadyForConnections(20 * time.Second) {
		assert.Fail(t, "nats server not ready")
	}
	nc, err := nats.Connect(server.ClientURL(), nats.InProcessServer(server))
	assert.NoError(t, err)
	assert.NotNil(t, nc)

	teardown := func() {
		nc.Drain()
		server.Shutdown()
	}
	return nc, teardown
}
