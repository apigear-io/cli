package cli

import (
	"fmt"
	"time"

	"github.com/apigear-io/cli/pkg/streams/config"
	"github.com/apigear-io/cli/pkg/streams/session"
	"github.com/spf13/cobra"
)

func newStreamShowCmd() *cobra.Command {
	var sessionID string

	cmd := &cobra.Command{
		Use:     "show",
		Short:   "show stream session details",
		Aliases: []string{"info"},
		GroupID: "session",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return withSessionManager(cmd.Context(), config.SessionBucket, func(mgr *session.SessionStore) error {
				meta, err := mgr.Info(sessionID)
				if err != nil {
					return err
				}

				duration := meta.End.Sub(meta.Start)
				if _, err := fmt.Fprintf(cmd.OutOrStdout(), "session:   %s\n", meta.SessionID); err != nil {
					return err
				}
				if _, err := fmt.Fprintf(cmd.OutOrStdout(), "device:    %s\n", meta.DeviceID); err != nil {
					return err
				}
				if _, err := fmt.Fprintf(cmd.OutOrStdout(), "stream:    %s\n", meta.Stream); err != nil {
					return err
				}
				if _, err := fmt.Fprintf(cmd.OutOrStdout(), "subject:   %s\n", meta.SourceSubject); err != nil {
					return err
				}
				if _, err := fmt.Fprintf(cmd.OutOrStdout(), "start:     %s\n", meta.Start.Format(time.RFC3339)); err != nil {
					return err
				}
				if _, err := fmt.Fprintf(cmd.OutOrStdout(), "end:       %s\n", meta.End.Format(time.RFC3339)); err != nil {
					return err
				}
				if _, err := fmt.Fprintf(cmd.OutOrStdout(), "duration:  %s\n", duration.Round(time.Millisecond)); err != nil {
					return err
				}
				if _, err := fmt.Fprintf(cmd.OutOrStdout(), "messages:  %d\n", meta.MessageCount); err != nil {
					return err
				}
				if meta.Retention != "" {
					if _, err := fmt.Fprintf(cmd.OutOrStdout(), "retention: %s\n", meta.Retention); err != nil {
						return err
					}
				}
				if meta.Note != "" {
					if _, err := fmt.Fprintf(cmd.OutOrStdout(), "note:      %s\n", meta.Note); err != nil {
						return err
					}
				}
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
