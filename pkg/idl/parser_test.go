package idl

import (
	"testing"

	"github.com/apigear-io/cli/pkg/model"

	"github.com/stretchr/testify/assert"
)

func parseModule(t *testing.T, doc string) *model.Module {
	system := model.NewSystem("test")
	parser := NewParser(system)
	assert.NoError(t, parser.ParseString(doc))
	assert.Equal(t, 1, len(system.Modules))
	assert.Len(t, system.Modules, 1)
	module := system.Modules[0]
	assert.NotNil(t, module)
	return module
}

var docEnum = `
module foo 1.0
enum Enum0 {
	Member0 = 0,
	Member1 = 1,
	Member2 = 2,
}
`

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

var docStruct = `
module foo 1.0
struct Struct0 {
	field0: bool
	field1: int
	field2: float
	field3: string
}
`

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

var docIface = `
module foo 1.0
interface Interface0 {
	prop0: bool
	prop1: int
	prop2: float
	prop3: string
}

interface Interface1 {
	operation0(param0: bool): bool
	operation1(param0: bool, param1: int): bool
	operation2(param0: bool, param1: int, param2: float): bool
}

interface Interface2 {
	signal signal0(param0: bool)
	signal signal1(param0: bool, param1: int)
	signal signal2(param0: bool, param1: int, param2: float)
}
`

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

func TestParseInterfaceOperation(t *testing.T) {
	module := parseModule(t, docIface)
	assert.Len(t, module.Interfaces, 3)
	assert.Equal(t, "Interface1", module.Interfaces[1].Name)
	iface := module.Interfaces[1]

	// operation1
	assert.Equal(t, "operation0", iface.Operations[0].Name)
	assert.Equal(t, "param0", iface.Operations[0].Params[0].Name)
	assert.Equal(t, "bool", iface.Operations[0].Params[0].Schema.Type)
	// operation2
	assert.Equal(t, "operation1", iface.Operations[1].Name)
	assert.Equal(t, "param0", iface.Operations[1].Params[0].Name)
	assert.Equal(t, "bool", iface.Operations[1].Params[0].Schema.Type)
	assert.Equal(t, "param1", iface.Operations[1].Params[1].Name)
	assert.Equal(t, "int", iface.Operations[1].Params[1].Schema.Type)
	// operation3
	assert.Equal(t, "operation2", iface.Operations[2].Name)
	assert.Equal(t, "param0", iface.Operations[2].Params[0].Name)
	assert.Equal(t, "bool", iface.Operations[2].Params[0].Schema.Type)
	assert.Equal(t, "param1", iface.Operations[2].Params[1].Name)
	assert.Equal(t, "int", iface.Operations[2].Params[1].Schema.Type)
	assert.Equal(t, "param2", iface.Operations[2].Params[2].Name)
	assert.Equal(t, "float", iface.Operations[2].Params[2].Schema.Type)
}

func TestParseInterfaceSignal(t *testing.T) {
	module := parseModule(t, docIface)
	assert.Len(t, module.Interfaces, 3)
	assert.Equal(t, "Interface2", module.Interfaces[2].Name)
	interface_ := module.Interfaces[2]
	assert.Equal(t, "signal0", interface_.Signals[0].Name)
	assert.Equal(t, "param0", interface_.Signals[0].Params[0].Name)
	assert.Equal(t, "bool", interface_.Signals[0].Params[0].Schema.Type)
	assert.Equal(t, "signal1", interface_.Signals[1].Name)
	assert.Equal(t, "param0", interface_.Signals[1].Params[0].Name)
	assert.Equal(t, "bool", interface_.Signals[1].Params[0].Schema.Type)
	assert.Equal(t, "param1", interface_.Signals[1].Params[1].Name)
	assert.Equal(t, "int", interface_.Signals[1].Params[1].Schema.Type)
}

var docSymbols = `
module foo 1.0
interface Interface0 { }
interface Interface1 {
	prop0: Enum0
	prop1: Struct0
	prop2: Interface0
	operation0(param0: Enum0): Enum0
	operation1(param0: Struct0): Struct0
	operation2(param0: Interface0): Interface0
	signal signal0(param0: Enum0)
	signal signal1(param0: Struct0)
	signal signal2(param0: Interface0)
}

enum Enum0 { }

struct Struct0 { }
`

