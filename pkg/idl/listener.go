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
	module       *model.Module
	iface        *model.Interface
	struct_      *model.Struct
	enum         *model.Enum
	method       *model.Method
	input        *model.Input
	signal       *model.Signal
	property     *model.Property
	field        *model.StructField
	schema       *model.Schema
	runningValue int
}

func itMust(b bool, msg string) {
	if !b {
		panic(msg)
	}
}

func (o *ObjectApiListener) VisitTerminal(node antlr.TerminalNode) {
	fmt.Printf("terminal: %s\n", node.GetText())
}

func (o *ObjectApiListener) VisitErrorNode(node antlr.ErrorNode) {
	fmt.Printf("error: %s\n", node.GetText())
}

func (o *ObjectApiListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("enter: %s\n", ctx.GetStart().GetText())
}

func (o *ObjectApiListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf("exit: %s\n", ctx.GetStart().GetText())
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
	itMust(o.module == nil, "module already defined")
	name := c.GetName().GetText()
	version := c.GetVersion().GetText()
	o.module = model.NewModule(name, version)
	o.System.Modules = append(o.System.Modules, o.module)
}

// EnterImportRule is called when entering the importRule production.
func (o *ObjectApiListener) EnterImportRule(c *parser.ImportRuleContext) {
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
	itMust(o.iface == nil, "interface already defined")
	name := c.GetName().GetText()
	o.iface = model.NewInterface(name)
	o.module.Interfaces = append(o.module.Interfaces, o.iface)
}

// EnterInterfaceMembersRule is called when entering the interfaceMembersRule production.
func (o *ObjectApiListener) EnterInterfaceMembersRule(c *parser.InterfaceMembersRuleContext) {
	// nothing todo
}

// EnterPropertyRule is called when entering the propertyRule production.
func (o *ObjectApiListener) EnterPropertyRule(c *parser.PropertyRuleContext) {
	itMust(o.property == nil, "property already defined")
	name := c.GetName().GetText()
	o.property = model.NewProperty(name)
	o.iface.Properties = append(o.iface.Properties, o.property)
}

// EnterMethodRule is called when entering the methodRule production.
func (o *ObjectApiListener) EnterMethodRule(c *parser.MethodRuleContext) {
	itMust(o.method == nil, "method already defined")
	name := c.GetName().GetText()
	o.method = model.NewMethod(name)
	o.iface.Methods = append(o.iface.Methods, o.method)
}

// EnterInputRule is called when entering the inputRule production.
func (o *ObjectApiListener) EnterInputRule(c *parser.InputRuleContext) {
	itMust(o.input == nil, "input already defined")
	name := c.GetName().GetText()
	o.input = model.NewMethodInput(name)
	if o.method != nil {
		o.method.Inputs = append(o.method.Inputs, o.input)
	}
	if o.signal != nil {
		o.signal.Inputs = append(o.signal.Inputs, o.input)
	}
}

// EnterSignalRule is called when entering the signalRule production.
func (o *ObjectApiListener) EnterSignalRule(c *parser.SignalRuleContext) {
	itMust(o.signal == nil, "signal already defined")
	name := c.GetName().GetText()
	o.signal = model.NewSignal(name)
	o.iface.Signals = append(o.iface.Signals, o.signal)
}

// EnterStructRule is called when entering the structRule production.
func (o *ObjectApiListener) EnterStructRule(c *parser.StructRuleContext) {
	name := c.GetName().GetText()
	o.struct_ = model.NewStruct(name)
	o.module.Structs = append(o.module.Structs, o.struct_)
}

// EnterStructFieldRule is called when entering the structFieldRule production.
func (o *ObjectApiListener) EnterStructFieldRule(c *parser.StructFieldRuleContext) {
	name := c.GetName().GetText()
	o.field = model.NewStructField(name)
	o.struct_.Fields = append(o.struct_.Fields, o.field)
}

