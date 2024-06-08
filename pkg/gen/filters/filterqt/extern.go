package filterqt

import (
	"github.com/apigear-io/cli/pkg/model"
)

type QtExtern struct {
	NameSpace string
	Include   string
	Name      string
	Package   string
	Component string
}

func parseQtExtern(schema *model.Schema) QtExtern {
	xe := schema.GetExtern()
	return qtExtern(xe)
}

func qtExtern(xe *model.Extern) QtExtern {
	ns := xe.Meta.GetString("qt.namespace")
	inc := xe.Meta.GetString("qt.include")
	name := xe.Meta.GetString("qt.type")
	pck := xe.Meta.GetString("qt.package")
	component := xe.Meta.GetString("qt.component")
	if name == "" {
		name = xe.Name
	}
	return QtExtern{
		NameSpace: ns,
		Include:   inc,
		Name:      name,
		Package:   pck,
		Component: component,
	}
}

func qtExterns(externs []*model.Extern) []QtExtern {
	var items = []QtExtern {}
	for _, ex := range externs { 
		items = append(items, qtExtern(ex))
	}
	return items
}
