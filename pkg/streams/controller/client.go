package controller

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

// SendCommand issues a controller command over NATS request/reply.
func SendCommand(ctx context.Context, nc *nats.Conn, subject string, cmd Command) (Response, error) {
	if nc == nil {
		return Response{}, errors.New("nats connection is nil")
	}
	if subject == "" {
		subject = DefaultCommandSubject
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return Response{}, err
	}

	msg, err := nc.RequestWithContext(ctx, subject, data)
	if err != nil {
		return Response{}, err
	}

	var resp Response
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return Response{}, err
	}
	return resp, nil
}

// FetchState retrieves a session state snapshot from the controller KV bucket.
func FetchState(js jetstream.JetStream, bucket, sessionID string) (StateSnapshot, error) {
	if js == nil {
		return StateSnapshot{}, errors.New("jetstream context is nil")
	}
	if bucket == "" {
		bucket = DefaultStateBucket
	}
	kv, err := js.KeyValue(context.Background(), bucket)
	if err != nil {
		return StateSnapshot{}, err
	}
	entry, err := kv.Get(context.Background(), sessionID)
	if err != nil {
		return StateSnapshot{}, err
	}
	var snap StateSnapshot
	err = json.Unmarshal(entry.Value(), &snap)
	if err != nil {
		return StateSnapshot{}, err
	}
	return snap, nil
}
