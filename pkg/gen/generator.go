package gen

import (
	"fmt"
	"io/ioutil"
	"objectapi/pkg/model"
	"path"

	"gopkg.in/yaml.v2"
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

// Generator applies template transformation on a set of files define in rules
type Generator struct {
	Engine RenderEngine
	Writer FileWriter
}

// NewGenerator creates a new processor
func NewGenerator(e RenderEngine, w FileWriter) *Generator {
	return &Generator{
		Engine: e,
		Writer: w,
	}
}

func NewDefaultGenerator(tplSearchDir string, outputDir string) *Generator {
	return NewGenerator(NewRenderer(tplSearchDir), NewFileWriter(outputDir))
}

func (g *Generator) ProcessFile(filename string, s *model.System) error {
	var bytes, err = ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading file %s: %s", filename, err)
	}
	var rules = RulesDoc{}
	err = yaml.Unmarshal(bytes, &rules)
	if err != nil {
		return fmt.Errorf("error parsing file %s: %s", filename, err)
	}
	return g.Process(rules, s)
}

// Process processes a set of rules from a rules document
func (g *Generator) Process(rules RulesDoc, s *model.System) error {
	if s == nil {
		return fmt.Errorf("system is nil")
	}
	for _, feature := range rules.Features {
		err := g.processFeature(feature, s)
		if err != nil {
			return fmt.Errorf("error processing feature %s: %s", feature.Name, err)
		}
	}
	return nil
}

func (g *Generator) processFeature(f *FeatureRule, s *model.System) error {
	// process system
	var ctx = Context{"system": s}
	scope := f.ScopeByMatch(ScopeSystem)
	err := g.processScope(scope, ctx)
	if err != nil {
		return fmt.Errorf("error processing system scope: %s", err)
	}
	for _, module := range s.Modules {
		// process module
		scope := f.ScopeByMatch(ScopeModule)
		ctx = Context{"system": s, "module": module}
		err := g.processScope(scope, ctx)
		if err != nil {
			return fmt.Errorf("error processing module %s: %s", module.Name, err)
		}
		for _, iface := range module.Interfaces {
			// process interface
			ctx["interface"] = iface
			scope := f.ScopeByMatch(ScopeInterface)
			err := g.processScope(scope, ctx)
			if err != nil {
				return fmt.Errorf("error processing interface %s: %s", iface.Name, err)
			}
		}
		for _, struct_ := range module.Structs {
			// process struct
			ctx["struct"] = struct_
			scope := f.ScopeByMatch(ScopeStruct)
			err := g.processScope(scope, ctx)
			if err != nil {
				return fmt.Errorf("error processing struct %s: %s", struct_.Name, err)
			}
		}
		for _, enum := range module.Enums {
			// process enum
			ctx["enum"] = enum
			scope := f.ScopeByMatch(ScopeEnum)
			err := g.processScope(scope, ctx)
			if err != nil {
				return fmt.Errorf("error processing enum %s: %s", enum.Name, err)
			}
		}
	}
	return nil
}

func (g *Generator) processScope(scope *ScopeRule, ctx Context) error {
	if scope == nil {
		return nil
	}
	for _, doc := range scope.Documents {
		g.processDocument(doc, ctx)
	}
	return nil
}

func (g *Generator) processDocument(doc *DocumentRule, ctx Context) error {
	var source = path.Clean(doc.Source)
	var target = path.Clean(doc.Target)
	if target == "" {
		target = source
	}
	// var force = doc.Force
	// var transform = doc.Transform
	log.Infof("processing document %s -> %s", source, target)
	content, err := g.Engine.RenderFile(source, ctx)
	if err != nil {
		return fmt.Errorf("error rendering file %s: %s", source, err)
	}
	g.Writer.WriteFile(target, content)
	return nil
}
