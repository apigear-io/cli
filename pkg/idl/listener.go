package idl

import (
	"fmt"
	"objectapi/pkg/idl/parser"
	"objectapi/pkg/model"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type ObjectApiListener struct {
	antlr.ParseTreeListener
	System       *model.System
	kind         model.Kind
	module       model.Module
	iface        model.Interface
	struct_      model.Struct
	enum         model.Enum
	method       model.Method
	input        model.TypedNode
	signal       model.Signal
	property     model.TypedNode
	field        model.TypedNode
	schema       model.Schema
	runningValue int
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
	fmt.Println("enter header")
	// nothing todo
}

// EnterModuleRule is called when entering the moduleRule production.
func (o *ObjectApiListener) EnterModuleRule(c *parser.ModuleRuleContext) {
	name := c.GetName().GetText()
	version := c.GetVersion().GetText()
	o.module = model.InitModule(name, version)
}

// ExitMethodRule is called when exiting the methodRule production.
func (o *ObjectApiListener) ExitMethodRule(c *parser.MethodRuleContext) {
	o.iface.Methods = append(o.iface.Methods, o.method)
	o.method = model.Method{}
}

// EnterImportRule is called when entering the importRule production.
func (o *ObjectApiListener) EnterImportRule(c *parser.ImportRuleContext) {
	name := c.GetName().GetText()
	version := c.GetVersion().GetText()
	import_ := model.InitImport(name, version)
	o.module.Imports = append(o.module.Imports, import_)
}

// EnterDeclarationsRule is called when entering the declarationsRule production.
func (o *ObjectApiListener) EnterDeclarationsRule(c *parser.DeclarationsRuleContext) {
	// nothing todo
}

// EnterInterfaceRule is called when entering the interfaceRule production.
func (o *ObjectApiListener) EnterInterfaceRule(c *parser.InterfaceRuleContext) {
	o.kind = model.KindInterface
	name := c.GetName().GetText()
	o.iface = model.InitInterface(name)
}

// ExitInterfaceRule is called when exiting the interfaceRule production.
func (o *ObjectApiListener) ExitInterfaceRule(c *parser.InterfaceRuleContext) {
	o.module.Interfaces = append(o.module.Interfaces, o.iface)
	o.iface = model.Interface{}
}

// EnterInterfaceMembersRule is called when entering the interfaceMembersRule production.
func (o *ObjectApiListener) EnterInterfaceMembersRule(c *parser.InterfaceMembersRuleContext) {
	// nothing todo
}

// EnterPropertyRule is called when entering the propertyRule production.
func (o *ObjectApiListener) EnterPropertyRule(c *parser.PropertyRuleContext) {
	name := c.GetName().GetText()
	o.kind = model.KindProperty
	o.property = model.InitTypeNode(name, model.KindProperty)
}

// EnterMethodRule is called when entering the methodRule production.
func (o *ObjectApiListener) EnterMethodRule(c *parser.MethodRuleContext) {
	name := c.GetName().GetText()
	o.kind = model.KindMethod
	o.method = model.InitMethod(name)
}

// EnterInputRule is called when entering the inputRule production.
func (o *ObjectApiListener) EnterInputRule(c *parser.InputRuleContext) {
	name := c.GetName().GetText()
	o.input = model.InitTypeNode(name, model.KindInput)
}

// ExitInputRule is called when exiting the inputRule production.
func (o *ObjectApiListener) ExitInputRule(c *parser.InputRuleContext) {
	o.input.Schema = o.schema
	if !o.method.IsEmpty() {
		o.method.Inputs = append(o.method.Inputs, o.input)
	} else if !o.signal.IsEmpty() {
		o.signal.Inputs = append(o.signal.Inputs, o.input)
	}
}

// EnterSignalRule is called when entering the signalRule production.
func (o *ObjectApiListener) EnterSignalRule(c *parser.SignalRuleContext) {
	name := c.GetName().GetText()
	o.signal = model.NewSignal(name)
}

// ExitSignalRule is called when exiting the signalRule production.
func (o *ObjectApiListener) ExitSignalRule(c *parser.SignalRuleContext) {
	o.iface.Signals = append(o.iface.Signals, o.signal)
	o.signal = model.Signal{}
}

// EnterStructRule is called when entering the structRule production.
func (o *ObjectApiListener) EnterStructRule(c *parser.StructRuleContext) {
	name := c.GetName().GetText()
	o.struct_ = model.NewStruct(name)
}

// ExitStructRule is called when exiting the structRule production.
func (o *ObjectApiListener) ExitStructRule(c *parser.StructRuleContext) {
	o.module.Structs = append(o.module.Structs, o.struct_)
	o.struct_ = model.Struct{}
}

// EnterStructFieldRule is called when entering the structFieldRule production.
func (o *ObjectApiListener) EnterStructFieldRule(c *parser.StructFieldRuleContext) {
	name := c.GetName().GetText()
	o.field = model.InitTypeNode(name, model.KindField)
}

// ExitStructFieldRule is called when exiting the structFieldRule production.
func (o *ObjectApiListener) ExitStructFieldRule(c *parser.StructFieldRuleContext) {
	o.field.Schema = o.schema
	o.struct_.Fields = append(o.struct_.Fields, o.field)
	o.field = model.TypedNode{}
}

// EnterEnumRule is called when entering the enumRule production.
func (o *ObjectApiListener) EnterEnumRule(c *parser.EnumRuleContext) {
	name := c.GetName().GetText()
	o.enum = model.InitEnum(name)
	o.runningValue = 0
}

// ExitEnumRule is called when exiting the enumRule production.
func (o *ObjectApiListener) ExitEnumRule(c *parser.EnumRuleContext) {
	o.module.Enums = append(o.module.Enums, o.enum)
	o.runningValue = 0
	o.enum = model.Enum{}
}

// EnterEnumMemberRule is called when entering the enumMemberRule production.
func (o *ObjectApiListener) EnterEnumMemberRule(c *parser.EnumMemberRuleContext) {
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
	o.schema = model.Schema{}
}

// ExitSchemaRule is called when exiting the schemaRule production.
func (o *ObjectApiListener) ExitSchemaRule(c *parser.SchemaRuleContext) {
	// schema is picked up and cleared by another rule
}

// EnterPrimitiveSchema is called when entering the primitiveSchema production.
func (o *ObjectApiListener) EnterPrimitiveSchema(c *parser.PrimitiveSchemaContext) {
	name := c.GetName().GetText()
	o.schema.Type = name
}

// EnterReferenceSchema is called when entering the referenceSchema production.
func (o *ObjectApiListener) EnterSymbolSchema(c *parser.SymbolSchemaContext) {
	name := c.GetName().GetText()
	o.schema.Type = name
}

// EnterArraySchema is called when entering the arraySchema production.
func (o *ObjectApiListener) EnterArraySchema(c *parser.ArraySchemaContext) {
	o.schema.Type += "[]"
}

// ExitDocumentRule is called when exiting the documentRule production.
func (o *ObjectApiListener) ExitDocumentRule(c *parser.DocumentRuleContext) {
	o.System.Modules = append(o.System.Modules, o.module)
	o.module = model.Module{}
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

// ExitPropertyRule is called when exiting the propertyRule production.
func (o *ObjectApiListener) ExitPropertyRule(c *parser.PropertyRuleContext) {
	o.property.Schema = o.schema
	o.iface.Properties = append(o.iface.Properties, o.property)
}

// ExitPrimitiveSchema is called when exiting the primitiveSchema production.
func (o *ObjectApiListener) ExitPrimitiveSchema(c *parser.PrimitiveSchemaContext) {
	// nothing todo
}

// ExitReferenceSchema is called when exiting the referenceSchema production.
func (o *ObjectApiListener) ExitSymbolSchema(c *parser.SymbolSchemaContext) {
	// nothing todo
}

// ExitArraySchema is called when exiting the arraySchema production.
func (o *ObjectApiListener) ExitArraySchema(c *parser.ArraySchemaContext) {
	// nothing todo
}

func NewObjectApiListener(system *model.System) parser.ObjectApiListener {
	return &ObjectApiListener{
		System: system,
	}
}
