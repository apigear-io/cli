package gen

import (
	"sync"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/sol"

	"github.com/spf13/cobra"
)

func NewSolutionCommand() *cobra.Command {
	var file string
	var watch bool
	var exec string
	var cmd = &cobra.Command{
		Use:     "solution [solution-file]",
		Short:   "Generate SDK using a solution document",
		Aliases: []string{"sol", "s"},
		Args:    cobra.ExactArgs(1),
		Long: `A solution is a yaml document which describes different layers. 
Each layer defines the input module files, output directory and the features to enable, 
as also the other options. To create a demo module or solution use the 'project create' command.`,
		Run: func(cmd *cobra.Command, args []string) {
			file = args[0]
			doc, err := sol.ReadSolutionDoc(file)
			if err != nil {
				log.Fatal().Err(err).Msg("read solution document")
			}
			runner := sol.NewRunner()
			if watch {
				var wg = sync.WaitGroup{}
				wg.Add(1)
				done, err := runner.StartWatch(file, doc)
				if err != nil {
					log.Fatal().Err(err).Msg("start watch")
				}
				wg.Wait()
				done <- true
			} else {
				err := runner.RunDoc(file, doc)
				if err != nil {
					log.Fatal().Err(err).Msgf("unable to run solution document %s", file)
				}
			}
		},
	}
	cmd.Flags().BoolVarP(&watch, "watch", "", false, "watch solution file for changes")
	cmd.Flags().StringVarP(&exec, "exec", "", "", "execute a command after generation")
	return cmd
}
