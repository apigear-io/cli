package gen

// A rules document defines a set of rules how to apply transformations
// to a set of documents.
// For this the rules document is separated into a set of features, which can be enabled independently.
// Each feature can depend on another feature, to form a dependency graph.
// Transformation are applied based on the symbol type. A symbol can be a
// system, module, interface, enum or struct.
// For this the feature has a set of scopes to match these symbol types.

type ScopeType string

const (
	ScopeSystem    ScopeType = "system"
	ScopeModule    ScopeType = "module"
	ScopeInterface ScopeType = "interface"
	ScopeStruct    ScopeType = "struct"
	ScopeEnum      ScopeType = "enum"
)

type RulesDoc struct {
	Features []*FeatureRule `json:"features" yaml:"features"`
}

type FeatureRule struct {
	Name     string       `json:"name" yaml:"name"`
	Requires []string     `json:"requires" yaml:"requires"`
	Scopes   []*ScopeRule `json:"scopes" yaml:"scopes"`
}

func NewFeatureRule(name string) *FeatureRule {
	return &FeatureRule{
		Name: name,
	}
}

func (s *FeatureRule) ScopeByMatch(match ScopeType) *ScopeRule {
	for _, scope := range s.Scopes {
		if scope.Match == match {
			return scope
		}
	}
	return nil
}

type ScopeRule struct {
	Match     ScopeType       `json:"match" yaml:"match"`
	Documents []*DocumentRule `json:"documents" yaml:"documents"`
}

func NewScopeRule(match ScopeType) *ScopeRule {
	return &ScopeRule{
		Match: match,
	}
}

type DocumentRule struct {
	Source    string `json:"source" yaml:"source"`
	Target    string `json:"target" yaml:"target"`
	Transform bool   `json:"transform" yaml:"transform"`
	Force     bool   `json:"force" yaml:"force"`
}

func NewDocumentRule(source, target string) *DocumentRule {
	return &DocumentRule{
		Source:    source,
		Target:    target,
		Transform: true,
		Force:     true,
	}
}
