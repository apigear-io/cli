package sol

import (
	"fmt"
	"path/filepath"
	"sync"

	"github.com/apigear-io/cli/pkg/gen"
	"github.com/apigear-io/cli/pkg/helper"
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
	sync.Mutex
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
	t.Lock()
	defer t.Unlock()
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
func (t *task) processLayer(layer *spec.SolutionLayer) error {
	log.Debug().Msgf("process layer %s", layer.Name)
	rootDir := t.doc.RootDir
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
	features := layer.Features
	if features == nil {
		features = []string{"all"}
	}

	return t.runGenerator(name, layer.Inputs, outputDir, templateDir, features, force)
}

func (t *task) runGenerator(name string, inputs []string, outputDir string, templateDir string, features []string, force bool) error {
	log.Debug().Msgf("run generator %s %v", name, inputs)
	var templatesDir = helper.Join(templateDir, "templates")
	var rulesFile = helper.Join(templateDir, "rules.yaml")
	expanded, err := expandInputs(t.doc.RootDir, inputs)
	if err != nil {
		return err
	}

	err = checkInputs(expanded)
	if err != nil {
		return err
	}
	t.deps = append(t.deps, expanded...)

	system := model.NewSystem(name)
	err = parseInputs(system, expanded)
	if err != nil {
		return err
	}

	err = helper.MakeDir(outputDir)
	if err != nil {
		return fmt.Errorf("error creating output directory: %w", err)
	}

	generator, err := gen.New(outputDir, templatesDir, system, features, force)
	if err != nil {
		return err
	}
	doc, err := gen.ReadRulesDoc(rulesFile)
	if err != nil {
		return err
	}
	return generator.ProcessRules(doc)
}

func (t *task) Clear() {
	t.done <- true
}
