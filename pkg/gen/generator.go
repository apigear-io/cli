package gen

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/apigear-io/cli/pkg/gen/filters"
	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/model"
	"github.com/apigear-io/cli/pkg/spec"
)

// Generator parses documents and applies
// template transformation on a set of files.

type GeneratorStats struct {
	FilesWritten int           `json:"files_written"`
	FilesSkipped int           `json:"files_skipped"`
	FilesCopied  int           `json:"files_copied"`
	FilesTouched []string      `json:"files_touched"`
	RunStart     time.Time     `json:"run_start"`
	RunEnd       time.Time     `json:"run_end"`
	Duration     time.Duration `json:"duration"`
}

func (g *GeneratorStats) Start() {
	g.RunStart = time.Now()
}

func (g *GeneratorStats) Stop() {
	g.RunEnd = time.Now()
	g.Duration = g.RunEnd.Sub(g.RunStart).Truncate(time.Millisecond)
	log.Info().Msgf("generated %d files in %s. (%d write, %d skip, %d copy)", g.TotalFiles(), g.Duration, g.FilesWritten, g.FilesSkipped, g.FilesCopied)
}

func (s *GeneratorStats) TotalFiles() int {
	return s.FilesWritten + s.FilesSkipped + s.FilesCopied
}

type GeneratorOptions struct {
	// OutputDir is the directory where files are written
	OutputDir string
	// TemplatesDir is the directory where templates are located
	TemplatesDir string
	// System is the root system model
	System *model.System
	// UserFeatures is a list of features defined by user
	UserFeatures []string
	// UserForce forces overwrite of existing files
	UserForce bool
	// Output is the output writer
	Output OutputWriter
	// DryRun does not write files
	DryRun bool
	// Meta is a map of metadata
	Meta map[string]any
}

// generator applies template transformation on a set of files define in rules
type generator struct {
	Template         *template.Template
	System           *model.System
	UserFeatures     []string // features defined by user
	ComputedFeatures map[string]bool
	UserForce        bool // force overwrite
	TemplatesDir     string
	OutputDir        string
	DryRun           bool
	Stats            GeneratorStats
	Output           OutputWriter
	Meta             map[string]any
}

func New(o GeneratorOptions) (*generator, error) {
	if o.Output == nil {
		o.Output = &FileOutput{}
	}
	if o.System == nil {
		return nil, fmt.Errorf("system is required")
	}
	if len(o.UserFeatures) == 0 {
		o.UserFeatures = []string{"all"}
	}
	g := &generator{
		OutputDir:    o.OutputDir,
		Template:     template.New(""),
		UserForce:    o.UserForce,
		System:       o.System,
		TemplatesDir: o.TemplatesDir,
		UserFeatures: o.UserFeatures,
		DryRun:       o.DryRun,
		Output:       o.Output,
		Meta:         o.Meta,
	}
	g.Template.Funcs(filters.PopulateFuncMap())
	err := g.ParseTemplatesDir(o.TemplatesDir)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (g *generator) ParseTemplate(path string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	tplName, err := filepath.Rel(g.TemplatesDir, path)
	if err != nil {
		return err
	}
	_, err = g.Template.New(tplName).Parse(string(b))
	return err
}

func (g *generator) ParseTemplatesDir(dir string) error {
	log.Debug().Msgf("parsing templates dir: %s", dir)
	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// ignore all dirs
		if d.IsDir() {
			return nil
		}
		// ignore files start starting with .
		if strings.HasPrefix(filepath.Base(path), ".") {
			return nil
		}
		if !strings.HasSuffix(filepath.Base(path), ".tpl") {
			return nil
		}
		return g.ParseTemplate(path)
	})
	if err != nil {
		return err
	}
	return nil
}

// ProcessRules processes a set of rules from a rules document
func (g *generator) ProcessRules(doc *spec.RulesDoc) error {
	log.Debug().Msgf("processing rules: %s", doc.Name)
	g.Stats = GeneratorStats{}
	g.Stats.Start()
	defer func() {
		g.Stats.Stop()
	}()
	if g.System == nil {
		return fmt.Errorf("system is nil")
	}
	err := doc.ComputeFeatures(g.UserFeatures)
	if err != nil {
		return err
	}
	g.ComputedFeatures = doc.FeatureNamesMap()
	for _, f := range doc.Features {
		if f.Skip {
			continue
		}
		err := g.processFeature(f)
		if err != nil {
			return err
		}
	}
	return nil
}

