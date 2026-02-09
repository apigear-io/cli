package filterue

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/objmodel"
	"github.com/ettle/strcase"
)

//TODO: add test including prefix for all filters

func ToReturnString(prefix string, schema *objmodel.Schema) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToReturnString schema is nil")
	}
	moduleId := strcase.ToPascal(schema.Module.Name)
	if schema.Import != "" {
		moduleId = strcase.ToPascal(schema.Import)
	}
	var text string
	switch schema.KindType {
	case objmodel.TypeString:
		text = "FString"
	case objmodel.TypeInt:
		text = "int32"
	case objmodel.TypeInt32:
		text = "int32"
	case objmodel.TypeInt64:
		text = "int64"
	case objmodel.TypeFloat:
		text = "float"
	case objmodel.TypeFloat32:
		text = "float"
	case objmodel.TypeFloat64:
		text = "double"
	case objmodel.TypeBool:
		text = "bool"
	case objmodel.TypeVoid:
		text = "void"
	case objmodel.TypeEnum:
		text = fmt.Sprintf("%sE%s%s", prefix, moduleId, schema.Type)
	case objmodel.TypeStruct:
		text = fmt.Sprintf("%sF%s%s", prefix, moduleId, schema.Type)
	case objmodel.TypeExtern:
		text = ueExtern(schema.GetExtern()).Name
	case objmodel.TypeInterface:
		text = fmt.Sprintf("TScriptInterface<%sI%s%sInterface>", prefix, moduleId, schema.Type)
	default:
		return "xxx", fmt.Errorf("ueReturn unknown schema %s", schema.Dump())
	}
	if schema.IsArray {
		text = fmt.Sprintf("TArray<%s>", text)
	}
	return text, nil
}

func ueReturn(prefix string, node *objmodel.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ueReturn called with nil node")
	}
	return ToReturnString(prefix, &node.Schema)
}