func TestParseSymbolProperties(t *testing.T) {
	module := parseModule(t, docSymbols)
	assert.Len(t, module.Interfaces, 2)
	assert.Len(t, module.Enums, 1)
	assert.Len(t, module.Structs, 1)
	iface0 := module.Interfaces[0]
	assert.Equal(t, "Interface0", iface0.Name)
	iface1 := module.Interfaces[1]
	assert.Equal(t, "prop0", iface1.Properties[0].Name)
	assert.Equal(t, "Enum0", iface1.Properties[0].Schema.Type)
	assert.Equal(t, "prop1", iface1.Properties[1].Name)
	assert.Equal(t, "Struct0", iface1.Properties[1].Schema.Type)
	assert.Equal(t, "prop2", iface1.Properties[2].Name)
	assert.Equal(t, "Interface0", iface1.Properties[2].Schema.Type)
}

func TestParseSymbolOperations(t *testing.T) {
	module := parseModule(t, docSymbols)
	assert.Len(t, module.Interfaces, 2)
	assert.Len(t, module.Enums, 1)
	assert.Len(t, module.Structs, 1)
	iface0 := module.Interfaces[0]
	assert.Equal(t, "Interface0", iface0.Name)
	iface1 := module.Interfaces[1]
	assert.Equal(t, "operation0", iface1.Operations[0].Name)
	assert.Equal(t, "param0", iface1.Operations[0].Params[0].Name)
	assert.Equal(t, "Enum0", iface1.Operations[0].Params[0].Schema.Type)
	assert.Equal(t, "operation1", iface1.Operations[1].Name)
	assert.Equal(t, "param0", iface1.Operations[1].Params[0].Name)
	assert.Equal(t, "Struct0", iface1.Operations[1].Params[0].Schema.Type)
	assert.Equal(t, "operation2", iface1.Operations[2].Name)
	assert.Equal(t, "param0", iface1.Operations[2].Params[0].Name)
	assert.Equal(t, "Interface0", iface1.Operations[2].Params[0].Schema.Type)
}

func TestParseSymbolSignals(t *testing.T) {
	module := parseModule(t, docSymbols)
	assert.Len(t, module.Interfaces, 2)
	assert.Len(t, module.Enums, 1)
	assert.Len(t, module.Structs, 1)
	iface0 := module.Interfaces[0]
	assert.Equal(t, "Interface0", iface0.Name)
	iface1 := module.Interfaces[1]
	assert.Equal(t, "signal0", iface1.Signals[0].Name)
	assert.Equal(t, "param0", iface1.Signals[0].Params[0].Name)
	assert.Equal(t, "Enum0", iface1.Signals[0].Params[0].Schema.Type)
	assert.Equal(t, "signal1", iface1.Signals[1].Name)
	assert.Equal(t, "param0", iface1.Signals[1].Params[0].Name)
	assert.Equal(t, "Struct0", iface1.Signals[1].Params[0].Schema.Type)
	assert.Equal(t, "signal2", iface1.Signals[2].Name)
	assert.Equal(t, "param0", iface1.Signals[2].Params[0].Name)
	assert.Equal(t, "Interface0", iface1.Signals[2].Params[0].Schema.Type)
}

var docPrimitiveArrays = `
module foo 1.0
interface Interface0 {}
interface Interface1 {
	prop0: bool[]
	prop1: int[]
	prop2: float[]
	prop3: string[]
	operation0(param0: bool[]): bool[]
	operation1(param0: int[]): int[]
	operation2(param0: float[]): float[]
	operation3(param0: string[]): string[]
	signal signal0(param0: bool[])
	signal signal1(param0: int[])
	signal signal2(param0: float[])
	signal signal3(param0: string[])
}
`

