package cli

import (
	"fmt"
	"time"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/session"
	"github.com/spf13/cobra"
)

func newStreamListCmd() *cobra.Command {
	bucket := config.SessionBucket

	cmd := &cobra.Command{
		Use:     "ls",
		Short:   "list recorded stream sessions",
		Aliases: []string{"list"},
		GroupID: "session",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return withSessionManager(cmd.Context(), bucket, func(mgr *session.SessionStore) error {
				metas, err := mgr.List()
				if err != nil {
					return err
				}

				if len(metas) == 0 {
					cmd.Println("no sessions found")
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

	cmd.Flags().StringVar(&bucket, "session-bucket", bucket, "Key-value bucket containing session metadata")
	return cmd
}
