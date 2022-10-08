package filterts

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToParamString(schema *model.Schema, name string, prefix string) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ToParamString schema is nil")
	}
	if schema.IsArray {
		inner := *schema
		inner.IsArray = false
		innerValue, err := ToReturnString(&inner, prefix)
		if err != nil {
			return "xxx", fmt.Errorf("ToParamString inner value error: %s", err)
		}
		return fmt.Sprintf("%s: %s[]", name, innerValue), nil
	}
	switch schema.KindType {
	case model.TypeString:
		return fmt.Sprintf("%s: string", name), nil
	case model.TypeInt:
		return fmt.Sprintf("%s: number", name), nil
	case model.TypeFloat:
		return fmt.Sprintf("%s: number", name), nil
	case model.TypeBool:
		return fmt.Sprintf("%s: boolean", name), nil
	case model.TypeEnum:
		e := schema.Module.LookupEnum(schema.Type)
		if e == nil {
			return "xxx", fmt.Errorf("ToParamString enum %s not found", schema.Type)
		}
		return fmt.Sprintf("%s: %s%s", name, prefix, e.Name), nil
	case model.TypeStruct:
		s := schema.Module.LookupStruct(schema.Type)
		if s == nil {
			return "xxx", fmt.Errorf("ToParamString struct %s not found", schema.Type)
		}
		return fmt.Sprintf("%s: %s%s", name, prefix, s.Name), nil
	case model.TypeInterface:
		i := schema.Module.LookupInterface(schema.Type)
		if i == nil {
			return "xxx", fmt.Errorf("ToParamString interface %s not found", schema.Type)
		}
		return fmt.Sprintf("%s: %s%s", name, prefix, i.Name), nil
	default:
		return "xxx", fmt.Errorf("unknown schema kind type: %s", schema.KindType)
	}
}

func tsParam(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("tsParam called with nil node")
	}
	return ToParamString(&node.Schema, node.Name, prefix)
}
