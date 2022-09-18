package spec

// A rules document defines a set of rules how to apply transformations
// to a set of documents.
// For this the rules document is separated into a set of features, which can be enabled independently.
// Each feature can depend on another feature, to form a dependency graph.
// Transformation are applied based on the symbol type. A symbol can be a
// system, module, interface, enum or struct.
// For this the feature has a set of scopes to match these symbol types.

// ScopeType is the type of a scope.
type ScopeType string

const (
	ScopeSystem    ScopeType = "system"
	ScopeModule    ScopeType = "module"
	ScopeInterface ScopeType = "interface"
	ScopeStruct    ScopeType = "struct"
	ScopeEnum      ScopeType = "enum"
)

type RulesDoc struct {
	Features []FeatureRule `json:"features" yaml:"features"`
}

// A feature rule defines a set of scopes to match a symbol type.
type FeatureRule struct {
	// Name of the feature.
	Name string `json:"name" yaml:"name"`
	// Which other features are required by this feature.
	Requires []string `json:"requires" yaml:"requires"`
	// Scopes to match.
	Scopes []ScopeRule `json:"scopes" yaml:"scopes"`
}

// FindScopeByMatch returns the first scope that matches the given match.
func (s *FeatureRule) FindScopeByMatch(match ScopeType) ScopeRule {
	for _, scope := range s.Scopes {
		if scope.Match == match {
			return scope
		}
	}
	return ScopeRule{}
}

// ScopeRule defines a scope rule which is matched based on the symbol type.
type ScopeRule struct {
	// Match is the type of the symbol to match
	Match ScopeType `json:"match" yaml:"match"`
	// Prefix is the prefix for all target documents
	Prefix string `json:"prefix" yaml:"prefix"`
	// Documents is a list of document rules to apply
	Documents []DocumentRule `json:"documents" yaml:"documents"`
}

// DocumentRule defines a document rule with a source and target document.
type DocumentRule struct {
	// Source is the source document to apply the transformation to.
	Source string `json:"source" yaml:"source"`
	// Target is the target document to write to after the transformation.
	Target string `json:"target" yaml:"target"`
	// Transform is true if the transformation should be applied.
	Raw bool `json:"raw" yaml:"raw"`
	// Force is true if the target file should be overwritten.
	Force bool `json:"force" yaml:"force"`
}
