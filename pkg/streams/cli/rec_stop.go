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

					// Identify unique parent sessions to stop
					// Multi-device sessions have format: {parent-session-id}-{device-id}
					// We want to stop the parent session, not individual device sessions
					parentSessions := make(map[string]bool)
					var sessionsToStop []string

					if deviceID != "" {
						// Stop all sessions for a specific device
						for _, state := range states {
							if state.Status == "running" && state.DeviceID == deviceID {
								sessionsToStop = append(sessionsToStop, state.SessionID)
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
						// Detect multi-device parent sessions by looking for device-specific sessions
						for _, state := range states {
							if state.Status != "running" {
								continue
							}
							// Check if this is a device-specific session (contains device ID in session ID)
							sessionID := state.SessionID
							// Device sessions have format: {parent}-{deviceID}
							// If we find such sessions, we want to stop the parent instead
							if lastDash := len(sessionID) - len(state.DeviceID) - 1; lastDash > 0 &&
								lastDash < len(sessionID) &&
								sessionID[lastDash] == '-' &&
								sessionID[lastDash+1:] == state.DeviceID {
								// This is a device-specific session from a multi-device recording
								parentSessionID := sessionID[:lastDash]
								if !parentSessions[parentSessionID] {
									parentSessions[parentSessionID] = true
									sessionsToStop = append(sessionsToStop, parentSessionID)
								}
							} else {
								// This is a standalone single-device session
								sessionsToStop = append(sessionsToStop, sessionID)
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
					for _, sessionID := range sessionsToStop {
						request := controller.RpcRequest{
							Action:    controller.ActionStop,
							SessionID: sessionID,
						}
						log.Debug().Str("session", sessionID).Msg("record stop request")

						resp, err := controller.SendCommand(ctx, nc, config.RecordRpcSubject, request)
						if err != nil {
							log.Error().Err(err).Str("session", sessionID).Msg("stop command failed")
							failedCount++
							continue
						}
						if !resp.OK {
							log.Error().Str("session", sessionID).Str("message", resp.Message).Msg("stop command failed")
							failedCount++
							continue
						}

						log.Debug().Str("session", resp.SessionID).Msg("recording stopped")
						cmd.Printf("stopped session=%s\n", resp.SessionID)
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
