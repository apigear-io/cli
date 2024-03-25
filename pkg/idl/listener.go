package idl

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"

	"github.com/antlr4-go/antlr/v4"
	"github.com/apigear-io/cli/pkg/idl/parser"
	"github.com/apigear-io/cli/pkg/log"
	"github.com/apigear-io/cli/pkg/model"
	"gopkg.in/yaml.v3"
)

type ObjectApiListener struct {
	antlr.ParseTreeListener
	System       *model.System
	kind         model.Kind
	module       *model.Module
	iface        *model.Interface
	struct_      *model.Struct
	enum         *model.Enum
	enumMember   *model.EnumMember
	operation    *model.Operation
	param        *model.TypedNode
	_return      *model.TypedNode
	signal       *model.Signal
	property     *model.TypedNode
	field        *model.TypedNode
	schema       *model.Schema
	runningValue int
}

func IsNil(v any) {
	if reflect.ValueOf(v).IsNil() {
		return
	}
	log.Error().Msgf("isNil: %v should be nil", v)
}

func IsNotNil(v any) {
	if !reflect.ValueOf(v).IsNil() {
		return
	}
	log.Error().Msgf("isNotNil: %v is nil", v)
}

func NewObjectApiListener(system *model.System) parser.ObjectApiListener {
	return &ObjectApiListener{
		System: system,
	}
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
	// nothing todo
}

// EnterHeaderRule is called when entering the headerRule production.
func (o *ObjectApiListener) EnterHeaderRule(c *parser.HeaderRuleContext) {
	// nothing todo
}

// EnterModuleRule is called when entering the moduleRule production.
func (o *ObjectApiListener) EnterModuleRule(c *parser.ModuleRuleContext) {
	IsNil(o.module)
	name := c.GetName().GetText()
	version := "1.0" // default version
	if c.GetVersion() != nil {
		version = c.GetVersion().GetText()
	}
	o.module = model.NewModule(name, version)
	o.module.System = o.System
}

// EnterImportRule is called when entering the importRule production.
func (o *ObjectApiListener) EnterImportRule(c *parser.ImportRuleContext) {
	IsNotNil(o.module)
	name := c.GetName().GetText()
	version := "1.0"
	if c.GetVersion() != nil {
		version = c.GetVersion().GetText()
	}
	import_ := model.NewImport(name, version)
	o.module.Imports = append(o.module.Imports, import_)
}

// EnterDeclarationsRule is called when entering the declarationsRule production.
func (o *ObjectApiListener) EnterDeclarationsRule(c *parser.DeclarationsRuleContext) {
	// nothing todo
}

// EnterInterfaceRule is called when entering the InterfaceRule production.
func (o *ObjectApiListener) EnterInterfaceRule(c *parser.InterfaceRuleContext) {
	IsNotNil(o.module)
	IsNil(o.iface)
	o.kind = model.KindInterface
	name := c.GetName().GetText()

	o.iface = model.NewInterface(name)

	// check if the interface extends another interface
	if c.GetExtends() != nil {
		extends := c.GetExtends().GetText()

		parts := strings.Split(extends, ".")
		// if there is a dot, then the first part is the import and the last part is the type
		if len(parts) > 1 {
			o.iface.Extends.Import = strings.Join(parts[:len(parts)-1], ".")
			o.iface.Extends.Name = parts[len(parts)-1]
		} else {
			o.iface.Extends.Import = ""
			o.iface.Extends.Name = extends
		}
	}

}

// ExitInterfaceRule is called when exiting the interfaceRule production.
func (o *ObjectApiListener) ExitInterfaceRule(c *parser.InterfaceRuleContext) {
	o.parseMeta(&o.iface.NamedNode, c.AllMetaRule())
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
	readOnly := false
	if c.GetReadonly() != nil {
		readOnly = true
	}
	o.kind = model.KindProperty
	o.property = model.NewTypedNode(name, model.KindProperty)
	o.property.IsReadOnly = readOnly
}

// ExitPropertyRule is called when exiting the propertyRule production.
func (o *ObjectApiListener) ExitPropertyRule(c *parser.PropertyRuleContext) {
	o.parseMeta(&o.property.NamedNode, c.AllMetaRule())
	o.property.Schema = *o.schema
	o.schema = nil
	o.iface.Properties = append(o.iface.Properties, o.property)
	o.property = nil
}

// EnterMethodRule is called when entering the methodRule production.
func (o *ObjectApiListener) EnterOperationRule(c *parser.OperationRuleContext) {
	IsNil(o.operation)
	IsNil(o.param)
	IsNil(o._return)
	name := c.GetName().GetText()
	o.kind = model.KindOperation
	o.operation = model.NewOperation(name)
}

// ExitOperationRule is called when exiting the operationRule production.
func (o *ObjectApiListener) ExitOperationRule(c *parser.OperationRuleContext) {
	o.parseMeta(&o.operation.NamedNode, c.AllMetaRule())
	o.operation.Return = o._return
	o.iface.Operations = append(o.iface.Operations, o.operation)
	o.operation = nil
	o.param = nil
	o._return = nil
	o.schema = nil
}

