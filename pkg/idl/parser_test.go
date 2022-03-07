package idl

import (
	"objectapi/pkg/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

var docEnum = `
module foo 1.0
enum Enum0 {
	Member0 = 0,
	Member1 = 1,
	Member2 = 2,
}
`

var docStruct = `
module foo 1.0
struct Struct0 {
	field0: bool
	field1: int
	field2: float
	field3: string
}
`

var docIface = `
module foo 1.0
interface Interface0 {
	prop0: bool
	prop1: int
	prop2: float
	prop3: string
}

interface Interface1 {
	method0(input0: bool): bool
	method1(input0: bool, input1: int): bool
	method2(input0: bool, input1: int, input2: float): bool
}

interface Interface2 {
	signal signal0(input0: bool)
	signal signal1(input0: bool, input1: int)
	signal signal2(input0: bool, input1: int, input2: float)
}
`

func parseModule(t *testing.T, doc string) *model.Module {
	system := model.NewSystem("system")
	parser := NewIDLParser(system)
	parser.ParseString(doc)
	assert.Equal(t, 1, len(system.Modules))
	assert.Len(t, system.Modules, 1)
	module := system.Modules[0]
	assert.NotNil(t, module)
	return module
}

func TestParseEnum(t *testing.T) {
	module := parseModule(t, docEnum)
	assert.Len(t, module.Enums, 1)
	assert.Equal(t, "Enum0", module.Enums[0].Name)
	enum := module.Enums[0]
	assert.Equal(t, "Member0", enum.Members[0].Name)
	assert.Equal(t, 0, enum.Members[0].Value)
	assert.Equal(t, "Member1", enum.Members[1].Name)
	assert.Equal(t, 1, enum.Members[1].Value)
	assert.Equal(t, "Member2", enum.Members[2].Name)
	assert.Equal(t, 2, enum.Members[2].Value)
}

func TestParseStruct(t *testing.T) {
	module := parseModule(t, docStruct)
	assert.Len(t, module.Structs, 1)
	assert.Equal(t, "Struct0", module.Structs[0].Name)
	struct_ := module.Structs[0]
	assert.Equal(t, "field0", struct_.Fields[0].Name)
	assert.Equal(t, "bool", struct_.Fields[0].Schema.Type)
	assert.Equal(t, "field1", struct_.Fields[1].Name)
	assert.Equal(t, "int", struct_.Fields[1].Schema.Type)
	assert.Equal(t, "field2", struct_.Fields[2].Name)
	assert.Equal(t, "float", struct_.Fields[2].Schema.Type)
	assert.Equal(t, "field3", struct_.Fields[3].Name)
	assert.Equal(t, "string", struct_.Fields[3].Schema.Type)
}

func TestParseInterfaceProperties(t *testing.T) {
	module := parseModule(t, docIface)
	assert.Equal(t, "Interface0", module.Interfaces[0].Name)
	interface_ := module.Interfaces[0]
	assert.Equal(t, "prop0", interface_.Properties[0].Name)
	assert.Equal(t, "bool", interface_.Properties[0].Schema.Type)
	assert.Equal(t, "prop1", interface_.Properties[1].Name)
	assert.Equal(t, "int", interface_.Properties[1].Schema.Type)
	assert.Equal(t, "prop2", interface_.Properties[2].Name)
	assert.Equal(t, "float", interface_.Properties[2].Schema.Type)
	assert.Equal(t, "prop3", interface_.Properties[3].Name)
	assert.Equal(t, "string", interface_.Properties[3].Schema.Type)
}

func TestParseInterfaceMethod(t *testing.T) {
	module := parseModule(t, docIface)
	assert.Len(t, module.Interfaces, 3)
	assert.Equal(t, "Interface1", module.Interfaces[1].Name)
	interface_ := module.Interfaces[1]

	// method1
	assert.Equal(t, "method0", interface_.Methods[0].Name)
	assert.Equal(t, "input0", interface_.Methods[0].Inputs[0].Name)
	assert.Equal(t, "bool", interface_.Methods[0].Inputs[0].Schema.Type)
	// method2
	assert.Equal(t, "method1", interface_.Methods[1].Name)
	assert.Equal(t, "input0", interface_.Methods[1].Inputs[0].Name)
	assert.Equal(t, "bool", interface_.Methods[1].Inputs[0].Schema.Type)
	assert.Equal(t, "input1", interface_.Methods[1].Inputs[1].Name)
	assert.Equal(t, "int", interface_.Methods[1].Inputs[1].Schema.Type)
	// method3
	assert.Equal(t, "method2", interface_.Methods[2].Name)
	assert.Equal(t, "input0", interface_.Methods[2].Inputs[0].Name)
	assert.Equal(t, "bool", interface_.Methods[2].Inputs[0].Schema.Type)
	assert.Equal(t, "input1", interface_.Methods[2].Inputs[1].Name)
	assert.Equal(t, "int", interface_.Methods[2].Inputs[1].Schema.Type)
	assert.Equal(t, "input2", interface_.Methods[2].Inputs[2].Name)
	assert.Equal(t, "float", interface_.Methods[2].Inputs[2].Schema.Type)
}

func TestParseInterfaceSignal(t *testing.T) {
	module := parseModule(t, docIface)
	assert.Len(t, module.Interfaces, 3)
	assert.Equal(t, "Interface2", module.Interfaces[2].Name)
	interface_ := module.Interfaces[2]
	assert.Equal(t, "signal0", interface_.Signals[0].Name)
	assert.Equal(t, "input0", interface_.Signals[0].Inputs[0].Name)
	assert.Equal(t, "bool", interface_.Signals[0].Inputs[0].Schema.Type)
	assert.Equal(t, "signal1", interface_.Signals[1].Name)
	assert.Equal(t, "input0", interface_.Signals[1].Inputs[0].Name)
	assert.Equal(t, "bool", interface_.Signals[1].Inputs[0].Schema.Type)
	assert.Equal(t, "input1", interface_.Signals[1].Inputs[1].Name)
	assert.Equal(t, "int", interface_.Signals[1].Inputs[1].Schema.Type)

}
