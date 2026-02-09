package filterpy

import (
	"github.com/apigear-io/cli/pkg/objmodel"
)

type PyExtern struct {
	Import  string
	Name    string
	Default string
}

func parsePyExtern(schema *objmodel.Schema) PyExtern {
	xe := schema.GetExtern()
	return pyExtern(xe)
}

func pyExtern(xe *objmodel.Extern) PyExtern {
	imp := xe.Meta.GetString("py.import")
	name := xe.Meta.GetString("py.name")
	dft := xe.Meta.GetString("py.default")
	if name == "" {
		name = xe.Name
	}
	return PyExtern{
		Import:  imp,
		Name:    name,
		Default: dft,
	}
}
