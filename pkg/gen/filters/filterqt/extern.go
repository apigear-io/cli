package filterqt

import (
	"github.com/apigear-io/cli/pkg/gen/filters/common"
	"github.com/apigear-io/cli/pkg/model"
)

type QtExtern struct {
	NameSpace string
	Include   string
	Name      string
	Library   string
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
	lib := xe.Meta.GetString("qt.library")
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
	  Library:   lib,
	  Package:   pck,
	  Component: component,
	}
}

func qtExterns(ex_list []*model.Extern) []QtExtern {
	var qtExternsList = []QtExtern {}
	for _, element := range ex_list { 
		qtExternsList = append(qtExternsList, qtExtern(element))
	}
	return qtExternsList
}

func qtMakeListOfFields_extern(inputList []QtExtern, fieldName string) ([]string, error){
	return common.MakeListOfFields(inputList, fieldName)
}
