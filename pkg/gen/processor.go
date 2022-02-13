package gen

import (
	"fmt"
	"objectapi/pkg/model"
	"path"
)

// Generator parses documents and applies
// template transformation on a set of files.

type Context = map[string]interface{}

// FileWriter writes a target file with content
type FileWriter interface {
	WriteFile(target string, content []byte) error
}

// RenderEngine renders to string from template or file using context
type RenderEngine interface {
	RenderString(template string, ctx Context) (string, error)
	RenderFile(name string, ctx Context) (string, error)
}

// Processor applies template transformation on a set of files define in rules
type Processor struct {
	Engine RenderEngine
	Writer FileWriter
}

// NewProcessor creates a new processor
func NewProcessor(e RenderEngine, w FileWriter) *Processor {
	return &Processor{Engine: e, Writer: w}
}

func (g *Processor) ProcessRules(rules []*FeatureRule, s *model.System) {
	for _, rule := range rules {
		g.processFeature(rule, s)
	}
}

func (g *Processor) processFeature(r *FeatureRule, s *model.System) {
	// process system
	var ctx = Context{"system": s}
	scope := r.ScopeByMatch(ScopeSystem)
	g.processScope(scope, ctx)
	for _, module := range s.Modules {
		// process module
		scope := r.ScopeByMatch(ScopeModule)
		ctx = Context{"system": s, "module": module}
		g.processScope(scope, ctx)
		for _, iface := range module.Interfaces {
			// process interface
			ctx["interface"] = iface
			scope := r.ScopeByMatch(ScopeInterface)
			g.processScope(scope, ctx)
		}
		for _, struct_ := range module.Structs {
			// process struct
			ctx["struct"] = struct_
			scope := r.ScopeByMatch(ScopeStruct)
			g.processScope(scope, ctx)
		}
		for _, enum := range module.Enums {
			// process enum
			ctx["enum"] = enum
			scope := r.ScopeByMatch(ScopeEnum)
			g.processScope(scope, ctx)
		}
	}
}

func (g *Processor) processScope(scope *ScopeRule, ctx Context) error {
	if scope == nil {
		return nil
	}
	for _, doc := range scope.Documents {
		g.processDocument(doc)
	}
	return nil
}

func (g *Processor) processDocument(doc *DocumentRule) error {
	var source = path.Clean(doc.Source)
	var target = path.Clean(doc.Target)
	if target == "" {
		target = source
	}
	var force = doc.Force
	var transform = doc.Transform
	fmt.Printf("Write %s to %s force=%T transform=%T", source, target, force, transform)
	return nil
}
