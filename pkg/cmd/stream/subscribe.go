package stream

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)

// subscribeOptions holds configuration for the subscribe command.
type subscribeOptions struct {
	Listen  string
	Echo    bool
	Format  string
	onReady func(addr string) // called when server is listening (for tests)
}

var subscribeUpgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

// runSubscribe starts a WebSocket server that prints received messages to stdout.
func runSubscribe(ctx context.Context, stdout, stderr io.Writer, opts subscribeOptions) error {
	u, err := url.Parse(opts.Listen)
	if err != nil {
		return fmt.Errorf("invalid listen URL: %w", err)
	}
	if u.Scheme != "ws" && u.Scheme != "wss" {
		return fmt.Errorf("listen URL scheme must be ws:// or wss://, got %q", u.Scheme)
	}

	host := u.Host
	path := u.Path
	if path == "" {
		path = "/"
	}

	mux := http.NewServeMux()
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		conn, err := subscribeUpgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			text := string(msg)

			if opts.Echo {
				fmt.Fprintf(stdout, "-> %s\n", text)
				if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
					return
				}
				fmt.Fprintf(stdout, "<- %s\n", text)
			} else if opts.Format == "json" {
				om := outputMessage{
					Timestamp: time.Now().UTC().Format(time.RFC3339Nano),
					Type:      "text",
					Data:      text,
				}
				data, _ := json.Marshal(om)
				fmt.Fprintf(stdout, "%s\n", data)
			} else {
				fmt.Fprintf(stdout, "%s\n", text)
			}
		}
	})

	ln, err := net.Listen("tcp", host)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", host, err)
	}

	addr := fmt.Sprintf("ws://%s%s", ln.Addr().String(), path)
	fmt.Fprintf(stderr, "Listening on %s\n", addr)

	if opts.onReady != nil {
		opts.onReady(addr)
	}

	server := &http.Server{Handler: mux}

	// Shutdown when context is cancelled
	go func() {
		<-ctx.Done()
		server.Close()
	}()

	err = server.Serve(ln)
	if err == http.ErrServerClosed {
		fmt.Fprintf(stderr, "Stopped\n")
		return nil
	}
	return err
}

// NewSubscribeCommand creates the subscribe command.
func NewSubscribeCommand() *cobra.Command {
	opts := &subscribeOptions{
		Listen: "ws://localhost:8888/ws",
		Format: "text",
	}

	cmd := &cobra.Command{
		Use:   "subscribe",
		Short: "Start a WebSocket server and print received messages",
		Long: `Start a lightweight WebSocket server that prints all received messages to stdout.

By default, messages are printed as raw text (one per line), making the output
pipeable. Use --echo to echo messages back to clients. Use --format json for
structured output.

Informational messages (Listening, Stopped) go to stderr; data goes to stdout.

Examples:
  # Start and print messages
  apigear stream subscribe

  # Custom listen address
  apigear stream subscribe --listen ws://localhost:9999/ws

  # Echo mode (print and echo back)
  apigear stream subscribe --echo

  # JSON output
  apigear stream subscribe --format json

  # Pipe to a file
  apigear stream subscribe > messages.log`,
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, stop := signal.NotifyContext(cmd.Context(), syscall.SIGINT, syscall.SIGTERM)
			defer stop()
			if err := runSubscribe(ctx, cmd.OutOrStdout(), cmd.ErrOrStderr(), *opts); err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "Error: %s\n", err)
				return err
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&opts.Listen, "listen", opts.Listen, "listen address (ws://host:port/path)")
	cmd.Flags().BoolVar(&opts.Echo, "echo", false, "echo messages back to clients")
	cmd.Flags().StringVar(&opts.Format, "format", opts.Format, `output format: "text" or "json"`)

	return cmd
}
