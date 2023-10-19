package spec

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/repos"
)

type SolTargetLinter interface {
	// Init initializes and validates the basic structure
	OnStart(doc *SolutionDoc, target *SolutionTarget) error
	// Compute computes variables and dependencies
	OnFix(doc *SolutionDoc, target *SolutionTarget) error
	// Validate validates the data structure
	OnFinish(doc *SolutionDoc, target *SolutionTarget) error
}

type solTargetFSLinter struct {
}

var _ SolTargetLinter = (*solTargetFSLinter)(nil)

func (v *solTargetFSLinter) OnStart(doc *SolutionDoc, target *SolutionTarget) error {
	if target.Name == "" {
		log.Warn().Msg("target name is empty")
	}
	if target.Output == "" {
		log.Warn().Msg("target output is empty")
		return fmt.Errorf("target %s output is empty", target.Name)
	}
	if target.Inputs == nil {
		target.Inputs = make([]string, 0)
	}
	if target.Features == nil {
		target.Features = []string{"all"}
	}
	return nil
}

func (v *solTargetFSLinter) OnFix(doc *SolutionDoc, target *SolutionTarget) error {
	// compute template dir
	tplDir := helper.Join(doc.RootDir, target.Template)
	// setup template dir
	if helper.IsDir(tplDir) {
		target.TemplateDir = tplDir
		target.TemplatesDir = helper.Join(tplDir, "templates")
		target.RulesFile = helper.Join(tplDir, "rules.yaml")
	} else {
		// try to find the template dir in the templates dir
		repoId, err := repos.GetOrInstallTemplateFromRepoID(target.Template)
		if err != nil {
			log.Err(err).Msgf("failed to get template %s", target.Template)
			return err
		}
		tplDir, err := repos.Cache.GetTemplateDir(repoId)
		if err != nil {
			log.Err(err).Msgf("failed to get template dir %s", target.Template)
			return err
		}
		target.Template = repoId
		target.TemplateDir = tplDir
		target.TemplatesDir = helper.Join(tplDir, "templates")
		target.RulesFile = helper.Join(tplDir, "rules.yaml")
	}

	// record dependencies
	if target.dependencies == nil {
		target.dependencies = make([]string, 0)
	}
	if len(target.dependencies) == 0 {
		if target.TemplatesDir != "" {
			target.dependencies = append(target.dependencies, target.TemplatesDir)
		}
		if target.RulesFile != "" {
			target.dependencies = append(target.dependencies, target.RulesFile)
		}
		target.dependencies = append(target.dependencies, target.expandedInputs...)
	}
	// expand inputs
	if target.expandedInputs == nil {
		target.expandedInputs = make([]string, 0)
	}
	if len(target.expandedInputs) == 0 {
		expanded, err := helper.ExpandInputs(doc.RootDir, target.Inputs...)
		if err != nil {
			return err
		}
		target.expandedInputs = append(target.expandedInputs, expanded...)
	}
	target.dependencies = append(target.dependencies, target.expandedInputs...)
	return nil
}

func (v *solTargetFSLinter) OnFinish(doc *SolutionDoc, target *SolutionTarget) error {
	if !helper.IsDir(target.TemplateDir) {
		return fmt.Errorf("target: %s: template directory %s does not exist", target.Name, target.TemplateDir)
	}
	if !helper.IsDir(target.TemplatesDir) {
		return fmt.Errorf("target: %s: templates directory %s does not exist", target.Name, target.TemplatesDir)
	}
	if !helper.IsFile(target.RulesFile) {
		return fmt.Errorf("target: %s: rules file %s does not exist", target.Name, target.RulesFile)
	}
	for _, input := range target.expandedInputs {
		result, err := CheckFile(input)
		if err != nil {
			return err
		}
		if !result.Valid() {
			for _, e := range result.Errors {
				log.Warn().Msg(e.String())
			}
			return fmt.Errorf("target %s: invalid file: %s", target.Name, input)
		}
	}
	target.computed = true
	return nil
}
