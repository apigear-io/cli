package filterjava

// toBoxedType converts a primitive Java type name to its boxed equivalent.
// Non-primitive types are returned unchanged.
func toBoxedType(primitiveType string) string {
	switch primitiveType {
	case "int":
		return "Integer"
	case "long":
		return "Long"
	case "float":
		return "Float"
	case "double":
		return "Double"
	case "boolean":
		return "Boolean"
	default:
		return primitiveType
	}
}
