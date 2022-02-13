package helper

import (
	"encoding/json"

	"gopkg.in/yaml.v2"
)

func ParseJson(data []byte, out interface{}) error {
	return json.Unmarshal(data, out)
}

func ParseYaml(data []byte, out interface{}) error {
	return yaml.Unmarshal(data, out)
}