func TestPrimitiveArrayProperties(t *testing.T) {
	module := parseModule(t, docPrimitiveArrays)
	assert.Len(t, module.Interfaces, 2)
	assert.Len(t, module.Enums, 0)
	assert.Len(t, module.Structs, 0)
	iface0 := module.Interfaces[0]
	assert.Equal(t, "Interface0", iface0.Name)
	iface1 := module.Interfaces[1]
	assert.Equal(t, "prop0", iface1.Properties[0].Name)
	assert.Equal(t, "bool", iface1.Properties[0].Schema.Type)
	assert.True(t, iface1.Properties[0].Schema.IsArray)
	assert.Equal(t, "prop1", iface1.Properties[1].Name)
	assert.Equal(t, "int", iface1.Properties[1].Schema.Type)
	assert.True(t, iface1.Properties[1].Schema.IsArray)
	assert.Equal(t, "prop2", iface1.Properties[2].Name)
	assert.Equal(t, "float", iface1.Properties[2].Schema.Type)
	assert.True(t, iface1.Properties[2].Schema.IsArray)
	assert.Equal(t, "prop3", iface1.Properties[3].Name)
	assert.Equal(t, "string", iface1.Properties[3].Schema.Type)
	assert.True(t, iface1.Properties[3].Schema.IsArray)
}

func TestPrimitiveArrayOperations(t *testing.T) {
	module := parseModule(t, docPrimitiveArrays)
	assert.Len(t, module.Interfaces, 2)
	assert.Len(t, module.Enums, 0)
	assert.Len(t, module.Structs, 0)
	iface0 := module.Interfaces[0]
	assert.Equal(t, "Interface0", iface0.Name)
	iface1 := module.Interfaces[1]
	assert.Equal(t, "operation0", iface1.Operations[0].Name)
	assert.Equal(t, "param0", iface1.Operations[0].Params[0].Name)
	assert.Equal(t, "bool", iface1.Operations[0].Params[0].Schema.Type)
	assert.True(t, iface1.Operations[0].Params[0].Schema.IsArray)
	assert.Equal(t, "operation1", iface1.Operations[1].Name)
	assert.Equal(t, "param0", iface1.Operations[1].Params[0].Name)
	assert.Equal(t, "int", iface1.Operations[1].Params[0].Schema.Type)
	assert.True(t, iface1.Operations[1].Params[0].Schema.IsArray)
	assert.Equal(t, "operation2", iface1.Operations[2].Name)
	assert.Equal(t, "param0", iface1.Operations[2].Params[0].Name)
	assert.Equal(t, "float", iface1.Operations[2].Params[0].Schema.Type)
	assert.True(t, iface1.Operations[2].Params[0].Schema.IsArray)
	assert.Equal(t, "operation3", iface1.Operations[3].Name)
	assert.Equal(t, "param0", iface1.Operations[3].Params[0].Name)
	assert.Equal(t, "string", iface1.Operations[3].Params[0].Schema.Type)
	assert.True(t, iface1.Operations[3].Params[0].Schema.IsArray)
}

func TestPrimitiveArraySignals(t *testing.T) {
	module := parseModule(t, docPrimitiveArrays)
	assert.Len(t, module.Interfaces, 2)
	assert.Len(t, module.Enums, 0)
	assert.Len(t, module.Structs, 0)
	iface0 := module.Interfaces[0]
	assert.Equal(t, "Interface0", iface0.Name)
	iface1 := module.Interfaces[1]
	signal := iface1.Signals[0]
	input := signal.Params[0]
	assert.Equal(t, "signal0", signal.Name)
	assert.Equal(t, "param0", input.Name)
	assert.Equal(t, "bool", input.Schema.Type)
	assert.True(t, input.Schema.IsArray)
	signal = iface1.Signals[1]
	input = signal.Params[0]
	assert.Equal(t, "signal1", signal.Name)
	assert.Equal(t, "param0", input.Name)
	assert.Equal(t, "int", input.Schema.Type)
	assert.True(t, input.Schema.IsArray)
	signal = iface1.Signals[2]
	input = signal.Params[0]
	assert.Equal(t, "signal2", signal.Name)
	assert.Equal(t, "param0", input.Name)
	assert.Equal(t, "float", input.Schema.Type)
	assert.True(t, input.Schema.IsArray)
	signal = iface1.Signals[3]
	input = signal.Params[0]
	assert.Equal(t, "signal3", signal.Name)
	assert.Equal(t, "param0", input.Name)
	assert.Equal(t, "string", input.Schema.Type)
	assert.True(t, input.Schema.IsArray)
}

