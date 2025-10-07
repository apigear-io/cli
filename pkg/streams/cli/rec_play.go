package cli

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/session"
	"github.com/spf13/cobra"
)

func newRecordingsPlayCmd() *cobra.Command {
	opts := &session.PlaybackOptions{
		Bucket: config.SessionBucket,
		Speed:  1,
	}

	cmd := &cobra.Command{
		Use:     "play",
		Short:   "Replay a recorded session",
		Aliases: []string{"replay"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
			defer cancel()

			opts.ServerURL = rootOpts.server
			opts.Verbose = rootOpts.verbose
			return session.Playback(ctx, *opts)
		},
	}

	cmd.Flags().StringVar(&opts.SessionID, "session-id", "", "Session identifier to replay")
	cmd.Flags().StringVar(&opts.TargetSubject, "target-subject", "", "Optional override subject to publish during playback")
	cmd.Flags().Float64Var(&opts.Speed, "speed", opts.Speed, "Playback speed multiplier (e.g. 0.25, 1, 5)")
	cmd.Flags().StringVar(&opts.Bucket, "session-bucket", opts.Bucket, "Key-value bucket containing session metadata")
	cmd.MarkFlagRequired("session-id")

	return cmd
}
