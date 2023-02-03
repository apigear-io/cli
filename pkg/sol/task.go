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

type genConfig struct {
	name         string
	templatesDir string
	rulesFile    string
	outputDir    string
	inputs       []string
	features     []string
	force        bool
	meta         map[string]interface{}
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
	var cfg genConfig
	log.Debug().Msgf("process layer %s", layer.Name)
	rootDir := t.doc.RootDir

	td, err := resolveTemplateDir(rootDir, layer.Template)
	if err != nil {
		return err
	}

	cfg.templatesDir = helper.Join(td, "templates")
	cfg.rulesFile = helper.Join(td, "rules.yaml")
	// add templates dir and rules file as watcher dependency
	t.deps = append(t.deps, cfg.templatesDir, cfg.rulesFile)

	cfg.outputDir = helper.Join(rootDir, layer.Output)

	cfg.name = layer.Name
	if layer.Name == "" {
		// if no layer name, name is the last part of the output directory
		cfg.name = filepath.Base(cfg.outputDir)
	}

	if layer.Inputs == nil {
		return fmt.Errorf("inputs are empty")
	}
	cfg.inputs = layer.Inputs

	cfg.features = layer.Features
	if layer.Features == nil {
		cfg.features = []string{"all"}
	}
	cfg.force = layer.Force

	cfg.meta = helper.MergeMaps(t.doc.Meta, layer.Meta)

	return t.runGenerator(cfg)
}

func (t *task) runGenerator(cfg genConfig) error {
	log.Debug().Msgf("run generator %s %v", cfg.name, cfg.inputs)
	expanded, err := expandInputs(t.doc.RootDir, cfg.inputs)
	if err != nil {
		return err
	}

	err = checkInputs(expanded)
	if err != nil {
		return err
	}
	t.deps = append(t.deps, expanded...)

	system := model.NewSystem(cfg.name)

	system.Meta = cfg.meta

	err = parseInputs(system, expanded)
	if err != nil {
		return err
	}

	err = helper.MakeDir(cfg.outputDir)
	if err != nil {
		return fmt.Errorf("error creating output directory: %w", err)
	}

	opts := gen.GeneratorOptions{
		OutputDir:    cfg.outputDir,
		TemplatesDir: cfg.templatesDir,
		System:       system,
		UserFeatures: cfg.features,
		UserForce:    cfg.force,
	}
	generator, err := gen.New(opts)
	if err != nil {
		return err
	}
	doc, err := gen.ReadRulesDoc(cfg.rulesFile)
	if err != nil {
		return err
	}
	return generator.ProcessRules(doc)
}

func (t *task) Clear() {
	t.done <- true
}
