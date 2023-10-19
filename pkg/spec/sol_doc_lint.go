package spec

import "fmt"

type SolutionLinter interface {
	// Init initializes and validates the basic structure
	OnStart(doc *SolutionDoc) error
	// Compute computes variables and dependencies
	OnFix(doc *SolutionDoc) error
	// Validate validates the data structure
	OnFinish(doc *SolutionDoc) error
}

type solFSLinter struct {
}

var _ SolutionLinter = (*solFSLinter)(nil)

func (v *solFSLinter) OnStart(doc *SolutionDoc) error {
	if doc.Targets == nil {
		doc.Targets = make([]*SolutionTarget, 0)
	}
	if len(doc.Layers) > 0 {
		log.Warn().Msg("layers inside solutions are deprecated, use targets instead")
		doc.Targets = append(doc.Targets, doc.Layers...)
		doc.Layers = nil
	}
	return nil
}

func (v *solFSLinter) OnFix(doc *SolutionDoc) error {
	// nothing to fix
	return nil
}

func (v *solFSLinter) OnFinish(doc *SolutionDoc) error {
	if doc.RootDir == "" {
		return fmt.Errorf("solution root dir is empty")
	}
	if doc.Name == "" {
		log.Warn().Msg("solution name is empty")
	}
	doc.computed = true
	return nil
}

func LintSolutionFS(doc *SolutionDoc) error {
	l := linterVisitor{
		solLinter:    &solFSLinter{},
		targetLinter: &solTargetFSLinter{},
	}
	return doc.Accept(&l)
}
