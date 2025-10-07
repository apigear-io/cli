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

func newDataPublishCmd() *cobra.Command {
	opts := &msgio.PublishOptions{
		Subject:  config.MonitorSubject,
		MaxLine:  8 * 1024 * 1024,
		Validate: true,
		Headers:  map[string]string{},
	}

	cmd := &cobra.Command{
		Use:     "publish",
		Short:   "Publish JSONL messages to a NATS monitor subject",
		Aliases: []string{"send", "pub"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
			defer cancel()

			opts.ServerURL = rootOpts.server
			opts.Verbose = rootOpts.verbose
			return msgio.PublishFromFile(ctx, *opts)
		},
	}

	cmd.Flags().StringVarP(&opts.FilePath, "file", "f", "", "Path to JSONL file to publish")
	cmd.Flags().StringVar(&opts.Subject, "subject", opts.Subject, "Base monitor subject name")
	cmd.Flags().StringVar(&opts.DeviceID, "device-id", "", "Device identifier used to segment streams")
	cmd.Flags().DurationVar(&opts.Interval, "interval", opts.Interval, "Optional delay between published messages")
	cmd.Flags().IntVar(&opts.MaxLine, "max-line-bytes", opts.MaxLine, "Maximum size of a single JSON line in bytes")
	cmd.Flags().BoolVar(&opts.Validate, "validate", opts.Validate, "Validate that each line contains valid JSON before publishing")
	cmd.Flags().StringToStringVar(&opts.Headers, "header", opts.Headers, "Additional NATS headers to include in each message")
	cmd.Flags().BoolVar(&opts.Echo, "echo", false, "Print each published message to stdout")

	cmd.MarkFlagRequired("file")
	cmd.MarkFlagRequired("device-id")

	return cmd
}