// EnterOperationReturnRule is called when entering the operationReturnRule production.
func (o *ObjectApiListener) EnterOperationReturnRule(c *parser.OperationReturnRuleContext) {
	IsNotNil(o.module)
	IsNotNil(o.iface)
	IsNotNil(o.operation)
	IsNil(o._return)
	IsNil(o.schema)
	o._return = model.NewTypedNode("", model.KindReturn)
}

// ExitOperationReturnRule is called when exiting the operationReturnRule production.
func (o *ObjectApiListener) ExitOperationReturnRule(c *parser.OperationReturnRuleContext) {
	o._return.Schema = *o.schema
	o.operation.Return = o._return
	o.schema = nil
}

// EnterOperationParamRule is called when entering the operationArgRule production.
func (o *ObjectApiListener) EnterOperationParamRule(c *parser.OperationParamRuleContext) {
	IsNotNil(o.module)
	IsNotNil(o.iface)
	IsNil(o.param)
	IsNil(o.schema)
	name := c.GetName().GetText()
	o.param = model.NewTypedNode(name, model.KindParam)
}

// ExitOperationParamRule is called when exiting the operationArgRule production.
func (o *ObjectApiListener) ExitOperationParamRule(c *parser.OperationParamRuleContext) {
	o.param.Schema = *o.schema
	if o.operation != nil {
		o.operation.Params = append(o.operation.Params, o.param)
	} else if o.signal != nil {
		o.signal.Params = append(o.signal.Params, o.param)
	}
	o.param = nil
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
	o.parseMeta(&o.signal.NamedNode, c.AllMetaRule())
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
	o.parseMeta(&o.struct_.NamedNode, c.AllMetaRule())
}

// ExitStructRule is called when exiting the structRule production.
func (o *ObjectApiListener) ExitStructRule(c *parser.StructRuleContext) {
	IsNotNil(o.struct_)
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
	readOnly := false
	if c.GetReadonly() != nil {
		readOnly = true
	}
	o.field = model.NewTypedNode(name, model.KindField)
	o.field.IsReadOnly = readOnly
}

// ExitStructFieldRule is called when exiting the structFieldRule production.
func (o *ObjectApiListener) ExitStructFieldRule(c *parser.StructFieldRuleContext) {
	IsNotNil(o.field)
	o.parseMeta(&o.field.NamedNode, c.AllMetaRule())
	o.field.Schema = *o.schema
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
	o.parseMeta(&o.enum.NamedNode, c.AllMetaRule())
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
	o.enumMember = model.NewEnumMember(name, value)
}

// ExitEnumMemberRule is called when exiting the enumMemberRule production.
func (o *ObjectApiListener) ExitEnumMemberRule(c *parser.EnumMemberRuleContext) {
	IsNotNil(o.enumMember)
	o.parseMeta(&o.enumMember.NamedNode, c.AllMetaRule())
	o.enum.Members = append(o.enum.Members, o.enumMember)
	o.enumMember = nil
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
	parts := strings.Split(name, ".")
	// if there is a dot, then the first part is the import and the last part is the type
	if len(parts) > 1 {
		o.schema.Import = strings.Join(parts[:len(parts)-1], ".")
		o.schema.Type = parts[len(parts)-1]
	} else {
		o.schema.Import = ""
		o.schema.Type = name

	}
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
	o.parseMeta(&o.module.NamedNode, c.AllMetaRule())
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

func (o *ObjectApiListener) EnterMetaRule(c *parser.MetaRuleContext) {
	// nothing todo
}

func (o *ObjectApiListener) ExitMetaRule(c *parser.MetaRuleContext) {
	// nothing todo

}

func (o *ObjectApiListener) parseMeta(node *model.NamedNode, ctxs []parser.IMetaRuleContext) {
	docLines := make([]string, 0)
	tagLines := make([]string, 0)
	ymlStart := 0
	ymlEnd := 0
	file := ""
	if len(ctxs) > 0 { // capture the input source
		file = ctxs[0].GetStart().GetInputStream().GetSourceName()
	}
	for _, ctx := range ctxs {
		if ctx.GetTagLine() != nil {
			if ymlStart == 0 { // update the start line, once
				ymlStart = ctx.GetTagLine().GetLine()
			}
			// update the end line, always
			ymlEnd = ctx.GetTagLine().GetLine()
			text := ctx.GetTagLine().GetText()
			line := strings.TrimSpace(strings.TrimLeft(text, "@"))
			if IsWordOnly(line) {
				line = fmt.Sprintf("%s: true", line)
			}
			tagLines = append(tagLines, line)
		}
		if ctx.GetDocLine() != nil {
			text := ctx.GetDocLine().GetText()
			line := strings.TrimSpace(strings.TrimLeft(text, "/"))
			docLines = append(docLines, line)
		}
	}
	if len(docLines) > 0 {
		node.Description = strings.Join(docLines, "\n")
	}
	if len(tagLines) > 0 {
		yml := strings.Join(tagLines, "\n")
		err := yaml.Unmarshal([]byte(yml), &node.Meta)

		if err != nil {
			log.Warn().Err(err).Msgf("failed to parse meta data in %s:%d-%d", file, ymlStart, ymlEnd)
		}
	}
}

func IsWordOnly(s string) bool {
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			continue
		}
		return false
	}
	return true
}
