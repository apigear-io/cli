package spec

// SolutionVisitor is a simple visitor interface for the solution.
// Each element has an accept method that accepts a visitor and knows
// how to iterate the tree and call the visitor.
type SolutionVisitor interface {
	VisitSolutionDoc(doc *SolutionDoc) error
	VisitSolutionTarget(doc *SolutionDoc, target *SolutionTarget) error
}

// linterVisitor is a visitor that calls the linter on each element.
// A linter is a visitor that checks for common issues, computes variables
// and dependencies and validates the data structure.
// it uses a OnStart, OnFix and OnFinish methods to do the work.
type linterVisitor struct {
	solLinter    SolutionLinter
	targetLinter SolTargetLinter
}

// NewLinterVisitor creates a new linter visitor.
func NewLinterVisitor(solLinter SolutionLinter, solTargetLinter SolTargetLinter) SolutionVisitor {
	return &linterVisitor{
		solLinter:    solLinter,
		targetLinter: solTargetLinter,
	}
}

// check that the linterVisitor implements the SolutionVisitor interface
var _ SolutionVisitor = (*linterVisitor)(nil)

// VisitSolutionDoc visits the solution doc.
func (v *linterVisitor) VisitSolutionDoc(doc *SolutionDoc) error {
	if doc.isValid {
		return nil
	}
	// first we try to check for common issues
	err := v.solLinter.OnStart(doc)
	if err != nil {
		return err
	}
	// then we compute variables and fix invalid values
	err = v.solLinter.OnFix(doc)
	if err != nil {
		return err
	}
	// then we validate to ensure the solution is valid
	err = v.solLinter.OnFinish(doc)
	if err != nil {
		return err
	}
	// mark the solution as valid
	doc.isValid = true
	return nil
}

func (v *linterVisitor) VisitSolutionTarget(doc *SolutionDoc, target *SolutionTarget) error {
	if target.isValid {
		return nil
	}

	// first we try to check for common issues
	err := v.targetLinter.OnStart(doc, target)
	if err != nil {
		return err
	}
	// then we compute variables and fix invalid values
	err = v.targetLinter.OnFix(doc, target)
	if err != nil {
		return err
	}
	// then we validate to ensure the target is valid
	err = v.targetLinter.OnFinish(doc, target)
	if err != nil {
		return err
	}
	// mark the target as valid
	target.isValid = true
	return nil
}
