package sol

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/apigear-io/cli/pkg/helper"
	"github.com/apigear-io/cli/pkg/idl"
	"github.com/apigear-io/cli/pkg/model"
)

// parseInputs parses the inputs from the layer.
// A input can be either a file or a directory.
// If the input is a directory, the files in the directory will be parsed.
func (t *task) parseInputs(s *model.System, inputs []string) error {
	log.Debugf("parse inputs %v", inputs)
	idlParser := idl.NewParser(s)
	dataParser := model.NewDataParser(s)
	files, err := t.expandInputs(t.doc.RootDir, inputs)
	if err != nil {
		log.Infof("error expanding inputs")
		return err
	}
	for _, file := range files {
		log.Debugf("parse input %s", file)
		switch filepath.Ext(file) {
		case ".yaml", ".yml", ".json":
			err := dataParser.ParseFile(file)
			if err != nil {
				log.Warnf("error parsing data file: %s. skip", err)
			}
		case ".idl":
			err := idlParser.ParseFile(file)
			if err != nil {
				log.Warnf("error parsing idl file: %s. skip", err)
			}
		default:
			log.Warnf("unknown file type %s. skip", file)
		}
	}
	err = s.ResolveAll()
	if err != nil {
		return fmt.Errorf("error resolving system: %w", err)
	}
	return nil
}

// ExpandInputs expands the input list to a list of files.
// If input entry is a file it is returned as a list.
// If input entry is a directory, all files in the directory are returned.

func (t *task) expandInputs(rootDir string, inputs []string) ([]string, error) {
	var files []string
	for _, input := range inputs {
		entry := helper.Join(rootDir, input)
		info, err := os.Stat(entry)
		if err != nil {
			log.Infof("error resolving input: %s", entry)
			continue
		}

		if info.IsDir() {
			// add every dir as dependency
			t.deps = append(t.deps, entry)
			err := filepath.WalkDir(entry, func(root string, d fs.DirEntry, err error) error {
				if err != nil {
					return fmt.Errorf("error resolving input: %w", err)
				}
				if d.IsDir() {
					return nil
				}
				if hasExtension(d.Name(), []string{"module.yaml", "module.yml", "module.json", ".odl"}) {
					files = append(files, root)
				}
				return nil
			})
			if err != nil {
				return nil, fmt.Errorf("error resolving input: %w", err)
			}
		} else {
			if hasExtension(entry, []string{"module.yaml", "module.yml", "module.json", ".odl"}) {
				// add every file as dependency
				t.deps = append(t.deps, entry)
				files = append(files, entry)
			}
		}
	}
	return files, nil
}
