package filterts

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/objmodel"
)

func ToParamString(schema *objmodel.Schema, name string, prefix string) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("tsParam schema is nil")
	}
	if schema.IsArray {
		inner := schema.InnerSchema()
		innerValue, err := ToReturnString(&inner, prefix)
		if err != nil {
			return "xxx", fmt.Errorf("tsParam inner value error: %s", err)
		}
		return fmt.Sprintf("%s: %s[]", name, innerValue), nil
	}
	switch schema.KindType {
	case objmodel.TypeString:
		return fmt.Sprintf("%s: string", name), nil
	case objmodel.TypeInt, objmodel.TypeInt32, objmodel.TypeInt64:
		return fmt.Sprintf("%s: number", name), nil
	case objmodel.TypeFloat, objmodel.TypeFloat32, objmodel.TypeFloat64:
		return fmt.Sprintf("%s: number", name), nil
	case objmodel.TypeBool:
		return fmt.Sprintf("%s: boolean", name), nil
	case objmodel.TypeEnum:
		e := schema.LookupEnum(schema.Import, schema.Type)
		if e == nil {
			return "xxx", fmt.Errorf("tsParam enum not found: %s", schema.Dump())
		}
		return fmt.Sprintf("%s: %s%s", name, prefix, e.Name), nil
	case objmodel.TypeStruct:
		s := schema.LookupStruct(schema.Import, schema.Type)
		if s == nil {
			return "xxx", fmt.Errorf("tsParam struct not found: %s", schema.Dump())
		}
		return fmt.Sprintf("%s: %s%s", name, prefix, s.Name), nil
	case objmodel.TypeInterface:
		i := schema.LookupInterface(schema.Import, schema.Type)
		if i == nil {
			return "xxx", fmt.Errorf("tsParam interface not found: %s", schema.Dump())
		}
		return fmt.Sprintf("%s: %s%s", name, prefix, i.Name), nil
	default:
		return "xxx", fmt.Errorf("tsParam unknown schema %s", schema.Dump())
	}
}

func tsParam(prefix string, node *objmodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("tsParam called with nil node")
	}
	return ToParamString(&node.Schema, node.Name, prefix)
}
