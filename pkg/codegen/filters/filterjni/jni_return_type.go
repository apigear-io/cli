package filterjni

import (
	"fmt"

	"github.com/apigear-io/cli/pkg/objmodel"
)

func ToType(schema *objmodel.Schema) (string, error) {
	if schema == nil {
		return "", fmt.Errorf("ToType schema is nil")
	}

	var text string
	switch schema.KindType {
	case objmodel.TypeString:
		text = "jstring"
	case objmodel.TypeInt:
		text = "jint"
	case objmodel.TypeInt32:
		text = "jint"
	case objmodel.TypeInt64:
		text = "jlong"
	case objmodel.TypeFloat:
		text = "jfloat"
	case objmodel.TypeFloat32:
		text = "jfloat"
	case objmodel.TypeFloat64:
		text = "jdouble"
	case objmodel.TypeBool:
		text = "jboolean"
	case objmodel.TypeVoid:
		text = "void"
	// enums are expected to passed as integers
	case objmodel.TypeEnum:
		text = "jobject"
	case objmodel.TypeStruct:
		text = "jobject"
	case objmodel.TypeExtern:
		text = "jobject"
	case objmodel.TypeInterface:
		text = "jobject"
	default:
		return "xxx", fmt.Errorf("jniToReturnType unknown schema %s", schema.Dump())
	}
	if schema.IsArray {
		if schema.KindType == objmodel.TypeString {
			text = "jobject"
		}
		text = fmt.Sprintf("%sArray", text)
	}
	return text, nil
}

func jniToReturnType(node *objmodel.TypedNode) (string, error) {
	return ToType(&node.Schema)
}
