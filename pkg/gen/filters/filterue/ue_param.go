package filterue

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
	"github.com/ettle/strcase"
)

func ToParamString(schema *model.Schema, name string, prefix string) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ToParamString schema is nil")
	}
	name = strcase.ToPascal(name)
	moduleId := strcase.ToPascal(schema.Module.Name)
	t := schema.Type
	if schema.IsArray {
		inner := *schema
		inner.IsArray = false
		innerValue, err := ToReturnString("", &inner)
		if err != nil {
			return "xxx", fmt.Errorf("ToParamString inner value error: %s", err)
		}
		return fmt.Sprintf("const TArray<%s>& %s%s", innerValue, prefix, name), nil
	}
	switch t {
	case "string":
		return fmt.Sprintf("const FString& %s%s", prefix, name), nil
	case "int":
		return fmt.Sprintf("int32 %s%s", prefix, name), nil
	case "float":
		return fmt.Sprintf("float %s%s", prefix, name), nil
	case "bool":
		return fmt.Sprintf("bool b%s%s", prefix, name), nil
	}

	if e := schema.Module.LookupEnum(t); e != nil {
		return fmt.Sprintf("const E%s%s& %s%s", moduleId, e.Name, prefix, name), nil
	} else if s := schema.Module.LookupStruct(t); s != nil {
		return fmt.Sprintf("const F%s%s& %s%s", moduleId, s.Name, prefix, name), nil
	} else if i := schema.Module.LookupInterface(t); i != nil {
		return fmt.Sprintf("F%s%s* %s%s", moduleId, i.Name, prefix, name), nil
	}
	return "xxx", fmt.Errorf("unknown type %s", t)
}

func ueParam(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("goParam called with nil node")
	}
	return ToParamString(&node.Schema, node.GetName(), prefix)
}
