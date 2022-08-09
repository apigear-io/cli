package spec

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/apigear-io/cli/pkg/model"

	"github.com/apigear-io/cli/pkg/idl"

	"github.com/gocarina/gocsv"
	"github.com/xeipuuv/gojsonschema"
)

func CheckFile(file string) (*gojsonschema.Result, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	if filepath.Ext(file) == ".yaml" || filepath.Ext(file) == ".yml" {
		data, err = YamlToJson(data)
		if err != nil {
			return nil, err
		}
	}
	t, err := GetDocumentType(file)
	if err != nil {
		return nil, err
	}
	return LintJsonDoc(t, data)
}

func CheckNdjsonFile(name string) error {
	// read file line by line with scanner
	file, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", name, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// decode each line into an empty interface using json
		var event interface{}
		err := json.NewDecoder(bufio.NewReader(strings.NewReader(line))).Decode(&event)
		if err != nil {
			return fmt.Errorf("failed to decode line %s: %w", line, err)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to read file %s: %w", name, err)
	}
	return nil
}

func CheckCsvFile(name string) error {
	// read file line by line using scanner
	file, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", name, err)
	}
	defer file.Close()
	var data []any
	err = gocsv.UnmarshalFile(file, &data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal file %s: %w", name, err)
	}
	return nil
}

func CheckIdlFile(name string) error {
	s := model.NewSystem("check")
	parser := idl.NewParser(s)
	err := parser.ParseFile(name)
	if err != nil {
		return fmt.Errorf("failed to parse file %s: %w", name, err)
	}
	return s.ResolveAll()
}
