// Generated from /Users/jryannel/dev/apigear/cli/pkg/idl/parser/ObjectApi.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link ObjectApiParser}.
 */
public interface ObjectApiListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link ObjectApiParser#documentRule}.
	 * @param ctx the parse tree
	 */
	void enterDocumentRule(ObjectApiParser.DocumentRuleContext ctx);
	/**
	 * Exit a parse tree produced by {@link ObjectApiParser#documentRule}.
	 * @param ctx the parse tree
	 */
	void exitDocumentRule(ObjectApiParser.DocumentRuleContext ctx);
	/**
	 * Enter a parse tree produced by {@link ObjectApiParser#headerRule}.
	 * @param ctx the parse tree
	 */
	void enterHeaderRule(ObjectApiParser.HeaderRuleContext ctx);
	/**
	 * Exit a parse tree produced by {@link ObjectApiParser#headerRule}.
	 * @param ctx the parse tree
	 */
	void exitHeaderRule(ObjectApiParser.HeaderRuleContext ctx);
	/**
	 * Enter a parse tree produced by {@link ObjectApiParser#moduleRule}.
	 * @param ctx the parse tree
	 */
	void enterModuleRule(ObjectApiParser.ModuleRuleContext ctx);
	/**
	 * Exit a parse tree produced by {@link ObjectApiParser#moduleRule}.
	 * @param ctx the parse tree
	 */
	void exitModuleRule(ObjectApiParser.ModuleRuleContext ctx);
	/**
	 * Enter a parse tree produced by {@link ObjectApiParser#importRule}.
	 * @param ctx the parse tree
	 */
	void enterImportRule(ObjectApiParser.ImportRuleContext ctx);
	/**
	 * Exit a parse tree produced by {@link ObjectApiParser#importRule}.
	 * @param ctx the parse tree
	 */
	void exitImportRule(ObjectApiParser.ImportRuleContext ctx);
	/**
	 * Enter a parse tree produced by {@link ObjectApiParser#declarationsRule}.
	 * @param ctx the parse tree
	 */
	void enterDeclarationsRule(ObjectApiParser.DeclarationsRuleContext ctx);
	/**
	 * Exit a parse tree produced by {@link ObjectApiParser#declarationsRule}.
	 * @param ctx the parse tree
	 */
	void exitDeclarationsRule(ObjectApiParser.DeclarationsRuleContext ctx);
	/**
	 * Enter a parse tree produced by {@link ObjectApiParser#interfaceRule}.
	 * @param ctx the parse tree
	 */
	void enterInterfaceRule(ObjectApiParser.InterfaceRuleContext ctx);
	/**
	 * Exit a parse tree produced by {@link ObjectApiParser#interfaceRule}.
	 * @param ctx the parse tree
	 */
	void exitInterfaceRule(ObjectApiParser.InterfaceRuleContext ctx);
	/**
	 * Enter a parse tree produced by {@link ObjectApiParser#interfaceMembersRule}.
	 * @param ctx the parse tree
	 */
	void enterInterfaceMembersRule(ObjectApiParser.InterfaceMembersRuleContext ctx);
	/**
	 * Exit a parse tree produced by {@link ObjectApiParser#interfaceMembersRule}.
	 * @param ctx the parse tree
	 */
	void exitInterfaceMembersRule(ObjectApiParser.InterfaceMembersRuleContext ctx);
	/**
	 * Enter a parse tree produced by {@link ObjectApiParser#propertyRule}.
	 * @param ctx the parse tree
	 */
	void enterPropertyRule(ObjectApiParser.PropertyRuleContext ctx);
	/**
	 * Exit a parse tree produced by {@link ObjectApiParser#propertyRule}.
	 * @param ctx the parse tree
	 */
	void exitPropertyRule(ObjectApiParser.PropertyRuleContext ctx);
	/**
	 * Enter a parse tree produced by {@link ObjectApiParser#operationRule}.
	 * @param ctx the parse tree
	 */
	void enterOperationRule(ObjectApiParser.OperationRuleContext ctx);
	/**
	 * Exit a parse tree produced by {@link ObjectApiParser#operationRule}.
	 * @param ctx the parse tree
	 */
	void exitOperationRule(ObjectApiParser.OperationRuleContext ctx);
	/**
	 * Enter a parse tree produced by {@link ObjectApiParser#operationReturnRule}.
	 * @param ctx the parse tree
	 */
	void enterOperationReturnRule(ObjectApiParser.OperationReturnRuleContext ctx);
	/**
	 * Exit a parse tree produced by {@link ObjectApiParser#operationReturnRule}.
	 * @param ctx the parse tree
	 */
	void exitOperationReturnRule(ObjectApiParser.OperationReturnRuleContext ctx);
	/**
	 * Enter a parse tree produced by {@link ObjectApiParser#operationParamRule}.
	 * @param ctx the parse tree
	 */
	void enterOperationParamRule(ObjectApiParser.OperationParamRuleContext ctx);
	/**
	 * Exit a parse tree produced by {@link ObjectApiParser#operationParamRule}.
	 * @param ctx the parse tree
	 */
	void exitOperationParamRule(ObjectApiParser.OperationParamRuleContext ctx);
	/**
	 * Enter a parse tree produced by {@link ObjectApiParser#signalRule}.
	 * @param ctx the parse tree
	 */
	void enterSignalRule(ObjectApiParser.SignalRuleContext ctx);
	/**
	 * Exit a parse tree produced by {@link ObjectApiParser#signalRule}.
	 * @param ctx the parse tree
	 */
	void exitSignalRule(ObjectApiParser.SignalRuleContext ctx);
	/**
	 * Enter a parse tree produced by {@link ObjectApiParser#structRule}.
	 * @param ctx the parse tree
	 */
	void enterStructRule(ObjectApiParser.StructRuleContext ctx);
	/**
	 * Exit a parse tree produced by {@link ObjectApiParser#structRule}.
	 * @param ctx the parse tree
	 */
	void exitStructRule(ObjectApiParser.StructRuleContext ctx);
	/**
	 * Enter a parse tree produced by {@link ObjectApiParser#structFieldRule}.
	 * @param ctx the parse tree
	 */
	void enterStructFieldRule(ObjectApiParser.StructFieldRuleContext ctx);
	/**
	 * Exit a parse tree produced by {@link ObjectApiParser#structFieldRule}.
	 * @param ctx the parse tree
	 */
	void exitStructFieldRule(ObjectApiParser.StructFieldRuleContext ctx);
	/**
	 * Enter a parse tree produced by {@link ObjectApiParser#enumRule}.
	 * @param ctx the parse tree
	 */
	void enterEnumRule(ObjectApiParser.EnumRuleContext ctx);
	/**
	 * Exit a parse tree produced by {@link ObjectApiParser#enumRule}.
	 * @param ctx the parse tree
	 */
	void exitEnumRule(ObjectApiParser.EnumRuleContext ctx);
	/**
	 * Enter a parse tree produced by {@link ObjectApiParser#enumMemberRule}.
	 * @param ctx the parse tree
	 */
	void enterEnumMemberRule(ObjectApiParser.EnumMemberRuleContext ctx);
	/**
	 * Exit a parse tree produced by {@link ObjectApiParser#enumMemberRule}.
	 * @param ctx the parse tree
	 */
	void exitEnumMemberRule(ObjectApiParser.EnumMemberRuleContext ctx);
	/**
	 * Enter a parse tree produced by {@link ObjectApiParser#schemaRule}.
	 * @param ctx the parse tree
	 */
	void enterSchemaRule(ObjectApiParser.SchemaRuleContext ctx);
	/**
	 * Exit a parse tree produced by {@link ObjectApiParser#schemaRule}.
	 * @param ctx the parse tree
	 */
	void exitSchemaRule(ObjectApiParser.SchemaRuleContext ctx);
	/**
	 * Enter a parse tree produced by {@link ObjectApiParser#arrayRule}.
	 * @param ctx the parse tree
	 */
	void enterArrayRule(ObjectApiParser.ArrayRuleContext ctx);
	/**
	 * Exit a parse tree produced by {@link ObjectApiParser#arrayRule}.
	 * @param ctx the parse tree
	 */
	void exitArrayRule(ObjectApiParser.ArrayRuleContext ctx);
	/**
	 * Enter a parse tree produced by {@link ObjectApiParser#primitiveSchema}.
	 * @param ctx the parse tree
	 */
	void enterPrimitiveSchema(ObjectApiParser.PrimitiveSchemaContext ctx);
	/**
	 * Exit a parse tree produced by {@link ObjectApiParser#primitiveSchema}.
	 * @param ctx the parse tree
	 */
	void exitPrimitiveSchema(ObjectApiParser.PrimitiveSchemaContext ctx);
	/**
	 * Enter a parse tree produced by {@link ObjectApiParser#symbolSchema}.
	 * @param ctx the parse tree
	 */
	void enterSymbolSchema(ObjectApiParser.SymbolSchemaContext ctx);
	/**
	 * Exit a parse tree produced by {@link ObjectApiParser#symbolSchema}.
	 * @param ctx the parse tree
	 */
	void exitSymbolSchema(ObjectApiParser.SymbolSchemaContext ctx);
	/**
	 * Enter a parse tree produced by {@link ObjectApiParser#metaRule}.
	 * @param ctx the parse tree
	 */
	void enterMetaRule(ObjectApiParser.MetaRuleContext ctx);
	/**
	 * Exit a parse tree produced by {@link ObjectApiParser#metaRule}.
	 * @param ctx the parse tree
	 */
	void exitMetaRule(ObjectApiParser.MetaRuleContext ctx);
}