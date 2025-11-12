package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/controller"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func newStreamStopCmd() *cobra.Command {
	var sessionID string
	var deviceID string

	cmd := &cobra.Command{
		Use:     "stop",
		Short:   "stop active stream recording(s)",
		Long:    "Stop one or more active recordings. Use --session to stop a specific session, --device to stop all sessions for a device, or omit both to stop all active recordings.",
		Aliases: []string{"end"},
		GroupID: "record",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return withSignalContext(cmd.Context(), func(ctx context.Context) error {
				// Case 1: Stop specific session
				if sessionID != "" {
					return stopSession(ctx, cmd, sessionID)
				}

				// Case 2 & 3: Stop by device or all sessions
				return withJetStream(ctx, func(js jetstream.JetStream) error {
					nc := js.Conn()

					// Get all active sessions
					states, err := controller.ListStates(js, config.StateBucket)
					if err != nil {
						return fmt.Errorf("list states: %w", err)
					}

					// Filter by device if specified
					var sessionsToStop []controller.StateSnapshot
					if deviceID != "" {
						for _, state := range states {
							if state.Status == "running" && state.DeviceID == deviceID {
								sessionsToStop = append(sessionsToStop, state)
							}
						}
						if len(sessionsToStop) == 0 {
							cmd.Printf("no active recordings found for device %s\n", deviceID)
							return nil
						}
						cmd.Printf("searching for sessions for device %s\n", deviceID)
						cmd.Printf("found %d active session(s)\n", len(sessionsToStop))
					} else {
						// Stop all running sessions
						for _, state := range states {
							if state.Status == "running" {
								sessionsToStop = append(sessionsToStop, state)
							}
						}
						if len(sessionsToStop) == 0 {
							cmd.Println("no active recordings found")
							return nil
						}
						cmd.Printf("found %d active session(s)\n", len(sessionsToStop))
					}

					// Stop each session
					stoppedCount := 0
					failedCount := 0
					for _, state := range sessionsToStop {
						request := controller.RpcRequest{
							Action:    controller.ActionStop,
							SessionID: state.SessionID,
						}
						log.Debug().Str("session", state.SessionID).Str("device", state.DeviceID).Msg("record stop request")

						resp, err := controller.SendCommand(ctx, nc, config.RecordRpcSubject, request)
						if err != nil {
							log.Error().Err(err).Str("session", state.SessionID).Msg("stop command failed")
							failedCount++
							continue
						}
						if !resp.OK {
							log.Error().Str("session", state.SessionID).Str("message", resp.Message).Msg("stop command failed")
							failedCount++
							continue
						}

						log.Debug().Str("session", resp.SessionID).Msg("recording stopped")
						cmd.Printf("stopped session=%s device=%s\n", resp.SessionID, state.DeviceID)
						stoppedCount++
					}

					cmd.Printf("\nstopped %d session(s)", stoppedCount)
					if failedCount > 0 {
						cmd.Printf(", %d failed", failedCount)
					}
					cmd.Println()
					return nil
				})
			})
		},
	}

	cmd.Flags().StringVar(&sessionID, "session", "", "Session identifier to stop")
	cmd.Flags().StringVar(&deviceID, "device", "", "Stop all sessions for this device")
	return cmd
}

func stopSession(ctx context.Context, cmd *cobra.Command, sessionID string) error {
	return withNATS(ctx, func(nc *nats.Conn) error {
		request := controller.RpcRequest{
			Action:    controller.ActionStop,
			SessionID: sessionID,
		}
		log.Debug().Str("session", sessionID).Msg("record stop request")

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

		log.Debug().Str("session", resp.SessionID).Msg("recording stopped")
		cmd.Printf("recording stopped session=%s\n", resp.SessionID)
		return nil
	})
}
