package spec

import (
	"github.com/apigear-io/cli/pkg/helper"
)

type SolutionTarget struct {
	Name        string   `json:"name" yaml:"name"`
	Description string   `json:"description" yaml:"description"`
	Inputs      []string `json:"inputs" yaml:"inputs"`
	Output      string   `json:"output" yaml:"output"`
	Template    string   `json:"template" yaml:"template"`
	Features    []string `json:"features" yaml:"features"`
	Force       bool     `json:"force" yaml:"force"`
	// computed fields
	computed bool `json:"-" yaml:"-"`
	// expandedInputs is the inputs with the variables expanded
	expandedInputs []string `json:"-"` // expanded inputs
	// dependencies are the dependencies of the target
	dependencies []string `json:"-"` // dependencies of the target
	// TemplateDir is the directory of the template
	TemplateDir string `json:"-" yaml:"-"`
	// TemplatesDir is the "templates" directory inside the template dir
	TemplatesDir string `json:"-" yaml:"-"`
	// RulesFile is the "rules.yaml" file inside the template dir
	RulesFile string `json:"-" yaml:"-"`
	isValid   bool   `json:"-" yaml:"-"` // is valid
}

func (t *SolutionTarget) Accept(doc *SolutionDoc, visitor SolutionVisitor) error {
	return visitor.VisitSolutionTarget(doc, t)
}

// GetOutputDir returns the output dir.
// The output dir can be relative to the root dir of the solution.
func (l *SolutionTarget) GetOutputDir(rootDir string) string {
	return helper.Join(rootDir, l.Output)
}

func (l *SolutionTarget) Dependencies() []string {
	if !l.computed {
		log.Error().Msg("target not computed, dependencies not available")
	}
	return l.dependencies
}

func (l *SolutionTarget) ExpandedInputs() []string {
	if !l.computed {
		log.Error().Msg("target not computed, expanded inputs not available")
	}
	return l.expandedInputs
}

// AddInput adds an input to the target.
func (l *SolutionTarget) AddInput(input string) {
	l.Inputs = append(l.Inputs, input)
}
