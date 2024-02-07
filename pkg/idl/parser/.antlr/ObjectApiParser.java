// Generated from /Users/jryannel/dev/apigear/cli/pkg/idl/parser/ObjectApi.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class ObjectApiParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		WHITESPACE=25, INTEGER=26, HEX=27, TYPE_IDENTIFIER=28, IDENTIFIER=29, 
		VERSION=30, DOCLINE=31, TAGLINE=32, COMMENT=33;
	public static final int
		RULE_documentRule = 0, RULE_headerRule = 1, RULE_moduleRule = 2, RULE_importRule = 3, 
		RULE_declarationsRule = 4, RULE_interfaceRule = 5, RULE_interfaceMembersRule = 6, 
		RULE_propertyRule = 7, RULE_operationRule = 8, RULE_operationReturnRule = 9, 
		RULE_operationParamRule = 10, RULE_signalRule = 11, RULE_structRule = 12, 
		RULE_structFieldRule = 13, RULE_enumRule = 14, RULE_enumMemberRule = 15, 
		RULE_schemaRule = 16, RULE_arrayRule = 17, RULE_primitiveSchema = 18, 
		RULE_symbolSchema = 19, RULE_metaRule = 20;
	private static String[] makeRuleNames() {
		return new String[] {
			"documentRule", "headerRule", "moduleRule", "importRule", "declarationsRule", 
			"interfaceRule", "interfaceMembersRule", "propertyRule", "operationRule", 
			"operationReturnRule", "operationParamRule", "signalRule", "structRule", 
			"structFieldRule", "enumRule", "enumMemberRule", "schemaRule", "arrayRule", 
			"primitiveSchema", "symbolSchema", "metaRule"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'module'", "'import'", "'interface'", "'{'", "'}'", "'readonly'", 
			"':'", "'('", "')'", "','", "'signal'", "'struct'", "'enum'", "'='", 
			"'['", "']'", "'bool'", "'int'", "'int32'", "'int64'", "'float'", "'float32'", 
			"'float64'", "'string'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, "WHITESPACE", "INTEGER", "HEX", "TYPE_IDENTIFIER", "IDENTIFIER", 
			"VERSION", "DOCLINE", "TAGLINE", "COMMENT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "ObjectApi.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public ObjectApiParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DocumentRuleContext extends ParserRuleContext {
		public HeaderRuleContext headerRule() {
			return getRuleContext(HeaderRuleContext.class,0);
		}
		public List<DeclarationsRuleContext> declarationsRule() {
			return getRuleContexts(DeclarationsRuleContext.class);
		}
		public DeclarationsRuleContext declarationsRule(int i) {
			return getRuleContext(DeclarationsRuleContext.class,i);
		}
		public DocumentRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_documentRule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).enterDocumentRule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).exitDocumentRule(this);
		}
	}

	public final DocumentRuleContext documentRule() throws RecognitionException {
		DocumentRuleContext _localctx = new DocumentRuleContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_documentRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(42);
			headerRule();
			setState(46);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 6442463240L) != 0)) {
				{
				{
				setState(43);
				declarationsRule();
				}
				}
				setState(48);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class HeaderRuleContext extends ParserRuleContext {
		public ModuleRuleContext moduleRule() {
			return getRuleContext(ModuleRuleContext.class,0);
		}
		public List<ImportRuleContext> importRule() {
			return getRuleContexts(ImportRuleContext.class);
		}
		public ImportRuleContext importRule(int i) {
			return getRuleContext(ImportRuleContext.class,i);
		}
		public HeaderRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_headerRule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).enterHeaderRule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).exitHeaderRule(this);
		}
	}

	public final HeaderRuleContext headerRule() throws RecognitionException {
		HeaderRuleContext _localctx = new HeaderRuleContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_headerRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(49);
			moduleRule();
			setState(53);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__1) {
				{
				{
				setState(50);
				importRule();
				}
				}
				setState(55);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ModuleRuleContext extends ParserRuleContext {
		public Token name;
		public Token version;
		public TerminalNode IDENTIFIER() { return getToken(ObjectApiParser.IDENTIFIER, 0); }
		public TerminalNode VERSION() { return getToken(ObjectApiParser.VERSION, 0); }
		public List<MetaRuleContext> metaRule() {
			return getRuleContexts(MetaRuleContext.class);
		}
		public MetaRuleContext metaRule(int i) {
			return getRuleContext(MetaRuleContext.class,i);
		}
		public ModuleRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_moduleRule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).enterModuleRule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).exitModuleRule(this);
		}
	}

	public final ModuleRuleContext moduleRule() throws RecognitionException {
		ModuleRuleContext _localctx = new ModuleRuleContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_moduleRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(59);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOCLINE || _la==TAGLINE) {
				{
				{
				setState(56);
				metaRule();
				}
				}
				setState(61);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(62);
			match(T__0);
			setState(63);
			((ModuleRuleContext)_localctx).name = match(IDENTIFIER);
			setState(64);
			((ModuleRuleContext)_localctx).version = match(VERSION);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ImportRuleContext extends ParserRuleContext {
		public Token name;
		public Token version;
		public TerminalNode IDENTIFIER() { return getToken(ObjectApiParser.IDENTIFIER, 0); }
		public TerminalNode VERSION() { return getToken(ObjectApiParser.VERSION, 0); }
		public ImportRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importRule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).enterImportRule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).exitImportRule(this);
		}
	}

	public final ImportRuleContext importRule() throws RecognitionException {
		ImportRuleContext _localctx = new ImportRuleContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_importRule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(66);
			match(T__1);
			setState(67);
			((ImportRuleContext)_localctx).name = match(IDENTIFIER);
			setState(68);
			((ImportRuleContext)_localctx).version = match(VERSION);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DeclarationsRuleContext extends ParserRuleContext {
		public InterfaceRuleContext interfaceRule() {
			return getRuleContext(InterfaceRuleContext.class,0);
		}
		public StructRuleContext structRule() {
			return getRuleContext(StructRuleContext.class,0);
		}
		public EnumRuleContext enumRule() {
			return getRuleContext(EnumRuleContext.class,0);
		}
		public DeclarationsRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declarationsRule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).enterDeclarationsRule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).exitDeclarationsRule(this);
		}
	}

	public final DeclarationsRuleContext declarationsRule() throws RecognitionException {
		DeclarationsRuleContext _localctx = new DeclarationsRuleContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_declarationsRule);
		try {
			setState(73);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(70);
				interfaceRule();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(71);
				structRule();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(72);
				enumRule();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class InterfaceRuleContext extends ParserRuleContext {
		public Token name;
		public TerminalNode IDENTIFIER() { return getToken(ObjectApiParser.IDENTIFIER, 0); }
		public List<MetaRuleContext> metaRule() {
			return getRuleContexts(MetaRuleContext.class);
		}
		public MetaRuleContext metaRule(int i) {
			return getRuleContext(MetaRuleContext.class,i);
		}
		public List<InterfaceMembersRuleContext> interfaceMembersRule() {
			return getRuleContexts(InterfaceMembersRuleContext.class);
		}
		public InterfaceMembersRuleContext interfaceMembersRule(int i) {
			return getRuleContext(InterfaceMembersRuleContext.class,i);
		}
		public InterfaceRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_interfaceRule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).enterInterfaceRule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).exitInterfaceRule(this);
		}
	}

	public final InterfaceRuleContext interfaceRule() throws RecognitionException {
		InterfaceRuleContext _localctx = new InterfaceRuleContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_interfaceRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(78);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOCLINE || _la==TAGLINE) {
				{
				{
				setState(75);
				metaRule();
				}
				}
				setState(80);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(81);
			match(T__2);
			setState(82);
			((InterfaceRuleContext)_localctx).name = match(IDENTIFIER);
			setState(83);
			match(T__3);
			setState(87);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 6979323968L) != 0)) {
				{
				{
				setState(84);
				interfaceMembersRule();
				}
				}
				setState(89);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(90);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class InterfaceMembersRuleContext extends ParserRuleContext {
		public PropertyRuleContext propertyRule() {
			return getRuleContext(PropertyRuleContext.class,0);
		}
		public OperationRuleContext operationRule() {
			return getRuleContext(OperationRuleContext.class,0);
		}
		public SignalRuleContext signalRule() {
			return getRuleContext(SignalRuleContext.class,0);
		}
		public InterfaceMembersRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_interfaceMembersRule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).enterInterfaceMembersRule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).exitInterfaceMembersRule(this);
		}
	}

	public final InterfaceMembersRuleContext interfaceMembersRule() throws RecognitionException {
		InterfaceMembersRuleContext _localctx = new InterfaceMembersRuleContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_interfaceMembersRule);
		try {
			setState(95);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(92);
				propertyRule();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(93);
				operationRule();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(94);
				signalRule();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PropertyRuleContext extends ParserRuleContext {
		public Token readonly;
		public Token name;
		public SchemaRuleContext schema;
		public TerminalNode IDENTIFIER() { return getToken(ObjectApiParser.IDENTIFIER, 0); }
		public SchemaRuleContext schemaRule() {
			return getRuleContext(SchemaRuleContext.class,0);
		}
		public List<MetaRuleContext> metaRule() {
			return getRuleContexts(MetaRuleContext.class);
		}
		public MetaRuleContext metaRule(int i) {
			return getRuleContext(MetaRuleContext.class,i);
		}
		public PropertyRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_propertyRule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).enterPropertyRule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).exitPropertyRule(this);
		}
	}

	public final PropertyRuleContext propertyRule() throws RecognitionException {
		PropertyRuleContext _localctx = new PropertyRuleContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_propertyRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(100);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOCLINE || _la==TAGLINE) {
				{
				{
				setState(97);
				metaRule();
				}
				}
				setState(102);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(104);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__5) {
				{
				setState(103);
				((PropertyRuleContext)_localctx).readonly = match(T__5);
				}
			}

			setState(106);
			((PropertyRuleContext)_localctx).name = match(IDENTIFIER);
			setState(107);
			match(T__6);
			setState(108);
			((PropertyRuleContext)_localctx).schema = schemaRule();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class OperationRuleContext extends ParserRuleContext {
		public Token name;
		public OperationParamRuleContext params;
		public TerminalNode IDENTIFIER() { return getToken(ObjectApiParser.IDENTIFIER, 0); }
		public List<MetaRuleContext> metaRule() {
			return getRuleContexts(MetaRuleContext.class);
		}
		public MetaRuleContext metaRule(int i) {
			return getRuleContext(MetaRuleContext.class,i);
		}
		public OperationReturnRuleContext operationReturnRule() {
			return getRuleContext(OperationReturnRuleContext.class,0);
		}
		public List<OperationParamRuleContext> operationParamRule() {
			return getRuleContexts(OperationParamRuleContext.class);
		}
		public OperationParamRuleContext operationParamRule(int i) {
			return getRuleContext(OperationParamRuleContext.class,i);
		}
		public OperationRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operationRule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).enterOperationRule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).exitOperationRule(this);
		}
	}

	public final OperationRuleContext operationRule() throws RecognitionException {
		OperationRuleContext _localctx = new OperationRuleContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_operationRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(113);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOCLINE || _la==TAGLINE) {
				{
				{
				setState(110);
				metaRule();
				}
				}
				setState(115);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(116);
			((OperationRuleContext)_localctx).name = match(IDENTIFIER);
			setState(117);
			match(T__7);
			setState(121);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENTIFIER) {
				{
				{
				setState(118);
				((OperationRuleContext)_localctx).params = operationParamRule();
				}
				}
				setState(123);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(124);
			match(T__8);
			setState(126);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__6) {
				{
				setState(125);
				operationReturnRule();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class OperationReturnRuleContext extends ParserRuleContext {
		public SchemaRuleContext schema;
		public SchemaRuleContext schemaRule() {
			return getRuleContext(SchemaRuleContext.class,0);
		}
		public OperationReturnRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operationReturnRule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).enterOperationReturnRule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).exitOperationReturnRule(this);
		}
	}

	public final OperationReturnRuleContext operationReturnRule() throws RecognitionException {
		OperationReturnRuleContext _localctx = new OperationReturnRuleContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_operationReturnRule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(128);
			match(T__6);
			setState(129);
			((OperationReturnRuleContext)_localctx).schema = schemaRule();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class OperationParamRuleContext extends ParserRuleContext {
		public Token name;
		public SchemaRuleContext schema;
		public TerminalNode IDENTIFIER() { return getToken(ObjectApiParser.IDENTIFIER, 0); }
		public SchemaRuleContext schemaRule() {
			return getRuleContext(SchemaRuleContext.class,0);
		}
		public OperationParamRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operationParamRule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).enterOperationParamRule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).exitOperationParamRule(this);
		}
	}

	public final OperationParamRuleContext operationParamRule() throws RecognitionException {
		OperationParamRuleContext _localctx = new OperationParamRuleContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_operationParamRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(131);
			((OperationParamRuleContext)_localctx).name = match(IDENTIFIER);
			setState(132);
			match(T__6);
			setState(133);
			((OperationParamRuleContext)_localctx).schema = schemaRule();
			setState(135);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__9) {
				{
				setState(134);
				match(T__9);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SignalRuleContext extends ParserRuleContext {
		public Token name;
		public OperationParamRuleContext params;
		public TerminalNode IDENTIFIER() { return getToken(ObjectApiParser.IDENTIFIER, 0); }
		public List<MetaRuleContext> metaRule() {
			return getRuleContexts(MetaRuleContext.class);
		}
		public MetaRuleContext metaRule(int i) {
			return getRuleContext(MetaRuleContext.class,i);
		}
		public List<OperationParamRuleContext> operationParamRule() {
			return getRuleContexts(OperationParamRuleContext.class);
		}
		public OperationParamRuleContext operationParamRule(int i) {
			return getRuleContext(OperationParamRuleContext.class,i);
		}
		public SignalRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_signalRule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).enterSignalRule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).exitSignalRule(this);
		}
	}

	public final SignalRuleContext signalRule() throws RecognitionException {
		SignalRuleContext _localctx = new SignalRuleContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_signalRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(140);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOCLINE || _la==TAGLINE) {
				{
				{
				setState(137);
				metaRule();
				}
				}
				setState(142);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(143);
			match(T__10);
			setState(144);
			((SignalRuleContext)_localctx).name = match(IDENTIFIER);
			setState(145);
			match(T__7);
			setState(149);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENTIFIER) {
				{
				{
				setState(146);
				((SignalRuleContext)_localctx).params = operationParamRule();
				}
				}
				setState(151);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(152);
			match(T__8);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StructRuleContext extends ParserRuleContext {
		public Token name;
		public TerminalNode IDENTIFIER() { return getToken(ObjectApiParser.IDENTIFIER, 0); }
		public List<MetaRuleContext> metaRule() {
			return getRuleContexts(MetaRuleContext.class);
		}
		public MetaRuleContext metaRule(int i) {
			return getRuleContext(MetaRuleContext.class,i);
		}
		public List<StructFieldRuleContext> structFieldRule() {
			return getRuleContexts(StructFieldRuleContext.class);
		}
		public StructFieldRuleContext structFieldRule(int i) {
			return getRuleContext(StructFieldRuleContext.class,i);
		}
		public StructRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structRule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).enterStructRule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).exitStructRule(this);
		}
	}

	public final StructRuleContext structRule() throws RecognitionException {
		StructRuleContext _localctx = new StructRuleContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_structRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(157);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOCLINE || _la==TAGLINE) {
				{
				{
				setState(154);
				metaRule();
				}
				}
				setState(159);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(160);
			match(T__11);
			setState(161);
			((StructRuleContext)_localctx).name = match(IDENTIFIER);
			setState(162);
			match(T__3);
			setState(166);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 6979321920L) != 0)) {
				{
				{
				setState(163);
				structFieldRule();
				}
				}
				setState(168);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(169);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StructFieldRuleContext extends ParserRuleContext {
		public Token readonly;
		public Token name;
		public SchemaRuleContext schema;
		public TerminalNode IDENTIFIER() { return getToken(ObjectApiParser.IDENTIFIER, 0); }
		public SchemaRuleContext schemaRule() {
			return getRuleContext(SchemaRuleContext.class,0);
		}
		public List<MetaRuleContext> metaRule() {
			return getRuleContexts(MetaRuleContext.class);
		}
		public MetaRuleContext metaRule(int i) {
			return getRuleContext(MetaRuleContext.class,i);
		}
		public StructFieldRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structFieldRule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).enterStructFieldRule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).exitStructFieldRule(this);
		}
	}

	public final StructFieldRuleContext structFieldRule() throws RecognitionException {
		StructFieldRuleContext _localctx = new StructFieldRuleContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_structFieldRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(174);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOCLINE || _la==TAGLINE) {
				{
				{
				setState(171);
				metaRule();
				}
				}
				setState(176);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(178);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__5) {
				{
				setState(177);
				((StructFieldRuleContext)_localctx).readonly = match(T__5);
				}
			}

			setState(180);
			((StructFieldRuleContext)_localctx).name = match(IDENTIFIER);
			setState(181);
			match(T__6);
			setState(182);
			((StructFieldRuleContext)_localctx).schema = schemaRule();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class EnumRuleContext extends ParserRuleContext {
		public Token name;
		public TerminalNode IDENTIFIER() { return getToken(ObjectApiParser.IDENTIFIER, 0); }
		public List<MetaRuleContext> metaRule() {
			return getRuleContexts(MetaRuleContext.class);
		}
		public MetaRuleContext metaRule(int i) {
			return getRuleContext(MetaRuleContext.class,i);
		}
		public List<EnumMemberRuleContext> enumMemberRule() {
			return getRuleContexts(EnumMemberRuleContext.class);
		}
		public EnumMemberRuleContext enumMemberRule(int i) {
			return getRuleContext(EnumMemberRuleContext.class,i);
		}
		public EnumRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enumRule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).enterEnumRule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).exitEnumRule(this);
		}
	}

	public final EnumRuleContext enumRule() throws RecognitionException {
		EnumRuleContext _localctx = new EnumRuleContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_enumRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(187);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOCLINE || _la==TAGLINE) {
				{
				{
				setState(184);
				metaRule();
				}
				}
				setState(189);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(190);
			match(T__12);
			setState(191);
			((EnumRuleContext)_localctx).name = match(IDENTIFIER);
			setState(192);
			match(T__3);
			setState(196);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 6979321856L) != 0)) {
				{
				{
				setState(193);
				enumMemberRule();
				}
				}
				setState(198);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(199);
			match(T__4);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class EnumMemberRuleContext extends ParserRuleContext {
		public Token name;
		public Token value;
		public TerminalNode IDENTIFIER() { return getToken(ObjectApiParser.IDENTIFIER, 0); }
		public List<MetaRuleContext> metaRule() {
			return getRuleContexts(MetaRuleContext.class);
		}
		public MetaRuleContext metaRule(int i) {
			return getRuleContext(MetaRuleContext.class,i);
		}
		public TerminalNode INTEGER() { return getToken(ObjectApiParser.INTEGER, 0); }
		public EnumMemberRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_enumMemberRule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).enterEnumMemberRule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).exitEnumMemberRule(this);
		}
	}

	public final EnumMemberRuleContext enumMemberRule() throws RecognitionException {
		EnumMemberRuleContext _localctx = new EnumMemberRuleContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_enumMemberRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(204);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOCLINE || _la==TAGLINE) {
				{
				{
				setState(201);
				metaRule();
				}
				}
				setState(206);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(207);
			((EnumMemberRuleContext)_localctx).name = match(IDENTIFIER);
			setState(210);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__13) {
				{
				setState(208);
				match(T__13);
				setState(209);
				((EnumMemberRuleContext)_localctx).value = match(INTEGER);
				}
			}

			setState(213);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__9) {
				{
				setState(212);
				match(T__9);
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SchemaRuleContext extends ParserRuleContext {
		public PrimitiveSchemaContext primitiveSchema() {
			return getRuleContext(PrimitiveSchemaContext.class,0);
		}
		public SymbolSchemaContext symbolSchema() {
			return getRuleContext(SymbolSchemaContext.class,0);
		}
		public ArrayRuleContext arrayRule() {
			return getRuleContext(ArrayRuleContext.class,0);
		}
		public SchemaRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_schemaRule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).enterSchemaRule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).exitSchemaRule(this);
		}
	}

	public final SchemaRuleContext schemaRule() throws RecognitionException {
		SchemaRuleContext _localctx = new SchemaRuleContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_schemaRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(217);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__16:
			case T__17:
			case T__18:
			case T__19:
			case T__20:
			case T__21:
			case T__22:
			case T__23:
				{
				setState(215);
				primitiveSchema();
				}
				break;
			case IDENTIFIER:
				{
				setState(216);
				symbolSchema();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(220);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__14) {
				{
				setState(219);
				arrayRule();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ArrayRuleContext extends ParserRuleContext {
		public ArrayRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrayRule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).enterArrayRule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).exitArrayRule(this);
		}
	}

	public final ArrayRuleContext arrayRule() throws RecognitionException {
		ArrayRuleContext _localctx = new ArrayRuleContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_arrayRule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(222);
			match(T__14);
			setState(223);
			match(T__15);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class PrimitiveSchemaContext extends ParserRuleContext {
		public Token name;
		public PrimitiveSchemaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primitiveSchema; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).enterPrimitiveSchema(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).exitPrimitiveSchema(this);
		}
	}

	public final PrimitiveSchemaContext primitiveSchema() throws RecognitionException {
		PrimitiveSchemaContext _localctx = new PrimitiveSchemaContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_primitiveSchema);
		try {
			setState(233);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__16:
				enterOuterAlt(_localctx, 1);
				{
				setState(225);
				((PrimitiveSchemaContext)_localctx).name = match(T__16);
				}
				break;
			case T__17:
				enterOuterAlt(_localctx, 2);
				{
				setState(226);
				((PrimitiveSchemaContext)_localctx).name = match(T__17);
				}
				break;
			case T__18:
				enterOuterAlt(_localctx, 3);
				{
				setState(227);
				((PrimitiveSchemaContext)_localctx).name = match(T__18);
				}
				break;
			case T__19:
				enterOuterAlt(_localctx, 4);
				{
				setState(228);
				((PrimitiveSchemaContext)_localctx).name = match(T__19);
				}
				break;
			case T__20:
				enterOuterAlt(_localctx, 5);
				{
				setState(229);
				((PrimitiveSchemaContext)_localctx).name = match(T__20);
				}
				break;
			case T__21:
				enterOuterAlt(_localctx, 6);
				{
				setState(230);
				((PrimitiveSchemaContext)_localctx).name = match(T__21);
				}
				break;
			case T__22:
				enterOuterAlt(_localctx, 7);
				{
				setState(231);
				((PrimitiveSchemaContext)_localctx).name = match(T__22);
				}
				break;
			case T__23:
				enterOuterAlt(_localctx, 8);
				{
				setState(232);
				((PrimitiveSchemaContext)_localctx).name = match(T__23);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SymbolSchemaContext extends ParserRuleContext {
		public Token name;
		public TerminalNode IDENTIFIER() { return getToken(ObjectApiParser.IDENTIFIER, 0); }
		public SymbolSchemaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_symbolSchema; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).enterSymbolSchema(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).exitSymbolSchema(this);
		}
	}

	public final SymbolSchemaContext symbolSchema() throws RecognitionException {
		SymbolSchemaContext _localctx = new SymbolSchemaContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_symbolSchema);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(235);
			((SymbolSchemaContext)_localctx).name = match(IDENTIFIER);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class MetaRuleContext extends ParserRuleContext {
		public Token tagLine;
		public Token docLine;
		public TerminalNode TAGLINE() { return getToken(ObjectApiParser.TAGLINE, 0); }
		public TerminalNode DOCLINE() { return getToken(ObjectApiParser.DOCLINE, 0); }
		public MetaRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_metaRule; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).enterMetaRule(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof ObjectApiListener ) ((ObjectApiListener)listener).exitMetaRule(this);
		}
	}

	public final MetaRuleContext metaRule() throws RecognitionException {
		MetaRuleContext _localctx = new MetaRuleContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_metaRule);
		try {
			setState(239);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TAGLINE:
				enterOuterAlt(_localctx, 1);
				{
				setState(237);
				((MetaRuleContext)_localctx).tagLine = match(TAGLINE);
				}
				break;
			case DOCLINE:
				enterOuterAlt(_localctx, 2);
				{
				setState(238);
				((MetaRuleContext)_localctx).docLine = match(DOCLINE);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\u0004\u0001!\u00f2\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0001\u0000\u0001\u0000"+
		"\u0005\u0000-\b\u0000\n\u0000\f\u00000\t\u0000\u0001\u0001\u0001\u0001"+
		"\u0005\u00014\b\u0001\n\u0001\f\u00017\t\u0001\u0001\u0002\u0005\u0002"+
		":\b\u0002\n\u0002\f\u0002=\t\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0003\u0004J\b\u0004\u0001\u0005\u0005\u0005"+
		"M\b\u0005\n\u0005\f\u0005P\t\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0005\u0005V\b\u0005\n\u0005\f\u0005Y\t\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006`\b\u0006"+
		"\u0001\u0007\u0005\u0007c\b\u0007\n\u0007\f\u0007f\t\u0007\u0001\u0007"+
		"\u0003\u0007i\b\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\b\u0005\bp\b\b\n\b\f\bs\t\b\u0001\b\u0001\b\u0001\b\u0005\bx\b"+
		"\b\n\b\f\b{\t\b\u0001\b\u0001\b\u0003\b\u007f\b\b\u0001\t\u0001\t\u0001"+
		"\t\u0001\n\u0001\n\u0001\n\u0001\n\u0003\n\u0088\b\n\u0001\u000b\u0005"+
		"\u000b\u008b\b\u000b\n\u000b\f\u000b\u008e\t\u000b\u0001\u000b\u0001\u000b"+
		"\u0001\u000b\u0001\u000b\u0005\u000b\u0094\b\u000b\n\u000b\f\u000b\u0097"+
		"\t\u000b\u0001\u000b\u0001\u000b\u0001\f\u0005\f\u009c\b\f\n\f\f\f\u009f"+
		"\t\f\u0001\f\u0001\f\u0001\f\u0001\f\u0005\f\u00a5\b\f\n\f\f\f\u00a8\t"+
		"\f\u0001\f\u0001\f\u0001\r\u0005\r\u00ad\b\r\n\r\f\r\u00b0\t\r\u0001\r"+
		"\u0003\r\u00b3\b\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\u000e\u0005\u000e"+
		"\u00ba\b\u000e\n\u000e\f\u000e\u00bd\t\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000e\u0005\u000e\u00c3\b\u000e\n\u000e\f\u000e\u00c6\t\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000f\u0005\u000f\u00cb\b\u000f\n\u000f"+
		"\f\u000f\u00ce\t\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0003\u000f"+
		"\u00d3\b\u000f\u0001\u000f\u0003\u000f\u00d6\b\u000f\u0001\u0010\u0001"+
		"\u0010\u0003\u0010\u00da\b\u0010\u0001\u0010\u0003\u0010\u00dd\b\u0010"+
		"\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0003\u0012"+
		"\u00ea\b\u0012\u0001\u0013\u0001\u0013\u0001\u0014\u0001\u0014\u0003\u0014"+
		"\u00f0\b\u0014\u0001\u0014\u0000\u0000\u0015\u0000\u0002\u0004\u0006\b"+
		"\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"$&(\u0000"+
		"\u0000\u0100\u0000*\u0001\u0000\u0000\u0000\u00021\u0001\u0000\u0000\u0000"+
		"\u0004;\u0001\u0000\u0000\u0000\u0006B\u0001\u0000\u0000\u0000\bI\u0001"+
		"\u0000\u0000\u0000\nN\u0001\u0000\u0000\u0000\f_\u0001\u0000\u0000\u0000"+
		"\u000ed\u0001\u0000\u0000\u0000\u0010q\u0001\u0000\u0000\u0000\u0012\u0080"+
		"\u0001\u0000\u0000\u0000\u0014\u0083\u0001\u0000\u0000\u0000\u0016\u008c"+
		"\u0001\u0000\u0000\u0000\u0018\u009d\u0001\u0000\u0000\u0000\u001a\u00ae"+
		"\u0001\u0000\u0000\u0000\u001c\u00bb\u0001\u0000\u0000\u0000\u001e\u00cc"+
		"\u0001\u0000\u0000\u0000 \u00d9\u0001\u0000\u0000\u0000\"\u00de\u0001"+
		"\u0000\u0000\u0000$\u00e9\u0001\u0000\u0000\u0000&\u00eb\u0001\u0000\u0000"+
		"\u0000(\u00ef\u0001\u0000\u0000\u0000*.\u0003\u0002\u0001\u0000+-\u0003"+
		"\b\u0004\u0000,+\u0001\u0000\u0000\u0000-0\u0001\u0000\u0000\u0000.,\u0001"+
		"\u0000\u0000\u0000./\u0001\u0000\u0000\u0000/\u0001\u0001\u0000\u0000"+
		"\u00000.\u0001\u0000\u0000\u000015\u0003\u0004\u0002\u000024\u0003\u0006"+
		"\u0003\u000032\u0001\u0000\u0000\u000047\u0001\u0000\u0000\u000053\u0001"+
		"\u0000\u0000\u000056\u0001\u0000\u0000\u00006\u0003\u0001\u0000\u0000"+
		"\u000075\u0001\u0000\u0000\u00008:\u0003(\u0014\u000098\u0001\u0000\u0000"+
		"\u0000:=\u0001\u0000\u0000\u0000;9\u0001\u0000\u0000\u0000;<\u0001\u0000"+
		"\u0000\u0000<>\u0001\u0000\u0000\u0000=;\u0001\u0000\u0000\u0000>?\u0005"+
		"\u0001\u0000\u0000?@\u0005\u001d\u0000\u0000@A\u0005\u001e\u0000\u0000"+
		"A\u0005\u0001\u0000\u0000\u0000BC\u0005\u0002\u0000\u0000CD\u0005\u001d"+
		"\u0000\u0000DE\u0005\u001e\u0000\u0000E\u0007\u0001\u0000\u0000\u0000"+
		"FJ\u0003\n\u0005\u0000GJ\u0003\u0018\f\u0000HJ\u0003\u001c\u000e\u0000"+
		"IF\u0001\u0000\u0000\u0000IG\u0001\u0000\u0000\u0000IH\u0001\u0000\u0000"+
		"\u0000J\t\u0001\u0000\u0000\u0000KM\u0003(\u0014\u0000LK\u0001\u0000\u0000"+
		"\u0000MP\u0001\u0000\u0000\u0000NL\u0001\u0000\u0000\u0000NO\u0001\u0000"+
		"\u0000\u0000OQ\u0001\u0000\u0000\u0000PN\u0001\u0000\u0000\u0000QR\u0005"+
		"\u0003\u0000\u0000RS\u0005\u001d\u0000\u0000SW\u0005\u0004\u0000\u0000"+
		"TV\u0003\f\u0006\u0000UT\u0001\u0000\u0000\u0000VY\u0001\u0000\u0000\u0000"+
		"WU\u0001\u0000\u0000\u0000WX\u0001\u0000\u0000\u0000XZ\u0001\u0000\u0000"+
		"\u0000YW\u0001\u0000\u0000\u0000Z[\u0005\u0005\u0000\u0000[\u000b\u0001"+
		"\u0000\u0000\u0000\\`\u0003\u000e\u0007\u0000]`\u0003\u0010\b\u0000^`"+
		"\u0003\u0016\u000b\u0000_\\\u0001\u0000\u0000\u0000_]\u0001\u0000\u0000"+
		"\u0000_^\u0001\u0000\u0000\u0000`\r\u0001\u0000\u0000\u0000ac\u0003(\u0014"+
		"\u0000ba\u0001\u0000\u0000\u0000cf\u0001\u0000\u0000\u0000db\u0001\u0000"+
		"\u0000\u0000de\u0001\u0000\u0000\u0000eh\u0001\u0000\u0000\u0000fd\u0001"+
		"\u0000\u0000\u0000gi\u0005\u0006\u0000\u0000hg\u0001\u0000\u0000\u0000"+
		"hi\u0001\u0000\u0000\u0000ij\u0001\u0000\u0000\u0000jk\u0005\u001d\u0000"+
		"\u0000kl\u0005\u0007\u0000\u0000lm\u0003 \u0010\u0000m\u000f\u0001\u0000"+
		"\u0000\u0000np\u0003(\u0014\u0000on\u0001\u0000\u0000\u0000ps\u0001\u0000"+
		"\u0000\u0000qo\u0001\u0000\u0000\u0000qr\u0001\u0000\u0000\u0000rt\u0001"+
		"\u0000\u0000\u0000sq\u0001\u0000\u0000\u0000tu\u0005\u001d\u0000\u0000"+
		"uy\u0005\b\u0000\u0000vx\u0003\u0014\n\u0000wv\u0001\u0000\u0000\u0000"+
		"x{\u0001\u0000\u0000\u0000yw\u0001\u0000\u0000\u0000yz\u0001\u0000\u0000"+
		"\u0000z|\u0001\u0000\u0000\u0000{y\u0001\u0000\u0000\u0000|~\u0005\t\u0000"+
		"\u0000}\u007f\u0003\u0012\t\u0000~}\u0001\u0000\u0000\u0000~\u007f\u0001"+
		"\u0000\u0000\u0000\u007f\u0011\u0001\u0000\u0000\u0000\u0080\u0081\u0005"+
		"\u0007\u0000\u0000\u0081\u0082\u0003 \u0010\u0000\u0082\u0013\u0001\u0000"+
		"\u0000\u0000\u0083\u0084\u0005\u001d\u0000\u0000\u0084\u0085\u0005\u0007"+
		"\u0000\u0000\u0085\u0087\u0003 \u0010\u0000\u0086\u0088\u0005\n\u0000"+
		"\u0000\u0087\u0086\u0001\u0000\u0000\u0000\u0087\u0088\u0001\u0000\u0000"+
		"\u0000\u0088\u0015\u0001\u0000\u0000\u0000\u0089\u008b\u0003(\u0014\u0000"+
		"\u008a\u0089\u0001\u0000\u0000\u0000\u008b\u008e\u0001\u0000\u0000\u0000"+
		"\u008c\u008a\u0001\u0000\u0000\u0000\u008c\u008d\u0001\u0000\u0000\u0000"+
		"\u008d\u008f\u0001\u0000\u0000\u0000\u008e\u008c\u0001\u0000\u0000\u0000"+
		"\u008f\u0090\u0005\u000b\u0000\u0000\u0090\u0091\u0005\u001d\u0000\u0000"+
		"\u0091\u0095\u0005\b\u0000\u0000\u0092\u0094\u0003\u0014\n\u0000\u0093"+
		"\u0092\u0001\u0000\u0000\u0000\u0094\u0097\u0001\u0000\u0000\u0000\u0095"+
		"\u0093\u0001\u0000\u0000\u0000\u0095\u0096\u0001\u0000\u0000\u0000\u0096"+
		"\u0098\u0001\u0000\u0000\u0000\u0097\u0095\u0001\u0000\u0000\u0000\u0098"+
		"\u0099\u0005\t\u0000\u0000\u0099\u0017\u0001\u0000\u0000\u0000\u009a\u009c"+
		"\u0003(\u0014\u0000\u009b\u009a\u0001\u0000\u0000\u0000\u009c\u009f\u0001"+
		"\u0000\u0000\u0000\u009d\u009b\u0001\u0000\u0000\u0000\u009d\u009e\u0001"+
		"\u0000\u0000\u0000\u009e\u00a0\u0001\u0000\u0000\u0000\u009f\u009d\u0001"+
		"\u0000\u0000\u0000\u00a0\u00a1\u0005\f\u0000\u0000\u00a1\u00a2\u0005\u001d"+
		"\u0000\u0000\u00a2\u00a6\u0005\u0004\u0000\u0000\u00a3\u00a5\u0003\u001a"+
		"\r\u0000\u00a4\u00a3\u0001\u0000\u0000\u0000\u00a5\u00a8\u0001\u0000\u0000"+
		"\u0000\u00a6\u00a4\u0001\u0000\u0000\u0000\u00a6\u00a7\u0001\u0000\u0000"+
		"\u0000\u00a7\u00a9\u0001\u0000\u0000\u0000\u00a8\u00a6\u0001\u0000\u0000"+
		"\u0000\u00a9\u00aa\u0005\u0005\u0000\u0000\u00aa\u0019\u0001\u0000\u0000"+
		"\u0000\u00ab\u00ad\u0003(\u0014\u0000\u00ac\u00ab\u0001\u0000\u0000\u0000"+
		"\u00ad\u00b0\u0001\u0000\u0000\u0000\u00ae\u00ac\u0001\u0000\u0000\u0000"+
		"\u00ae\u00af\u0001\u0000\u0000\u0000\u00af\u00b2\u0001\u0000\u0000\u0000"+
		"\u00b0\u00ae\u0001\u0000\u0000\u0000\u00b1\u00b3\u0005\u0006\u0000\u0000"+
		"\u00b2\u00b1\u0001\u0000\u0000\u0000\u00b2\u00b3\u0001\u0000\u0000\u0000"+
		"\u00b3\u00b4\u0001\u0000\u0000\u0000\u00b4\u00b5\u0005\u001d\u0000\u0000"+
		"\u00b5\u00b6\u0005\u0007\u0000\u0000\u00b6\u00b7\u0003 \u0010\u0000\u00b7"+
		"\u001b\u0001\u0000\u0000\u0000\u00b8\u00ba\u0003(\u0014\u0000\u00b9\u00b8"+
		"\u0001\u0000\u0000\u0000\u00ba\u00bd\u0001\u0000\u0000\u0000\u00bb\u00b9"+
		"\u0001\u0000\u0000\u0000\u00bb\u00bc\u0001\u0000\u0000\u0000\u00bc\u00be"+
		"\u0001\u0000\u0000\u0000\u00bd\u00bb\u0001\u0000\u0000\u0000\u00be\u00bf"+
		"\u0005\r\u0000\u0000\u00bf\u00c0\u0005\u001d\u0000\u0000\u00c0\u00c4\u0005"+
		"\u0004\u0000\u0000\u00c1\u00c3\u0003\u001e\u000f\u0000\u00c2\u00c1\u0001"+
		"\u0000\u0000\u0000\u00c3\u00c6\u0001\u0000\u0000\u0000\u00c4\u00c2\u0001"+
		"\u0000\u0000\u0000\u00c4\u00c5\u0001\u0000\u0000\u0000\u00c5\u00c7\u0001"+
		"\u0000\u0000\u0000\u00c6\u00c4\u0001\u0000\u0000\u0000\u00c7\u00c8\u0005"+
		"\u0005\u0000\u0000\u00c8\u001d\u0001\u0000\u0000\u0000\u00c9\u00cb\u0003"+
		"(\u0014\u0000\u00ca\u00c9\u0001\u0000\u0000\u0000\u00cb\u00ce\u0001\u0000"+
		"\u0000\u0000\u00cc\u00ca\u0001\u0000\u0000\u0000\u00cc\u00cd\u0001\u0000"+
		"\u0000\u0000\u00cd\u00cf\u0001\u0000\u0000\u0000\u00ce\u00cc\u0001\u0000"+
		"\u0000\u0000\u00cf\u00d2\u0005\u001d\u0000\u0000\u00d0\u00d1\u0005\u000e"+
		"\u0000\u0000\u00d1\u00d3\u0005\u001a\u0000\u0000\u00d2\u00d0\u0001\u0000"+
		"\u0000\u0000\u00d2\u00d3\u0001\u0000\u0000\u0000\u00d3\u00d5\u0001\u0000"+
		"\u0000\u0000\u00d4\u00d6\u0005\n\u0000\u0000\u00d5\u00d4\u0001\u0000\u0000"+
		"\u0000\u00d5\u00d6\u0001\u0000\u0000\u0000\u00d6\u001f\u0001\u0000\u0000"+
		"\u0000\u00d7\u00da\u0003$\u0012\u0000\u00d8\u00da\u0003&\u0013\u0000\u00d9"+
		"\u00d7\u0001\u0000\u0000\u0000\u00d9\u00d8\u0001\u0000\u0000\u0000\u00da"+
		"\u00dc\u0001\u0000\u0000\u0000\u00db\u00dd\u0003\"\u0011\u0000\u00dc\u00db"+
		"\u0001\u0000\u0000\u0000\u00dc\u00dd\u0001\u0000\u0000\u0000\u00dd!\u0001"+
		"\u0000\u0000\u0000\u00de\u00df\u0005\u000f\u0000\u0000\u00df\u00e0\u0005"+
		"\u0010\u0000\u0000\u00e0#\u0001\u0000\u0000\u0000\u00e1\u00ea\u0005\u0011"+
		"\u0000\u0000\u00e2\u00ea\u0005\u0012\u0000\u0000\u00e3\u00ea\u0005\u0013"+
		"\u0000\u0000\u00e4\u00ea\u0005\u0014\u0000\u0000\u00e5\u00ea\u0005\u0015"+
		"\u0000\u0000\u00e6\u00ea\u0005\u0016\u0000\u0000\u00e7\u00ea\u0005\u0017"+
		"\u0000\u0000\u00e8\u00ea\u0005\u0018\u0000\u0000\u00e9\u00e1\u0001\u0000"+
		"\u0000\u0000\u00e9\u00e2\u0001\u0000\u0000\u0000\u00e9\u00e3\u0001\u0000"+
		"\u0000\u0000\u00e9\u00e4\u0001\u0000\u0000\u0000\u00e9\u00e5\u0001\u0000"+
		"\u0000\u0000\u00e9\u00e6\u0001\u0000\u0000\u0000\u00e9\u00e7\u0001\u0000"+
		"\u0000\u0000\u00e9\u00e8\u0001\u0000\u0000\u0000\u00ea%\u0001\u0000\u0000"+
		"\u0000\u00eb\u00ec\u0005\u001d\u0000\u0000\u00ec\'\u0001\u0000\u0000\u0000"+
		"\u00ed\u00f0\u0005 \u0000\u0000\u00ee\u00f0\u0005\u001f\u0000\u0000\u00ef"+
		"\u00ed\u0001\u0000\u0000\u0000\u00ef\u00ee\u0001\u0000\u0000\u0000\u00f0"+
		")\u0001\u0000\u0000\u0000\u001c.5;INW_dhqy~\u0087\u008c\u0095\u009d\u00a6"+
		"\u00ae\u00b2\u00bb\u00c4\u00cc\u00d2\u00d5\u00d9\u00dc\u00e9\u00ef";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}