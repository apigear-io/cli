package controller_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/apigear-io/cli/pkg/streams/buffer"
	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/controller"
	"github.com/apigear-io/cli/pkg/streams/natsutil"
	"github.com/apigear-io/cli/pkg/streams/session"
	"github.com/apigear-io/cli/pkg/streams/store"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/stretchr/testify/require"
)

func TestControllerStartStop(t *testing.T) {
	h := newControllerHarness(t)
	t.Cleanup(h.Close)

	client := h.NewClientConn()
	defer client.Close()

	sessionID := "test-session"
	subject := "monitor"
	deviceID := "device-1"

	startResp := sendStart(t, client, controller.Command{
		Action:        controller.ActionStart,
		Subject:       subject,
		DeviceID:      deviceID,
		SessionID:     sessionID,
		SessionBucket: session.DefaultBucket,
		DeviceBucket:  store.DefaultDeviceBucket,
	})
	require.True(t, startResp.OK, "start response should succeed: %s", startResp.Message)

	// Allow recorder goroutine to establish subscriptions.
	time.Sleep(100 * time.Millisecond)

	publisher := h.NewClientConn()
	defer publisher.Close()

	for i := 0; i < 3; i++ {
		msg := fmt.Sprintf("{\"i\":%d}", i)
		require.NoError(t, publisher.Publish(subject+"."+deviceID, []byte(msg)))
	}
	require.NoError(t, publisher.Flush())

	require.Eventually(t, func() bool {
		snap, err := controller.FetchState(h.ctrlJS, controller.DefaultStateBucket, sessionID)
		if err != nil {
			return false
		}
		return snap.Status == "running" && snap.MessageCount == 3
	}, 3*time.Second, 100*time.Millisecond, "controller did not capture messages")

	stopResp := sendStop(t, client, sessionID)
	require.True(t, stopResp.OK, "stop response should succeed: %s", stopResp.Message)

	require.Eventually(t, func() bool {
		snap, err := controller.FetchState(h.ctrlJS, controller.DefaultStateBucket, sessionID)
		if err != nil {
			return false
		}
		return snap.Status == "stopped" && snap.MessageCount == 3
	}, 3*time.Second, 100*time.Millisecond, "controller state not updated after stop")

	mgr, err := session.NewSessionStore(h.ctrlJS, session.DefaultBucket)
	require.NoError(t, err)
	meta, _, err := mgr.Load(sessionID)
	require.NoError(t, err)
	require.Equal(t, 3, meta.MessageCount)
}

func TestControllerDuplicateStart(t *testing.T) {
	h := newControllerHarness(t)
	t.Cleanup(h.Close)

	client := h.NewClientConn()
	defer client.Close()

	cmd := controller.Command{
		Action:        controller.ActionStart,
		Subject:       "monitor",
		DeviceID:      "device-1",
		SessionID:     "dup-session",
		SessionBucket: session.DefaultBucket,
		DeviceBucket:  store.DefaultDeviceBucket,
	}

	resp := sendStart(t, client, cmd)
	require.True(t, resp.OK)

	dup := sendStart(t, client, cmd)
	require.False(t, dup.OK)
	require.Contains(t, dup.Message, "already running")
}

