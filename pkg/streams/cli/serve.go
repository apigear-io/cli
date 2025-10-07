package cli

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/apigear-io/cli/pkg/streams/buffer"
	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/controller"
	"github.com/apigear-io/cli/pkg/streams/natsutil"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

type serviceAllOptions struct {
	Host           string
	Port           int
	StoreDir       string
	Embedded       bool
	NoNATS         bool
	CommandSubject string
	StateBucket    string
	DeviceBucket   string
	MonitorSubject string
	BufferRefresh  time.Duration
}

func newServeCmd() *cobra.Command {
	opts := &serviceAllOptions{
		Host:           "127.0.0.1",
		Port:           4222,
		CommandSubject: config.CommandSubject,
		StateBucket:    config.StateBucket,
		DeviceBucket:   config.DeviceBucket,
		MonitorSubject: config.MonitorSubject,
		BufferRefresh:  15 * time.Second,
	}

	cmd := &cobra.Command{
		Use:     "serve",
		Short:   "Serve controller and buffer services and optional NATS server",
		Aliases: []string{"run"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runServiceAll(cmd, opts)
		},
	}

	cmd.Flags().StringVar(&opts.Host, "host", opts.Host, "Host interface for the embedded NATS server")
	cmd.Flags().IntVar(&opts.Port, "port", opts.Port, "Port for embedded NATS (use -1 for random)")
	cmd.Flags().StringVar(&opts.StoreDir, "store", "", "Directory for JetStream storage (defaults to temp)")
	cmd.Flags().BoolVar(&opts.Embedded, "embedded", false, "Use in-process client connections when running embedded NATS")
	cmd.Flags().BoolVar(&opts.NoNATS, "external", false, "Use an external NATS server instead of starting one")
	cmd.Flags().StringVar(&opts.CommandSubject, "command-subject", opts.CommandSubject, "Subject for controller commands")
	cmd.Flags().StringVar(&opts.StateBucket, "state-bucket", opts.StateBucket, "KV bucket for controller state")
	cmd.Flags().StringVar(&opts.DeviceBucket, "device-bucket", opts.DeviceBucket, "Device metadata bucket")
	cmd.Flags().StringVar(&opts.MonitorSubject, "monitor-subject", opts.MonitorSubject, "Base monitor subject to buffer")
	cmd.Flags().DurationVar(&opts.BufferRefresh, "buffer-refresh", opts.BufferRefresh, "Interval for refreshing buffer configuration")

	return cmd
}

func runServiceAll(cmd *cobra.Command, opts *serviceAllOptions) error {
	ctx, cancel := signal.NotifyContext(cmd.Context(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	var (
		srv       *natsutil.ServerHandle
		err       error
		serverURL string
	)

	if !opts.NoNATS {
		srvOpts := &server.Options{
			Host:      opts.Host,
			Port:      opts.Port,
			JetStream: true,
		}
		if opts.StoreDir != "" {
			err := os.MkdirAll(opts.StoreDir, 0o755)
			if err != nil {
				return err
			}
			srvOpts.StoreDir = opts.StoreDir
		}

		serverCfg := natsutil.ServerConfig{Options: srvOpts, Embedded: opts.Embedded}
		if opts.StoreDir == "" {
			serverCfg.TempDir = filepath.Join(os.TempDir(), "streams-service")
		}

		srv, err = natsutil.StartServer(serverCfg)
		if err != nil {
			return err
		}
		serverURL = srv.ClientURL()
		log.Info().Str("url", serverURL).Msg("nats server started")
		cmd.Printf("NATS server listening at %s\n", serverURL)
		defer srv.Shutdown()
	} else {
		serverURL = rootOpts.server
		log.Info().Str("url", serverURL).Msg("using external nats")
	}

	var js jetstream.JetStream
	if !opts.NoNATS && opts.Embedded {
		js, err = natsutil.ConnectJetStream(srv.ClientURL(), srv.InProcessOption())
	} else {
		js, err = natsutil.ConnectJetStream(serverURL)
	}
	if err != nil {
		return err
	}
	defer js.Conn().Drain()

	controllerOpts := controller.Options{
		ServerURL:      serverURL,
		CommandSubject: opts.CommandSubject,
		StateBucket:    opts.StateBucket,
	}

	group, groupCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		return controller.Run(groupCtx, js, controllerOpts)
	})
	group.Go(func() error {
		return buffer.RunBuffer(groupCtx, js, buffer.BufferOptions{
			DeviceBucket:    opts.DeviceBucket,
			MonitorSubject:  opts.MonitorSubject,
			RefreshInterval: opts.BufferRefresh,
		})
	})

	log.Info().Msg("services running (controller + buffer)")
	cmd.Printf("services running (controller subject=%s)\n", controllerOpts.CommandSubject)
	cmd.Printf("press Ctrl+C to stop\n")

	err = group.Wait()
	if err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	return nil
}