// EnterEnumRule is called when entering the enumRule production.
func (o *ObjectApiListener) EnterEnumRule(c *parser.EnumRuleContext) {
	itMust(o.enum == nil, "enum already defined")
	name := c.GetName().GetText()
	o.enum = model.NewEnum(name)
	o.module.Enums = append(o.module.Enums, o.enum)
	o.runningValue = 0
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

// EnterSchemaRule is called when entering the schemaRule production.
func (o *ObjectApiListener) EnterSchemaRule(c *parser.SchemaRuleContext) {
	itMust(o.schema == nil, "schema already defined")
	o.schema = model.NewSchema()
}

// EnterPrimitiveSchema is called when entering the primitiveSchema production.
func (o *ObjectApiListener) EnterPrimitiveSchema(c *parser.PrimitiveSchemaContext) {
	name := c.GetName().GetText()
	if o.schema.Type == "array" {
		o.schema.Items = name
	} else {
		o.schema.Type = name
	}
}

// EnterReferenceSchema is called when entering the referenceSchema production.
func (o *ObjectApiListener) EnterReferenceSchema(c *parser.ReferenceSchemaContext) {
	name := c.GetName().GetText()
	if o.schema.Type == "array" {
		// place type inside items when schema is an array
		o.schema.Items = name
	} else {
		o.schema.Type = name
	}
}

// EnterArraySchema is called when entering the arraySchema production.
func (o *ObjectApiListener) EnterArraySchema(c *parser.ArraySchemaContext) {
	// set schema type to array
	o.schema.Type = "array"
}

// ExitDocumentRule is called when exiting the documentRule production.
func (o *ObjectApiListener) ExitDocumentRule(c *parser.DocumentRuleContext) {
	// nothing todo
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

// ExitInterfaceRule is called when exiting the interfaceRule production.
func (o *ObjectApiListener) ExitInterfaceRule(c *parser.InterfaceRuleContext) {
	o.iface = nil
}

// ExitInterfaceMembersRule is called when exiting the interfaceMembersRule production.
func (o *ObjectApiListener) ExitInterfaceMembersRule(c *parser.InterfaceMembersRuleContext) {
	// nothing todo
}

// ExitPropertyRule is called when exiting the propertyRule production.
func (o *ObjectApiListener) ExitPropertyRule(c *parser.PropertyRuleContext) {
	itMust(o.schema != nil, fmt.Sprintf("property has no schema: %s", o.property.Name))
	o.property.Schema = o.schema
	o.schema = nil
	o.property = nil
}

// ExitMethodRule is called when exiting the methodRule production.
func (o *ObjectApiListener) ExitMethodRule(c *parser.MethodRuleContext) {
	if o.schema != nil {
		o.method.Output = model.NewMethodOutput(o.schema)
	}
	o.schema = nil
	o.method = nil
}

// ExitInputRule is called when exiting the inputRule production.
func (o *ObjectApiListener) ExitInputRule(c *parser.InputRuleContext) {
	itMust(o.schema != nil, fmt.Sprintf("input has no schema (%s)", o.input.Name))
	o.input.Schema = o.schema
	o.schema = nil
	o.input = nil
}

// ExitSignalRule is called when exiting the signalRule production.
func (o *ObjectApiListener) ExitSignalRule(c *parser.SignalRuleContext) {
	o.signal = nil
}

// ExitStructRule is called when exiting the structRule production.
func (o *ObjectApiListener) ExitStructRule(c *parser.StructRuleContext) {
	o.struct_ = nil
}

// ExitStructFieldRule is called when exiting the structFieldRule production.
func (o *ObjectApiListener) ExitStructFieldRule(c *parser.StructFieldRuleContext) {
	itMust(o.schema != nil, fmt.Sprintf("struct field has no schema: %s", o.field.Name))
	o.field.Schema = o.schema
	o.schema = nil
	o.field = nil
}

// ExitEnumRule is called when exiting the enumRule production.
func (o *ObjectApiListener) ExitEnumRule(c *parser.EnumRuleContext) {
	o.enum = nil
	o.runningValue = 0
}

// ExitEnumMemberRule is called when exiting the enumMemberRule production.
func (o *ObjectApiListener) ExitEnumMemberRule(c *parser.EnumMemberRuleContext) {
	// nothing todo
}

// ExitSchemaRule is called when exiting the schemaRule production.
func (o *ObjectApiListener) ExitSchemaRule(c *parser.SchemaRuleContext) {
	// schema is picked up and cleared by another rule
}

// ExitPrimitiveSchema is called when exiting the primitiveSchema production.
func (o *ObjectApiListener) ExitPrimitiveSchema(c *parser.PrimitiveSchemaContext) {
	// nothing todo
}

// ExitReferenceSchema is called when exiting the referenceSchema production.
func (o *ObjectApiListener) ExitReferenceSchema(c *parser.ReferenceSchemaContext) {
	// nothing todo
}

// ExitArraySchema is called when exiting the arraySchema production.
func (o *ObjectApiListener) ExitArraySchema(c *parser.ArraySchemaContext) {
	// nothing todo
}

func NewObjectApiListener(system *model.System) parser.ObjectApiListener {
	if system == nil {
		system = model.NewSystem("system")
	}
	return &ObjectApiListener{
		System: system,
	}
}
