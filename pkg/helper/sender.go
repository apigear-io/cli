package helper

import (
	"time"
)

type SenderControl[T any] struct {
	sleep  time.Duration
	repeat int
}

func NewSenderControl[T any](repeat int, sleep time.Duration) *SenderControl[T] {
	return &SenderControl[T]{
		sleep:  sleep,
		repeat: repeat,
	}
}

func (t *SenderControl[T]) Run(items []T, send func(T) error) error {
	if t.repeat == 0 {
		t.repeat = 1
	}
	for i := 0; i < t.repeat; i++ {
		for _, item := range items {
			err := send(item)
			if err != nil {
				return err
			}
			if t.sleep > 0 {
				time.Sleep(t.sleep)
			}
		}
	}
	return nil
}
