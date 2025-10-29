package net

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/gorilla/websocket"
)

// WSClientOptions configure the wscat client behaviour.
type WSClientOptions struct {
	URL        string
	Headers    http.Header
	Dialer     *websocket.Dialer
	File       string
	DecodeJSON bool
	Interval   time.Duration
	Repeat     int
	OnMessage  func(messageType int, payload []byte)
	OnClose    func(error)
	OnSend     func([]byte)
}

// RunWSClient connects to the WebSocket endpoint and either reads from stdin or the configured file.
func RunWSClient(ctx context.Context, opts WSClientOptions) error {
	if opts.URL == "" {
		return fmt.Errorf("wscat: url cannot be empty")
	}

	dialer := opts.Dialer
	if dialer == nil {
		dialer = websocket.DefaultDialer
	}

	conn, _, err := dialer.DialContext(ctx, opts.URL, opts.Headers)
	if err != nil {
		return err
	}
	defer func() {
		_ = conn.Close()
	}()

	recvErr := make(chan error, 1)
	go func() {
		for {
			typ, payload, err := conn.ReadMessage()
			if err != nil {
				recvErr <- err
				return
			}
			if opts.OnMessage != nil {
				opts.OnMessage(typ, payload)
			}
		}
	}()

	sendErr := make(chan error, 1)
	go func() {
		var err error
		if opts.File != "" {
			err = sendFile(ctx, conn, opts.File, opts.Interval, opts.Repeat, opts.DecodeJSON, opts.OnSend)
		} else {
			err = sendInteractive(ctx, conn, opts.DecodeJSON, opts.OnSend)
		}
		sendErr <- err
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-recvErr:
		if opts.OnClose != nil {
			opts.OnClose(err)
		}
		return err
	case err := <-sendErr:
		if opts.OnClose != nil {
			opts.OnClose(err)
		}
		return err
	}
}

func sendInteractive(ctx context.Context, conn *websocket.Conn, decodeJSON bool, onSend func([]byte)) error {
	reader := bufio.NewReader(os.Stdin)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		payload := append([]byte{}, processPayload(line, decodeJSON)...)
		if onSend != nil {
			onSend(payload)
		}
		if err := conn.WriteMessage(websocket.TextMessage, payload); err != nil {
			return err
		}
	}
}

func sendFile(ctx context.Context, conn *websocket.Conn, path string, interval time.Duration, repeat int, decodeJSON bool, onSend func([]byte)) error {
	repeatCount := repeat
	if repeatCount == 0 {
		repeatCount = 1
	}

	for pass := 0; repeatCount < 0 || pass < repeatCount; pass++ {
		f, err := os.Open(path)
		if err != nil {
			return err
		}

		scanErr := func() error {
			defer func() {
				_ = f.Close()
			}()
			scanner := helper.NewNDJSONScanner(interval, 1)
			return scanner.Scan(f, func(line []byte) error {
				select {
				case <-ctx.Done():
					return ctx.Err()
				default:
				}
				payload := append([]byte{}, processPayload(line, decodeJSON)...)
				if onSend != nil {
					onSend(payload)
				}
				return conn.WriteMessage(websocket.TextMessage, payload)
			})
		}()
		if scanErr != nil {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}
			if errors.Is(scanErr, context.Canceled) {
				return ctx.Err()
			}
			return scanErr
		}
	}
	return nil
}

func processPayload(input []byte, decode bool) []byte {
	if !decode {
		return input
	}
	var anyVal any
	if err := json.Unmarshal(input, &anyVal); err != nil {
		return input
	}
	if normalized, err := json.Marshal(anyVal); err == nil {
		return normalized
	}
	return input
}
