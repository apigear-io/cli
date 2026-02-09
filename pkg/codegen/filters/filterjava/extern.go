package filterjava

import "github.com/apigear-io/cli/pkg/apimodel"

type JavaExtern struct {
	Package         string
	Name            string
	Default         string
	Version         string
	DownloadPackage string
}

func parseJavaExtern(schema *apimodel.Schema) JavaExtern {
	xe := schema.GetExtern()
	return javaExtern(xe)
}

func MakeJavaExtern(schema *apimodel.Schema) JavaExtern {
	return parseJavaExtern(schema)
}

func javaExtern(xe *apimodel.Extern) JavaExtern {
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
