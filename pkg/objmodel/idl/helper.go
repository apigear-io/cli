package idl

import "github.com/apigear-io/cli/pkg/objmodel"

func LoadIdlFromString(name string, content string) (*objmodel.System, error) {
	system := objmodel.NewSystem(name)
	parser := NewParser(system)
	err := parser.ParseString(content)
	if err != nil {
		return nil, err
	}
	return system, nil
}

func LoadIdlFromFiles(name string, files []string) (*objmodel.System, error) {
	system := objmodel.NewSystem(name)
	for _, file := range files {
		parser := NewParser(system)
		err := parser.ParseFile(file)
		if err != nil {
			return nil, err
		}
	}
	return system, nil
}
