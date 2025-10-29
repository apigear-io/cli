package cli

import (
	"context"
	"errors"
	"time"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/controller"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type recordStartOptions struct {
	Subject       string
	DeviceID      string
	SessionID     string
	Retention     time.Duration
	SessionBucket string
	DeviceBucket  string
	DeviceDesc    string
	DeviceLoc     string
	DeviceOwner   string
	PreRoll       time.Duration
}

func newRecordingsStartCmd() *cobra.Command {
	opts := &recordStartOptions{
		Subject:       config.MonitorSubject,
		SessionBucket: config.SessionBucket,
		DeviceBucket:  config.DeviceBucket,
	}

	cmd := &cobra.Command{
		Use:     "start",
		Short:   "Start recording messages for a device",
		Aliases: []string{"begin"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return withSignalContext(cmd.Context(), func(ctx context.Context) error {
				return runRecordingStart(ctx, cmd, opts)
			})
		},
	}

	cmd.Flags().StringVar(&opts.Subject, "subject", opts.Subject, "Base subject to record from")
	cmd.Flags().StringVar(&opts.DeviceID, "device-id", "", "Device identifier to record")
	cmd.Flags().StringVar(&opts.SessionID, "session-id", "", "Optional session identifier (defaults to UUID)")
	cmd.Flags().DurationVar(&opts.Retention, "retention", 0, "Optional JetStream retention (e.g. 24h)")
	cmd.Flags().StringVar(&opts.SessionBucket, "session-bucket", opts.SessionBucket, "Key-value bucket for session metadata")
	cmd.Flags().StringVar(&opts.DeviceBucket, "device-bucket", opts.DeviceBucket, "Key-value bucket for device profiles")
	cmd.Flags().StringVar(&opts.DeviceDesc, "device-desc", "", "Optional device description")
	cmd.Flags().StringVar(&opts.DeviceLoc, "device-location", "", "Optional device location")
	cmd.Flags().StringVar(&opts.DeviceOwner, "device-owner", "", "Optional device owner")
	cmd.Flags().DurationVar(&opts.PreRoll, "pre-roll", 0, "Optional buffer window to include before start (e.g. 5m)")
	if err := cmd.MarkFlagRequired("device-id"); err != nil {
		cobra.CheckErr(err)
	}

	return cmd
}

func runRecordingStart(ctx context.Context, cmd *cobra.Command, opts *recordStartOptions) error {
	retention := ""
	if opts.Retention > 0 {
		retention = opts.Retention.String()
	}
	preRoll := ""
	if opts.PreRoll > 0 {
		preRoll = opts.PreRoll.String()
	}

	request := controller.RpcRequest{
		Action:        controller.ActionStart,
		Subject:       opts.Subject,
		DeviceID:      opts.DeviceID,
		SessionID:     opts.SessionID,
		Retention:     retention,
		SessionBucket: opts.SessionBucket,
		DeviceBucket:  opts.DeviceBucket,
		DeviceDesc:    opts.DeviceDesc,
		DeviceLoc:     opts.DeviceLoc,
		DeviceOwner:   opts.DeviceOwner,
		PreRoll:       preRoll,
		Verbose:       rootOpts.verbose,
	}

	return withNATS(ctx, func(nc *nats.Conn) error {
		log.Info().Str("device", opts.DeviceID).Str("subject", opts.Subject).Msg("record start request")

		resp, err := controller.SendCommand(ctx, nc, config.RecordRpcSubject, request)
		if err != nil {
			return err
		}
		if !resp.OK {
			if resp.Message == "" {
				return errors.New("record command failed")
			}
			return errors.New(resp.Message)
		}

		log.Info().Str("session", resp.SessionID).Str("device", opts.DeviceID).Msg("recording started")
		cmd.Printf("recording started session=%s\n", resp.SessionID)
		if rootOpts.verbose && resp.State != nil {
			cmd.Printf("state: %s (subject=%s device=%s messages=%d)\n",
				resp.State.Status, resp.State.Subject, resp.State.DeviceID, resp.State.MessageCount)
			if !resp.State.StartedAt.IsZero() {
				cmd.Printf("started: %s\n", resp.State.StartedAt.Format(time.RFC3339))
			}
		}
		return nil
	})
}
