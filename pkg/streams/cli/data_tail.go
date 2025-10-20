package cli

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/msgio"
	"github.com/spf13/cobra"
)

func newDataTailCmd() *cobra.Command {
	opts := &msgio.TailOptions{
		Subject: config.MonitorSubject,
	}

	cmd := &cobra.Command{
		Use:     "tail",
		Short:   "Tail a monitor subject for a given device ID",
		Aliases: []string{"follow", "watch"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
			defer cancel()

			opts.ServerURL = rootOpts.server
			opts.Verbose = rootOpts.verbose
			return msgio.Tail(ctx, *opts)
		},
	}

	cmd.Flags().StringVar(&opts.Subject, "subject", opts.Subject, "Base monitor subject name")
	cmd.Flags().StringVar(&opts.DeviceID, "device-id", "", "Device identifier to subscribe to")
	cmd.Flags().BoolVar(&opts.Pretty, "pretty", false, "Pretty print JSON payloads")
	cmd.Flags().BoolVar(&opts.Headers, "headers", false, "Print message headers")

	return cmd
}
