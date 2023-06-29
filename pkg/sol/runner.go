package sol

import (
	"context"
	"fmt"

	"github.com/apigear-io/cli/pkg/gen"
	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/model"
	"github.com/apigear-io/cli/pkg/repos"
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

func (r *Runner) WatchSource(ctx context.Context, source string, force bool) error {
	doc, err := ReadSolutionDoc(source)
	if err != nil {
		return err
	}
	deps := doc.ComputeDependencies()
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

func RunSolutionSource(ctx context.Context, source string, force bool) error {
	doc, err := ReadSolutionDoc(source)
	if err != nil {
		return err
	}
	if force {
		for _, layer := range doc.Layers {
			layer.Force = true
		}
	}
	return runSolution(doc)
}

func runSolution(doc *spec.SolutionDoc) error {
	log.Debug().Msgf("run solution %s", doc.RootDir)
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
		// check for local template
		tplDir := helper.Join(rootDir, layer.Template)
		if helper.IsDir(tplDir) {
			log.Info().Msgf("using local template %s", tplDir)
		} else {
			log.Info().Msgf("try to detect registered template %s", layer.Template)
			repoId, err := repos.GetOrInstallTemplateFromRepoID(layer.Template)
			if err != nil {
				return err
			}
			tplDir, err = repos.Cache.GetTemplateDir(repoId)
			if err != nil {
				return fmt.Errorf("can't find template %s", layer.Template)
			}
			// update template id based on the resolved repo id
			layer.Template = repoId
			log.Debug().Msgf("using registered template %s", tplDir)
		}
		if tplDir == "" {
			// we don't have a template
			return fmt.Errorf("template is neither local nor registry template: %s", layer.Template)
		}
		log.Info().Msgf("using template from: %s", tplDir)
		rulesFile := helper.Join(tplDir, "rules.yaml")
		if !helper.IsFile(rulesFile) {
			return fmt.Errorf("rules document not found: %s", rulesFile)
		}
		log.Debug().Msgf("using rules document %s", rulesFile)
		layer.UpdateTemplateDependencies(tplDir, rulesFile)
		err = checkInputs(layer.ComputeExpandedInputs(rootDir))
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
			TemplatesDir: helper.Join(tplDir, "templates"),
			System:       system,
			UserFeatures: layer.Features,
			UserForce:    layer.Force,
			Meta:         doc.Meta,
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
