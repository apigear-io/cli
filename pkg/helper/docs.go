package helper

import (
	"encoding/json"
	"strings"

	"gopkg.in/yaml.v3"
)

func ParseJson(data []byte, out any) error {
	return json.Unmarshal(data, out)
}

func ParseYaml(data []byte, out any) error {
	return yaml.Unmarshal(data, out)
}

func GetDocumentType(path string) string {
	if strings.HasSuffix(path, ".idl") {
		return "module"
	}
	if strings.HasSuffix(path, ".module.yaml") {
		return "module"
	}
	if strings.HasSuffix(path, ".solution.yaml") {
		return "solution"
	}
	if strings.HasSuffix(path, ".js") {
		return "simulation"
	}
	return "unknown"
}
