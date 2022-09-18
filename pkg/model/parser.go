package model

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type DataParser struct {
	s *System
}

func (p *DataParser) ParseFile(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("error reading file %s: %s", file, err)
	}
	switch filepath.Ext(file) {
	case ".json":
		return p.ParseJson(data)
	case ".yaml", ".yml":
		return p.ParseYaml(data)
	default:
		return fmt.Errorf("unknown file extension %s", filepath.Ext(file))
	}
}

func (p *DataParser) ParseYaml(data []byte) error {
	var module Module
	err := yaml.Unmarshal(data, &module)
	if err != nil {
		return fmt.Errorf("error parsing yaml string: %s", err)
	}
	p.s.Modules = append(p.s.Modules, &module)
	return nil
}

func (p *DataParser) ParseJson(data []byte) error {
	var module Module
	err := json.Unmarshal(data, &module)
	if err != nil {
		return fmt.Errorf("error parsing json string: %s", err)
	}
	p.s.Modules = append(p.s.Modules, &module)
	return nil
}

func (p *DataParser) LintModule(module *Module) error {
	return nil
}

func NewDataParser(s *System) *DataParser {
	return &DataParser{
		s: s,
	}
}
