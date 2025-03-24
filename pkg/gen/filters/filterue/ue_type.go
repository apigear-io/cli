package filterue

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
	"github.com/ettle/strcase"
)

func ToTypeString(prefix string, schema *model.Schema) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ueType schema is nil")
	}
	moduleId := strcase.ToPascal(schema.Module.Name)
	if schema.Import != "" {
		moduleId = strcase.ToPascal(schema.Import)
	}
	var text string
	switch schema.KindType {
	case model.TypeString:
		text = "FString"
	case model.TypeInt:
		text = "int32"
	case model.TypeInt32:
		text = "int32"
	case model.TypeInt64:
		text = "int64"
	case model.TypeFloat:
		text = "float"
	case model.TypeFloat32:
		text = "float"
	case model.TypeFloat64:
		text = "double"
	case model.TypeBool:
		text = "bool"
	case model.TypeVoid:
		text = "void"
	case model.TypeEnum:
		text = fmt.Sprintf("%sE%s%s", prefix, moduleId, schema.Type)
	case model.TypeStruct:
		text = fmt.Sprintf("%sF%s%s", prefix, moduleId, schema.Type)
	case model.TypeExtern:
		text = ueExtern(schema.GetExtern()).Name
	case model.TypeInterface:
		text = fmt.Sprintf("TScriptInterface<%sI%s%sInterface>", prefix, moduleId, schema.Type)
	default:
		return "xxx", fmt.Errorf("ueType unknown schema %s", schema.Dump())
	}
	if schema.IsArray {
		switch schema.KindType {
		case model.TypeString:
			text = "TArray<FString>"
		case model.TypeInt:
			text = "TArray<int32>"
		case model.TypeInt32:
			text = "TArray<int32>"
		case model.TypeInt64:
			text = "TArray<int64>"
		case model.TypeFloat:
			text = "TArray<float>"
		case model.TypeFloat32:
			text = "TArray<float>"
		case model.TypeFloat64:
			text = "TArray<double>"
		case model.TypeBool:
			text = "TArray<bool>"
		case model.TypeEnum:
			text = fmt.Sprintf("TArray<%sE%s%s>", prefix, moduleId, schema.Type)
		case model.TypeStruct:
			text = fmt.Sprintf("TArray<%sF%s%s>", prefix, moduleId, schema.Type)
		case model.TypeExtern:
			text = fmt.Sprintf("TArray<%s>", ueExtern(schema.GetExtern()).Name)
		case model.TypeInterface:
			text = fmt.Sprintf("TArray<TScriptInterface<%sI%s%sInterface>>", prefix, moduleId, schema.Type)
		default:
			return "xxx", fmt.Errorf("ueType unknown array schema %s", schema.Dump())
		}
	}
	return text, nil
}

func ueType(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ueType node is nil")
	}
	return ToTypeString(prefix, &node.Schema)
}
