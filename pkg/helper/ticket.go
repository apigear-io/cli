package helper

import (
	"context"
	"time"
)

// EachTicked calls fn for each item in items every dt.
func EachTicked[T any](ctx context.Context, items []T, fn func(T), dt time.Duration) {
	ticker := time.NewTicker(dt)
	defer ticker.Stop()
	it := NewIterator(items)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if it.HasNext() {
				item, ok := it.Next()
				if !ok {
					return
				}
				fn(item)
			} else {
				return
			}
		}
	}
}
