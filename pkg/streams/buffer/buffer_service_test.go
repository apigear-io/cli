package buffer_test

import (
	"context"
	"testing"
	"time"

	"github.com/apigear-io/cli/pkg/streams/buffer"
	"github.com/apigear-io/cli/pkg/streams/natsutil"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/require"
)

func TestRunBufferMirrorsMessages(t *testing.T) {
	srv, err := natsutil.StartServer(natsutil.ServerConfig{Options: &server.Options{JetStream: true, StoreDir: t.TempDir()}})
	require.NoError(t, err)
	t.Cleanup(srv.Shutdown)

	js, err := natsutil.ConnectJetStream(srv.ClientURL())
	require.NoError(t, err)
	t.Cleanup(js.Conn().Close)

	// No need to create device metadata - buffering is now always on
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan error, 1)
	go func() {
		done <- buffer.RunBuffer(ctx, js, buffer.BufferOptions{MonitorSubject: "monitor"})
	}()

	// Brief delay to let buffer service start
	time.Sleep(50 * time.Millisecond)

	pub, err := nats.Connect(srv.ClientURL())
	require.NoError(t, err)
	t.Cleanup(pub.Close)

	require.NoError(t, pub.Publish("monitor.device-a", []byte(`{"temp":21}`)))
	require.NoError(t, pub.Flush())

	// Wait for append
	require.Eventually(t, func() bool {
		_, _, err := buffer.EnsureStream(js, "device-a", 2*time.Minute)
		if err != nil {
			return false
		}
		stream, err := js.Stream(context.Background(), "STREAMS_BUFFER_DEVICE_A")
		if err != nil {
			return false
		}
		info, err := stream.Info(context.Background())
		if err != nil {
			return false
		}
		return info.State.Msgs > 0
	}, 2*time.Second, 100*time.Millisecond)

	cancel()
	require.Eventually(t, func() bool {
		select {
		case err := <-done:
			return err == context.Canceled || err == nil
		default:
			return false
		}
	}, time.Second, 50*time.Millisecond)
}
