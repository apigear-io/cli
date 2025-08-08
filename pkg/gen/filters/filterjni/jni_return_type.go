package filterjni

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/model"
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
		text = "jobject"
	case model.TypeStruct:
		text = "jobject"
	case model.TypeExtern:
		text = "jobject"
	case model.TypeInterface:
		text = "jobject"
	default:
		return "xxx", fmt.Errorf("jniToReturnType unknown schema %s", schema.Dump())
	}
	if schema.IsArray {
		if schema.KindType == model.TypeString {
			text = "jobject"
		}
		text = fmt.Sprintf("%sArray", text)
	}
	return text, nil
}

func jniToReturnType(node *model.TypedNode) (string, error) {
	return ToType(&node.Schema)
}
