package natsutil

import "github.com/nats-io/nats.go"

// CloneHeader returns a deep copy of a NATS header map.
func CloneHeader(h nats.Header) nats.Header {
	if len(h) == 0 {
		return nats.Header{}
	}
	out := nats.Header{}
	for key, values := range h {
		for _, value := range values {
			out.Add(key, value)
		}
	}
	return out
}
