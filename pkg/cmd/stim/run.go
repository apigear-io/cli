package stim

import (
	"fmt"
	"os"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/stim"
	"github.com/apigear-io/cli/pkg/stim/model"
	"github.com/spf13/cobra"
)

func NewRunCmd() *cobra.Command {
	var script string
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run a script",
		Long:  `Run a script`,
		RunE: func(cmd *cobra.Command, args []string) error {
			data, err := os.ReadFile(script)
			if err != nil {
				return err
			}
			man := stim.New()
			_, err = man.RunScript("", model.Script{
				Name:   "shell",
				Source: string(data),
			})
			if err != nil {
				log.Error().Err(err).Msg("run script")
				return err
			}
			fmt.Printf("press enter to exit\n")
			_, err = fmt.Scanln()
			if err != nil {
				log.Error().Err(err).Msg("read stdin")
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&script, "script", "f", "", "Execute a script")
	return cmd
}
