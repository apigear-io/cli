package gen

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"objectapi/pkg/gen/filtercpp"
	"objectapi/pkg/model"
	"objectapi/pkg/spec"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

// Generator parses documents and applies
// template transformation on a set of files.

type DataMap = map[string]interface{}

// IFileWriter writes a target file with content
type IFileWriter interface {
	WriteFile(fn string, buf []byte, force bool) error
}

type GeneratorOptions struct {
	System       *model.System
	UserForce    bool
	TemplatesDir string
	OutputDir    string
}

// Generator applies template transformation on a set of files define in rules
type Generator struct {
	Template     *template.Template
	Writer       IFileWriter
	System       *model.System
	UserForce    bool
	TemplatesDir string
	OutputDir    string
}

func NewGenerator(outputDir string, templatesDir string) (*Generator, error) {
	g := &Generator{
		Writer:       NewFileWriter(outputDir),
		Template:     template.New(""),
		UserForce:    false,
		System:       model.NewSystem(""),
		TemplatesDir: templatesDir,
	}
	err := g.ParseTemplateDir(templatesDir)
	if err != nil {
		return nil, err
	}
	g.Template.Funcs(filtercpp.GetFuncMap())
	return g, nil
}

func (g *Generator) ParseTemplate(path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	_, err = g.Template.New(filepath.Base(path)).Parse(string(b))
	return err
}

func (g *Generator) ParseTemplateDir(dir string) error {
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			fmt.Println("error walking dir:", err)
		}
		// ignore all dirs
		if d.IsDir() {
			return nil
		}
		// ignore files start starting with .
		if strings.HasPrefix(filepath.Base(path), ".") {
			return nil
		}
		return g.ParseTemplate(path)
	})
	return err
}

func (g *Generator) ProcessFile(filename string) error {
	var bytes, err = ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading file %s: %s", filename, err)
	}
	var rules = spec.RulesDoc{}
	err = yaml.Unmarshal(bytes, &rules)
	if err != nil {
		return fmt.Errorf("error parsing file %s: %s", filename, err)
	}
	return g.ProcessRulesDoc(rules)
}

// ProcessRulesDoc processes a set of rules from a rules document
func (g *Generator) ProcessRulesDoc(rules spec.RulesDoc) error {
	if g.System == nil {
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
func (g *Generator) processFeature(f spec.FeatureRule) error {
	// process system
	var data = DataMap{"System": g.System}
	scope := f.FindScopeByMatch(spec.ScopeSystem)
	err := g.processScope(scope, data)
	if err != nil {
		return fmt.Errorf("error processing system scope: %s", err)
	}
	for _, module := range g.System.Modules {
		// process module
		scope := f.FindScopeByMatch(spec.ScopeModule)
		data = DataMap{"System": g.System, "Module": module}
		err := g.processScope(scope, data)
		if err != nil {
			return fmt.Errorf("error processing module %s: %s", module.Name, err)
		}
		for _, iface := range module.Interfaces {
			// process interface
			data["Interface"] = iface
			scope := f.FindScopeByMatch(spec.ScopeInterface)
			err := g.processScope(scope, data)
			if err != nil {
				return fmt.Errorf("error processing interface %s: %s", iface.Name, err)
			}
		}
		for _, struct_ := range module.Structs {
			// process struct
			data["Struct"] = struct_
			scope := f.FindScopeByMatch(spec.ScopeStruct)
			err := g.processScope(scope, data)
			if err != nil {
				return fmt.Errorf("error processing struct %s: %s", struct_.Name, err)
			}
		}
		for _, enum := range module.Enums {
			// process enum
			data["Enum"] = enum
			scope := f.FindScopeByMatch(spec.ScopeEnum)
			err := g.processScope(scope, data)
			if err != nil {
				return fmt.Errorf("error processing enum %s: %s", enum.Name, err)
			}
		}
	}
	return nil
}

// processScope processes a scope rule (e.g. system, modules, ...) with the given context
func (g *Generator) processScope(scope spec.ScopeRule, ctx DataMap) error {
	for _, doc := range scope.Documents {
		err := g.processDocument(doc, ctx)
		if err != nil {
			return fmt.Errorf("error processing document %s: %s", doc.Source, err)
		}
	}
	return nil
}

// processDocument processes a document rule with the given context
func (g *Generator) processDocument(doc spec.DocumentRule, ctx DataMap) error {
	// the source file to render
	var source = path.Clean(doc.Source)
	// the target destination file
	var target = path.Clean(doc.Target)
	// either user can force an overwrite or the target or the rules document
	force := doc.Force || g.UserForce
	if target == "" {
		target = source
	}
	// var force = doc.Force
	// var transform = doc.Transform
	log.Infof("transform %s -> %s", source, target)
	// render the template using the context
	buf := bytes.NewBuffer(nil)
	err := g.Template.ExecuteTemplate(buf, source, ctx)
	if err != nil {
		return fmt.Errorf("error rendering file %s: %s", source, err)
	}
	// write the file
	err = g.Writer.WriteFile(target, buf.Bytes(), force)
	if err != nil {
		return fmt.Errorf("error writing file %s: %s", target, err)
	}
	return nil
}
