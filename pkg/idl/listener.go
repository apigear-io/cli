package idl

import (
	"fmt"
	"objectapi/pkg/idl/parser"
	"objectapi/pkg/log"
	"objectapi/pkg/model"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type ObjectApiListener struct {
	antlr.ParseTreeListener
	System       *model.System
	kind         model.Kind
	module       *model.Module
	iface        *model.Interface
	struct_      *model.Struct
	enum         *model.Enum
	method       *model.Method
	input        *model.TypedNode
	output       *model.TypedNode
	signal       *model.Signal
	property     *model.TypedNode
	field        *model.TypedNode
	schema       *model.Schema
	runningValue int
}

func IsNil(v interface{}) {
	if reflect.ValueOf(v).IsNil() {
		return
	}
	log.Fatalf("isNil: %v should be nil", v)
}

func IsNotNil(v interface{}) {
	if !reflect.ValueOf(v).IsNil() {
		return
	}
	log.Fatalf("isNotNil: %v is nil", v)
}

func (o *ObjectApiListener) VisitTerminal(node antlr.TerminalNode) {
	// fmt.Printf("terminal: %s\n", node.GetText())
}

func (o *ObjectApiListener) VisitErrorNode(node antlr.ErrorNode) {
	fmt.Printf("error: %s\n", node.GetText())
}

func (o *ObjectApiListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	// fmt.Printf("enter: %s\n", ctx.GetStart().GetText())
}

func (o *ObjectApiListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	// fmt.Printf("exit: %s\n", ctx.GetStart().GetText())
}

// EnterDocumentRule is called when entering the documentRule production.
func (o *ObjectApiListener) EnterDocumentRule(c *parser.DocumentRuleContext) {
}

// EnterHeaderRule is called when entering the headerRule production.
func (o *ObjectApiListener) EnterHeaderRule(c *parser.HeaderRuleContext) {
	// fmt.Println("enter header")
	// nothing todo
}

// EnterModuleRule is called when entering the moduleRule production.
func (o *ObjectApiListener) EnterModuleRule(c *parser.ModuleRuleContext) {
	IsNil(o.module)
	name := c.GetName().GetText()
	version := c.GetVersion().GetText()
	o.module = model.NewModule(name, version)
}

// EnterImportRule is called when entering the importRule production.
func (o *ObjectApiListener) EnterImportRule(c *parser.ImportRuleContext) {
	IsNotNil(o.module)
	name := c.GetName().GetText()
	version := c.GetVersion().GetText()
	import_ := model.NewImport(name, version)
	o.module.Imports = append(o.module.Imports, import_)
}

// EnterDeclarationsRule is called when entering the declarationsRule production.
func (o *ObjectApiListener) EnterDeclarationsRule(c *parser.DeclarationsRuleContext) {
	// nothing todo
}

// EnterInterfaceRule is called when entering the interfaceRule production.
func (o *ObjectApiListener) EnterInterfaceRule(c *parser.InterfaceRuleContext) {
	IsNotNil(o.module)
	IsNil(o.iface)
	o.kind = model.KindInterface
	name := c.GetName().GetText()
	o.iface = model.NewInterface(name)
}

// ExitInterfaceRule is called when exiting the interfaceRule production.
func (o *ObjectApiListener) ExitInterfaceRule(c *parser.InterfaceRuleContext) {
	o.module.Interfaces = append(o.module.Interfaces, o.iface)
	o.iface = nil
	IsNil(o.iface)
	IsNotNil(o.module)
}

// EnterInterfaceMembersRule is called when entering the interfaceMembersRule production.
func (o *ObjectApiListener) EnterInterfaceMembersRule(c *parser.InterfaceMembersRuleContext) {
	// nothing todo
}

// EnterPropertyRule is called when entering the propertyRule production.
func (o *ObjectApiListener) EnterPropertyRule(c *parser.PropertyRuleContext) {
	IsNotNil(o.iface)
	IsNotNil(o.module)
	IsNil(o.property)
	name := c.GetName().GetText()
	o.kind = model.KindProperty
	o.property = model.NewTypeNode(name, model.KindProperty)
}

// ExitPropertyRule is called when exiting the propertyRule production.
func (o *ObjectApiListener) ExitPropertyRule(c *parser.PropertyRuleContext) {
	o.property.Schema = o.schema
	o.schema = nil
	o.iface.Properties = append(o.iface.Properties, o.property)
	o.property = nil
}

// EnterMethodRule is called when entering the methodRule production.
func (o *ObjectApiListener) EnterMethodRule(c *parser.MethodRuleContext) {
	IsNil(o.method)
	IsNil(o.input)
	IsNil(o.output)
	name := c.GetName().GetText()
	o.kind = model.KindMethod
	o.method = model.NewMethod(name)
}

// ExitMethodRule is called when exiting the methodRule production.
func (o *ObjectApiListener) ExitMethodRule(c *parser.MethodRuleContext) {
	o.method.Output = o.output
	o.iface.Methods = append(o.iface.Methods, o.method)
	o.method = nil
	o.input = nil
	o.output = nil
	o.schema = nil
}

func (o *ObjectApiListener) EnterOutputRule(c *parser.OutputRuleContext) {
	IsNotNil(o.module)
	IsNotNil(o.iface)
	IsNotNil(o.method)
	IsNil(o.output)
	IsNil(o.schema)
	o.output = model.NewTypeNode("", model.KindOutput)
}

func (o *ObjectApiListener) ExitOutputRule(c *parser.OutputRuleContext) {
	o.output.Schema = o.schema
	o.method.Output = o.output
	o.schema = nil
}

// EnterInputRule is called when entering the inputRule production.
func (o *ObjectApiListener) EnterInputRule(c *parser.InputRuleContext) {
	IsNotNil(o.module)
	IsNotNil(o.iface)
	IsNil(o.input)
	IsNil(o.schema)
	name := c.GetName().GetText()
	o.input = model.NewTypeNode(name, model.KindInput)
}

// ExitInputRule is called when exiting the inputRule production.
func (o *ObjectApiListener) ExitInputRule(c *parser.InputRuleContext) {
	o.input.Schema = o.schema
	if o.method != nil {
		o.method.Inputs = append(o.method.Inputs, o.input)
	} else if o.signal != nil {
		o.signal.Inputs = append(o.signal.Inputs, o.input)
	}
	o.input = nil
	o.schema = nil
}

// EnterSignalRule is called when entering the signalRule production.
func (o *ObjectApiListener) EnterSignalRule(c *parser.SignalRuleContext) {
	IsNotNil(o.module)
	IsNotNil(o.iface)
	IsNil(o.signal)
	IsNil(o.schema)
	name := c.GetName().GetText()
	o.signal = model.NewSignal(name)
}

// ExitSignalRule is called when exiting the signalRule production.
func (o *ObjectApiListener) ExitSignalRule(c *parser.SignalRuleContext) {
	o.iface.Signals = append(o.iface.Signals, o.signal)
	o.schema = nil
	o.signal = nil
}

// EnterStructRule is called when entering the structRule production.
func (o *ObjectApiListener) EnterStructRule(c *parser.StructRuleContext) {
	IsNotNil(o.module)
	IsNil(o.struct_)
	IsNil(o.schema)
	name := c.GetName().GetText()
	o.kind = model.KindStruct
	o.struct_ = model.NewStruct(name)
}

// ExitStructRule is called when exiting the structRule production.
func (o *ObjectApiListener) ExitStructRule(c *parser.StructRuleContext) {
	o.module.Structs = append(o.module.Structs, o.struct_)
	o.schema = nil
	o.struct_ = nil
}

// EnterStructFieldRule is called when entering the structFieldRule production.
func (o *ObjectApiListener) EnterStructFieldRule(c *parser.StructFieldRuleContext) {
	IsNotNil(o.struct_)
	IsNil(o.schema)
	IsNil(o.field)
	name := c.GetName().GetText()
	o.field = model.NewTypeNode(name, model.KindField)
}

// ExitStructFieldRule is called when exiting the structFieldRule production.
func (o *ObjectApiListener) ExitStructFieldRule(c *parser.StructFieldRuleContext) {
	o.field.Schema = o.schema
	o.struct_.Fields = append(o.struct_.Fields, o.field)
	o.field = nil
	o.schema = nil
}

// EnterEnumRule is called when entering the enumRule production.
func (o *ObjectApiListener) EnterEnumRule(c *parser.EnumRuleContext) {
	IsNotNil(o.module)
	IsNil(o.enum)
	IsNil(o.schema)
	name := c.GetName().GetText()
	o.enum = model.NewEnum(name)
	o.kind = model.KindEnum
	o.runningValue = 0
}

// ExitEnumRule is called when exiting the enumRule production.
func (o *ObjectApiListener) ExitEnumRule(c *parser.EnumRuleContext) {
	IsNotNil(o.enum)
	o.module.Enums = append(o.module.Enums, o.enum)
	o.runningValue = 0
	o.enum = nil
}

// EnterEnumMemberRule is called when entering the enumMemberRule production.
func (o *ObjectApiListener) EnterEnumMemberRule(c *parser.EnumMemberRuleContext) {
	IsNotNil(o.enum)
	name := c.GetName().GetText()
	var value int
	if c.GetValue() != nil {
		text := c.GetValue().GetText()
		i, err := strconv.Atoi(text)
		if err != nil {
			panic(err)
		}
		value = i
	} else {
		value = o.runningValue
		o.runningValue++
	}
	member := model.NewEnumMember(name, value)
	o.enum.Members = append(o.enum.Members, member)
}

// ExitEnumMemberRule is called when exiting the enumMemberRule production.
func (o *ObjectApiListener) ExitEnumMemberRule(c *parser.EnumMemberRuleContext) {
}

// EnterSchemaRule is called when entering the schemaRule production.
func (o *ObjectApiListener) EnterSchemaRule(c *parser.SchemaRuleContext) {
	IsNil(o.schema)
	o.schema = &model.Schema{}
}

// ExitSchemaRule is called when exiting the schemaRule production.
func (o *ObjectApiListener) ExitSchemaRule(c *parser.SchemaRuleContext) {
	IsNotNil(o.schema)
	IsNotNil(o.module)
	// schema is picked up and cleared by another rule
	o.schema.Module = o.module
}

// EnterPrimitiveSchema is called when entering the primitiveSchema production.
func (o *ObjectApiListener) EnterPrimitiveSchema(c *parser.PrimitiveSchemaContext) {
	IsNotNil(o.schema)
	name := c.GetName().GetText()
	o.schema.Type = name
	o.schema.IsPrimitive = true
}

// EnterReferenceSchema is called when entering the referenceSchema production.
func (o *ObjectApiListener) EnterSymbolSchema(c *parser.SymbolSchemaContext) {
	IsNotNil(o.schema)
	name := c.GetName().GetText()
	o.schema.Type = name
	o.schema.IsSymbol = true
}

// EnterArraySchema is called when entering the arraySchema production.
func (o *ObjectApiListener) EnterArrayRule(c *parser.ArrayRuleContext) {
	IsNotNil(o.schema)
	o.schema.IsArray = true
}

// ExitArraySchema is called when exiting the arraySchema production.
func (o *ObjectApiListener) ExitArrayRule(c *parser.ArrayRuleContext) {
	// nothing todo
}

// ExitDocumentRule is called when exiting the documentRule production.
func (o *ObjectApiListener) ExitDocumentRule(c *parser.DocumentRuleContext) {
	o.System.Modules = append(o.System.Modules, o.module)
	o.module = nil
}

// ExitHeaderRule is called when exiting the headerRule production.
func (o *ObjectApiListener) ExitHeaderRule(c *parser.HeaderRuleContext) {
	// nothing todo
}

// ExitModuleRule is called when exiting the moduleRule production.
func (o *ObjectApiListener) ExitModuleRule(c *parser.ModuleRuleContext) {
	// nothing todo
}

// ExitImportRule is called when exiting the importRule production.
func (o *ObjectApiListener) ExitImportRule(c *parser.ImportRuleContext) {
	// nothing todo
}

// ExitDeclarationsRule is called when exiting the declarationsRule production.
func (o *ObjectApiListener) ExitDeclarationsRule(c *parser.DeclarationsRuleContext) {
	// nothing todo
}

// ExitInterfaceMembersRule is called when exiting the interfaceMembersRule production.
func (o *ObjectApiListener) ExitInterfaceMembersRule(c *parser.InterfaceMembersRuleContext) {
	// nothing todo
}

// ExitPrimitiveSchema is called when exiting the primitiveSchema production.
func (o *ObjectApiListener) ExitPrimitiveSchema(c *parser.PrimitiveSchemaContext) {
	// nothing todo
}

// ExitReferenceSchema is called when exiting the referenceSchema production.
func (o *ObjectApiListener) ExitSymbolSchema(c *parser.SymbolSchemaContext) {
	// nothing todo
}

func NewObjectApiListener(system *model.System) parser.ObjectApiListener {
	return &ObjectApiListener{
		System: system,
	}
}
