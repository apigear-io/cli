package helper

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func IsDir(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func IsFile(path string) bool {
	fi, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !fi.IsDir()
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func ReadDocument(path string, v interface{}) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	ext := filepath.Ext(path)
	switch ext {
	case ".json":
		return json.Unmarshal(data, v)
	case ".yaml", ".yml":
		return yaml.Unmarshal(data, v)
	default:
		return fmt.Errorf("unsupported file extension: %s", ext)
	}
}

func WriteDocument(path string, v interface{}) error {
	ext := filepath.Ext(path)
	switch ext {
	case ".json":
		data, err := json.MarshalIndent(v, "", "  ")
		if err != nil {
			return err
		}
		return os.WriteFile(path, data, 0644)
	case ".yaml", ".yml":
		data, err := yaml.Marshal(v)
		if err != nil {
			return err
		}
		return os.WriteFile(path, data, 0644)
	default:
		return fmt.Errorf("unsupported file extension: %s", ext)
	}
}

func IsDocument(path string) bool {
	ext := filepath.Ext(path)
	switch ext {
	case ".json", ".yaml", ".yml":
		return true
	default:
		return false
	}
}

func FindDocuments(path string) ([]string, error) {
	result := []string{}
	files, err := os.ReadDir(path)
	if err != nil {
		return []string{}, err
	}
	for _, file := range files {
		name := file.Name()
		if file.IsDir() {
			continue
		}
		if !IsDocument(filepath.Ext(name)) {
			continue
		}
		result = append(result, filepath.Join(path, name))
	}
	return result, nil
}

func ReadYamlFromData(in []byte, out interface{}) error {
	return yaml.Unmarshal(in, out)
}

func YamlToJson(in []byte) ([]byte, error) {
	out := make(map[string]interface{})
	err := yaml.Unmarshal(in, &out)
	if err != nil {
		return nil, fmt.Errorf("error un marshalling yaml: %w", err)
	}
	return json.Marshal(out)
}
