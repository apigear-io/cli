package filterjava

import "github.com/apigear-io/cli/pkg/model"

type JavaExtern struct {
	Package         string
	Name            string
	Default         string
	Version         string
	DownloadPackage string
}

func parseJavaExtern(schema *model.Schema) JavaExtern {
	xe := schema.GetExtern()
	return javaExtern(xe)
}

func MakeJavaExtern(schema *model.Schema) JavaExtern {
	return parseJavaExtern(schema)
}

func javaExtern(xe *model.Extern) JavaExtern {
	ns := xe.Meta.GetString("java.package")
	name := xe.Meta.GetString("java.name")
	dft := xe.Meta.GetString("java.default")
	ver := xe.Meta.GetString("java.version")
	dwnld_pck := xe.Meta.GetString("java.download_package")

	if name == "" {
		name = xe.Name
	}
	return JavaExtern{
		Package:         ns,
		Name:            name,
		Default:         dft,
		Version:         ver,
		DownloadPackage: dwnld_pck,
	}
}
