package sol

import (
	"context"

	"github.com/apigear-io/cli/pkg/cfg"
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
	task := func(ctx context.Context) error {
		return r.runSolutionFromSource(ctx, source, force)
	}
	meta := map[string]interface{}{
		"solution": source,
	}
	r.tm.Register(source, meta, task)
	return r.tm.Run(ctx, source)
}

// RunDoc runs the given file task once.
// It should not act on a cached value.
func (r *Runner) RunDoc(ctx context.Context, file string, doc *spec.SolutionDoc) error {
	task := func(ctx context.Context) error {
		return runSolution(doc)
	}
	meta := map[string]interface{}{
		"solution": file,
	}
	r.tm.Register(file, meta, task)
	return r.tm.Run(ctx, file)
}

func (r *Runner) WatchSource(ctx context.Context, source string, force bool) error {
	doc, err := ReadSolutionDoc(source)
	if err != nil {
		return err
	}
	deps := doc.AggregateDependencies()
	deps = append(deps, source)
	task := func(ctx context.Context) error {
		return r.runSolutionFromSource(ctx, source, force)
	}
	meta := map[string]interface{}{
		"solution": source,
	}
	r.tm.Register(source, meta, task)
	return r.tm.Watch(ctx, source, deps...)
}

// WatchDoc starts the watch of the given file task.
func (r *Runner) WatchDoc(ctx context.Context, file string, doc *spec.SolutionDoc) error {
	if err := doc.Validate(); err != nil {
		return err
	}
	deps := doc.AggregateDependencies()
	deps = append(deps, file)
	task := func(ctx context.Context) error {
		return runSolution(doc)
	}
	meta := map[string]interface{}{
		"solution": file,
	}
	r.tm.Register(file, meta, task)
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

func (r *Runner) runSolutionFromSource(_ context.Context, source string, force bool) error {
	doc, err := ReadSolutionDoc(source)
	if err != nil {
		return err
	}
	if force {
		for _, target := range doc.Targets {
			target.Force = true
		}
	}
	return runSolution(doc)
}

func runSolution(doc *spec.SolutionDoc) error {
	log.Debug().Msgf("run solution %s", doc.RootDir)
	if err := doc.Validate(); err != nil {
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
		doc.Meta["Layer"] = target
		doc.Meta["App"] = cfg.GetBuildInfo("cli")
		system.Meta = helper.JoinMaps(doc.Meta, target.Meta)
		if err := parseInputs(system, target.ExpandedInputs()); err != nil {
			return err
		}
		applyMetaDocument(target, system)
		if err := helper.MakeDir(outDir); err != nil {
			return err
		}
		opts := gen.Options{
			OutputDir:    outDir,
			TemplatesDir: target.TemplatesDir,
			System:       system,
			Features:     target.Features,
			Force:        target.Force,
			Meta:         helper.JoinMaps(doc.Meta, target.Meta),
		}
		g, err := gen.New(opts)
		if err != nil {
			return err
		}
		doc, err := gen.ReadRulesDoc(target.RulesFile)
		if err != nil {
			return err
		}
		bi := cfg.GetBuildInfo("cli")
		ok, errs := doc.CheckEngines(bi.Version)
		if !ok {
			log.Warn().Msgf("template requires cli version %s. Only found %s", doc.Engines.Cli, bi.Version)
			for _, err := range errs {
				log.Warn().Err(err).Msg("cli version mismatch")
			}
			return nil
		}
		// check keywords according to the rules languages
		system.CheckReservedWords(doc.Languages)
		err = g.ProcessRules(doc)
		if err != nil {
			return err
		}
	}
	return nil
}

func applyMetaDocument(t *spec.SolutionTarget, s *model.System) {
	for k, v := range t.MetaImports {
		log.Warn().Msgf("import %s %v", k, v)
		node := s.LookupNode(k)
		if node == nil {
			log.Warn().Msgf("node %s not found", k)
			continue
		}
		meta, ok := v.(map[string]interface{})
		if !ok {
			log.Warn().Msgf("meta for %s is not a map", k)
			continue
		}
		log.Info().Msgf("apply meta to node %s", k)
		node.Meta = helper.JoinMaps(node.Meta, meta)
	}
}
