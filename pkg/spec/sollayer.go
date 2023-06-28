package spec

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/helper"
)

type SolutionLayer struct {
	Name           string   `json:"name" yaml:"name"`
	Description    string   `json:"description" yaml:"description"`
	Inputs         []string `json:"inputs" yaml:"inputs"`
	Output         string   `json:"output" yaml:"output"`
	Template       string   `json:"template" yaml:"template"`
	Features       []string `json:"features" yaml:"features"`
	Force          bool     `json:"force" yaml:"force"`
	expandedInputs []string `json:"-"` // expanded inputs
	dependencies   []string `json:"-"` // dependencies of the layer
	templateDir    string   `json:"-"` // template dir
	rulesFile      string   `json:"-"` // rules file
}

// GetOutputDir returns the output dir.
// The output dir can be relative to the root dir of the solution.
func (l *SolutionLayer) GetOutputDir(rootDir string) string {
	return helper.Join(rootDir, l.Output)
}

func (l *SolutionLayer) Validate() error {
	if l.Output == "" {
		return fmt.Errorf("layer output is required")
	}
	if l.Template == "" {
		return fmt.Errorf("layer template is required")
	}
	if l.Inputs == nil {
		l.Inputs = make([]string, 0)
	}
	if l.Features == nil {
		// if no features, use all
		l.Features = []string{"all"}
	}
	return nil
}

// ComputeDependencies computes the dependencies of a layer.
// The dependencies are used for file system watchers.
// The dependencies of a layer are the rules file, the templates dir and the expanded inputs.
// The rules file is a file in the template dir.
// The templates dir is a sub directory of the template dir.
// The expanded inputs are the inputs with the variables expanded.
func (l *SolutionLayer) ComputeDependencies(rootDir string) []string {
	if l.dependencies == nil {
		l.dependencies = make([]string, 0)
	}
	if len(l.dependencies) == 0 {
		if l.templateDir != "" {
			l.dependencies = append(l.dependencies, l.templateDir)
		}
		if l.rulesFile != "" {
			l.dependencies = append(l.dependencies, l.rulesFile)
		}
		inputs := l.ComputeExpandedInputs(rootDir)
		l.dependencies = append(l.dependencies, inputs...)
	}
	return l.dependencies
}

// ComputeExpandedInputs computes the expanded inputs of a layer.
// The expanded inputs are the inputs with the variables expanded.
func (l *SolutionLayer) ComputeExpandedInputs(rootDir string) []string {
	if l.expandedInputs == nil {
		l.expandedInputs = make([]string, 0)
	}
	if len(l.expandedInputs) == 0 {
		inputs, err := helper.ExpandInputs(rootDir, l.Inputs...)
		if err != nil {
			log.Error().Err(err).Msg("failed to expand inputs")
		}
		l.expandedInputs = append(l.expandedInputs, inputs...)
	}
	return l.expandedInputs
}

// Compute computes the dependencies and expanded inputs of a layer.
func (l *SolutionLayer) Compute(rootDir string) error {
	l.ComputeDependencies(rootDir)
	l.ComputeExpandedInputs(rootDir)
	return nil
}

// UpdateTemplateDependencies updates the template dir and rules file of a layer.
func (l *SolutionLayer) UpdateTemplateDependencies(templateDir, rulesFile string) {
	l.templateDir = templateDir
	l.rulesFile = rulesFile
}
