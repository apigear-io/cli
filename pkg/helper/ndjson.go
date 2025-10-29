package helper

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"os"
	"time"
)

// NDJSONScanner streams NDJSON content line by line to a callback.
type NDJSONScanner struct {
	Sleep  time.Duration
	Repeat int
}

// NewNDJSONScanner creates a new NDJSON scanner.
func NewNDJSONScanner(sleep time.Duration, repeat int) *NDJSONScanner {
	return &NDJSONScanner{Sleep: sleep, Repeat: repeat}
}

// OnLineFunc is invoked for each NDJSON line. Returning io.EOF stops the scan gracefully.
type OnLineFunc func(line []byte) error

// Scan streams lines from the reader to the callback.
func (s *NDJSONScanner) Scan(r io.Reader, fn OnLineFunc) error {
	if fn == nil {
		return errors.New("ndjson: callback cannot be nil")
	}

	repeat := s.Repeat
	if repeat == 0 {
		repeat = 1
	}

	run := func(reader io.Reader) error {
		scanner := bufio.NewScanner(reader)
		for scanner.Scan() {
			line := scanner.Bytes()
			dup := append([]byte(nil), line...)
			if err := fn(dup); err != nil {
				if errors.Is(err, io.EOF) {
					return nil
				}
				return err
			}
			if s.Sleep > 0 {
				time.Sleep(s.Sleep)
			}
		}
		return scanner.Err()
	}

	seeker, seekable := r.(io.Seeker)

	for pass := 0; repeat < 0 || pass < repeat; pass++ {
		if pass > 0 {
			if !seekable {
				return errors.New("ndjson: repeat requires seekable reader")
			}
			if _, err := seeker.Seek(0, io.SeekStart); err != nil {
				return err
			}
		}
		if err := run(r); err != nil {
			return err
		}
	}
	return nil
}

// ScanFile streams a file's NDJSON content to the callback.
func (s *NDJSONScanner) ScanFile(path string, fn OnLineFunc) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()
	return s.Scan(f, fn)
}

// ReadNDJSON reads all NDJSON entries from the reader into a slice.
func ReadNDJSON[T any](r io.Reader) ([]T, error) {
	var out []T
	scanner := NewNDJSONScanner(0, 1)
	err := scanner.Scan(r, func(line []byte) error {
		var item T
		if err := json.Unmarshal(line, &item); err != nil {
			return err
		}
		out = append(out, item)
		return nil
	})
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, err
	}
	return out, nil
}

// ReadNDJSONFile reads all NDJSON entries from the file into a slice.
func ReadNDJSONFile[T any](path string) ([]T, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()
	return ReadNDJSON[T](f)
}
