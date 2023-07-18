package spec

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/repos"
)

type SolutionLayer struct {
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
	// dependencies are the dependencies of the layer
	dependencies []string `json:"-"` // dependencies of the layer
	// TemplateDir is the directory of the template
	TemplateDir string `json:"-" yaml:"-"`
	// TemplatesDir is the "templates" directory inside the template dir
	TemplatesDir string `json:"-" yaml:"-"`
	// RulesFile is the "rules.yaml" file inside the template dir
	RulesFile string `json:"-" yaml:"-"`
}

// GetOutputDir returns the output dir.
// The output dir can be relative to the root dir of the solution.
func (l *SolutionLayer) GetOutputDir(rootDir string) string {
	return helper.Join(rootDir, l.Output)
}

func (l *SolutionLayer) Validate() error {
	// basic validation
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
	// check for advanced validation
	if l.computed {
		if !helper.IsDir(l.TemplateDir) {
			return fmt.Errorf("template dir not found: %s", l.TemplateDir)
		}
		if !helper.IsDir(l.TemplatesDir) {
			return fmt.Errorf("templates dir not found: %s", l.TemplatesDir)
		}
		if !helper.IsFile(l.RulesFile) {
			return fmt.Errorf("rules file not found: %s", l.RulesFile)
		}
		// check inputs
		for _, input := range l.expandedInputs {
			switch helper.Ext(input) {
			case ".yaml", ".yml", ".json":
				result, err := CheckFile(input)
				if err != nil {
					return err
				}
				if !result.Valid() {
					for _, e := range result.Errors() {
						log.Warn().Msg(e.String())
					}
					return fmt.Errorf("invalid file: %s", input)
				}
			case ".idl":
				err := CheckIdlFile(input)
				if err != nil {
					return err
				}
			default:
				return fmt.Errorf("unknown type %s", input)
			}
		}
	}
	return nil
}

// Compute computes the dependencies and expanded inputs of a layer.
func (l *SolutionLayer) Compute(doc *SolutionDoc) error {
	if l.computed {
		return nil
	}
	l.computed = true

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
	return nil
}

func (l *SolutionLayer) Dependencies() []string {
	if !l.computed {
		log.Error().Msg("layer not computed, dependencies not available")
	}
	return l.dependencies
}

func (l *SolutionLayer) ExpandedInputs() []string {
	if !l.computed {
		log.Error().Msg("layer not computed, expanded inputs not available")
	}
	return l.expandedInputs
}
