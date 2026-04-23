package natsutil

import (
	"context"
	"errors"

	"github.com/nats-io/nats.go/jetstream"
)

// EnsureKeyValue returns a key-value bucket, creating it if missing.
func EnsureKeyValue(ctx context.Context, js jetstream.JetStream, bucket string) (jetstream.KeyValue, error) {
	kv, err := js.KeyValue(ctx, bucket)
	if errors.Is(err, jetstream.ErrBucketNotFound) {
		kv, err = js.CreateKeyValue(ctx, jetstream.KeyValueConfig{Bucket: bucket})
	}
	return kv, err
}
