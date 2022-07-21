package sol

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/config"
	"github.com/apigear-io/cli/pkg/gen"
	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/idl"
	"github.com/apigear-io/cli/pkg/model"
	"github.com/apigear-io/cli/pkg/spec"
)

type runner struct {
	doc     spec.SolutionDoc
	rootDir string
}

func (r *runner) Run() {
	log.Debug("run solution")
	for _, layer := range r.doc.Layers {
		err := r.processLayer(layer)
		if err != nil {
			log.Errorf("error processing layer: %s", err)
		}
	}
}

// processLayer processes a layer from the solution.
// A layer contains information about the inputs, used template and output.
func (r *runner) processLayer(layer spec.SolutionLayer) error {
	log.Debugf("process layer %s", layer.Name)
	// TODO: template can be a dir or a name of a template
	var templateDir string

	if helper.IsDir(filepath.Join(r.rootDir, layer.Template)) {
		templateDir = filepath.Join(r.rootDir, layer.Template)
	} else if helper.IsDir(filepath.Join(config.GetPackageDir(), layer.Template)) {
		templateDir = filepath.Join(config.GetPackageDir(), layer.Template)
	} else {
		return fmt.Errorf("template %s not found", layer.Template)
	}
	var templatesDir = filepath.Join(templateDir, "templates")
	var rulesFile = filepath.Join(templateDir, "rules.yaml")
	var outputDir = filepath.Join(r.rootDir, layer.Output)
	var force = layer.Force
	name := layer.Name
	if name == "" {
		// if no layer name, name is the last part of the output directory
		name = filepath.Base(outputDir)
	}
	system := model.NewSystem(name)
	err := r.parseInputs(system, layer.Inputs)
	if err != nil {
		return fmt.Errorf("error parsing inputs: %w", err)
	}

	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating output directory: %w", err)
	}

	generator, err := gen.New(outputDir, templatesDir, system, force)
	if err != nil {
		return fmt.Errorf("error creating generator: %w", err)
	}
	return generator.Run(rulesFile)
}

// parseInputs parses the inputs from the layer.
// A input can be either a file or a directory.
// If the input is a directory, the files in the directory will be parsed.
func (r *runner) parseInputs(s *model.System, inputs []string) error {
	log.Debugf("parse inputs %v", inputs)
	idlParser := idl.NewParser(s)
	dataParser := model.NewDataParser(s)
	files, err := r.expandInputs(r.rootDir, inputs)
	if err != nil {
		log.Infof("error expanding inputs")
		return err
	}
	for _, file := range files {
		log.Debugf("parse input %s", file)
		switch filepath.Ext(file) {
		case ".yaml", ".yml", ".json":
			err := dataParser.ParseFile(file)
			if err != nil {
				log.Warnf("error parsing data file: %s. skip", err)
			}
		case ".idl":
			err := idlParser.ParseFile(file)
			if err != nil {
				log.Warnf("error parsing idl file: %s. skip", err)
			}
		default:
			log.Warnf("unknown file type %s. skip", file)
		}
	}
	err = s.ResolveAll()
	if err != nil {
		return fmt.Errorf("error resolving system: %w", err)
	}
	return nil
}

// ExpandInputs expands the input list to a list of files.
// If input entry is a file it is returned as a list.
// If input entry is a directory, all files in the directory are returned.

func (r *runner) expandInputs(rootDir string, inputs []string) ([]string, error) {
	var files []string
	for _, input := range inputs {
		entry := filepath.Join(rootDir, input)
		info, err := os.Stat(entry)
		if err != nil {
			log.Infof("error resolving input: %s", entry)
			continue
		}

		if info.IsDir() {
			err := filepath.WalkDir(entry, func(root string, d fs.DirEntry, err error) error {
				if err != nil {
					return fmt.Errorf("error resolving input: %w", err)
				}
				if d.IsDir() {
					return nil
				}
				files = append(files, root)
				return nil
			})
			if err != nil {
				return nil, fmt.Errorf("error resolving input: %w", err)
			}
		} else {
			files = append(files, entry)
		}
	}
	return files, nil
}

func NewSolutionRunner(rootDir string, doc spec.SolutionDoc) *runner {
	return &runner{
		doc:     doc,
		rootDir: rootDir,
	}
}
