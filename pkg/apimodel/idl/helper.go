package idl

import "github.com/apigear-io/cli/pkg/apimodel"

func LoadIdlFromString(name string, content string) (*apimodel.System, error) {
	system := apimodel.NewSystem(name)
	parser := NewParser(system)
	err := parser.ParseString(content)
	if err != nil {
		return nil, err
	}
	return system, nil
}

func LoadIdlFromFiles(name string, files []string) (*apimodel.System, error) {
	system := apimodel.NewSystem(name)
	for _, file := range files {
		parser := NewParser(system)
		err := parser.ParseFile(file)
		if err != nil {
			return nil, err
		}
	}
	return system, nil
}