// processFeature processes a feature rule
func (g *generator) processFeature(f *spec.FeatureRule) error {
	log.Debug().Msgf("processing feature %s", f.Name)
	// process system
	ctx := model.SystemScope{
		System:   g.System,
		Features: g.ComputedFeatures,
		Meta:     g.Meta,
	}
	scopes := f.FindScopesByMatch(spec.ScopeSystem)
	for _, scope := range scopes {
		err := g.processScope(scope, ctx)
		if err != nil {
			return err
		}
	}
	for _, module := range g.System.Modules {
		// process module
		scopes := f.FindScopesByMatch(spec.ScopeModule)
		ctx := model.ModuleScope{
			System:   g.System,
			Module:   module,
			Features: g.ComputedFeatures,
			Meta:     g.Meta,
		}
		for _, scope := range scopes {
			err := g.processScope(scope, ctx)
			if err != nil {
				return err
			}
		}
		for _, iface := range module.Interfaces {
			// process interface
			ctx := model.InterfaceScope{
				System:    g.System,
				Module:    module,
				Interface: iface,
				Features:  g.ComputedFeatures,
				Meta:      g.Meta,
			}
			scopes := f.FindScopesByMatch(spec.ScopeInterface)
			for _, scope := range scopes {
				err := g.processScope(scope, ctx)
				if err != nil {
					return err
				}
			}
		}
		for _, struct_ := range module.Structs {
			// process struct
			ctx := model.StructScope{
				System:   g.System,
				Module:   module,
				Struct:   struct_,
				Features: g.ComputedFeatures,
				Meta:     g.Meta,
			}
			scopes := f.FindScopesByMatch(spec.ScopeStruct)
			for _, scope := range scopes {
				err := g.processScope(scope, ctx)
				if err != nil {
					return err
				}
			}
		}
		for _, enum := range module.Enums {
			// process enum
			ctx := model.EnumScope{
				System:   g.System,
				Module:   module,
				Enum:     enum,
				Features: g.ComputedFeatures,
				Meta:     g.Meta,
			}
			scopes := f.FindScopesByMatch(spec.ScopeEnum)
			for _, scope := range scopes {
				err := g.processScope(scope, ctx)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// processScope processes a scope rule (e.g. system, modules, ...) with the given context
func (g *generator) processScope(scope *spec.ScopeRule, ctx any) error {
	prefix := scope.Prefix
	for _, doc := range scope.Documents {
		// clean doc target
		if doc.Target == "" {
			doc.Target = doc.Source
		}
		// apply prefix if set
		if prefix != "" {
			doc.Target = prefix + doc.Target
		}
		err := g.processDocument(doc, ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

// processDocument processes a document rule with the given context
func (g *generator) processDocument(doc spec.DocumentRule, ctx any) error {
	log.Debug().Msgf("processing document %s", doc.Source)
	// the source file to render
	var source = filepath.Clean(doc.Source)
	// the docTarget destination file
	var docTarget = filepath.Clean(doc.Target)
	// either user can force an overwrite or the target or the rules document
	force := doc.Force || g.UserForce
	// transform the target name using the context
	target, err := g.RenderString(docTarget, ctx)
	if err != nil {
		return fmt.Errorf("render rules target %s: %s", docTarget, err)
	}
	// TODO: when doc.Raw is set, we should just copy it to the target
	if doc.Raw {
		// copy the source to the target
		err := g.CopyFile(source, target, force)
		if err != nil {
			log.Warn().Msgf("copy file %s to %s: %s", source, target, err)
			return err
		}
	} else {
		// render the source file to the target
		err := g.RenderFile(source, target, ctx, force)
		if err != nil {
			return err
		}
	}
	return nil
}

// Renders a string using the given context
func (g *generator) RenderString(s string, ctx any) (string, error) {
	var buf = bytes.NewBuffer(nil)
	t := template.New("target")
	t.Funcs(filters.PopulateFuncMap())
	_, err := t.Parse(s)
	if err != nil {
		log.Warn().Msgf("render string: %s: %s", s, err)
		return "", err
	}
	err = t.Execute(buf, ctx)
	if err != nil {
		log.Warn().Msgf("exec template %s: %s", s, err)
		return "", err
	}
	return buf.String(), nil
}

func (g *generator) CopyFile(source, target string, force bool) error {
	g.Stats.FilesCopied++
	if g.DryRun {
		log.Info().Msgf("dry run: copying file %s to %s", source, target)
		g.Stats.FilesTouched = append(g.Stats.FilesTouched, target)
		return nil
	}
	source = helper.Join(g.TemplatesDir, source)
	target = helper.Join(g.OutputDir, target)
	return g.Output.Copy(source, target, force)
}

func (g *generator) RenderFile(source, target string, ctx any, force bool) error {
	// var force = doc.Force
	// var transform = doc.Transform
	log.Debug().Msgf("render %s -> %s", source, target)
	// render the template using the context
	buf := bytes.NewBuffer(nil)
	err := g.Template.ExecuteTemplate(buf, source, ctx)
	if err != nil {
		log.Warn().Msgf("exec template %s: %s", source, err)
		return fmt.Errorf("render template %s: %w", source, err)
	}
	// write the file
	log.Debug().Msgf("write %s", target)
	err = g.WriteFile(buf.Bytes(), target, force)
	if err != nil {
		log.Warn().Msgf("write file %s: %s", target, err)
		return fmt.Errorf("write file %s: %w", target, err)
	}
	return nil
}

func (g *generator) WriteFile(input []byte, target string, force bool) error {
	target = helper.Join(g.OutputDir, target)
	if !force {
		same, err := g.Output.Compare(input, target)
		if err != nil {
			return err
		}

		if same {
			g.Stats.FilesSkipped++
			log.Info().Msgf("skipping file %s", target)
			return nil
		}
	}
	log.Debug().Msgf("write file %s", target)
	g.Stats.FilesTouched = append(g.Stats.FilesTouched, target)
	g.Stats.FilesWritten++
	if g.DryRun {
		log.Info().Msgf("dry run: writing file %s", target)
		return nil
	}
	return g.Output.Write(input, target, force)
}
