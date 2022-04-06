package gen

import (
	"fmt"
	"io/ioutil"
	"objectapi/pkg/model"
	"objectapi/pkg/spec"
	"path"

	"gopkg.in/yaml.v2"
)

// Generator parses documents and applies
// template transformation on a set of files.

type Context = map[string]interface{}

// IFileWriter writes a target file with content
type IFileWriter interface {
	WriteFile(fn string, content string, force bool) error
}

// IRenderEngine renders to string from template or file using context
type IRenderEngine interface {
	RenderString(template string, ctx Context) (string, error)
	RenderFile(name string, ctx Context) (string, error)
}

type GeneratorOptions struct {
	System    *model.System
	UserForce bool `yaml:"force"`
}

// Generator applies template transformation on a set of files define in rules
type Generator struct {
	Engine  IRenderEngine
	Writer  IFileWriter
	Options *GeneratorOptions
}

// NewGenerator creates a new processor
func NewGenerator(e IRenderEngine, w IFileWriter) *Generator {
	return &Generator{
		Engine: e,
		Writer: w,
	}
}

func NewDefaultGenerator(tplSearchDir string, outputDir string) *Generator {
	return NewGenerator(NewRenderer(tplSearchDir), NewFileWriter(outputDir))
}

func (g *Generator) ProcessFile(filename string, o *GeneratorOptions) error {
	var bytes, err = ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading file %s: %s", filename, err)
	}
	var rules = spec.RulesDoc{}
	err = yaml.Unmarshal(bytes, &rules)
	if err != nil {
		return fmt.Errorf("error parsing file %s: %s", filename, err)
	}
	return g.Process(rules, o)
}

// Process processes a set of rules from a rules document
func (g *Generator) Process(rules spec.RulesDoc, o *GeneratorOptions) error {
	g.Options = o
	if g.Options.System == nil {
		return fmt.Errorf("system is nil")
	}
	for _, feature := range rules.Features {
		err := g.processFeature(feature)
		if err != nil {
			return fmt.Errorf("error processing feature %s: %s", feature.Name, err)
		}
	}
	return nil
}

// processFeature processes a feature rule
func (g *Generator) processFeature(f *spec.FeatureRule) error {
	s := g.Options.System
	// process system
	var ctx = Context{"System": s}
	scope := f.FindScopeByMatch(spec.ScopeSystem)
	err := g.processScope(scope, ctx)
	if err != nil {
		return fmt.Errorf("error processing system scope: %s", err)
	}
	for _, module := range s.Modules {
		// process module
		scope := f.FindScopeByMatch(spec.ScopeModule)
		ctx = Context{"System": s, "Module": module}
		err := g.processScope(scope, ctx)
		if err != nil {
			return fmt.Errorf("error processing module %s: %s", module.Name, err)
		}
		for _, iface := range module.Interfaces {
			// process interface
			ctx["Interface"] = iface
			scope := f.FindScopeByMatch(spec.ScopeInterface)
			err := g.processScope(scope, ctx)
			if err != nil {
				return fmt.Errorf("error processing interface %s: %s", iface.Name, err)
			}
		}
		for _, struct_ := range module.Structs {
			// process struct
			ctx["Struct"] = struct_
			scope := f.FindScopeByMatch(spec.ScopeStruct)
			err := g.processScope(scope, ctx)
			if err != nil {
				return fmt.Errorf("error processing struct %s: %s", struct_.Name, err)
			}
		}
		for _, enum := range module.Enums {
			// process enum
			ctx["Enum"] = enum
			scope := f.FindScopeByMatch(spec.ScopeEnum)
			err := g.processScope(scope, ctx)
			if err != nil {
				return fmt.Errorf("error processing enum %s: %s", enum.Name, err)
			}
		}
	}
	return nil
}

// processScope processes a scope rule (e.g. system, modules, ...) with the given context
func (g *Generator) processScope(scope *spec.ScopeRule, ctx Context) error {
	if scope == nil {
		return nil
	}
	for _, doc := range scope.Documents {
		g.processDocument(doc, ctx)
	}
	return nil
}

// processDocument processes a document rule with the given context
func (g *Generator) processDocument(doc *spec.DocumentRule, ctx Context) error {
	// the source file to render
	var source = path.Clean(doc.Source)
	// the target destination file
	var target = path.Clean(doc.Target)
	// either user can force an overwrite or the target or the rules document
	force := doc.Force || g.Options.UserForce
	if target == "" {
		target = source
	}
	// var force = doc.Force
	// var transform = doc.Transform
	log.Infof("processing document %s -> %s", source, target)
	// render the template using the context
	content, err := g.Engine.RenderFile(source, ctx)
	if err != nil {
		return fmt.Errorf("error rendering file %s: %s", source, err)
	}
	// write the file
	err = g.Writer.WriteFile(target, content, force)
	if err != nil {
		return fmt.Errorf("error writing file %s: %s", target, err)
	}
	return nil
}