func TestControllerPreRoll(t *testing.T) {
	h := newControllerHarness(t)
	t.Cleanup(h.Close)

	// Configure device buffer
	devStore, err := store.NewDeviceStore(h.ctrlJS, store.DefaultDeviceBucket)
	require.NoError(t, err)
	require.NoError(t, devStore.Upsert("preroll-device", store.DeviceInfo{BufferDuration: "5m"}))

	_, subject, err := buffer.EnsureStream(h.ctrlJS, "preroll-device", 5*time.Minute)
	require.NoError(t, err)

	recordedAt := time.Now().Add(-30 * time.Second).UTC()
	msg := &nats.Msg{
		Subject: subject,
		Header:  nats.Header{},
		Data:    []byte(`{"preroll":true}`),
	}
	msg.Header.Set(config.HeaderBufferedAt, recordedAt.Format(time.RFC3339Nano))
	_, err = h.ctrlJS.PublishMsg(context.Background(), msg)
	require.NoError(t, err)

	client := h.NewClientConn()
	defer client.Close()

	resp := sendStart(t, client, controller.Command{
		Action:        controller.ActionStart,
		Subject:       "monitor",
		DeviceID:      "preroll-device",
		SessionID:     "preroll-session",
		SessionBucket: session.DefaultBucket,
		DeviceBucket:  store.DefaultDeviceBucket,
		PreRoll:       "2m",
	})
	require.True(t, resp.OK, resp.Message)

	require.Eventually(t, func() bool {
		snap, err := controller.FetchState(h.ctrlJS, controller.DefaultStateBucket, "preroll-session")
		return err == nil && snap.Status == "running" && snap.MessageCount >= 1
	}, 2*time.Second, 100*time.Millisecond, "pre-roll data not observed")

	stop := sendStop(t, client, "preroll-session")
	require.True(t, stop.OK)

	mgr, err := session.NewSessionStore(h.ctrlJS, session.DefaultBucket)
	require.NoError(t, err)
	meta, _, err := mgr.Load("preroll-session")
	require.NoError(t, err)
	require.GreaterOrEqual(t, meta.MessageCount, 1)

	require.Eventually(t, func() bool {
		snap, err := controller.FetchState(h.ctrlJS, controller.DefaultStateBucket, "preroll-session")
		return err == nil && snap.Status == "stopped" && snap.MessageCount >= 1
	}, 2*time.Second, 100*time.Millisecond)
}

func TestControllerStopWithoutStart(t *testing.T) {
	h := newControllerHarness(t)
	t.Cleanup(h.Close)

	client := h.NewClientConn()
	defer client.Close()

	resp := sendStop(t, client, "missing-session")
	require.True(t, resp.OK)
	require.Equal(t, "no active recording", resp.Message)
}

func TestControllerInvalidAction(t *testing.T) {
	h := newControllerHarness(t)
	t.Cleanup(h.Close)

	client := h.NewClientConn()
	defer client.Close()

	resp, err := controller.SendCommand(context.Background(), client, controller.DefaultCommandSubject, controller.Command{Action: "bogus"})
	require.NoError(t, err)
	require.False(t, resp.OK)
	require.Contains(t, resp.Message, "unknown action")
}

type controllerHarness struct {
	t         *testing.T
	srv       *natsutil.ServerHandle
	ctrlJS    jetstream.JetStream
	serverURL string
	cancel    context.CancelFunc
	errCh     chan error
}

func newControllerHarness(t *testing.T) *controllerHarness {
	t.Helper()

	srv, err := natsutil.StartServer(natsutil.ServerConfig{
		Options: &server.Options{
			JetStream: true,
			StoreDir:  t.TempDir(),
		},
	})
	require.NoError(t, err)

	js, err := natsutil.ConnectJetStream(srv.ClientURL())
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	errCh := make(chan error, 1)
	go func() {
		errCh <- controller.Run(ctx, js, controller.Options{ServerURL: srv.ClientURL()})
	}()

	harness := &controllerHarness{
		t:         t,
		srv:       srv,
		ctrlJS:    js,
		serverURL: srv.ClientURL(),
		cancel:    cancel,
		errCh:     errCh,
	}

	// Give the subscription a moment to be registered.
	time.Sleep(50 * time.Millisecond)

	return harness
}

func (h *controllerHarness) NewClientConn() *nats.Conn {
	h.t.Helper()
	conn, err := nats.Connect(h.serverURL)
	require.NoError(h.t, err)
	h.t.Cleanup(func() {
		conn.Drain()
	})
	return conn
}

func (h *controllerHarness) Close() {
	h.t.Helper()
	h.cancel()
	select {
	case err := <-h.errCh:
		if err != nil && err != context.Canceled {
			h.t.Fatalf("controller run error: %v", err)
		}
	case <-time.After(2 * time.Second):
		h.t.Fatal("controller did not shut down")
	}
	h.ctrlJS.Conn().Drain()
	h.srv.Shutdown()
}

func sendStart(t *testing.T, nc *nats.Conn, cmd controller.Command) controller.Response {
	t.Helper()
	resp, err := controller.SendCommand(context.Background(), nc, controller.DefaultCommandSubject, cmd)
	require.NoError(t, err)
	return resp
}

func sendStop(t *testing.T, nc *nats.Conn, sessionID string) controller.Response {
	t.Helper()
	resp, err := controller.SendCommand(context.Background(), nc, controller.DefaultCommandSubject, controller.Command{
		Action:    controller.ActionStop,
		SessionID: sessionID,
	})
	require.NoError(t, err)
	return resp
}
