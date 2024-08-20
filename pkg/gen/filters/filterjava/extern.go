package filterjava

import "github.com/apigear-io/cli/pkg/model"

type JavaExtern struct {
	Package string
	Name    string
	Default string
}

func parseJavaExtern(schema *model.Schema) JavaExtern {
	xe := schema.GetExtern()
	return javaExtern(xe)
}

func javaExtern(xe *model.Extern) JavaExtern {
	ns := xe.Meta.GetString("java.package")
	name := xe.Meta.GetString("java.name")
	dft := xe.Meta.GetString("java.default")
	if name == "" {
		name = xe.Name
	}
	return JavaExtern{
		Package: ns,
		Name:    name,
		Default: dft,
	}
}
