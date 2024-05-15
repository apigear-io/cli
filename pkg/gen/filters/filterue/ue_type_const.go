package filterue

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
	"github.com/ettle/strcase"
)

func ToConstTypeString(prefix string, schema *model.Schema) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToReturnString schema is nil")
	}
	moduleId := strcase.ToPascal(schema.Module.Name)
	if schema.Import != "" {
		moduleId = strcase.ToPascal(schema.Import)
	}
	var text string
	switch schema.KindType {
	case model.TypeString:
		text = "const FString&"
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
		text = fmt.Sprintf("const %sF%s%s&", prefix, moduleId, schema.Type)
	case model.TypeExtern:
		text = fmt.Sprintf("const %s&", ueExtern(schema.GetExtern()).Name)
	case model.TypeInterface:
		text = fmt.Sprintf("%sF%s%s*", prefix, moduleId, schema.Type)
	default:
		return "xxx", fmt.Errorf("ueConstType unknown schema %s", schema.Dump())
	}
	if schema.IsArray {
		switch schema.KindType {
		case model.TypeString:
			text = "const TArray<FString>&"
		case model.TypeInt:
			text = "const TArray<int32>&"
		case model.TypeInt32:
			text = "const TArray<int32>&"
		case model.TypeInt64:
			text = "const TArray<int64>&"
		case model.TypeFloat:
			text = "const TArray<float>&"
		case model.TypeFloat32:
			text = "const TArray<float>&"
		case model.TypeFloat64:
			text = "const TArray<double>&"
		case model.TypeBool:
			text = "const TArray<bool>&"
		case model.TypeVoid:
			text = "const TArray<void>&"
		case model.TypeEnum:
			text = fmt.Sprintf("const TArray<%sE%s%s>&", prefix, moduleId, schema.Type)
		case model.TypeStruct:
			text = fmt.Sprintf("const TArray<%sF%s%s>&", prefix, moduleId, schema.Type)
		case model.TypeExtern:
			text = fmt.Sprintf("const TArray<%s>&", ueExtern(schema.GetExtern()).Name)
		case model.TypeInterface:
			text = fmt.Sprintf("const TArray<%sF%s%s*>&", prefix, moduleId, schema.Type)
		default:
			return "xxx", fmt.Errorf("ueConstType unknown schema %s", schema.Dump())
		}
	}
	return text, nil
}

func ueConstType(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ueConstType node is nil")
	}
	return ToConstTypeString(prefix, &node.Schema)
}
