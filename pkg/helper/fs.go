package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

func Join(elem ...string) string {
	// if last elem isAbs then return it
	if filepath.IsAbs(elem[len(elem)-1]) {
		return elem[len(elem)-1]
	}
	// otherwise join all elem
	return filepath.Join(elem...)
}

// BaseName returns the last element of path.
func BaseName(src string) string {
	return filepath.Base(src)
}

func IsDir(elem string) bool {
	fi, err := os.Stat(elem)
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

func ReadDocument(path string, v any) error {
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

func WriteDocument(path string, v any) error {
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
		result = append(result, Join(path, name))
	}
	return result, nil
}

func ReadYamlFromData(in []byte, out any) error {
	return yaml.Unmarshal(in, out)
}

func YamlToJson(in []byte) ([]byte, error) {
	out := make(map[string]any)
	err := yaml.Unmarshal(in, &out)
	if err != nil {
		return nil, fmt.Errorf("error un marshalling yaml: %w", err)
	}
	return json.Marshal(out)
}

func RemoveDir(dst string) error {
	return os.RemoveAll(dst)
}

func MakeDir(dst string) error {
	return os.MkdirAll(dst, 0755)
}

func Dir(path string) string {
	return filepath.Dir(path)
}

func WriteFile(dst string, data []byte) error {
	return os.WriteFile(dst, data, 0644)
}

func HasExt(file string, exts ...string) bool {
	for _, ext := range exts {
		if strings.HasSuffix(file, ext) {
			return true
		}
	}
	return false
}

func Ext(file string) string {
	return filepath.Ext(file)
}

func ListDir(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Printf("error reading dir: %s", err)
		return
	}
	for _, file := range files {
		log.Printf("file: %s", file.Name())
	}
}

// FallbackDir returns the first dir that exists.
func FallbackDir(name string, dirs ...string) (string, error) {
	for _, dir := range dirs {
		if IsDir(Join(dir, name)) {
			return Join(dir, name), nil
		}
	}
	return "", fmt.Errorf("dir %s not found", name)
}

type ExtFilter func(string) bool

func ExpandFiles(rootDir string, filter ExtFilter, inputs ...string) ([]string, error) {
	result := make([]string, 0)
	for _, input := range inputs {
		input = Join(rootDir, input)
		if IsDir(input) {
			entries, err := os.ReadDir(input)
			if err != nil {
				return nil, err
			}
			for _, entry := range entries {
				if entry.IsDir() {
					continue
				}
				if filter != nil && !filter(entry.Name()) {
					result = append(result, Join(input, entry.Name()))
				}
			}
		} else {
			result = append(result, input)
		}
	}
	return result, nil
}

func ExpandInputs(rootDir string, inputs ...string) ([]string, error) {
	filter := func(s string) bool {
		return HasExt(s, "module.yaml", "module.yml", "module.json", ".idl")
	}
	return ExpandFiles(rootDir, filter, inputs...)
}
