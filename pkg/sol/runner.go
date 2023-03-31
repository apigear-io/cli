package sol

import (
	"context"
	"fmt"

	"github.com/apigear-io/cli/pkg/gen"
	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/model"
	"github.com/apigear-io/cli/pkg/spec"
	"github.com/apigear-io/cli/pkg/tasks"
)

type Runner struct {
	tm *tasks.TaskManager
	// tasks map[string]*task
}

func NewRunner() *Runner {
	return &Runner{
		// tasks: make(map[string]*task),
		tm: tasks.NewTaskManager(),
	}
}

func (r *Runner) HasTask(name string) bool {
	return r.tm.Has(name)
}

func (r *Runner) TaskFiles() []string {
	return r.tm.Names()
}

func (r *Runner) OnTask(fn func(*tasks.TaskEvent)) {
	r.tm.On(fn)
}

func (r *Runner) RunSource(ctx context.Context, source string) error {
	run := func(ctx context.Context) error {
		return RunSolutionSource(ctx, source)
	}
	meta := map[string]interface{}{
		"solution": source,
	}
	r.tm.Register(source, meta, run)
	return r.tm.Run(ctx, source)
}

// RunDoc runs the given file task once.
// TODO: Run should always act on a source of truth, such as a file.
// It should not act on a cached value.
func (r *Runner) RunDoc(ctx context.Context, file string, doc *spec.SolutionDoc) error {
	run := func(ctx context.Context) error {
		return runSolution(doc)
	}
	meta := map[string]interface{}{
		"solution": file,
	}
	r.tm.Register(file, meta, run)
	return r.tm.Run(ctx, file)
}

func (r *Runner) WatchSource(ctx context.Context, source string) error {
	doc, err := ReadSolutionDoc(source)
	if err != nil {
		return err
	}
	deps := doc.ComputeDependencies()
	deps = append(deps, source)
	run := func(ctx context.Context) error {
		return RunSolutionSource(ctx, source)
	}
	meta := map[string]interface{}{
		"solution": source,
	}
	r.tm.Register(source, meta, run)
	return r.tm.Watch(ctx, source, deps...)
}

// WatchDoc starts the watch of the given file task.
func (r *Runner) WatchDoc(ctx context.Context, file string, doc *spec.SolutionDoc) error {
	deps := doc.ComputeDependencies()
	deps = append(deps, file)
	run := func(ctx context.Context) error {
		return runSolution(doc)
	}
	meta := map[string]interface{}{
		"solution": file,
	}
	r.tm.Register(file, meta, run)
	return r.tm.Watch(ctx, file, deps...)
}

// StopWatch stops the watch of the given file task.
func (r *Runner) StopWatch(file string) {
	err := r.tm.Cancel(file)
	if err != nil {
		log.Error().Err(err).Msgf("failed to stop watch %s", file)
	}
}

func (r *Runner) Clear() {
	r.tm.CancelAll()
}

func RunSolutionSource(ctx context.Context, source string) error {
	doc, err := ReadSolutionDoc(source)
	if err != nil {
		return err
	}
	return runSolution(doc)
}

func runSolution(doc *spec.SolutionDoc) error {
	log.Info().Msgf("run solution %s", doc.RootDir)
	err := doc.Compute()
	if err != nil {
		return err
	}
	rootDir := doc.RootDir
	if rootDir == "" {
		return fmt.Errorf("root dir is empty")
	}
	for _, layer := range doc.Layers {
		name := layer.Name
		outDir := layer.GetOutputDir(rootDir)
		if name == "" {
			name = helper.BaseName(outDir)
		}

		tplDir := layer.GetTemplatesDir(rootDir)
		if tplDir == "" {
			return fmt.Errorf("templates dir is empty")
		}
		rulesFile := layer.GetRulesFile(rootDir)
		if rulesFile == "" {
			return fmt.Errorf("rules file is empty")
		}
		err := checkInputs(layer.ComputeExpandedInputs(rootDir))
		if err != nil {
			return err
		}
		system := model.NewSystem(name)
		err = parseInputs(system, layer.ComputeExpandedInputs(rootDir))
		if err != nil {
			return err
		}
		err = helper.MakeDir(outDir)
		if err != nil {
			return err
		}
		opts := gen.GeneratorOptions{
			OutputDir:    outDir,
			TemplatesDir: tplDir,
			System:       system,
			UserFeatures: layer.Features,
			UserForce:    layer.Force,
		}
		g, err := gen.New(opts)
		if err != nil {
			return err
		}
		doc, err := gen.ReadRulesDoc(rulesFile)
		if err != nil {
			return err
		}
		err = g.ProcessRules(doc)
		if err != nil {
			return err
		}
	}
	return nil
}
