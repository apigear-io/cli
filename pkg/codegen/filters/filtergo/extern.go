package filtergo

import (
	"strings"

	"github.com/apigear-io/cli/pkg/apimodel"
)

type GoExtern struct {
	Module string
	Import string
	Name   string
}

func parseGoExtern(schema *apimodel.Schema) GoExtern {
	xe := schema.GetExtern()
	return goExtern(xe)
}

func shortGoImport(name string) string {
	parts := strings.Split(name, "/")
	return parts[len(parts)-1]
}

func goExtern(xe *apimodel.Extern) GoExtern {
	mod := xe.Meta.GetString("go.module")
	imp := shortGoImport(mod)
	name := xe.Meta.GetString("go.name")
	if name == "" {
		name = xe.Name
	}
	return GoExtern{mod, imp, name}
}
