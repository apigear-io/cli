package filterue

import (
	"fmt"
	"strings"

	"github.com/apigear-io/cli/pkg/model"
	"github.com/ettle/strcase"
)

func ToType(schema *model.Schema) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToType schema is nil")
	}

	var text string
	switch schema.KindType {
	case model.TypeString:
		text = "jstring"
	case model.TypeInt:
		text = "jint"
	case model.TypeInt32:
		text = "jint"
	case model.TypeInt64:
		text = "jlong"
	case model.TypeFloat:
		text = "jfloat"
	case model.TypeFloat32:
		text = "jfloat"
	case model.TypeFloat64:
		text = "jdouble"
	case model.TypeBool:
		text = "jboolean"
	case model.TypeVoid:
		text = "void"
	// enums are expected to passed as integers
	case model.TypeEnum:
		text = "jint"
	case model.TypeStruct:
		text = "jobject"
	case model.TypeExtern:
		text = "jobject"
	case model.TypeInterface:
		text = "jobject"
	default:
		return "xxx", fmt.Errorf("ueReturn unknown schema %s", schema.Dump())
	}
	if schema.IsArray {
		if schema.KindType == model.TypeString {
			text = "jobject"
		}
		text = fmt.Sprintf("%sArray", text)
	}
	return text, nil
}

func ueToReturnType(node *model.TypedNode) (string, error) {
	return ToType(&node.Schema)
}

func ToUeJniJavaParamString(schema *model.Schema, name string, prefix string) (string, error) {
	if schema == nil {
		return "xxx", fmt.Errorf("ueParam schema is nil")
	}
	name = strcase.ToPascal(name)

	t, err := ToType(schema)
	if err == nil {
		return fmt.Sprintf("%s %s%s", t, prefix, name), nil
	}

	return "xxx", fmt.Errorf("ueJniJavaParam: unknown schema %s", schema.Dump())
}

func ueJniJavaParam(prefix string, node *model.TypedNode) (string, error) {
	if node == nil {
		return "xxx", fmt.Errorf("ueJniJavaParam called with nil node")
	}
	return ToUeJniJavaParamString(&node.Schema, node.Name, prefix)
}

func ueJniJavaParams(prefix string, nodes []*model.TypedNode) (string, error) {
	if nodes == nil {
		return "", fmt.Errorf("ueJniJavaParams called with nil nodes")
	}
	var params []string
	for _, p := range nodes {
		r, err := ToUeJniJavaParamString(&p.Schema, p.Name, prefix)
		if err != nil {
			return "xxx", err
		}
		params = append(params, r)
	}
	return strings.Join(params, ", "), nil
}

func ueJniSignatureType(schema *model.Schema) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToType schema is nil")
	}

	var text string
	switch schema.KindType {
	case model.TypeString:
		text = "Ljava/lang/String;"
	case model.TypeInt:
		text = "I"
	case model.TypeInt32:
		text = "I"
	case model.TypeInt64:
		text = "J"
	case model.TypeFloat:
		text = "F"
	case model.TypeFloat32:
		text = "F"
	case model.TypeFloat64:
		text = "D"
	case model.TypeBool:
		text = "Z"
	case model.TypeVoid:
		text = "V"
	// enums are expected to passed as integers
	case model.TypeEnum:
		text = "I"

	case model.TypeStruct:
		s := schema.LookupStruct(schema.Import, schema.Type)
		if s != nil {
			packageName := ueJavaPath(s.Module.Name, s.Module.Name, "common")
			text = "L" + packageName + "/" + s.Name + ";"
		} else {
			return "xxx", fmt.Errorf("ToSignatureType interface not found %s", schema.Dump())
		}
	case model.TypeExtern:
		return "xxx", fmt.Errorf("ToSignatureType TypeExtern not supported yet %s", schema.Dump())
	case model.TypeInterface:
		i := schema.LookupInterface(schema.Import, schema.Type)
		if i != nil {
			packageName := ueJavaPath(i.Module.Name, i.Module.Name, "common")
			text = "L" + packageName + "/" + i.Name + ";"
		} else {
			return "xxx", fmt.Errorf("ToSignatureType interface not found %s", schema.Dump())
		}
	default:
		return "xxx", fmt.Errorf("ueReturn unknown schema %s", schema.Dump())
	}
	if schema.IsArray {
		text = fmt.Sprintf("[%", text)
	}
	return text, nil
}

func ueJniJavaSignatureParam(node *model.TypedNode) (string, error) {
	if node == nil {
		return "", fmt.Errorf("ueJniJavaParam called with nil nodes")
	}
	return ueJniSignatureType(&node.Schema)
}

func ueJniJavaSignatureParams(nodes []*model.TypedNode) (string, error) {
	if nodes == nil {
		return "", fmt.Errorf("ueJniJavaParams called with nil nodes")
	}
	var text = ""
	for _, p := range nodes {
		r, err := ueJniSignatureType(&p.Schema)
		if err != nil {
			return "xxx", err
		}
		text += r
	}
	return text, nil
}

func ToEnvNameType(schema *model.Schema) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToType schema is nil")
	}

	var text string
	switch schema.KindType {
	case model.TypeString:
		text = "Object"
	case model.TypeInt:
		text = "Int"
	case model.TypeInt32:
		text = "Int"
	case model.TypeInt64:
		text = "Long"
	case model.TypeFloat:
		text = "Float"
	case model.TypeFloat32:
		text = "Float"
	case model.TypeFloat64:
		text = "Double"
	case model.TypeBool:
		text = "Boolean"
	// enums are expected to passed as integers
	case model.TypeEnum:
		text = "Int"
	case model.TypeStruct:
		text = "Object"
	case model.TypeExtern:
		text = "Object"
	case model.TypeInterface:
		text = "Object"
	default:
		return "xxx", fmt.Errorf("ToEnvNameType unknown schema %s", schema.Dump())
	}
	if schema.IsArray {
		text = "Object"
	}
	return text, nil
}

func ueToEnvNameType(node *model.TypedNode) (string, error) {
	return ToEnvNameType(&node.Schema)
}
