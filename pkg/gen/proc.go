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
	WriteFile(fn string, content string) error
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

// ProcessRules processes a set of rules from a rules document
func (g *Processor) ProcessRules(rules []*FeatureRule, s *model.System) error {
	if s == nil {
		return fmt.Errorf("system is nil")
	}
	for _, rule := range rules {
		err := g.processFeature(rule, s)
		if err != nil {
			return fmt.Errorf("error processing feature %s: %s", rule.Name, err)
		}
	}
	return nil
}

func (g *Processor) processFeature(r *FeatureRule, s *model.System) error {
	// process system
	var ctx = Context{"system": s}
	scope := r.ScopeByMatch(ScopeSystem)
	err := g.processScope(scope, ctx)
	if err != nil {
		return fmt.Errorf("error processing system scope: %s", err)
	}
	for _, module := range s.Modules {
		// process module
		scope := r.ScopeByMatch(ScopeModule)
		ctx = Context{"system": s, "module": module}
		err := g.processScope(scope, ctx)
		if err != nil {
			return fmt.Errorf("error processing module %s: %s", module.Name, err)
		}
		for _, iface := range module.Interfaces {
			// process interface
			ctx["interface"] = iface
			scope := r.ScopeByMatch(ScopeInterface)
			err := g.processScope(scope, ctx)
			if err != nil {
				return fmt.Errorf("error processing interface %s: %s", iface.Name, err)
			}
		}
		for _, struct_ := range module.Structs {
			// process struct
			ctx["struct"] = struct_
			scope := r.ScopeByMatch(ScopeStruct)
			err := g.processScope(scope, ctx)
			if err != nil {
				return fmt.Errorf("error processing struct %s: %s", struct_.Name, err)
			}
		}
		for _, enum := range module.Enums {
			// process enum
			ctx["enum"] = enum
			scope := r.ScopeByMatch(ScopeEnum)
			err := g.processScope(scope, ctx)
			if err != nil {
				return fmt.Errorf("error processing enum %s: %s", enum.Name, err)
			}
		}
	}
	return nil
}

func (g *Processor) processScope(scope *ScopeRule, ctx Context) error {
	if scope == nil {
		return nil
	}
	for _, doc := range scope.Documents {
		g.processDocument(doc, ctx)
	}
	return nil
}

func (g *Processor) processDocument(doc *DocumentRule, ctx Context) error {
	var source = path.Clean(doc.Source)
	var target = path.Clean(doc.Target)
	if target == "" {
		target = source
	}
	var force = doc.Force
	var transform = doc.Transform
	fmt.Printf("Write %s to %s force=%T transform=%T", source, target, force, transform)
	content, err := g.Engine.RenderFile(source, ctx)
	if err != nil {
		return fmt.Errorf("error rendering file %s: %s", source, err)
	}
	g.Writer.WriteFile(target, content)
	return nil
}
