package spec

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/cfg"
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
}

// ResolveTemplateDir resolves the template dir.
// The template dir can be a relative path to the root dir of the solution.
// If the template dir is not a relative path, it is considered as a template name.
// The template name is used to find the template dir in the template cache dir.
func (l *SolutionLayer) ResolveTemplateDir(rootDir string) string {
	if helper.IsDir(helper.Join(rootDir, l.Template)) {
		return helper.Join(rootDir, l.Template)
	}
	if helper.IsDir(helper.Join(cfg.CacheDir(), l.Template)) {
		return helper.Join(cfg.CacheDir(), l.Template)
	}
	return ""
}

// GetOutputDir returns the output dir.
// The output dir can be relative to the root dir of the solution.
func (l *SolutionLayer) GetOutputDir(rootDir string) string {
	return helper.Join(rootDir, l.Output)
}

// GetTemplatesDir returns the templates dir.
// The templates dir is a sub directory of the template dir.
func (l *SolutionLayer) GetTemplatesDir(rootDir string) string {
	tDir := l.ResolveTemplateDir(rootDir)
	if tDir == "" {
		return ""
	}
	return helper.Join(tDir, "templates")
}

// GetRulesFile returns the rules file.
// The rules file is a file in the template dir.
func (l *SolutionLayer) GetRulesFile(rootDir string) string {
	tDir := l.ResolveTemplateDir(rootDir)
	if tDir == "" {
		return ""
	}
	return helper.Join(tDir, "rules.yaml")
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

func (l *SolutionLayer) ComputeDependencies(rootDir string) []string {
	if l.dependencies == nil {
		l.dependencies = make([]string, 0)
	}
	if len(l.dependencies) == 0 {
		rulesFile := l.GetRulesFile(rootDir)
		if rulesFile != "" {
			l.dependencies = append(l.dependencies, rulesFile)
		}
		tplsDir := l.GetTemplatesDir(rootDir)
		if tplsDir != "" {
			l.dependencies = append(l.dependencies, tplsDir)
		}
		inputs := l.ComputeExpandedInputs(rootDir)
		l.dependencies = append(l.dependencies, inputs...)
	}
	return l.dependencies
}

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
