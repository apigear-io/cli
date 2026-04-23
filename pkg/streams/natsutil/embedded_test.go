package natsutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestEmbeddedServer(t *testing.T) {
	es, err := NewEmbeddedServer()
	assert.NoError(t, err)
	defer es.Close()
	assert.DirExists(t, es.storeDir)

	nc := es.NatsConn()
	assert.NotNil(t, nc)
	assert.False(t, nc.IsClosed())

	srv := es.Server()
	assert.NotNil(t, srv)
	assert.True(t, srv.ReadyForConnections(5*time.Second))

	es.Close()
	assert.NoDirExists(t, es.storeDir)
}
