package filterpy

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/gen/filters/common"
	"github.com/apigear-io/cli/pkg/model"
)

func ToParamString(schema *model.Schema, name string, prefix string) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ToParamString schema is nil")
	}
	name = common.SnakeCaseLower(name)
	if schema.IsArray {
		inner := *schema
		inner.IsArray = false
		innerValue, err := ToReturnString(&inner, prefix)
		if err != nil {
			return "xxx", fmt.Errorf("ToParamString inner value error: %s", err)
		}
		return fmt.Sprintf("%s: list[%s]", name, innerValue), nil
	}
	switch schema.KindType {
	case model.TypeString:
		return fmt.Sprintf("%s: str", name), nil
	case model.TypeInt:
		return fmt.Sprintf("%s: int", name), nil
	case model.TypeInt32:
		return fmt.Sprintf("%s: int32", name), nil
	case model.TypeInt64:
		return fmt.Sprintf("%s: int64", name), nil
	case model.TypeFloat:
		return fmt.Sprintf("%s: float", name), nil
	case model.TypeFloat32:
		return fmt.Sprintf("%s: float32", name), nil
	case model.TypeFloat64:
		return fmt.Sprintf("%s: float64", name), nil
	case model.TypeBool:
		return fmt.Sprintf("%s: bool", name), nil
	case model.TypeEnum:
		e := schema.Module.LookupEnum(schema.Type)
		if e == nil {
			return "xxx", fmt.Errorf("ToParamString enum %s not found", schema.Type)
		}
		ident := common.CamelTitleCase(e.Name)
		return fmt.Sprintf("%s: %s%s", name, prefix, ident), nil
	case model.TypeStruct:
		s := schema.Module.LookupStruct(schema.Type)
		if s == nil {
			return "xxx", fmt.Errorf("ToParamString struct %s not found", schema.Type)
		}
		ident := common.CamelTitleCase(s.Name)
		return fmt.Sprintf("%s: %s%s", name, prefix, ident), nil
	case model.TypeInterface:
		i := schema.Module.LookupInterface(schema.Type)
		if i == nil {
			return "xxx", fmt.Errorf("ToParamString interface %s not found", schema.Type)
		}
		ident := common.CamelTitleCase(i.Name)
		return fmt.Sprintf("%s: %s%s", name, prefix, ident), nil
	default:
		return "xxx", fmt.Errorf("unknown schema kind type: %s", schema.KindType)
	}
}

func pyParam(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("called with nil node")
	}
	return ToParamString(&node.Schema, node.Name, prefix)
}
