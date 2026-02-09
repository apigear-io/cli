package filterqt

import (
	"github.com/apigear-io/cli/pkg/apimodel"
)

type QtExtern struct {
	NameSpace string
	Include   string
	Name      string
	Package   string
	Component string
	Default   string
}

func parseQtExtern(schema *apimodel.Schema) QtExtern {
	xe := schema.GetExtern()
	return qtExtern(xe)
}

func qtExtern(xe *apimodel.Extern) QtExtern {
	ns := xe.Meta.GetString("qt.namespace")
	inc := xe.Meta.GetString("qt.include")
	name := xe.Meta.GetString("qt.type")
	pck := xe.Meta.GetString("qt.package")
	component := xe.Meta.GetString("qt.component")
	dft := xe.Meta.GetString("qt.default")
	if name == "" {
		name = xe.Name
	}
	return QtExtern{
		NameSpace: ns,
		Include:   inc,
		Name:      name,
		Package:   pck,
		Component: component,
		Default:   dft,
	}
}

func qtExterns(externs []*apimodel.Extern) []QtExtern {
	var items = []QtExtern{}
	for _, ex := range externs {
		items = append(items, qtExtern(ex))
	}
	return items
}
