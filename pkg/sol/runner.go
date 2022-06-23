package sol

import (
	"apigear/pkg/gen"
	"apigear/pkg/idl"
	"apigear/pkg/log"
	"apigear/pkg/model"
	"apigear/pkg/spec"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

type runner struct {
	doc     spec.SolutionDoc
	rootDir string
}

func (r *runner) Run() {
	log.Info("run solution")
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
	log.Infof("process layer %s", layer.Name)
	var templateDir = path.Join(r.rootDir, layer.Template)
	var templatesDir = path.Join(templateDir, "templates")
	var rulesFile = path.Join(templateDir, "rules.yaml")
	var outputDir = path.Join(r.rootDir, layer.Output)
	var force = layer.Force
	name := layer.Name
	if name == "" {
		// if no layer name, name is the last part of the output directory
		name = path.Base(outputDir)
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
	log.Infof("parse inputs %v", inputs)
	idlParser := idl.NewParser(s)
	dataParser := model.NewDataParser(s)
	files, err := r.expandInputs(r.rootDir, inputs)
	if err != nil {
		log.Infof("error expanding inputs")
		return err
	}
	for _, file := range files {
		log.Infof("parse input %s", file)
		switch path.Ext(file) {
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
	s.ResolveAll()
	return nil
}

// ExpandInputs expands the input list to a list of files.
// If input entry is a file it is returned as a list.
// If input entry is a directory, all files in the directory are returned.

func (r *runner) expandInputs(rootDir string, inputs []string) ([]string, error) {
	var files []string
	for _, input := range inputs {
		entry := path.Join(rootDir, input)
		info, err := os.Stat(entry)
		if err != nil {
			log.Infof("error resolving input: %s", entry)
			continue
		}

		if info.IsDir() {
			filepath.WalkDir(entry, func(root string, d fs.DirEntry, err error) error {
				if err != nil {
					return fmt.Errorf("error resolving input: %w", err)
				}
				if d.IsDir() {
					return nil
				}
				files = append(files, root)
				return nil
			})
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