var docSymbolArrays = `
module foo 1.0
interface Interface0 { }
interface Interface1 {ƒ
	prop0: Enum0[]
	prop1: Struct0[]
	prop2: Interface0[]
	operation0(param0: Enum0[]): Enum0[]
	operation1(param0: Struct0[]): Struct0[]
	operation2(param0: Interface0[]): Interface0[]
	signal signal0(param0: Enum0[])
	signal signal1(param0: Struct0[])
	signal signal2(param0: Interface0[])
}

enum Enum0 { }

struct Struct0 { }
`

func TestSymbolArrayProperties(t *testing.T) {
	module := parseModule(t, docSymbolArrays)
	assert.Len(t, module.Interfaces, 2)
	assert.Len(t, module.Enums, 1)
	assert.Len(t, module.Structs, 1)
	iface0 := module.Interfaces[0]
	assert.Equal(t, "Interface0", iface0.Name)
	iface1 := module.Interfaces[1]
	prop := iface1.Properties[0]
	assert.Equal(t, "prop0", prop.Name)
	assert.Equal(t, "Enum0", prop.Schema.Type)
	assert.True(t, prop.Schema.IsArray)
	prop = iface1.Properties[1]
	assert.Equal(t, "prop1", prop.Name)
	assert.Equal(t, "Struct0", prop.Schema.Type)
	assert.True(t, prop.Schema.IsArray)
	prop = iface1.Properties[2]
	assert.Equal(t, "prop2", prop.Name)
	assert.Equal(t, "Interface0", prop.Schema.Type)
	assert.True(t, prop.Schema.IsArray)
}

func TestSymbolArrayOperations(t *testing.T) {
	module := parseModule(t, docSymbolArrays)
	assert.Len(t, module.Interfaces, 2)
	assert.Len(t, module.Enums, 1)
	assert.Len(t, module.Structs, 1)
	iface0 := module.Interfaces[0]
	assert.Equal(t, "Interface0", iface0.Name)
	iface1 := module.Interfaces[1]
	op := iface1.Operations[0]
	input := op.Params[0]
	assert.Equal(t, "operation0", op.Name)
	assert.Equal(t, "param0", input.Name)
	assert.Equal(t, "Enum0", input.Schema.Type)
	assert.True(t, input.Schema.IsArray)
	op = iface1.Operations[1]
	input = op.Params[0]
	assert.Equal(t, "operation1", op.Name)
	assert.Equal(t, "param0", input.Name)
	assert.Equal(t, "Struct0", input.Schema.Type)
	assert.True(t, input.Schema.IsArray)
	op = iface1.Operations[2]
	input = op.Params[0]
	assert.Equal(t, "operation2", op.Name)
	assert.Equal(t, "param0", input.Name)
	assert.Equal(t, "Interface0", input.Schema.Type)
	assert.True(t, input.Schema.IsArray)
}

func TestSymbolArraySignals(t *testing.T) {
	module := parseModule(t, docSymbolArrays)
	assert.Len(t, module.Interfaces, 2)
	assert.Len(t, module.Enums, 1)
	assert.Len(t, module.Structs, 1)
	iface0 := module.Interfaces[0]
	assert.Equal(t, "Interface0", iface0.Name)
	iface1 := module.Interfaces[1]
	assert.Equal(t, "signal0", iface1.Signals[0].Name)
	assert.Equal(t, "param0", iface1.Signals[0].Params[0].Name)
	assert.Equal(t, "Enum0", iface1.Signals[0].Params[0].Schema.Type)
	assert.True(t, iface1.Signals[0].Params[0].Schema.IsArray)
	assert.Equal(t, "signal1", iface1.Signals[1].Name)
	assert.Equal(t, "param0", iface1.Signals[1].Params[0].Name)
	assert.Equal(t, "Struct0", iface1.Signals[1].Params[0].Schema.Type)
	assert.True(t, iface1.Signals[2].Params[0].Schema.IsArray)
	assert.Equal(t, "signal2", iface1.Signals[2].Name)
	assert.Equal(t, "param0", iface1.Signals[2].Params[0].Name)
	assert.Equal(t, "Interface0", iface1.Signals[2].Params[0].Schema.Type)
	assert.True(t, iface1.Signals[2].Params[0].Schema.IsArray)
}
