package helper

import (
	"context"
	"time"
)

// EachTicked calls fn for each item in items every dt.
func EachTicked[T any](ctx context.Context, items []T, fn func(T), dt time.Duration) {
	ticker := time.NewTicker(dt)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			for _, item := range items {
				fn(item)
			}
		}
	}
}
