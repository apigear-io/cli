package cli

import (
	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/session"
	"github.com/spf13/cobra"
)

func newRecordingsDeleteCmd() *cobra.Command {
	var sessionID string
	bucket := config.SessionBucket

	cmd := &cobra.Command{
		Use:     "delete",
		Short:   "Delete a recorded session",
		Aliases: []string{"rm"},
		RunE: func(cmd *cobra.Command, _ []string) error {
			return withSessionManager(cmd.Context(), bucket, func(mgr *session.SessionStore) error {
				err := mgr.Delete(sessionID)
				if err != nil {
					return err
				}
				cmd.Printf("session %s deleted\n", sessionID)
				return nil
			})
		},
	}

	cmd.Flags().StringVar(&sessionID, "session-id", "", "Session identifier")
	cmd.Flags().StringVar(&bucket, "session-bucket", bucket, "Key-value bucket containing session metadata")
	if err := cmd.MarkFlagRequired("session-id"); err != nil {
		cobra.CheckErr(err)
	}

	return cmd
}
