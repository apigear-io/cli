package cli

import (
	"fmt"
	"time"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/session"
	"github.com/spf13/cobra"
)

func newStreamListCmd() *cobra.Command {
	var deviceID string

	cmd := &cobra.Command{
		Use:     "ls",
		Short:   "list recorded stream sessions",
		Aliases: []string{"list"},
		GroupID: "session",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return withSessionManager(cmd.Context(), config.SessionBucket, func(mgr *session.SessionStore) error {
				metas, err := mgr.List()
				if err != nil {
					return err
				}

				// Filter by device if specified
				if deviceID != "" {
					filtered := make([]session.Metadata, 0)
					for _, meta := range metas {
						if meta.DeviceID == deviceID {
							filtered = append(filtered, meta)
						}
					}
					metas = filtered
				}

				if len(metas) == 0 {
					if deviceID != "" {
						cmd.Printf("no sessions found for device %s\n", deviceID)
					} else {
						cmd.Println("no sessions found")
					}
					return nil
				}

				if _, err := fmt.Fprintf(cmd.OutOrStdout(), "%-36s  %-12s  %-25s  %-25s  %-9s  %s\n",
					"SESSION", "DEVICE", "START", "END", "DURATION", "MESSAGES"); err != nil {
					return err
				}
				for _, meta := range metas {
					duration := meta.End.Sub(meta.Start)
					if _, err := fmt.Fprintf(cmd.OutOrStdout(), "%-36s  %-12s  %-25s  %-25s  %-9s  %d\n",
						meta.SessionID,
						meta.DeviceID,
						meta.Start.Format(time.RFC3339),
						meta.End.Format(time.RFC3339),
						duration.Round(time.Millisecond),
						meta.MessageCount,
					); err != nil {
						return err
					}
				}
				return nil
			})
		},
	}

	cmd.Flags().StringVar(&deviceID, "device", "", "Filter sessions by device identifier")

	return cmd
}
