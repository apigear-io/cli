package stream

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math/rand"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"text/template"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)

// publishOptions holds configuration for the publish command.
type publishOptions struct {
	Count      int
	Interval   time.Duration
	Wait       time.Duration
	NoResponse bool
	File       string
}

// outputMessage represents a message in JSON output format.
type outputMessage struct {
	Timestamp string `json:"timestamp"`
	Type      string `json:"type"`
	Data      string `json:"data"`
}

// templateData is passed to message templates on each iteration.
type templateData struct {
	Index     int    // 0-based iteration index
	Count     int    // total number of messages
	Timestamp string // current UTC timestamp (RFC3339)
}

// newTemplateFuncMap returns template functions for message generation.
func newTemplateFuncMap() template.FuncMap {
	return template.FuncMap{
		// Identifiers
		"uuid": gofakeit.UUID,
		"seq": func() string {
			return fmt.Sprintf("%d", rand.Int63())
		},

		// Time
		"now":       func() string { return time.Now().UTC().Format(time.RFC3339) },
		"timestamp": func() string { return time.Now().UTC().Format(time.RFC3339Nano) },
		"unixMilli": func() int64 { return time.Now().UnixMilli() },

		// Numbers
		"intRange": func(min, max int) int { return gofakeit.IntRange(min, max) },
		"float":    func(min, max float64) float64 { return gofakeit.Float64Range(min, max) },
		"boolean":  gofakeit.Bool,

		// Person
		"name":      gofakeit.Name,
		"firstName": gofakeit.FirstName,
		"lastName":  gofakeit.LastName,
		"email":     gofakeit.Email,
		"phone":     gofakeit.Phone,
		"username":  gofakeit.Username,
		"gender":    gofakeit.Gender,

		// Address
		"city":    gofakeit.City,
		"country": gofakeit.Country,
		"street":  gofakeit.Street,
		"zip":     gofakeit.Zip,
		"state":   gofakeit.State,
		"lat":     gofakeit.Latitude,
		"lon":     gofakeit.Longitude,

		// Internet
		"url":       gofakeit.URL,
		"domain":    gofakeit.DomainName,
		"ipv4":      gofakeit.IPv4Address,
		"ipv6":      gofakeit.IPv6Address,
		"userAgent": gofakeit.UserAgent,

		// Company
		"company":  gofakeit.Company,
		"jobTitle": gofakeit.JobTitle,
		"buzzWord": gofakeit.BuzzWord,

		// Text
		"word":     gofakeit.Word,
		"sentence": func() string { return gofakeit.Sentence(6) },
		"phrase":   gofakeit.Phrase,
		"question": gofakeit.Question,
		"noun":     gofakeit.Noun,
		"verb":     gofakeit.Verb,

		// Product
		"productName":     gofakeit.ProductName,
		"productCategory": gofakeit.ProductCategory,
		"price": func(min, max float64) float64 {
			return gofakeit.Price(min, max)
		},

		// Misc
		"hexColor": gofakeit.HexColor,
		"color":    gofakeit.SafeColor,
		"emoji":    gofakeit.Emoji,
		"animal":   gofakeit.Animal,
		"appName":  gofakeit.AppName,
	}
}

// renderTemplate parses and executes a Go template with the given data.
func renderTemplate(tmpl *template.Template, data templateData) (string, error) {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("template error: %w", err)
	}
	return buf.String(), nil
}

// validateWSURL checks that the URL has a ws:// or wss:// scheme.
func validateWSURL(rawURL string) error {
	u, err := url.Parse(rawURL)
	if err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}
	if u.Scheme != "ws" && u.Scheme != "wss" {
		return fmt.Errorf("URL scheme must be ws:// or wss://, got %q", u.Scheme)
	}
	return nil
}

// resolveMessage returns the message template string from either the argument or a file.
func resolveMessage(message string, opts publishOptions) (string, error) {
	if opts.File != "" {
		data, err := os.ReadFile(opts.File)
		if err != nil {
			return "", fmt.Errorf("failed to read template file: %w", err)
		}
		return string(data), nil
	}
	return message, nil
}

