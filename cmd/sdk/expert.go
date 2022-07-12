package sdk

import (
	"apigear/pkg/log"
	"apigear/pkg/sol"
	"apigear/pkg/spec"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
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
	Must(cmd.MarkFlagRequired("input"))
	Must(cmd.MarkFlagRequired("output"))
	Must(cmd.MarkFlagRequired("template"))
	return cmd
}

func runExpert(options *ExpertOptions) error {
	log.Debug("run expert code generation")
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
	err := runExpert(options)
	if err != nil {
		log.Errorf("failed to run expert mode: %s", err)
	}
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
					err := runExpert(options)
					if err != nil {
						log.Errorf("failed to run expert mode: %s", err)
					}
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
	// add directories of template dir recursively
	err = filepath.Walk(options.templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			err = watcher.Add(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Fatalf("failed to watch template directory: %s", err)
	}
	<-done
}
