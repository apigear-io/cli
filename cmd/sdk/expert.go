package sdk

import (
	"objectapi/pkg/log"
	"objectapi/pkg/sol"
	"objectapi/pkg/spec"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
)

type ExpertOptions struct {
	inputs      []string
	outputDir   string
	features    []string
	force       bool
	watch       bool
	templateDir string
}

func runExpert(options *ExpertOptions) error {
	log.Info("run generator from expert mode")
	doc := spec.SolutionDoc{
		Schema: "apigear.solution/1.0",
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
	log.Debugf("solution doc: %v", doc)
	rootDir, err := os.Getwd()
	log.Debugf("rootDir: %s", rootDir)
	if err != nil {
		log.Fatalf("failed to get current directory: %s", err)
	}
	proc := sol.NewSolutionRunner(rootDir, doc)
	proc.Run()
	return nil
}

func watchExpert(options *ExpertOptions) {
	runExpert(options)
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Infof("[%s] modified", event.Name)
					runExpert(options)
				}
			case err := <-watcher.Errors:
				log.Error(err)
			}
		}
	}()
	for _, input := range options.inputs {
		err = watcher.Add(input)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = watcher.Add(options.templateDir)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func NewExpertCommand() *cobra.Command {
	options := &ExpertOptions{}

	cmd := &cobra.Command{
		Use:     "x",
		Aliases: []string{"expert"},
		Short:   "generate code using expert mode",
		Long:    `In expert mode you can individually set your generator options. This is helpful when you do not have a solution document.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.Debugf("expert mode: %v", options)
			if options.watch {
				watchExpert(options)
			} else {
				err := runExpert(options)
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
	cmd.MarkFlagRequired("input")
	cmd.MarkFlagRequired("output")
	cmd.MarkFlagRequired("template")
	return cmd
}
