package gen

import (
	"objectapi/pkg/logger"
	"objectapi/pkg/sol"
	"path/filepath"

	"github.com/spf13/cobra"
)

var log = logger.Get()

type GenSolutionOptions struct {
	file string
}

func NewGenSolution() *cobra.Command {
	o := GenSolutionOptions{}
	// genSolCmd represents the genRun command
	var cmd = &cobra.Command{
		Use:     "sol [file to run]",
		Short:   "generate code using a solution",
		Aliases: []string{"solution", "s"},
		Long: `A solution is a yaml document which describes different layers. 
Each layer defines the input module files, output directory and the features to enable, 
as also the other options. To create a demo module or solution use the 'project create' command.`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			o.file = args[0]
			log.Info("run generator from solution ", o.file)
			doc, err := sol.ReadSolutionDoc(o.file)
			if err != nil {
				panic(err)
			}
			rootDir := filepath.Dir(o.file)
			proc := sol.NewSolutionRunner(rootDir, doc)
			proc.Run()
		},
	}

	return cmd
}
