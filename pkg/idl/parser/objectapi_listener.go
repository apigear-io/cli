// Code generated from pkg/idl/parser/ObjectApi.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // ObjectApi

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ObjectApiListener is a complete listener for a parse tree produced by ObjectApiParser.
type ObjectApiListener interface {
	antlr.ParseTreeListener

	// EnterDocumentRule is called when entering the documentRule production.
	EnterDocumentRule(c *DocumentRuleContext)

	// EnterHeaderRule is called when entering the headerRule production.
	EnterHeaderRule(c *HeaderRuleContext)

	// EnterModuleRule is called when entering the moduleRule production.
	EnterModuleRule(c *ModuleRuleContext)

	// EnterImportRule is called when entering the importRule production.
	EnterImportRule(c *ImportRuleContext)

	// EnterDeclarationsRule is called when entering the declarationsRule production.
	EnterDeclarationsRule(c *DeclarationsRuleContext)

	// EnterInterfaceRule is called when entering the interfaceRule production.
	EnterInterfaceRule(c *InterfaceRuleContext)

	// EnterInterfaceMembersRule is called when entering the interfaceMembersRule production.
	EnterInterfaceMembersRule(c *InterfaceMembersRuleContext)

	// EnterPropertyRule is called when entering the propertyRule production.
	EnterPropertyRule(c *PropertyRuleContext)

	// EnterMethodRule is called when entering the methodRule production.
	EnterMethodRule(c *MethodRuleContext)

	// EnterInputRule is called when entering the inputRule production.
	EnterInputRule(c *InputRuleContext)

	// EnterSignalRule is called when entering the signalRule production.
	EnterSignalRule(c *SignalRuleContext)

	// EnterStructRule is called when entering the structRule production.
	EnterStructRule(c *StructRuleContext)

	// EnterStructFieldRule is called when entering the structFieldRule production.
	EnterStructFieldRule(c *StructFieldRuleContext)

	// EnterEnumRule is called when entering the enumRule production.
	EnterEnumRule(c *EnumRuleContext)

	// EnterEnumMemberRule is called when entering the enumMemberRule production.
	EnterEnumMemberRule(c *EnumMemberRuleContext)

	// EnterSchemaRule is called when entering the schemaRule production.
	EnterSchemaRule(c *SchemaRuleContext)

	// EnterPrimitiveSchema is called when entering the primitiveSchema production.
	EnterPrimitiveSchema(c *PrimitiveSchemaContext)

	// EnterReferenceSchema is called when entering the referenceSchema production.
	EnterReferenceSchema(c *ReferenceSchemaContext)

	// EnterArraySchema is called when entering the arraySchema production.
	EnterArraySchema(c *ArraySchemaContext)

	// ExitDocumentRule is called when exiting the documentRule production.
	ExitDocumentRule(c *DocumentRuleContext)

	// ExitHeaderRule is called when exiting the headerRule production.
	ExitHeaderRule(c *HeaderRuleContext)

	// ExitModuleRule is called when exiting the moduleRule production.
	ExitModuleRule(c *ModuleRuleContext)

	// ExitImportRule is called when exiting the importRule production.
	ExitImportRule(c *ImportRuleContext)

	// ExitDeclarationsRule is called when exiting the declarationsRule production.
	ExitDeclarationsRule(c *DeclarationsRuleContext)

	// ExitInterfaceRule is called when exiting the interfaceRule production.
	ExitInterfaceRule(c *InterfaceRuleContext)

	// ExitInterfaceMembersRule is called when exiting the interfaceMembersRule production.
	ExitInterfaceMembersRule(c *InterfaceMembersRuleContext)

	// ExitPropertyRule is called when exiting the propertyRule production.
	ExitPropertyRule(c *PropertyRuleContext)

	// ExitMethodRule is called when exiting the methodRule production.
	ExitMethodRule(c *MethodRuleContext)

	// ExitInputRule is called when exiting the inputRule production.
	ExitInputRule(c *InputRuleContext)

	// ExitSignalRule is called when exiting the signalRule production.
	ExitSignalRule(c *SignalRuleContext)

	// ExitStructRule is called when exiting the structRule production.
	ExitStructRule(c *StructRuleContext)

	// ExitStructFieldRule is called when exiting the structFieldRule production.
	ExitStructFieldRule(c *StructFieldRuleContext)

	// ExitEnumRule is called when exiting the enumRule production.
	ExitEnumRule(c *EnumRuleContext)

	// ExitEnumMemberRule is called when exiting the enumMemberRule production.
	ExitEnumMemberRule(c *EnumMemberRuleContext)

	// ExitSchemaRule is called when exiting the schemaRule production.
	ExitSchemaRule(c *SchemaRuleContext)

	// ExitPrimitiveSchema is called when exiting the primitiveSchema production.
	ExitPrimitiveSchema(c *PrimitiveSchemaContext)

	// ExitReferenceSchema is called when exiting the referenceSchema production.
	ExitReferenceSchema(c *ReferenceSchemaContext)

	// ExitArraySchema is called when exiting the arraySchema production.
	ExitArraySchema(c *ArraySchemaContext)
}
