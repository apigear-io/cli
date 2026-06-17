package controller

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

// SendCommand issues a controller command over NATS request/reply.
func SendCommand(ctx context.Context, nc *nats.Conn, subject string, cmd RpcRequest) (RpcResponse, error) {
	if nc == nil {
		return RpcResponse{}, errors.New("nats connection is nil")
	}
	if subject == "" {
		subject = DefaultCommandSubject
	}

	data, err := json.Marshal(cmd)
	if err != nil {
		return RpcResponse{}, err
	}

	msg, err := nc.RequestWithContext(ctx, subject, data)
	if err != nil {
		return RpcResponse{}, err
	}

	var resp RpcResponse
	err = json.Unmarshal(msg.Data, &resp)
	if err != nil {
		return RpcResponse{}, err
	}
	return resp, nil
}

// FetchState retrieves a session state snapshot from the controller KV bucket.
func FetchState(js jetstream.JetStream, bucket, sessionID string) (StateSnapshot, error) {
	snap := StateSnapshot{SessionID: sessionID}
	if js == nil {
		return snap, errors.New("jetstream context is nil")
	}
	if bucket == "" {
		bucket = DefaultStateBucket
	}
	kv, err := js.KeyValue(context.Background(), bucket)
	if err != nil {
		return snap, err
	}
	entry, err := kv.Get(context.Background(), sessionID)
	if err != nil {
		return snap, err
	}
	err = json.Unmarshal(entry.Value(), &snap)
	if err != nil {
		return snap, err
	}
	return snap, nil
}

// ListStates retrieves all session state snapshots from the controller KV bucket.
func ListStates(js jetstream.JetStream, bucket string) ([]StateSnapshot, error) {
	if js == nil {
		return nil, errors.New("jetstream context is nil")
	}
	if bucket == "" {
		bucket = DefaultStateBucket
	}
	kv, err := js.KeyValue(context.Background(), bucket)
	if err != nil {
		return nil, err
	}

	keys, err := kv.Keys(context.Background())
	if err != nil {
		if errors.Is(err, jetstream.ErrNoKeysFound) {
			return []StateSnapshot{}, nil
		}
		return nil, err
	}

	states := make([]StateSnapshot, 0, len(keys))
	for _, key := range keys {
		entry, err := kv.Get(context.Background(), key)
		if err != nil {
			continue
		}
		var snap StateSnapshot
		if err := json.Unmarshal(entry.Value(), &snap); err != nil {
			continue
		}
		if snap.SessionID == "" {
			snap.SessionID = key
		}
		states = append(states, snap)
	}
	return states, nil
}
