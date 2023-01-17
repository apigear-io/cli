// Code generated from java-escape by ANTLR 4.11.1. DO NOT EDIT.

package parser // ObjectApi

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

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

	// EnterOperationRule is called when entering the operationRule production.
	EnterOperationRule(c *OperationRuleContext)

	// EnterOperationReturnRule is called when entering the operationReturnRule production.
	EnterOperationReturnRule(c *OperationReturnRuleContext)

	// EnterOperationParamRule is called when entering the operationParamRule production.
	EnterOperationParamRule(c *OperationParamRuleContext)

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

	// EnterArrayRule is called when entering the arrayRule production.
	EnterArrayRule(c *ArrayRuleContext)

	// EnterPrimitiveSchema is called when entering the primitiveSchema production.
	EnterPrimitiveSchema(c *PrimitiveSchemaContext)

	// EnterSymbolSchema is called when entering the symbolSchema production.
	EnterSymbolSchema(c *SymbolSchemaContext)

	// EnterMetaRule is called when entering the metaRule production.
	EnterMetaRule(c *MetaRuleContext)

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

	// ExitOperationRule is called when exiting the operationRule production.
	ExitOperationRule(c *OperationRuleContext)

	// ExitOperationReturnRule is called when exiting the operationReturnRule production.
	ExitOperationReturnRule(c *OperationReturnRuleContext)

	// ExitOperationParamRule is called when exiting the operationParamRule production.
	ExitOperationParamRule(c *OperationParamRuleContext)

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

	// ExitArrayRule is called when exiting the arrayRule production.
	ExitArrayRule(c *ArrayRuleContext)

	// ExitPrimitiveSchema is called when exiting the primitiveSchema production.
	ExitPrimitiveSchema(c *PrimitiveSchemaContext)

	// ExitSymbolSchema is called when exiting the symbolSchema production.
	ExitSymbolSchema(c *SymbolSchemaContext)

	// ExitMetaRule is called when exiting the metaRule production.
	ExitMetaRule(c *MetaRuleContext)
}
