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

func newStreamStopCmd() *cobra.Command {
	var sessionID string

	cmd := &cobra.Command{
		Use:     "stop",
		Short:   "stop an active stream recording",
		Aliases: []string{"end"},
		GroupID: "record",
		RunE: func(cmd *cobra.Command, _ []string) error {
			if sessionID == "" {
				return errors.New("session cannot be empty")
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

	cmd.Flags().StringVar(&sessionID, "session", "", "Session identifier")
	if err := cmd.MarkFlagRequired("session"); err != nil {
		cobra.CheckErr(err)
	}
	return cmd
}
