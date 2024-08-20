package filterjava

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
)

func ToParamString(prefix string, schema *model.Schema, name string) (string, error) {
	if schema.IsArray {
		inner := schema.InnerSchema()
		ret, err := ToReturnString(prefix, &inner)
		if err != nil {
			return "xxx", fmt.Errorf("javaParam inner value error: %s", err)
		}
		return fmt.Sprintf("%s[] %s", ret, name), nil
	}
	switch schema.KindType {
	case model.TypeString:
		return fmt.Sprintf("String %s", name), nil
	case model.TypeInt:
		return fmt.Sprintf("int %s", name), nil
	case model.TypeInt32:
		return fmt.Sprintf("int %s", name), nil
	case model.TypeInt64:
		return fmt.Sprintf("long %s", name), nil
	case model.TypeFloat:
		return fmt.Sprintf("float %s", name), nil
	case model.TypeFloat32:
		return fmt.Sprintf("float %s", name), nil
	case model.TypeFloat64:
		return fmt.Sprintf("double %s", name), nil
	case model.TypeBool:
		return fmt.Sprintf("boolean %s", name), nil
	case model.TypeExtern:
		xe := parseJavaExtern(schema)
		return fmt.Sprintf("%s %s", xe.Name, name), nil
	case model.TypeEnum:
		symbol := schema.GetEnum()
		return fmt.Sprintf("%s%s %s", prefix, symbol.Name, name), nil
	case model.TypeStruct:
		symbol := schema.GetStruct()
		return fmt.Sprintf("%s%s %s", prefix, symbol.Name, name), nil
	case model.TypeInterface:
		symbol := schema.GetInterface()
		return fmt.Sprintf("%s%s %s", prefix, symbol.Name, name), nil
	}
	return "xxx", fmt.Errorf("javaParam unknown schema %s", schema.Dump())
}

func javaParam(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("javaParam node is nil")
	}
	return ToParamString(prefix, &node.Schema, node.Name)
}
