package sol

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/spec"
)

type Runner struct {
	tasks map[string]*task
}

func NewRunner() *Runner {
	return &Runner{
		tasks: make(map[string]*task),
	}
}

func (r *Runner) HasTask(file string) bool {
	return r.task(file) != nil
}

func (r *Runner) TaskFiles() []string {
	files := make([]string, 0, len(r.tasks))
	for file := range r.tasks {
		files = append(files, file)
	}
	return files
}

func (r *Runner) task(file string) *task {
	return r.tasks[file]
}

// RunDoc runs the given file task once.
func (r *Runner) RunDoc(file string, doc *spec.SolutionDoc) error {
	result, err := spec.CheckFile(file)
	if err != nil {
		log.Warn().Msgf("check document %s: %s", file, err)
		return fmt.Errorf("check document %s: %s", file, err)
	}
	if !result.Valid() {
		log.Warn().Msgf("document %s is invalid", file)
		for _, desc := range result.Errors() {
			log.Warn().Msgf("\t%s", desc)
		}
		return fmt.Errorf("document %s is invalid", file)
	}
	t, err := newTask(file, doc)
	if err != nil {
		return err
	}
	return t.run()
}

func (r *Runner) ensureTask(file string, doc *spec.SolutionDoc) (*task, error) {
	t := r.task(file)
	if t != nil {
		return t, nil
	}
	t, err := newTask(file, doc)
	if err != nil {
		return nil, err
	}
	r.tasks[file] = t
	return t, nil
}

// StartWatch starts the watch of the given file task.
func (r *Runner) StartWatch(file string, doc *spec.SolutionDoc) (chan<- bool, error) {
	t, err := r.ensureTask(file, doc)
	if err != nil {
		return nil, err
	}
	err = t.run()
	if err != nil {
		log.Warn().Msgf("task %s: %s", file, err)
	}
	return t.startWatch()
}

// StopWatch stops the watch of the given file task.
func (r *Runner) StopWatch(file string) {
	t := r.task(file)
	if t != nil {
		t.stopWatch()
		// remove task from runner
		delete(r.tasks, file)
	}
}

func (r *Runner) Clear() {
	for file := range r.tasks {
		r.StopWatch(file)
	}
	r.tasks = make(map[string]*task)
}