// runPublish connects to a WebSocket, sends messages, and prints responses.
func runPublish(ctx context.Context, stdout, stderr io.Writer, wsURL, message string, opts publishOptions) error {
	if err := validateWSURL(wsURL); err != nil {
		return err
	}

	// Resolve template source
	tmplSource, err := resolveMessage(message, opts)
	if err != nil {
		return err
	}

	// Parse template
	tmpl, err := template.New("msg").Funcs(newTemplateFuncMap()).Parse(tmplSource)
	if err != nil {
		return fmt.Errorf("invalid template: %w", err)
	}

	dialer := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
	}

	// Retry connection with backoff
	var conn *websocket.Conn
	delays := []time.Duration{0, 500 * time.Millisecond, 1 * time.Second, 2 * time.Second, 4 * time.Second}
	for i, delay := range delays {
		if delay > 0 {
			fmt.Fprintf(stderr, "Retrying in %s (%d/%d)...\n", delay, i, len(delays)-1)
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(delay):
			}
		}
		var dialErr error
		conn, _, dialErr = dialer.DialContext(ctx, wsURL, nil)
		if dialErr == nil {
			break
		}
		if i == len(delays)-1 {
			return fmt.Errorf("failed to connect after %d attempts: %w", len(delays), dialErr)
		}
	}
	defer conn.Close()

	fmt.Fprintf(stderr, "Connected to %s\n", wsURL)

	// Start reading responses immediately so the server's write buffer
	// doesn't fill up and close the connection during bulk sends.
	// The channel also signals if the server closes the connection.
	readDone := make(chan error, 1)
	if !opts.NoResponse {
		go func() {
			for {
				_, msg, err := conn.ReadMessage()
				if err != nil {
					// Extract close reason if available
					if closeErr, ok := err.(*websocket.CloseError); ok {
						readDone <- fmt.Errorf("server closed connection: %s (code %d)", closeErr.Text, closeErr.Code)
					} else {
						readDone <- nil
					}
					return
				}
				fmt.Fprintf(stdout, "< %s\n", string(msg))
			}
		}()
	}
	// When NoResponse, readDone is never written to — selects fall through to default.

	for i := 0; i < opts.Count; i++ {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		// Check if server closed the connection
		select {
		case err := <-readDone:
			if err != nil {
				return err
			}
			return fmt.Errorf("server closed connection after %d/%d messages", i, opts.Count)
		default:
		}

		data := templateData{
			Index:     i,
			Count:     opts.Count,
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		}
		rendered, err := renderTemplate(tmpl, data)
		if err != nil {
			return err
		}

		if err := conn.WriteMessage(websocket.TextMessage, []byte(rendered)); err != nil {
			// Check if server sent a close reason
			select {
			case readErr := <-readDone:
				if readErr != nil {
					return readErr
				}
			default:
			}
			return fmt.Errorf("failed to send message %d/%d: %w", i+1, opts.Count, err)
		}
		fmt.Fprintf(stdout, "> %s\n", rendered)

		if i < opts.Count-1 && opts.Interval > 0 {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case err := <-readDone:
				if err != nil {
					return err
				}
				return fmt.Errorf("server closed connection after %d/%d messages", i+1, opts.Count)
			case <-time.After(opts.Interval):
			}
		}
	}

	if opts.NoResponse {
		// Send close frame, best effort
		conn.WriteMessage(websocket.CloseMessage,
			websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
		fmt.Fprintf(stderr, "Disconnected\n")
		return nil
	}

	// Wait for remaining responses after all sends complete
	select {
	case <-ctx.Done():
	case <-time.After(opts.Wait):
	case <-readDone:
	}

	// Send close frame, best effort
	conn.WriteMessage(websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	fmt.Fprintf(stderr, "Disconnected\n")
	return nil
}

// NewPublishCommand creates the publish command.
func NewPublishCommand() *cobra.Command {
	opts := &publishOptions{
		Count:    1,
		Wait:     2 * time.Second,
		Interval: 0,
	}

	cmd := &cobra.Command{
		Use:   "publish <url> [message]",
		Short: "Publish messages to a WebSocket server",
		Long: `Connect to a WebSocket server, send one or more messages, and print responses.

The message supports Go template syntax with built-in functions for
generating dynamic data. Use --count and --interval for mass messaging.

Template variables:
  {{.Index}}      0-based iteration index
  {{.Count}}      total number of messages
  {{.Timestamp}}  current UTC timestamp (RFC3339)

Template functions (subset):
  {{uuid}}            random UUID
  {{name}}            full name          {{email}}     email address
  {{username}}        username           {{phone}}     phone number
  {{city}}            city name          {{country}}   country name
  {{company}}         company name       {{jobTitle}}  job title
  {{sentence}}        random sentence    {{word}}      random word
  {{intRange 1 100}}  random int         {{float 0 1}} random float
  {{boolean}}         random bool        {{ipv4}}      IPv4 address
  {{now}}             UTC timestamp      {{unixMilli}} unix millis
  {{productName}}     product name       {{price 1 100}} random price

Output uses > for sent messages and < for received messages.
Informational messages go to stderr; data goes to stdout (pipeable).

Examples:
  # Send a simple message
  apigear stream publish ws://localhost:8888/ws "hello"

  # Send with template (generates unique JSON each time)
  apigear stream publish ws://localhost:8888/ws \
    '{"id":"{{uuid}}","user":"{{name}}","index":{{.Index}}}' --count 100

  # Mass send with interval
  apigear stream publish ws://localhost:8888/ws \
    '{"seq":{{.Index}},"data":"{{sentence}}"}' --count 1000 --interval 10ms

  # Read template from file
  apigear stream publish ws://localhost:8888/ws --file message.tmpl --count 50

  # Fire-and-forget (no response wait)
  apigear stream publish ws://localhost:8888/ws "ping" --no-response`,
		Args:          cobra.RangeArgs(1, 2),
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			var message string
			if len(args) > 1 {
				message = args[1]
			}
			if message == "" && opts.File == "" {
				return fmt.Errorf("provide a message argument or use --file")
			}
			ctx, stop := signal.NotifyContext(cmd.Context(), syscall.SIGINT, syscall.SIGTERM)
			defer stop()
			if err := runPublish(ctx, cmd.OutOrStdout(), cmd.ErrOrStderr(), args[0], message, *opts); err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "Error: %s\n", err)
				return err
			}
			return nil
		},
	}

	cmd.Flags().IntVarP(&opts.Count, "count", "n", opts.Count, "send message N times")
	cmd.Flags().DurationVarP(&opts.Interval, "interval", "i", opts.Interval, "delay between sends when count > 1")
	cmd.Flags().DurationVar(&opts.Wait, "wait", opts.Wait, "how long to wait for responses")
	cmd.Flags().BoolVar(&opts.NoResponse, "no-response", opts.NoResponse, "don't wait for response, just send and exit")
	cmd.Flags().StringVarP(&opts.File, "file", "f", "", "read message template from file")

	return cmd
}
