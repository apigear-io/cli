package cli

import (
	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/session"
	"github.com/spf13/cobra"
)

func newStreamRemoveCmd() *cobra.Command {
	var sessionID string
	var purgeAll bool

	cmd := &cobra.Command{
		Use:     "rm",
		Short:   "remove a recorded stream session",
		Aliases: []string{"rm"},
		GroupID: "session",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return withSessionManager(cmd.Context(), config.SessionBucket, func(mgr *session.SessionStore) error {
				if purgeAll {
					// Delete all sessions
					sessions, err := mgr.List()
					if err != nil {
						return err
					}
					if len(sessions) == 0 {
						cmd.Println("no sessions to delete")
						return nil
					}

					deletedCount := 0
					failedCount := 0
					for _, meta := range sessions {
						err := mgr.Delete(meta.SessionID)
						if err != nil {
							cmd.Printf("failed to delete session %s: %v\n", meta.SessionID, err)
							failedCount++
							continue
						}
						cmd.Printf("deleted session %s\n", meta.SessionID)
						deletedCount++
					}

					cmd.Printf("\ndeleted %d session(s)", deletedCount)
					if failedCount > 0 {
						cmd.Printf(", %d failed", failedCount)
					}
					cmd.Println()
					return nil
				}

				// Delete single session
				err := mgr.Delete(sessionID)
				if err != nil {
					return err
				}
				cmd.Printf("session %s deleted\n", sessionID)
				return nil
			})
		},
	}

	cmd.Flags().StringVar(&sessionID, "session", "", "Session identifier")
	cmd.Flags().BoolVar(&purgeAll, "purge-all", false, "Delete all sessions")

	// Make session flag required only when purge-all is not set
	cmd.MarkFlagsOneRequired("session", "purge-all")
	cmd.MarkFlagsMutuallyExclusive("session", "purge-all")

	return cmd
}
