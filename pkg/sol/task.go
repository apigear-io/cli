package sol

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/gen"
	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/model"
	"github.com/apigear-io/cli/pkg/spec"
	"github.com/fsnotify/fsnotify"
)

type task struct {
	file    string
	doc     *spec.SolutionDoc
	deps    []string
	watcher *fsnotify.Watcher
	done    chan bool
}

func newTask(file string, doc *spec.SolutionDoc) (*task, error) {
	t := &task{
		file: file,
		doc:  doc,
		deps: make([]string, 0),
		done: make(chan bool),
	}
	if doc == nil {
		return nil, fmt.Errorf("doc is nil")
	}
	if doc.RootDir == "" {
		return nil, fmt.Errorf("root dir is empty")
	}
	t.addDep(file)
	return t, nil
}

func (t *task) addDep(dep string) {
	t.deps = append(t.deps, dep)
}

func (t *task) run() error {
	log.Debug().Msgf("run task %s", t.file)
	// reset deps
	t.deps = make([]string, 0)
	for _, layer := range t.doc.Layers {
		err := t.processLayer(layer)
		if err != nil {
			return err
		}
	}
	return nil
}

// processLayer processes a layer from the solution.
// A layer contains information about the inputs, used template and output.
func (t *task) processLayer(layer spec.SolutionLayer) error {
	rootDir := t.doc.RootDir
	log.Debug().Msgf("process layer %s", layer.Name)
	// TODO: template can be a dir or a name of a template
	var templateDir string
	td, err := GetTemplateDir(rootDir, layer.Template)
	if err != nil {
		return err
	}
	templateDir = td
	var templatesDir = helper.Join(templateDir, "templates")
	var rulesFile = helper.Join(templateDir, "rules.yaml")
	var outputDir = helper.Join(rootDir, layer.Output)
	// add templates dir and rules file as dependency
	t.deps = append(t.deps, templatesDir, rulesFile)
	var force = layer.Force
	name := layer.Name
	if name == "" {
		// if no layer name, name is the last part of the output directory
		name = filepath.Base(outputDir)
	}
	if layer.Inputs == nil {
		return fmt.Errorf("inputs are empty")
	}
	return t.runGenerator(name, layer.Inputs, outputDir, templateDir, force)
}

func (t *task) runGenerator(name string, inputs []string, outputDir string, templateDir string, force bool) error {
	var templatesDir = helper.Join(templateDir, "templates")
	var rulesFile = helper.Join(templateDir, "rules.yaml")
	system := model.NewSystem(name)
	err := t.parseInputs(system, inputs)
	if err != nil {
		return fmt.Errorf("parsing inputs: %w", err)
	}

	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating output directory: %w", err)
	}

	generator, err := gen.New(outputDir, templatesDir, system, force)
	if err != nil {
		return fmt.Errorf("error creating generator: %w", err)
	}
	doc, err := gen.ReadRulesDoc(rulesFile)
	if err != nil {
		return fmt.Errorf("error reading rules file: %w", err)
	}
	return generator.ProcessRules(doc)
}

func (t *task) Clear() {
	t.done <- true
}
