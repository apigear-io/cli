package rpc

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/log"
	"github.com/stretchr/testify/assert"
)

func CreateHubWithServer(ctx context.Context) (*Hub, int, error) {
	port, err := helper.GetFreePort()
	if err != nil {
		return nil, 0, err
	}
	hub := NewHub(ctx)
	s := &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: hub}
	go func() {
		err = s.ListenAndServe()
		if err != nil {
			log.Errorf("server error: %v", err)
		}
	}()
	go func() {
		<-ctx.Done()
		err := s.Shutdown(ctx)
		if err != nil {
			log.Errorf("error shutting down server: %v", err)
		}
	}()
	return hub, port, nil
}

func TestHubCreate(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	NewHub(ctx)
	cancel()
}

func TestHubRegister(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	hub, _, err := CreateHubWithServer(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, hub)
	cancel()
}

func TestHubWithClient(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	hub, port, err := CreateHubWithServer(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, hub)
	conn, err := Dial(ctx, fmt.Sprintf("ws://localhost:%d", port))
	assert.NoError(t, err)
	assert.NotNil(t, conn)
	cancel()
}

func TestSendMessage(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	hub, port, err := CreateHubWithServer(ctx)
	assert.NoError(t, err)
	assert.NotNil(t, hub)
	conn, err := Dial(ctx, fmt.Sprintf("ws://localhost:%d", port))
	assert.NoError(t, err)
	assert.NotNil(t, conn)
	msg := MakeCall("test", 0, map[string]any{})
	go func(msg Message) {
		for req := range hub.Requests() {
			var m Message
			err := req.AsJSON(&m)
			assert.NoError(t, err)
			assert.Equal(t, msg, m)
			err = req.ReplyJSON(MakeResult(m.Id, "ok"))
			assert.NoError(t, err)
		}
	}(msg)
	err = conn.WriteJSON(msg)
	assert.NoError(t, err)
	var reply Message
	err = conn.ReadJSON(&reply)
	assert.NoError(t, err)
	cancel()
}
