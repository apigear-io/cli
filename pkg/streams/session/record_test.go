package session_test

import (
	"context"
	"testing"
	"time"

	"github.com/apigear-io/cli/pkg/streams/natsutil"
	"github.com/apigear-io/cli/pkg/streams/session"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/require"
)

func TestRecordProgressCallback(t *testing.T) {
	srv, err := natsutil.StartServer(natsutil.ServerConfig{
		Options: &server.Options{
			JetStream: true,
			StoreDir:  t.TempDir(),
		},
	})
	require.NoError(t, err)
	t.Cleanup(srv.Shutdown)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	progressCh := make(chan session.Metadata, 4)
	opts := session.RecordOptions{
		ServerURL:     srv.ClientURL(),
		Subject:       "monitor",
		DeviceID:      "device-progress",
		SessionBucket: session.DefaultBucket,
		Progress: func(meta session.Metadata) {
			progressCh <- meta
		},
	}

	metaCh := make(chan *session.Metadata, 1)
	errCh := make(chan error, 1)

	go func() {
		meta, err := session.Record(ctx, opts)
		metaCh <- meta
		errCh <- err
	}()

	time.Sleep(100 * time.Millisecond)

	publisher, err := nats.Connect(srv.ClientURL())
	require.NoError(t, err)
	t.Cleanup(publisher.Close)

	require.NoError(t, publisher.Publish("monitor.device-progress", []byte(`{"hello":true}`)))
	require.NoError(t, publisher.Flush())

	var update session.Metadata
	require.Eventually(t, func() bool {
		select {
		case update = <-progressCh:
			return update.MessageCount >= 1
		default:
			return false
		}
	}, 2*time.Second, 50*time.Millisecond, "expected progress update")
	require.GreaterOrEqual(t, update.MessageCount, 1)
	require.Equal(t, "device-progress", update.DeviceID)

	cancel()

	select {
	case err := <-errCh:
		require.NoError(t, err)
	case <-time.After(2 * time.Second):
		t.Fatal("record did not stop")
	}

	select {
	case meta := <-metaCh:
		require.NotNil(t, meta)
		require.GreaterOrEqual(t, meta.MessageCount, 1)
	default:
		t.Fatal("expected metadata result")
	}
}
