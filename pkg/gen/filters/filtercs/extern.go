package filtercs

import "github.com/apigear-io/cli/pkg/model"

// CsExtern holds the C# binding information for an external type, read from
// the extern's "cs.*" meta keys.
type CsExtern struct {
	Namespace string
	Name      string
	Default   string
	Using     string
	Package   string
	Version   string
}

func parseCsExtern(schema *model.Schema) CsExtern {
	xe := schema.GetExtern()
	return csExtern(xe)
}

// MakeCsExtern exposes the parsed extern info to templates via a schema.
func MakeCsExtern(schema *model.Schema) CsExtern {
	return parseCsExtern(schema)
}

func csExtern(xe *model.Extern) CsExtern {
	ns := xe.Meta.GetString("cs.namespace")
	name := xe.Meta.GetString("cs.name")
	dft := xe.Meta.GetString("cs.default")
	using := xe.Meta.GetString("cs.using")
	pkg := xe.Meta.GetString("cs.package")
	ver := xe.Meta.GetString("cs.version")
	if name == "" {
		name = xe.Name
	}
	return CsExtern{
		Namespace: ns,
		Name:      name,
		Default:   dft,
		Using:     using,
		Package:   pkg,
		Version:   ver,
	}
}

func csExterns(externs []*model.Extern) []CsExtern {
	items := []CsExtern{}
	for _, ex := range externs {
		items = append(items, csExtern(ex))
	}
	return items
}
