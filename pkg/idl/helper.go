package idl

import "github.com/apigear-io/cli/pkg/model"

func LoadIdlFromString(name string, content string) (*model.System, error) {
	system := model.NewSystem(name)
	parser := NewParser(system)
	err := parser.ParseString(content)
	if err != nil {
		return nil, err
	}
	return system, nil
}

func LoadIdlFromFiles(name string, files []string) (*model.System, error) {
	system := model.NewSystem(name)
	for _, file := range files {
		parser := NewParser(system)
		err := parser.ParseFile(file)
		if err != nil {
			return nil, err
		}
	}
	return system, nil
}
