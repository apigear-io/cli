package filterjs

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/objmodel"
)

func ToParamString(schema *objmodel.Schema, name string, prefix string) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("jsParam schema is nil")
	}
	if schema.IsArray {
		return name, nil
	}
	switch schema.KindType {
	case objmodel.TypeString:
		return name, nil
	case objmodel.TypeInt, objmodel.TypeInt32, objmodel.TypeInt64:
		return name, nil
	case objmodel.TypeFloat, objmodel.TypeFloat32, objmodel.TypeFloat64:
		return name, nil
	case objmodel.TypeBool:
		return name, nil
	case objmodel.TypeEnum:
		e := schema.LookupEnum(schema.Import, schema.Type)
		if e == nil {
			return "xxx", fmt.Errorf("jsParam enum not found: %s", schema.Dump())
		}
		return name, nil
	case objmodel.TypeStruct:
		s := schema.LookupStruct(schema.Import, schema.Type)
		if s == nil {
			return "xxx", fmt.Errorf("jsParam struct not found: %s", schema.Dump())
		}
		return name, nil
	case objmodel.TypeInterface:
		i := schema.LookupInterface(schema.Import, schema.Type)
		if i == nil {
			return "xxx", fmt.Errorf("jsParam interface not found: %s", schema.Dump())
		}
		return name, nil
	default:
		return "xxx", fmt.Errorf("jsParam unknown schema %s", schema.Dump())
	}
}

func jsParam(prefix string, node *objmodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("jsParam called with nil node")
	}
	return ToParamString(&node.Schema, node.Name, prefix)
}
