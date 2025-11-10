package cli

import (
	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/session"
	"github.com/spf13/cobra"
)

func newStreamRemoveCmd() *cobra.Command {
	var sessionID string

	cmd := &cobra.Command{
		Use:     "rm",
		Short:   "remove a recorded stream session",
		Aliases: []string{"rm"},
		GroupID: "session",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return withSessionManager(cmd.Context(), config.SessionBucket, func(mgr *session.SessionStore) error {
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
	if err := cmd.MarkFlagRequired("session"); err != nil {
		cobra.CheckErr(err)
	}

	return cmd
}
