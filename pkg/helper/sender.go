package helper

import (
	"time"
)

type SenderControl[T any] struct {
	interval time.Duration
	repeat   int
	batch    int
}

func NewSenderControl[T any](repeat int, interval time.Duration, batch int) *SenderControl[T] {
	if repeat < 0 {
		repeat = 1
	}
	if batch < 0 {
		batch = 1
	}
	if interval < 0 {
		interval = 100 * time.Millisecond
	}
	return &SenderControl[T]{
		interval: interval,
		repeat:   repeat,
		batch:    batch,
	}
}

func (t *SenderControl[T]) Run(items []T, send func(T) error) error {
	for i := 0; i < t.repeat; i++ {
		for j := 0; j < len(items); j += t.batch {
			end := j + t.batch
			if end > len(items) {
				end = len(items)
			}
			batch := items[j:end]
			for _, item := range batch {
				err := send(item)
				if err != nil {
					return err
				}
			}
			if t.interval > 0 {
				time.Sleep(t.interval)
			}
		}
	}
	return nil
}
