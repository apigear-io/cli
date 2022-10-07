package filterpy

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/model"
)

func ToParamString(schema *model.Schema, name string, prefix string) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToParamString schema is nil")
	}
	if schema.IsArray {
		inner := *schema
		inner.IsArray = false
		innerValue, err := ToReturnString(&inner, prefix)
		if err != nil {
			return "", fmt.Errorf("ToParamString inner value error: %s", err)
		}
		return fmt.Sprintf("%s: list[%s]", name, innerValue), nil
	}
	switch schema.KindType {
	case model.TypeString:
		return fmt.Sprintf("%s: str", name), nil
	case model.TypeInt:
		return fmt.Sprintf("%s: int", name), nil
	case model.TypeFloat:
		return fmt.Sprintf("%s: float", name), nil
	case model.TypeBool:
		return fmt.Sprintf("%s: bool", name), nil
	case model.TypeEnum:
		e := schema.Module.LookupEnum(schema.Type)
		if e == nil {
			return "", fmt.Errorf("ToParamString enum %s not found", schema.Type)
		}
		return fmt.Sprintf("%s: %s%s", name, prefix, e.Name), nil
	case model.TypeStruct:
		s := schema.Module.LookupStruct(schema.Type)
		if s == nil {
			return "", fmt.Errorf("ToParamString struct %s not found", schema.Type)
		}
		return fmt.Sprintf("%s: %s%s", name, prefix, s.Name), nil
	case model.TypeInterface:
		i := schema.Module.LookupInterface(schema.Type)
		if i == nil {
			return "", fmt.Errorf("ToParamString interface %s not found", schema.Type)
		}
		return fmt.Sprintf("%s: %s%s", name, prefix, i.Name), nil
	default:
		return "", fmt.Errorf("unknown schema kind type: %s", schema.KindType)
	}
}

func pyParam(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "", fmt.Errorf("called with nil node")
	}
	log.Debug().Msgf("pyParam called with node: %s", node.Name)
	return ToParamString(&node.Schema, node.Name, prefix)
}
