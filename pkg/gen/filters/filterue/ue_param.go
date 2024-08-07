package filterue

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
	"github.com/ettle/strcase"
)

func ToParamString(schema *model.Schema, name string, prefix string) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ueParam schema is nil")
	}
	name = strcase.ToPascal(name)
	moduleId := strcase.ToPascal(schema.Module.Name)
	if schema.Import != "" {
		moduleId = strcase.ToPascal(schema.Import)
	}
	t := schema.Type
	if schema.IsArray {
		inner := schema.InnerSchema()
		innerValue, err := ToReturnString("", &inner)
		if err != nil {
			return "xxx", fmt.Errorf("ueParam inner value error: %s", err)
		}
		return fmt.Sprintf("const TArray<%s>& %s%s", innerValue, prefix, name), nil
	}
	switch t {
	case "string":
		return fmt.Sprintf("const FString& %s%s", prefix, name), nil
	case "int":
		return fmt.Sprintf("int32 %s%s", prefix, name), nil
	case "int32":
		return fmt.Sprintf("int32 %s%s", prefix, name), nil
	case "int64":
		return fmt.Sprintf("int64 %s%s", prefix, name), nil
	case "float":
		return fmt.Sprintf("float %s%s", prefix, name), nil
	case "float32":
		return fmt.Sprintf("float %s%s", prefix, name), nil
	case "float64":
		return fmt.Sprintf("double %s%s", prefix, name), nil
	case "bool":
		return fmt.Sprintf("bool b%s%s", prefix, name), nil
	}

	e := schema.LookupEnum(schema.Import, schema.Type)
	if e != nil {
		return fmt.Sprintf("E%s%s %s%s", moduleId, e.Name, prefix, name), nil
	}
	s := schema.LookupStruct(schema.Import, schema.Type)
	if s != nil {
		return fmt.Sprintf("const F%s%s& %s%s", moduleId, s.Name, prefix, name), nil
	}
	ex := schema.LookupExtern(schema.Import, schema.Type)
	if ex != nil {
		return fmt.Sprintf("const %s& %s%s", ueExtern(schema.GetExtern()).Name, prefix, name), nil
	}
	i := schema.LookupInterface(schema.Import, schema.Type)
	if i != nil {
		return fmt.Sprintf("F%s%s* %s%s", moduleId, i.Name, prefix, name), nil
	}
	return "xxx", fmt.Errorf("ueParam: unknown schema %s", schema.Dump())
}

func ueParam(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ueParam called with nil node")
	}
	return ToParamString(&node.Schema, node.Name, prefix)
}
