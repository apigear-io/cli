package sdk

import (
	"os"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/sol"
	"github.com/apigear-io/cli/pkg/spec"

	"github.com/spf13/cobra"
)

func Must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type ExpertOptions struct {
	inputs      []string
	outputDir   string
	features    []string
	force       bool
	watch       bool
	templateDir string
}

func NewExpertCommand() *cobra.Command {
	options := &ExpertOptions{}

	cmd := &cobra.Command{
		Use:     "expert",
		Aliases: []string{"x"},
		Short:   "generate code using expert mode",
		Long:    `In expert mode you can individually set your generator options. This is helpful when you do not have a solution document.`,
		Run: func(cmd *cobra.Command, args []string) {
			doc := makeSolution(options)

			if options.watch {
				sol.WatchSolution(doc.RootDir)
				watchExpert(options)
			} else {
				_, err := sol.RunSolutionDocument(doc)
				if err != nil {
					log.Fatalf("failed to run expert mode: %s", err)
				}
			}
		},
	}
	cmd.Flags().StringVarP(&options.templateDir, "template", "t", "tpl", "template directory")
	cmd.Flags().StringSliceVarP(&options.inputs, "input", "i", []string{"api"}, "input files")
	cmd.Flags().StringVarP(&options.outputDir, "output", "o", "out", "output directory")
	cmd.Flags().StringSliceVarP(&options.features, "feature", "f", []string{"core"}, "features to enable")
	cmd.Flags().BoolVarP(&options.force, "force", "", false, "force overwrite")
	cmd.Flags().BoolVarP(&options.watch, "watch", "", false, "watch for changes")
	Must(cmd.MarkFlagRequired("input"))
	Must(cmd.MarkFlagRequired("output"))
	Must(cmd.MarkFlagRequired("template"))
	return cmd
}

func makeSolution(options *ExpertOptions) *spec.SolutionDoc {
	rootDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return &spec.SolutionDoc{
		Schema:  "apigear.solution/1.0",
		RootDir: rootDir,
		Layers: []spec.SolutionLayer{
			{
				Inputs:   options.inputs,
				Output:   options.outputDir,
				Template: options.templateDir,
				Features: options.features,
				Force:    options.force,
			},
		},
	}
}

// func runExpert(options *ExpertOptions) error {
// 	log.Debug("run expert code generation")
// 	rootDir, err := os.Getwd()
// 	if err != nil {
// 		return err
// 	}
// 	doc := spec.SolutionDoc{
// 		Schema:  "apigear.solution/1.0",
// 		RootDir: rootDir,
// 		Layers: []spec.SolutionLayer{
// 			{
// 				Inputs:   options.inputs,
// 				Output:   options.outputDir,
// 				Template: options.templateDir,
// 				Features: options.features,
// 				Force:    options.force,
// 			},
// 		},
// 	}
// 	return doc
// }

// func watchExpert(options *ExpertOptions) {
// 	err := runExpert(options)
// 	if err != nil {
// 		log.Errorf("failed to run expert mode: %s", err)
// 	}
// 	watcher, err := fsnotify.NewWatcher()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer watcher.Close()
// 	done := make(chan bool)
// 	go func() {
// 		for {
// 			select {
// 			case event := <-watcher.Events:
// 				if event.Op&fsnotify.Write == fsnotify.Write {
// 					log.Infof("[%s] modified", event.Name)
// 					err := runExpert(options)
// 					if err != nil {
// 						log.Errorf("failed to run expert mode: %s", err)
// 					}
// 				}
// 			case err := <-watcher.Errors:
// 				log.Error(err)
// 			}
// 		}
// 	}()
// 	for _, input := range options.inputs {
// 		err = watcher.Add(input)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// 	// add directories of template dir recursively
// 	err = filepath.Walk(options.templateDir, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}
// 		if info.IsDir() {
// 			err = watcher.Add(path)
// 			if err != nil {
// 				return err
// 			}
// 		}
// 		return nil
// 	})
// 	if err != nil {
// 		log.Fatalf("failed to watch template directory: %s", err)
// 	}
// 	<-done
// }
