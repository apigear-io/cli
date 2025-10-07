package store

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

const DefaultDeviceBucket = config.DeviceBucket

// DeviceInfo captures descriptive information about a device being monitored.
type DeviceInfo struct {
	Description    string    `json:"description,omitempty"`
	Location       string    `json:"location,omitempty"`
	Owner          string    `json:"owner,omitempty"`
	Updated        time.Time `json:"updated,omitempty"`
	BufferDuration string    `json:"buffer_duration,omitempty"`
}

// IsZero reports whether the info carries any user-supplied metadata.
func (info DeviceInfo) IsZero() bool {
	return info.Description == "" && info.Location == "" && info.Owner == "" && info.BufferDuration == ""
}

// DeviceEntry represents a stored device profile.
type DeviceEntry struct {
	DeviceID string     `json:"device_id"`
	Info     DeviceInfo `json:"info"`
}

// DeviceStore helps manage device metadata in KV.
type DeviceStore struct {
	js     jetstream.JetStream
	bucket string
	kv     jetstream.KeyValue
}

// NewDeviceStore ensures the device bucket exists and returns a store instance.
func NewDeviceStore(js jetstream.JetStream, bucket string) (*DeviceStore, error) {
	if bucket == "" {
		bucket = config.DeviceBucket
	}
	ctx := context.Background()
	kv, err := natsutil.EnsureKeyValue(ctx, js, bucket)
	if err != nil {
		return nil, fmt.Errorf("device bucket %s: %w", bucket, err)
	}
	return &DeviceStore{js: js, bucket: bucket, kv: kv}, nil
}

// Bucket returns the bucket name.
func (s *DeviceStore) Bucket() string { return s.bucket }

func key(deviceID string) string {
	return strings.ToLower(strings.TrimSpace(deviceID))
}

// Upsert creates or updates a device profile.
func (s *DeviceStore) Upsert(deviceID string, update DeviceInfo) error {
	if deviceID = strings.TrimSpace(deviceID); deviceID == "" {
		return errors.New("device-id cannot be empty")
	}
	existing, rev, err := s.loadInternal(deviceID)
	if err != nil && !errors.Is(err, jetstream.ErrKeyNotFound) {
		return err
	}

	merged := mergeInfo(existing, update)
	if merged.IsZero() {
		return errors.New("no device information provided")
	}
	merged.Updated = time.Now().UTC()

	data, err := json.Marshal(merged)
	if err != nil {
		return err
	}

	k := key(deviceID)
	ctx := context.Background()
	if rev == 0 {
		_, err := s.kv.Create(ctx, k, data)
		if err == nil {
			return nil
		} else if err != nil && !errors.Is(err, jetstream.ErrKeyExists) {
			return err
		}
	}
	_, err = s.kv.Put(ctx, k, data)
	return err
}

// Ensure guarantees a device entry exists, creating a minimal placeholder when missing.
func (s *DeviceStore) Ensure(deviceID string) error {
	if deviceID = strings.TrimSpace(deviceID); deviceID == "" {
		return errors.New("device-id cannot be empty")
	}
	_, _, err := s.loadInternal(deviceID)
	if err == nil {
		return nil
	} else if !errors.Is(err, jetstream.ErrKeyNotFound) {
		return err
	}

	info := DeviceInfo{Updated: time.Now().UTC()}
	data, err := json.Marshal(info)
	if err != nil {
		return err
	}
	_, err = s.kv.Create(context.Background(), key(deviceID), data)
	if err != nil && !errors.Is(err, jetstream.ErrKeyExists) {
		return err
	}
	return nil
}

// Get fetches a device info entry.
func (s *DeviceStore) Get(deviceID string) (DeviceInfo, error) {
	info, _, err := s.loadInternal(deviceID)
	return info, err
}

// Delete removes a device profile.
func (s *DeviceStore) Delete(deviceID string) error {
	if deviceID = strings.TrimSpace(deviceID); deviceID == "" {
		return errors.New("device-id cannot be empty")
	}
	return s.kv.Delete(context.Background(), key(deviceID))
}

// List returns all device profiles.
func (s *DeviceStore) List() ([]DeviceEntry, error) {
	keys, err := s.kv.Keys(context.Background())
	if err != nil {
		if errors.Is(err, jetstream.ErrNoKeysFound) {
			return nil, nil
		}
		return nil, err
	}
	entries := make([]DeviceEntry, 0, len(keys))
	for _, k := range keys {
		info, _, err := s.loadInternal(k)
		if err != nil {
			continue
		}
		entries = append(entries, DeviceEntry{DeviceID: k, Info: info})
	}
	return entries, nil
}

func (s *DeviceStore) loadInternal(deviceID string) (DeviceInfo, uint64, error) {
	entry, err := s.kv.Get(context.Background(), key(deviceID))
	if err != nil {
		return DeviceInfo{}, 0, err
	}
	var info DeviceInfo
	err = json.Unmarshal(entry.Value(), &info)
	if err != nil {
		return DeviceInfo{}, 0, err
	}
	return info, entry.Revision(), nil
}

func mergeInfo(base, update DeviceInfo) DeviceInfo {
	info := base
	if update.Description != "" {
		info.Description = update.Description
	}
	if update.Location != "" {
		info.Location = update.Location
	}
	if update.Owner != "" {
		info.Owner = update.Owner
	}
	if update.BufferDuration != "" {
		info.BufferDuration = update.BufferDuration
	}
	return info
}
