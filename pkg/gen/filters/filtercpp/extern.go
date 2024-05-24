package filtercpp

import (
	"github.com/apigear-io/cli/pkg/model"
)

type CppExtern struct {
	NameSpace    string
	Include      string
	Name         string
	Default      string
	Component    string
	Package      string
	ConanPackage string
	ConanVersion string
}

func parseCppExtern(schema *model.Schema) CppExtern {
	xe := schema.GetExtern()
	return cppExtern(xe)
}

func cppExtern(xe *model.Extern) CppExtern {
	ns := xe.Meta.GetString("cpp.namespace")
	inc := xe.Meta.GetString("cpp.include")
	name := xe.Meta.GetString("cpp.name")
	dft := xe.Meta.GetString("cpp.default")
	cmakePackage := xe.Meta.GetString("cpp.package")
	cmakeComponent := xe.Meta.GetString("cpp.component")
	conanPackage := xe.Meta.GetString("cpp.conanpackage")
	conanVersion := xe.Meta.GetString("cpp.conanversion")
	if name == "" {
		name = xe.Name
	}
	return CppExtern{
		NameSpace:    ns,
		Include:      inc,
		Name:         name,
		Default:      dft,
		Package:      cmakePackage,
		Component:    cmakeComponent,
		ConanPackage: conanPackage,
		ConanVersion: conanVersion,
	}
}
