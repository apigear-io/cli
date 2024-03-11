package spec

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/repos"
)

type SolutionTarget struct {
	Name        string                 `json:"name" yaml:"name"`
	Description string                 `json:"description" yaml:"description"`
	Inputs      []string               `json:"inputs" yaml:"inputs"`
	Output      string                 `json:"output" yaml:"output"`
	Template    string                 `json:"template" yaml:"template"`
	Features    []string               `json:"features" yaml:"features"`
	Force       bool                   `json:"force" yaml:"force"`
	Meta        map[string]interface{} `json:"meta" yaml:"meta"`
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
}

// GetOutputDir returns the output dir.
// The output dir can be relative to the root dir of the solution.
func (l *SolutionTarget) GetOutputDir(rootDir string) string {
	return helper.Join(rootDir, l.Output)
}

func (l *SolutionTarget) Validate(doc *SolutionDoc) error {
	if l.Meta == nil {
		l.Meta = make(map[string]interface{})
	}
	// basic validation
	if l.Output == "" {
		return fmt.Errorf("target %s: output is required", l.Name)
	}
	if l.Template == "" {
		return fmt.Errorf("target %s: template is required", l.Name)
	}
	if l.Inputs == nil {
		l.Inputs = make([]string, 0)
	}
	if l.Features == nil {
		// if no features, use all
		l.Features = []string{"all"}
	}
	// compute derived fields
	if err := l.compute(doc); err != nil {
		return err
	}
	// extended validation
	if !helper.IsDir(l.TemplateDir) {
		return fmt.Errorf("target %s: template dir not found: %s", l.Name, l.TemplateDir)
	}
	if !helper.IsDir(l.TemplatesDir) {
		return fmt.Errorf("target %s: templates dir not found: %s", l.Name, l.TemplatesDir)
	}
	if !helper.IsFile(l.RulesFile) {
		return fmt.Errorf("target %s: rules file not found: %s", l.Name, l.RulesFile)
	}
	// check inputs
	for _, input := range l.expandedInputs {
		result, err := CheckFile(input)
		if err != nil {
			return err
		}
		if !result.Valid() {
			for _, e := range result.Errors {
				log.Warn().Msg(e.String())
			}
			return fmt.Errorf("target %s: invalid file: %s", l.Name, input)
		}
	}
	return nil
}

// compute computes the dependencies and expanded inputs of a target.
func (l *SolutionTarget) compute(doc *SolutionDoc) error {
	if l.computed {
		return nil
	}
	// compute template dir
	tplDir := helper.Join(doc.RootDir, l.Template)
	if helper.IsDir(tplDir) {
		l.TemplateDir = tplDir
		l.TemplatesDir = helper.Join(tplDir, "templates")
		l.RulesFile = helper.Join(tplDir, "rules.yaml")
	} else {
		// try to find the template dir in the templates dir
		repoId, err := repos.GetOrInstallTemplateFromRepoID(l.Template)
		if err != nil {
			log.Err(err).Msgf("failed to get template %s", l.Template)
			return err
		}
		tplDir, err := repos.Cache.GetTemplateDir(repoId)
		if err != nil {
			log.Err(err).Msgf("failed to get template dir %s", l.Template)
			return err
		}
		l.Template = repoId
		l.TemplateDir = tplDir
		l.TemplatesDir = helper.Join(tplDir, "templates")
		l.RulesFile = helper.Join(tplDir, "rules.yaml")
	}

	// record dependencies
	if l.dependencies == nil {
		l.dependencies = make([]string, 0)
	}
	if len(l.dependencies) == 0 {
		if l.TemplatesDir != "" {
			l.dependencies = append(l.dependencies, l.TemplatesDir)
		}
		if l.RulesFile != "" {
			l.dependencies = append(l.dependencies, l.RulesFile)
		}
		l.dependencies = append(l.dependencies, l.expandedInputs...)
	}
	// expand inputs
	if l.expandedInputs == nil {
		l.expandedInputs = make([]string, 0)
	}
	if len(l.expandedInputs) == 0 {
		expanded, err := helper.ExpandInputs(doc.RootDir, l.Inputs...)
		if err != nil {
			return err
		}
		l.expandedInputs = append(l.expandedInputs, expanded...)
	}
	l.dependencies = append(l.dependencies, l.expandedInputs...)
	l.computed = true
	return nil
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
