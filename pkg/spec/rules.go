package spec

import (
	"fmt"
	"sort"
)

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

func containsString(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

type RulesDoc struct {
	Name     string         `json:"name" yaml:"name"`
	Features []*FeatureRule `json:"features" yaml:"features"`
}

// FeatureByName returns the feature with the given name.
func (r *RulesDoc) FeatureByName(name string) *FeatureRule {
	for _, f := range r.Features {
		if f.Name == name {
			return f
		}
	}
	return nil
}

func (d *RulesDoc) Validate() error {
	if d.Features == nil {
		d.Features = make([]*FeatureRule, 0)
	}
	for _, f := range d.Features {
		if err := f.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// ComputeFeatures returns a filtered set of features based on the given features.
// And the features that are required by the given features.
func (r *RulesDoc) ComputeFeatures(wanted []string) error {
	log.Debug().Msgf("computing features: %v", wanted)
	// we skip all features first
	for _, f := range r.Features {
		f.Skip = true
	}
	for _, name := range wanted {
		// return all features if the wanted feature is "all"
		if name == "all" {
			for _, f := range r.Features {
				f.Skip = false
			}
			// we are done
			return nil
		}
	}
	// check that all wanted features exist
	for _, name := range wanted {
		// check if the wanted feature exists at all
		f := r.FeatureByName(name)
		if f == nil {
			return fmt.Errorf("unknown wanted feature %s", name)
		}
	}
	// walk the dependency graph of the wanted features
	// avoid circular dependencies
	features := map[string]bool{}
	for _, f := range r.Features {
		if !containsString(wanted, f.Name) {
			// skip features that are not wanted
			continue
		}
		features[f.Name] = true
		err := r.walkRequired(f.Requires, features)
		if err != nil {
			return err
		}
	}
	// enable all features that are required
	for name := range features {
		f := r.FeatureByName(name)
		if f == nil {
			// required feature does not exist, should not happen
			return fmt.Errorf("feature %s not found", name)
		}
		f.Skip = false
	}
	return nil
}

func (r *RulesDoc) walkRequired(required []string, features map[string]bool) error {
	for _, name := range required {
		if features[name] {
			// feature is already enabled
			// this avoids circular dependencies
			continue
		}
		features[name] = true
		f := r.FeatureByName(name)
		if f == nil {
			// required feature does not exist
			return fmt.Errorf("feature %s not found", name)
		}
		err := r.walkRequired(f.Requires, features)
		if err != nil {
			return err
		}
	}
	return nil
}

// // walkWantedFeatures walks the dependency graph of the given features.
// func (r *RulesDoc) walkWantedFeatures(features []string) error {
// 	// make a set of wanted features
// 	if len(features) == 0 {
// 		return nil
// 	}
// 	// if no features are given, then no features are wanted
// 	for _, name := range features {
// 		// resolve feature by name
// 		f := r.FeatureByName(name)
// 		if f == nil {
// 			return fmt.Errorf("feature %s not found", name)
// 		}
// 		// mark feature as wanted
// 		f.Skip = false
// 		// recursively walk the dependency graph
// 		err := r.walkWantedFeatures(f.Requires)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

func (r *RulesDoc) FeatureNamesMap() map[string]bool {
	m := make(map[string]bool, len(r.Features))
	for _, f := range r.Features {
		m[f.Name] = !f.Skip
	}
	return m
}

// A feature rule defines a set of scopes to match a symbol type.
type FeatureRule struct {
	// Name of the feature.
	Name string `json:"name" yaml:"name"`
	// Which other features are required by this feature.
	Requires []string `json:"requires" yaml:"requires"`
	// Scopes to match.
	Scopes []*ScopeRule `json:"scopes" yaml:"scopes"`
	Skip   bool         `json:"-" yaml:"-"`
}

func (r *FeatureRule) Validate() error {
	if r.Scopes == nil {
		r.Scopes = make([]*ScopeRule, 0)
	}
	for _, s := range r.Scopes {
		if err := s.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// FindScopeByMatch returns the first scope that matches the given match.
func (s *FeatureRule) FindScopesByMatch(match ScopeType) []*ScopeRule {
	var scopes []*ScopeRule
	for _, scope := range s.Scopes {
		if scope.Match == match {
			scopes = append(scopes, scope)
		}
	}
	return scopes
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

func (r *ScopeRule) Validate() error {
	if r.Documents == nil {
		r.Documents = make([]DocumentRule, 0)
	}
	for _, d := range r.Documents {
		if err := d.Validate(); err != nil {
			return err
		}
	}
	return nil
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

func (r *DocumentRule) Validate() error {
	return nil
}

func FeatureRulesToStrings(features []*FeatureRule) []string {
	result := []string{}
	for _, f := range features {
		result = append(result, f.Name)
	}
	sort.Strings(result)
	return result
}

func FeatureRulesToStringMap(features []*FeatureRule) map[string]bool {
	result := map[string]bool{}
	for _, f := range features {
		result[f.Name] = true
	}
	return result
}
