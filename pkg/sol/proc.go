package sol

import (
	"fmt"
	"io/ioutil"
	"objectapi/pkg/gen"
	"objectapi/pkg/idl"
	"objectapi/pkg/logger"
	"objectapi/pkg/model"
	"objectapi/pkg/spec"
	"path"

	"gopkg.in/yaml.v2"
)

var log = logger.Get()

type Processor struct {
	doc     spec.SolutionDoc
	rootDir string
}

func (r *Processor) ProcessFile(file string) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	var sol spec.SolutionDoc
	err = yaml.Unmarshal(data, &sol)
	if err != nil {
		panic(err)
	}
	rootDir := path.Dir(file)
	sol.RootDir = rootDir
	r.rootDir = rootDir
	r.doc = sol
	r.Process(sol)
}

func (p *Processor) Process(sol spec.SolutionDoc) {
	for _, layer := range sol.Layers {
		p.ProcessLayer(layer)
	}
}

func (p *Processor) ProcessLayer(layer spec.Layer) error {
	var templateDir = path.Join(p.rootDir, layer.Template)
	var searchDir = path.Join(templateDir, "templates")
	var rulesFile = path.Join(templateDir, "rules.yaml")
	renderer := gen.NewRenderer(searchDir)
	system := model.NewSystem(layer.Name)
	err := p.ParseInputs(system, layer.Inputs)
	if err != nil {
		return fmt.Errorf("error parsing inputs: %w", err)
	}

	proc := gen.NewProcessor(renderer, p)
	return proc.ProcessFile(rulesFile, system)
}

func (p *Processor) ParseInputs(s *model.System, inputs []string) error {
	idlParser := idl.NewIDLParser(s)
	dataParser := model.NewDataParser(s)
	for _, input := range inputs {
		switch path.Ext(input) {
		case ".yaml", ".yml", ".json":
			err := dataParser.ParseFile(input)
			if err != nil {
				log.Warnf("error parsing data file: %s. skip", err)
			}
		case ".idl":
			err := idlParser.ParseFile(input)
			if err != nil {
				log.Warnf("error parsing idl file: %s. skip", err)
			}
		default:
			log.Warnf("unknown file type %s. skip", input)
		}
	}
	return nil
}

func (p *Processor) WriteFile(fn string, content string) error {
	fmt.Println("write file ", fn)
	return ioutil.WriteFile(fn, []byte(content), 0644)
}

func NewSolutionProcessor() *Processor {
	return &Processor{}
}
