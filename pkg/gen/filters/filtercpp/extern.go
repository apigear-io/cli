package filtercpp

import (
	"github.com/apigear-io/cli/pkg/model"
)

type CppExtern struct {
	NameSpace string
	Include   string
	Name      string
	Library   string
}

func parseCppExtern(schema *model.Schema) CppExtern {
	xe := schema.GetExtern()
	return cppExtern(xe)
}

func cppExtern(xe *model.Extern) CppExtern {
	ns := xe.Meta.GetString("cpp.namespace")
	inc := xe.Meta.GetString("cpp.include")
	lib := xe.Meta.GetString("cpp.library")
	name := xe.Meta.GetString("cpp.name")
	if name == "" {
		name = xe.Name
	}
	return CppExtern{
		NameSpace: ns,
		Include:   inc,
		Name:      name,
		Library:   lib,
	}
}
