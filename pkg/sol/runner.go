package sol

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/spec"
)

type runner struct {
	tasks map[string]*task
}

func NewRunner() *runner {
	return &runner{
		tasks: make(map[string]*task),
	}
}

func (r *runner) HasTask(file string) bool {
	return r.task(file) != nil
}

func (r *runner) TaskFiles() []string {
	files := make([]string, len(r.tasks))
	i := 0
	for k := range r.tasks {
		files[i] = k
		i++
	}
	return files
}

func (r *runner) task(file string) *task {
	return r.tasks[file]
}

// RunDoc runs the given file task once.
func (r *runner) RunDoc(file string, doc *spec.SolutionDoc) error {
	t, err := newTask(file, doc)
	if err != nil {
		return err
	}
	return t.run()
}

func (r *runner) ensureTask(file string, doc *spec.SolutionDoc) (*task, error) {
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
func (r *runner) StartWatch(file string, doc *spec.SolutionDoc) (chan<- bool, error) {
	t, err := r.ensureTask(file, doc)
	if err != nil {
		return nil, err
	}
	err = t.run()
	if err != nil {
		return nil, fmt.Errorf("error running watched doc: %s", err)
	}
	return t.startWatch()
}

// StopWatch stops the watch of the given file task.
func (r *runner) StopWatch(file string) {
	t := r.task(file)
	if t != nil {
		t.stopWatch()
	}
}
