package solution

import (
	"context"
	"time"

	"github.com/apigear-io/cli/pkg/codegen"
	"github.com/apigear-io/cli/pkg/foundation"
	"github.com/apigear-io/cli/pkg/foundation/config"
	"github.com/apigear-io/cli/pkg/foundation/tasks"
	"github.com/apigear-io/cli/pkg/objmodel"
	"github.com/apigear-io/cli/pkg/objmodel/spec"
)

// RunStats accumulates code generation statistics across all targets.
type RunStats struct {
	FilesWritten int      `json:"filesWritten"`
	FilesSkipped int      `json:"filesSkipped"`
	FilesCopied  int      `json:"filesCopied"`
	TotalFiles   int      `json:"totalFiles"`
	TargetCount  int      `json:"targetCount"`
	DurationMs   int64    `json:"durationMs"`
}

type Runner struct {
	tm    *tasks.TaskManager
	Stats RunStats
}

func NewRunner() *Runner {
	return &Runner{
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
	r.tm.AddHook(fn)
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
		return r.runSolution(doc)
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
		return r.runSolution(doc)
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
	return r.runSolution(doc)
}

func (r *Runner) runSolution(doc *spec.SolutionDoc) error {
	log.Info().Msgf("run solution %s", doc.RootDir)
	if err := doc.Validate(); err != nil {
		return err
	}
	rootDir := doc.RootDir

	start := time.Now()
	r.Stats = RunStats{}

	for _, target := range doc.Targets {
		name := target.Name
		outDir := target.GetOutputDir(rootDir)
		if name == "" {
			name = foundation.BaseName(outDir)
		}
		system := objmodel.NewSystem(name)
		doc.Meta["Layer"] = target
		doc.Meta["App"] = config.GetBuildInfo("cli")
		system.Meta = foundation.JoinMaps(doc.Meta, target.Meta)
		if err := parseInputs(system, target.ExpandedInputs()); err != nil {
			return err
		}
		applyMetaDocument(target, system)
		if err := foundation.MakeDir(outDir); err != nil {
			return err
		}
		opts := codegen.Options{
			OutputDir:    outDir,
			TemplatesDir: target.TemplatesDir,
			System:       system,
			Features:     target.Features,
			Force:        target.Force,
			Meta:         foundation.JoinMaps(doc.Meta, target.Meta),
		}
		g, err := codegen.New(opts)
		if err != nil {
			return err
		}
		doc, err := codegen.ReadRulesDoc(target.RulesFile)
		if err != nil {
			return err
		}
		bi := config.GetBuildInfo("cli")
		ok, errs := doc.CheckEngines(bi.Version)
		if !ok {
			// a warning should be enough
			log.Warn().Msgf("template requires cli version %s. Only found %s", doc.Engines.Cli, bi.Version)
			for _, err := range errs {
				log.Warn().Err(err).Msg("cli version mismatch")
			}
		}
		// check keywords according to the rules languages
		system.CheckReservedWords(doc.Languages)
		err = g.ProcessRules(doc)
		if err != nil {
			return err
		}
		// Accumulate stats from this target
		r.Stats.FilesWritten += g.Stats.FilesWritten
		r.Stats.FilesSkipped += g.Stats.FilesSkipped
		r.Stats.FilesCopied += g.Stats.FilesCopied
		r.Stats.TargetCount++
	}

	r.Stats.TotalFiles = r.Stats.FilesWritten + r.Stats.FilesSkipped + r.Stats.FilesCopied
	r.Stats.DurationMs = time.Since(start).Milliseconds()

	return nil
}

func applyMetaDocument(t *spec.SolutionTarget, s *objmodel.System) {
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
		node.Meta = foundation.JoinMaps(node.Meta, meta)
	}
}
