package filterue

import (
	"github.com/apigear-io/cli/pkg/model"
)

type UeExtern struct {
	NameSpace string
	Include   string
	Name      string
	Library   string
	Plugin    string
}

func parseUeExtern(schema *model.Schema) UeExtern {
	xe := schema.GetExtern()
	return ueExtern(xe)
}

func ueExtern(xe *model.Extern) UeExtern {
	ns := xe.Meta.GetString("ue.namespace")
	inc := xe.Meta.GetString("ue.include")
	lib := xe.Meta.GetString("ue.module")
	name := xe.Meta.GetString("ue.type")
	plugin := xe.Meta.GetString("ue.plugin")
	if name == "" {
		name = xe.Name
	}
	return UeExtern{
		NameSpace: ns,
		Include:   inc,
		Name:      name,
		Library:   lib,
		Plugin:    plugin,
	}
}
