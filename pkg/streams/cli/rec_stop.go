package cli

import (
	"context"
	"errors"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/controller"
	"github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func newRecordingsStopCmd() *cobra.Command {
	var sessionID string

	cmd := &cobra.Command{
		Use:     "stop",
		Short:   "Stop an active recording",
		Aliases: []string{"end"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			if sessionID == "" {
				return errors.New("session-id cannot be empty")
			}

			return withSignalContext(cmd.Context(), func(ctx context.Context) error {
				return withNATS(ctx, func(nc *nats.Conn) error {
					request := controller.RpcRequest{
						Action:    controller.ActionStop,
						SessionID: sessionID,
					}
					log.Info().Str("session", sessionID).Msg("record stop request")

					resp, err := controller.SendCommand(ctx, nc, config.RecordRpcSubject, request)
					if err != nil {
						return err
					}
					if !resp.OK {
						if resp.Message == "" {
							return errors.New("stop command failed")
						}
						return errors.New(resp.Message)
					}

					log.Info().Str("session", resp.SessionID).Msg("recording stopped")
					cmd.Printf("recording stopped session=%s\n", resp.SessionID)
					return nil
				})
			})
		},
	}

	cmd.Flags().StringVar(&sessionID, "session-id", "", "Session identifier")
	if err := cmd.MarkFlagRequired("session-id"); err != nil {
		cobra.CheckErr(err)
	}
	return cmd
}
