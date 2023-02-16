package net

import (
	"bufio"
	"io"
	"os"
	"time"

	"github.com/apigear-io/cli/pkg/log"
)

// NDJSONScanner scans a reader line by line and writes to the writer.
type NDJSONScanner struct {
	sleep  time.Duration
	repeat int
}

// NewNDJSONScanner creates a new NDJSON scanner.
func NewNDJSONScanner(sleep time.Duration, repeat int) *NDJSONScanner {
	return &NDJSONScanner{
		sleep:  sleep,
		repeat: repeat,
	}
}

// Scan scans a reader line by line and writes to the writer.
func (s *NDJSONScanner) Scan(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	for i := 0; i < s.repeat; i++ {
		for scanner.Scan() {
			line := scanner.Bytes()
			log.Debug().Msgf("write: %s", line)
			_, err := w.Write(line)
			if err != nil {
				return err
			}
			if s.sleep > 0 {
				time.Sleep(s.sleep)
			}
		}
	}
	return scanner.Err()
}

// ScanFile scans a file line by line and writes to the writer.
func (s *NDJSONScanner) ScanFile(path string, w io.Writer) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return s.Scan(f, w)
}
