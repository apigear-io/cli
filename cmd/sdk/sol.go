package sdk

import (
	"objectapi/pkg/logger"
	"objectapi/pkg/sol"
	"path/filepath"

	"github.com/spf13/cobra"
)

var log = logger.Get()

type SolutionOptions struct {
	file string
}

func NewSolutionCommand() *cobra.Command {
	var options = &SolutionOptions{}
	var cmd = &cobra.Command{
		Use:     "sol [file to run]",
		Short:   "generate code using a solution",
		Aliases: []string{"solution", "s"},
		Long: `A solution is a yaml document which describes different layers. 
Each layer defines the input module files, output directory and the features to enable, 
as also the other options. To create a demo module or solution use the 'project create' command.`,
		Run: func(cmd *cobra.Command, args []string) {
			options.file = args[0]
			log.Info("run generator from solution ", options.file)
			doc, err := sol.ReadSolutionDoc(options.file)
			if err != nil {
				panic(err)
			}
			rootDir := filepath.Dir(options.file)
			proc := sol.NewSolutionRunner(rootDir, doc)
			proc.Run()
		},
	}
	return cmd
}
