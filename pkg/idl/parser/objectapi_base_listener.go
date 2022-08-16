// Code generated from pkg/idl/parser/ObjectApi.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // ObjectApi

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseObjectApiListener is a complete listener for a parse tree produced by ObjectApiParser.
type BaseObjectApiListener struct{}

var _ ObjectApiListener = &BaseObjectApiListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseObjectApiListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseObjectApiListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseObjectApiListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseObjectApiListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocumentRule is called when production documentRule is entered.
func (s *BaseObjectApiListener) EnterDocumentRule(ctx *DocumentRuleContext) {}

// ExitDocumentRule is called when production documentRule is exited.
func (s *BaseObjectApiListener) ExitDocumentRule(ctx *DocumentRuleContext) {}

// EnterHeaderRule is called when production headerRule is entered.
func (s *BaseObjectApiListener) EnterHeaderRule(ctx *HeaderRuleContext) {}

// ExitHeaderRule is called when production headerRule is exited.
func (s *BaseObjectApiListener) ExitHeaderRule(ctx *HeaderRuleContext) {}

// EnterModuleRule is called when production moduleRule is entered.
func (s *BaseObjectApiListener) EnterModuleRule(ctx *ModuleRuleContext) {}

// ExitModuleRule is called when production moduleRule is exited.
func (s *BaseObjectApiListener) ExitModuleRule(ctx *ModuleRuleContext) {}

// EnterImportRule is called when production importRule is entered.
func (s *BaseObjectApiListener) EnterImportRule(ctx *ImportRuleContext) {}

// ExitImportRule is called when production importRule is exited.
func (s *BaseObjectApiListener) ExitImportRule(ctx *ImportRuleContext) {}

// EnterDeclarationsRule is called when production declarationsRule is entered.
func (s *BaseObjectApiListener) EnterDeclarationsRule(ctx *DeclarationsRuleContext) {}

// ExitDeclarationsRule is called when production declarationsRule is exited.
func (s *BaseObjectApiListener) ExitDeclarationsRule(ctx *DeclarationsRuleContext) {}

// EnterInterfaceRule is called when production interfaceRule is entered.
func (s *BaseObjectApiListener) EnterInterfaceRule(ctx *InterfaceRuleContext) {}

// ExitInterfaceRule is called when production interfaceRule is exited.
func (s *BaseObjectApiListener) ExitInterfaceRule(ctx *InterfaceRuleContext) {}

// EnterInterfaceMembersRule is called when production interfaceMembersRule is entered.
func (s *BaseObjectApiListener) EnterInterfaceMembersRule(ctx *InterfaceMembersRuleContext) {}

// ExitInterfaceMembersRule is called when production interfaceMembersRule is exited.
func (s *BaseObjectApiListener) ExitInterfaceMembersRule(ctx *InterfaceMembersRuleContext) {}

// EnterPropertyRule is called when production propertyRule is entered.
func (s *BaseObjectApiListener) EnterPropertyRule(ctx *PropertyRuleContext) {}

// ExitPropertyRule is called when production propertyRule is exited.
func (s *BaseObjectApiListener) ExitPropertyRule(ctx *PropertyRuleContext) {}

// EnterOperationRule is called when production operationRule is entered.
func (s *BaseObjectApiListener) EnterOperationRule(ctx *OperationRuleContext) {}

// ExitOperationRule is called when production operationRule is exited.
func (s *BaseObjectApiListener) ExitOperationRule(ctx *OperationRuleContext) {}

// EnterOperationReturnRule is called when production operationReturnRule is entered.
func (s *BaseObjectApiListener) EnterOperationReturnRule(ctx *OperationReturnRuleContext) {}

// ExitOperationReturnRule is called when production operationReturnRule is exited.
func (s *BaseObjectApiListener) ExitOperationReturnRule(ctx *OperationReturnRuleContext) {}

// EnterOperationParamRule is called when production operationParamRule is entered.
func (s *BaseObjectApiListener) EnterOperationParamRule(ctx *OperationParamRuleContext) {}

// ExitOperationParamRule is called when production operationParamRule is exited.
func (s *BaseObjectApiListener) ExitOperationParamRule(ctx *OperationParamRuleContext) {}

// EnterSignalRule is called when production signalRule is entered.
func (s *BaseObjectApiListener) EnterSignalRule(ctx *SignalRuleContext) {}

// ExitSignalRule is called when production signalRule is exited.
func (s *BaseObjectApiListener) ExitSignalRule(ctx *SignalRuleContext) {}

// EnterStructRule is called when production structRule is entered.
func (s *BaseObjectApiListener) EnterStructRule(ctx *StructRuleContext) {}

// ExitStructRule is called when production structRule is exited.
func (s *BaseObjectApiListener) ExitStructRule(ctx *StructRuleContext) {}

// EnterStructFieldRule is called when production structFieldRule is entered.
func (s *BaseObjectApiListener) EnterStructFieldRule(ctx *StructFieldRuleContext) {}

// ExitStructFieldRule is called when production structFieldRule is exited.
func (s *BaseObjectApiListener) ExitStructFieldRule(ctx *StructFieldRuleContext) {}

// EnterEnumRule is called when production enumRule is entered.
func (s *BaseObjectApiListener) EnterEnumRule(ctx *EnumRuleContext) {}

// ExitEnumRule is called when production enumRule is exited.
func (s *BaseObjectApiListener) ExitEnumRule(ctx *EnumRuleContext) {}

// EnterEnumMemberRule is called when production enumMemberRule is entered.
func (s *BaseObjectApiListener) EnterEnumMemberRule(ctx *EnumMemberRuleContext) {}

// ExitEnumMemberRule is called when production enumMemberRule is exited.
func (s *BaseObjectApiListener) ExitEnumMemberRule(ctx *EnumMemberRuleContext) {}

// EnterSchemaRule is called when production schemaRule is entered.
func (s *BaseObjectApiListener) EnterSchemaRule(ctx *SchemaRuleContext) {}

// ExitSchemaRule is called when production schemaRule is exited.
func (s *BaseObjectApiListener) ExitSchemaRule(ctx *SchemaRuleContext) {}

// EnterArrayRule is called when production arrayRule is entered.
func (s *BaseObjectApiListener) EnterArrayRule(ctx *ArrayRuleContext) {}

// ExitArrayRule is called when production arrayRule is exited.
func (s *BaseObjectApiListener) ExitArrayRule(ctx *ArrayRuleContext) {}

// EnterPrimitiveSchema is called when production primitiveSchema is entered.
func (s *BaseObjectApiListener) EnterPrimitiveSchema(ctx *PrimitiveSchemaContext) {}

// ExitPrimitiveSchema is called when production primitiveSchema is exited.
func (s *BaseObjectApiListener) ExitPrimitiveSchema(ctx *PrimitiveSchemaContext) {}

// EnterSymbolSchema is called when production symbolSchema is entered.
func (s *BaseObjectApiListener) EnterSymbolSchema(ctx *SymbolSchemaContext) {}

// ExitSymbolSchema is called when production symbolSchema is exited.
func (s *BaseObjectApiListener) ExitSymbolSchema(ctx *SymbolSchemaContext) {}
