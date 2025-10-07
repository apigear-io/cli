package cli

import "github.com/spf13/cobra"

func newDataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "data",
		Short:   "Work with live monitor traffic and sample payloads",
		Aliases: []string{"d"},
	}

	cmd.AddCommand(
		newDataTailCmd(),
		newDataPublishCmd(),
		newDataGenerateCmd(),
	)

	return cmd
}
