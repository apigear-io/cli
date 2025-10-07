package msgio

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
)

// PublishOptions controls how a JSONL file is streamed to NATS.
type PublishOptions struct {
	ServerURL string
	Subject   string
	DeviceID  string
	FilePath  string
	Interval  time.Duration
	MaxLine   int
	Validate  bool
	Headers   map[string]string
	Verbose   bool
	Echo      bool
}

// PublishFromFile reads a JSONL file and publishes each line to the derived subject.
func PublishFromFile(ctx context.Context, opts PublishOptions) error {
	if opts.FilePath == "" {
		return errors.New("file path cannot be empty")
	}

	baseSubject := strings.TrimSpace(opts.Subject)
	if baseSubject == "" {
		return errors.New("subject cannot be empty")
	}

	deviceID := strings.TrimSpace(opts.DeviceID)
	if deviceID == "" {
		return errors.New("device-id cannot be empty")
	}

	if opts.ServerURL == "" {
		return errors.New("server URL cannot be empty")
	}

	fullSubject := config.DeviceSubject(baseSubject, deviceID)

	nc, err := nats.Connect(opts.ServerURL)
	if err != nil {
		return fmt.Errorf("connect to NATS: %w", err)
	}
	defer nc.Drain()

	file, err := os.Open(opts.FilePath)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	info, err := file.Stat()
	if err != nil {
		file.Close()
		return fmt.Errorf("stat file: %w", err)
	}
	if info.IsDir() {
		file.Close()
		return fmt.Errorf("%s is a directory", opts.FilePath)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 1024*1024)
	maxLine := opts.MaxLine
	if maxLine <= 0 {
		maxLine = 8 * 1024 * 1024
	}
	scanner.Buffer(buf, maxLine)

	lineNumber := 0
	for scanner.Scan() {
		select {
		case <-ctx.Done():
			if opts.Verbose {
				log.Warn().Int("lines", lineNumber).Msg("publish interrupted")
			}
			return ctx.Err()
		default:
		}

		rawLine := strings.TrimSpace(scanner.Text())
		if rawLine == "" {
			continue
		}
		lineNumber++

		if opts.Validate {
			var jsRaw json.RawMessage
			err := json.Unmarshal([]byte(rawLine), &jsRaw)
			if err != nil {
				return fmt.Errorf("line %d: invalid JSON: %w", lineNumber, err)
			}
		}

		msg := &nats.Msg{
			Subject: fullSubject,
			Header:  nats.Header{},
			Data:    []byte(rawLine),
		}

		msg.Header.Set("Content-Type", "application/json")
		msg.Header.Set(config.HeaderDevice, deviceID)
		msg.Header.Set(config.HeaderFile, filepath.Base(opts.FilePath))
		for k, v := range opts.Headers {
			msg.Header.Set(k, v)
		}

		err := nc.PublishMsg(msg)
		if err != nil {
			return fmt.Errorf("publish line %d: %w", lineNumber, err)
		}

		if opts.Verbose {
			log.Info().Int("line", lineNumber).Str("subject", fullSubject).Msg("published line")
		}

		if opts.Echo {
			fmt.Fprintln(os.Stdout, rawLine)
		}

		if opts.Interval > 0 {
			select {
			case <-ctx.Done():
				if opts.Verbose {
					log.Warn().Int("line", lineNumber).Msg("publish interrupted during interval")
				}
				return ctx.Err()
			case <-time.After(opts.Interval):
			}
		}
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("scan file: %w", err)
	}

	err = nc.Flush()
	if err != nil {
		return fmt.Errorf("flush connection: %w", err)
	}
	err = nc.LastError()
	if err != nil {
		return fmt.Errorf("nats error: %w", err)
	}

	if opts.Verbose {
		log.Info().Int("messages", lineNumber).Str("subject", fullSubject).Msg("completed publishing")
	}

	return nil
}
