package session

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/natsutil"
	"github.com/nats-io/nats.go/jetstream"
)

const DefaultBucket = config.SessionBucket

// Metadata captures information about a recorded session.
type Metadata struct {
	SessionID      string    `json:"session_id"`
	DeviceID       string    `json:"device_id"`
	SourceSubject  string    `json:"source_subject"`
	SessionSubject string    `json:"session_subject"`
	Stream         string    `json:"stream"`
	Bucket         string    `json:"bucket"`
	Start          time.Time `json:"start"`
	End            time.Time `json:"end"`
	MessageCount   int       `json:"message_count"`
	Retention      string    `json:"retention,omitempty"`
}

// SessionStore provides helper methods around session metadata backed by JetStream KV.
type SessionStore struct {
	js     jetstream.JetStream
	bucket string
	kv     jetstream.KeyValue
}

// NewSessionStore returns a Manager ensuring the session bucket exists.
func NewSessionStore(js jetstream.JetStream, bucket string) (*SessionStore, error) {
	if bucket == "" {
		bucket = config.SessionBucket
	}
	ctx := context.Background()
	kv, err := natsutil.EnsureKeyValue(ctx, js, bucket)
	if err != nil {
		return nil, fmt.Errorf("sessions bucket %s: %w", bucket, err)
	}
	return &SessionStore{js: js, bucket: bucket, kv: kv}, nil
}

// Bucket returns the configured bucket name.
func (m *SessionStore) Bucket() string {
	return m.bucket
}

// Put stores or updates session metadata in the bucket.
func (m *SessionStore) Put(meta *Metadata, revision uint64) (uint64, error) {
	if meta == nil {
		return revision, errors.New("metadata is nil")
	}
	data, err := json.Marshal(meta)
	if err != nil {
		return revision, err
	}
	ctx := context.Background()
	if revision == 0 {
		rev, err := m.kv.Create(ctx, meta.SessionID, data)
		if err != nil {
			return revision, err
		}
		return rev, nil
	}
	rev, err := m.kv.Update(ctx, meta.SessionID, data, revision)
	if err != nil {
		return revision, err
	}
	return rev, nil
}

// Load retrieves metadata for a session ID along with its revision.
func (m *SessionStore) Load(sessionID string) (*Metadata, uint64, error) {
	entry, err := m.kv.Get(context.Background(), sessionID)
	if err != nil {
		return nil, 0, err
	}
	var meta Metadata
	err = json.Unmarshal(entry.Value(), &meta)
	if err != nil {
		return nil, 0, err
	}
	if meta.SessionID == "" {
		meta.SessionID = sessionID
	}
	if meta.Bucket == "" {
		meta.Bucket = m.bucket
	}
	return &meta, entry.Revision(), nil
}

// Info retrieves metadata for a session without revision details.
func (m *SessionStore) Info(sessionID string) (*Metadata, error) {
	meta, _, err := m.Load(sessionID)
	return meta, err
}

// List returns all metadata entries in the bucket.
func (m *SessionStore) List() ([]Metadata, error) {
	keys, err := m.kv.Keys(context.Background())
	if err != nil {
		if errors.Is(err, jetstream.ErrNoKeysFound) {
			return nil, nil
		}
		return nil, err
	}

	sessions := make([]Metadata, 0, len(keys))
	for _, key := range keys {
		entry, err := m.kv.Get(context.Background(), key)
		if err != nil {
			continue
		}
		var meta Metadata
		err = json.Unmarshal(entry.Value(), &meta)
		if err != nil {
			continue
		}
		if meta.SessionID == "" {
			meta.SessionID = key
		}
		if meta.Bucket == "" {
			meta.Bucket = m.bucket
		}
		sessions = append(sessions, meta)
	}
	return sessions, nil
}

// Delete removes a session's metadata and JetStream stream.
func (m *SessionStore) Delete(sessionID string) error {
	meta, _, err := m.Load(sessionID)
	if err != nil {
		return err
	}
	ctx := context.Background()
	err = m.js.DeleteStream(ctx, meta.Stream)
	if err != nil && !errors.Is(err, jetstream.ErrStreamNotFound) {
		return fmt.Errorf("delete stream: %w", err)
	}
	return m.kv.Delete(ctx, sessionID)
}

// StreamName derives a sanitized stream name for a session identifier.
func StreamName(sessionID string) string {
	upper := strings.ToUpper(sessionID)
	upper = strings.ReplaceAll(upper, "-", "_")
	upper = strings.ReplaceAll(upper, ".", "_")
	return "STREAMS_" + upper
}
