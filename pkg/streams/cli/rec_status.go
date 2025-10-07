package cli

import (
	"context"
	"errors"
	"time"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/controller"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func newRecordingsStatusCmd() *cobra.Command {
	var sessionID string

	cmd := &cobra.Command{
		Use:     "status",
		Short:   "Show the latest controller state",
		Aliases: []string{"state"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			if sessionID == "" {
				return errors.New("session-id cannot be empty")
			}
			return withSignalContext(cmd.Context(), func(ctx context.Context) error {
				return withJetStream(ctx, func(js jetstream.JetStream) error {
					state, err := controller.FetchState(js, config.StateBucket, sessionID)
					if err != nil {
						return err
					}

					log.Debug().Str("session", state.SessionID).Str("status", state.Status).Int("messages", state.MessageCount).Msg("record status")

					cmd.Printf("session: %s\n", state.SessionID)
					cmd.Printf("status:  %s\n", state.Status)
					cmd.Printf("device:  %s\n", state.DeviceID)
					cmd.Printf("subject: %s\n", state.Subject)
					cmd.Printf("messages:%d\n", state.MessageCount)
					if !state.StartedAt.IsZero() {
						cmd.Printf("started: %s\n", state.StartedAt.Format(time.RFC3339))
					}
					if !state.LastMessageAt.IsZero() {
						cmd.Printf("last-message: %s\n", state.LastMessageAt.Format(time.RFC3339))
					}
					if state.LastError != "" {
						cmd.Printf("error:   %s\n", state.LastError)
					}
					cmd.Printf("updated: %s\n", state.UpdatedAt.Format(time.RFC3339))
					return nil
				})
			})
		},
	}

	cmd.Flags().StringVar(&sessionID, "session-id", "", "Session identifier")
	cmd.MarkFlagRequired("session-id")
	return cmd
}
