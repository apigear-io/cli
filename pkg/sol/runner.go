package sol

import (
	"context"

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

func (r *Runner) RunSource(ctx context.Context, source string, force bool) error {
	run := func(ctx context.Context) error {
		return RunSolutionSource(ctx, source, force)
	}
	meta := map[string]interface{}{
		"solution": source,
	}
	r.tm.Register(source, meta, run)
	return r.tm.Run(ctx, source)
}

// RunDoc runs the given file task once.
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

func (r *Runner) WatchSource(ctx context.Context, source string, force bool) error {
	doc, err := ReadSolutionDoc(source)
	if err != nil {
		return err
	}
	deps := doc.AggregateDependencies()
	deps = append(deps, source)
	run := func(ctx context.Context) error {
		return RunSolutionSource(ctx, source, force)
	}
	meta := map[string]interface{}{
		"solution": source,
	}
	r.tm.Register(source, meta, run)
	return r.tm.Watch(ctx, source, deps...)
}

// WatchDoc starts the watch of the given file task.
func (r *Runner) WatchDoc(ctx context.Context, file string, doc *spec.SolutionDoc) error {
	err := spec.LintSolutionFS(doc)
	if err != nil {
		return err
	}
	deps := doc.AggregateDependencies()
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
		log.Warn().Err(err).Msgf("stop watch %s", file)
	}
}

func (r *Runner) Clear() {
	r.tm.CancelAll()
}

func RunSolutionSource(ctx context.Context, source string, force bool) error {
	doc, err := ReadSolutionDoc(source)
	if err != nil {
		return err
	}
	if force {
		for _, layer := range doc.Targets {
			layer.Force = true
		}
	}
	return runSolution(doc)
}

func runSolution(doc *spec.SolutionDoc) error {
	log.Debug().Msgf("run solution %s", doc.RootDir)
	err := spec.LintSolutionFS(doc)
	if err != nil {
		return err
	}
	rootDir := doc.RootDir

	for _, target := range doc.Targets {
		name := target.Name
		outDir := target.GetOutputDir(rootDir)
		if name == "" {
			name = helper.BaseName(outDir)
		}
		system := model.NewSystem(name)
		err = parseInputs(system, target.ExpandedInputs())
		if err != nil {
			return err
		}
		err = helper.MakeDir(outDir)
		if err != nil {
			return err
		}
		opts := gen.GeneratorOptions{
			OutputDir:    outDir,
			TemplatesDir: target.TemplatesDir,
			System:       system,
			UserFeatures: target.Features,
			UserForce:    target.Force,
			Meta:         doc.Meta,
		}
		g, err := gen.New(opts)
		if err != nil {
			return err
		}
		doc, err := gen.ReadRulesDoc(target.RulesFile)
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
