// Generated from /Users/jryannel/dev/apigear/cli/pkg/idl/parser/ObjectApi.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ObjectApiParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.9.2", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, WHITESPACE=20, INTEGER=21, HEX=22, TYPE_IDENTIFIER=23, 
		IDENTIFIER=24, VERSION=25, DOCLINE=26, TAGLINE=27, COMMENT=28;
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
			null, "'module'", "'import'", "'interface'", "'{'", "'}'", "':'", "'('", 
			"')'", "','", "'signal'", "'struct'", "'enum'", "'='", "'['", "']'", 
			"'bool'", "'int'", "'float'", "'string'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, "WHITESPACE", "INTEGER", 
			"HEX", "TYPE_IDENTIFIER", "IDENTIFIER", "VERSION", "DOCLINE", "TAGLINE", 
			"COMMENT"
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
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__2) | (1L << T__10) | (1L << T__11) | (1L << DOCLINE) | (1L << TAGLINE))) != 0)) {
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

	public static class ImportRuleContext extends ParserRuleContext {
		public Token name;
		public Token version;
		public TerminalNode IDENTIFIER() { return getToken(ObjectApiParser.IDENTIFIER, 0); }
		public TerminalNode VERSION() { return getToken(ObjectApiParser.VERSION, 0); }
		public ImportRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importRule; }
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
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__9) | (1L << IDENTIFIER) | (1L << DOCLINE) | (1L << TAGLINE))) != 0)) {
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

	public static class PropertyRuleContext extends ParserRuleContext {
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
			setState(103);
			((PropertyRuleContext)_localctx).name = match(IDENTIFIER);
			setState(104);
			match(T__5);
			setState(105);
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
	}

	public final OperationRuleContext operationRule() throws RecognitionException {
		OperationRuleContext _localctx = new OperationRuleContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_operationRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(110);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOCLINE || _la==TAGLINE) {
				{
				{
				setState(107);
				metaRule();
				}
				}
				setState(112);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(113);
			((OperationRuleContext)_localctx).name = match(IDENTIFIER);
			setState(114);
			match(T__6);
			setState(118);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENTIFIER) {
				{
				{
				setState(115);
				((OperationRuleContext)_localctx).params = operationParamRule();
				}
				}
				setState(120);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(121);
			match(T__7);
			setState(123);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__5) {
				{
				setState(122);
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

	public static class OperationReturnRuleContext extends ParserRuleContext {
		public SchemaRuleContext schema;
		public SchemaRuleContext schemaRule() {
			return getRuleContext(SchemaRuleContext.class,0);
		}
		public OperationReturnRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_operationReturnRule; }
	}

	public final OperationReturnRuleContext operationReturnRule() throws RecognitionException {
		OperationReturnRuleContext _localctx = new OperationReturnRuleContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_operationReturnRule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(125);
			match(T__5);
			setState(126);
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
	}

	public final OperationParamRuleContext operationParamRule() throws RecognitionException {
		OperationParamRuleContext _localctx = new OperationParamRuleContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_operationParamRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(128);
			((OperationParamRuleContext)_localctx).name = match(IDENTIFIER);
			setState(129);
			match(T__5);
			setState(130);
			((OperationParamRuleContext)_localctx).schema = schemaRule();
			setState(132);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__8) {
				{
				setState(131);
				match(T__8);
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
	}

	public final SignalRuleContext signalRule() throws RecognitionException {
		SignalRuleContext _localctx = new SignalRuleContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_signalRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(137);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOCLINE || _la==TAGLINE) {
				{
				{
				setState(134);
				metaRule();
				}
				}
				setState(139);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(140);
			match(T__9);
			setState(141);
			((SignalRuleContext)_localctx).name = match(IDENTIFIER);
			setState(142);
			match(T__6);
			setState(146);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENTIFIER) {
				{
				{
				setState(143);
				((SignalRuleContext)_localctx).params = operationParamRule();
				}
				}
				setState(148);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(149);
			match(T__7);
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
	}

	public final StructRuleContext structRule() throws RecognitionException {
		StructRuleContext _localctx = new StructRuleContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_structRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(154);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOCLINE || _la==TAGLINE) {
				{
				{
				setState(151);
				metaRule();
				}
				}
				setState(156);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(157);
			match(T__10);
			setState(158);
			((StructRuleContext)_localctx).name = match(IDENTIFIER);
			setState(159);
			match(T__3);
			setState(163);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << IDENTIFIER) | (1L << DOCLINE) | (1L << TAGLINE))) != 0)) {
				{
				{
				setState(160);
				structFieldRule();
				}
				}
				setState(165);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(166);
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

	public static class StructFieldRuleContext extends ParserRuleContext {
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
	}

	public final StructFieldRuleContext structFieldRule() throws RecognitionException {
		StructFieldRuleContext _localctx = new StructFieldRuleContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_structFieldRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(171);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOCLINE || _la==TAGLINE) {
				{
				{
				setState(168);
				metaRule();
				}
				}
				setState(173);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(174);
			((StructFieldRuleContext)_localctx).name = match(IDENTIFIER);
			setState(175);
			match(T__5);
			setState(176);
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
	}

	public final EnumRuleContext enumRule() throws RecognitionException {
		EnumRuleContext _localctx = new EnumRuleContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_enumRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(181);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOCLINE || _la==TAGLINE) {
				{
				{
				setState(178);
				metaRule();
				}
				}
				setState(183);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(184);
			match(T__11);
			setState(185);
			((EnumRuleContext)_localctx).name = match(IDENTIFIER);
			setState(186);
			match(T__3);
			setState(190);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << IDENTIFIER) | (1L << DOCLINE) | (1L << TAGLINE))) != 0)) {
				{
				{
				setState(187);
				enumMemberRule();
				}
				}
				setState(192);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(193);
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
	}

	public final EnumMemberRuleContext enumMemberRule() throws RecognitionException {
		EnumMemberRuleContext _localctx = new EnumMemberRuleContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_enumMemberRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(198);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOCLINE || _la==TAGLINE) {
				{
				{
				setState(195);
				metaRule();
				}
				}
				setState(200);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(201);
			((EnumMemberRuleContext)_localctx).name = match(IDENTIFIER);
			setState(204);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__12) {
				{
				setState(202);
				match(T__12);
				setState(203);
				((EnumMemberRuleContext)_localctx).value = match(INTEGER);
				}
			}

			setState(207);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__8) {
				{
				setState(206);
				match(T__8);
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
	}

	public final SchemaRuleContext schemaRule() throws RecognitionException {
		SchemaRuleContext _localctx = new SchemaRuleContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_schemaRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(211);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__15:
			case T__16:
			case T__17:
			case T__18:
				{
				setState(209);
				primitiveSchema();
				}
				break;
			case IDENTIFIER:
				{
				setState(210);
				symbolSchema();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(214);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__13) {
				{
				setState(213);
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

	public static class ArrayRuleContext extends ParserRuleContext {
		public ArrayRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrayRule; }
	}

	public final ArrayRuleContext arrayRule() throws RecognitionException {
		ArrayRuleContext _localctx = new ArrayRuleContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_arrayRule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(216);
			match(T__13);
			setState(217);
			match(T__14);
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

	public static class PrimitiveSchemaContext extends ParserRuleContext {
		public Token name;
		public PrimitiveSchemaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primitiveSchema; }
	}

	public final PrimitiveSchemaContext primitiveSchema() throws RecognitionException {
		PrimitiveSchemaContext _localctx = new PrimitiveSchemaContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_primitiveSchema);
		try {
			setState(223);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__15:
				enterOuterAlt(_localctx, 1);
				{
				setState(219);
				((PrimitiveSchemaContext)_localctx).name = match(T__15);
				}
				break;
			case T__16:
				enterOuterAlt(_localctx, 2);
				{
				setState(220);
				((PrimitiveSchemaContext)_localctx).name = match(T__16);
				}
				break;
			case T__17:
				enterOuterAlt(_localctx, 3);
				{
				setState(221);
				((PrimitiveSchemaContext)_localctx).name = match(T__17);
				}
				break;
			case T__18:
				enterOuterAlt(_localctx, 4);
				{
				setState(222);
				((PrimitiveSchemaContext)_localctx).name = match(T__18);
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

	public static class SymbolSchemaContext extends ParserRuleContext {
		public Token name;
		public TerminalNode IDENTIFIER() { return getToken(ObjectApiParser.IDENTIFIER, 0); }
		public SymbolSchemaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_symbolSchema; }
	}

	public final SymbolSchemaContext symbolSchema() throws RecognitionException {
		SymbolSchemaContext _localctx = new SymbolSchemaContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_symbolSchema);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(225);
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

	public static class MetaRuleContext extends ParserRuleContext {
		public Token tagLine;
		public Token docLine;
		public TerminalNode TAGLINE() { return getToken(ObjectApiParser.TAGLINE, 0); }
		public TerminalNode DOCLINE() { return getToken(ObjectApiParser.DOCLINE, 0); }
		public MetaRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_metaRule; }
	}

	public final MetaRuleContext metaRule() throws RecognitionException {
		MetaRuleContext _localctx = new MetaRuleContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_metaRule);
		try {
			setState(229);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TAGLINE:
				enterOuterAlt(_localctx, 1);
				{
				setState(227);
				((MetaRuleContext)_localctx).tagLine = match(TAGLINE);
				}
				break;
			case DOCLINE:
				enterOuterAlt(_localctx, 2);
				{
				setState(228);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\36\u00ea\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\3\2\3\2\7\2/\n\2\f\2\16\2\62"+
		"\13\2\3\3\3\3\7\3\66\n\3\f\3\16\39\13\3\3\4\7\4<\n\4\f\4\16\4?\13\4\3"+
		"\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\6\3\6\3\6\5\6L\n\6\3\7\7\7O\n\7\f\7\16"+
		"\7R\13\7\3\7\3\7\3\7\3\7\7\7X\n\7\f\7\16\7[\13\7\3\7\3\7\3\b\3\b\3\b\5"+
		"\bb\n\b\3\t\7\te\n\t\f\t\16\th\13\t\3\t\3\t\3\t\3\t\3\n\7\no\n\n\f\n\16"+
		"\nr\13\n\3\n\3\n\3\n\7\nw\n\n\f\n\16\nz\13\n\3\n\3\n\5\n~\n\n\3\13\3\13"+
		"\3\13\3\f\3\f\3\f\3\f\5\f\u0087\n\f\3\r\7\r\u008a\n\r\f\r\16\r\u008d\13"+
		"\r\3\r\3\r\3\r\3\r\7\r\u0093\n\r\f\r\16\r\u0096\13\r\3\r\3\r\3\16\7\16"+
		"\u009b\n\16\f\16\16\16\u009e\13\16\3\16\3\16\3\16\3\16\7\16\u00a4\n\16"+
		"\f\16\16\16\u00a7\13\16\3\16\3\16\3\17\7\17\u00ac\n\17\f\17\16\17\u00af"+
		"\13\17\3\17\3\17\3\17\3\17\3\20\7\20\u00b6\n\20\f\20\16\20\u00b9\13\20"+
		"\3\20\3\20\3\20\3\20\7\20\u00bf\n\20\f\20\16\20\u00c2\13\20\3\20\3\20"+
		"\3\21\7\21\u00c7\n\21\f\21\16\21\u00ca\13\21\3\21\3\21\3\21\5\21\u00cf"+
		"\n\21\3\21\5\21\u00d2\n\21\3\22\3\22\5\22\u00d6\n\22\3\22\5\22\u00d9\n"+
		"\22\3\23\3\23\3\23\3\24\3\24\3\24\3\24\5\24\u00e2\n\24\3\25\3\25\3\26"+
		"\3\26\5\26\u00e8\n\26\3\26\2\2\27\2\4\6\b\n\f\16\20\22\24\26\30\32\34"+
		"\36 \"$&(*\2\2\2\u00f2\2,\3\2\2\2\4\63\3\2\2\2\6=\3\2\2\2\bD\3\2\2\2\n"+
		"K\3\2\2\2\fP\3\2\2\2\16a\3\2\2\2\20f\3\2\2\2\22p\3\2\2\2\24\177\3\2\2"+
		"\2\26\u0082\3\2\2\2\30\u008b\3\2\2\2\32\u009c\3\2\2\2\34\u00ad\3\2\2\2"+
		"\36\u00b7\3\2\2\2 \u00c8\3\2\2\2\"\u00d5\3\2\2\2$\u00da\3\2\2\2&\u00e1"+
		"\3\2\2\2(\u00e3\3\2\2\2*\u00e7\3\2\2\2,\60\5\4\3\2-/\5\n\6\2.-\3\2\2\2"+
		"/\62\3\2\2\2\60.\3\2\2\2\60\61\3\2\2\2\61\3\3\2\2\2\62\60\3\2\2\2\63\67"+
		"\5\6\4\2\64\66\5\b\5\2\65\64\3\2\2\2\669\3\2\2\2\67\65\3\2\2\2\678\3\2"+
		"\2\28\5\3\2\2\29\67\3\2\2\2:<\5*\26\2;:\3\2\2\2<?\3\2\2\2=;\3\2\2\2=>"+
		"\3\2\2\2>@\3\2\2\2?=\3\2\2\2@A\7\3\2\2AB\7\32\2\2BC\7\33\2\2C\7\3\2\2"+
		"\2DE\7\4\2\2EF\7\32\2\2FG\7\33\2\2G\t\3\2\2\2HL\5\f\7\2IL\5\32\16\2JL"+
		"\5\36\20\2KH\3\2\2\2KI\3\2\2\2KJ\3\2\2\2L\13\3\2\2\2MO\5*\26\2NM\3\2\2"+
		"\2OR\3\2\2\2PN\3\2\2\2PQ\3\2\2\2QS\3\2\2\2RP\3\2\2\2ST\7\5\2\2TU\7\32"+
		"\2\2UY\7\6\2\2VX\5\16\b\2WV\3\2\2\2X[\3\2\2\2YW\3\2\2\2YZ\3\2\2\2Z\\\3"+
		"\2\2\2[Y\3\2\2\2\\]\7\7\2\2]\r\3\2\2\2^b\5\20\t\2_b\5\22\n\2`b\5\30\r"+
		"\2a^\3\2\2\2a_\3\2\2\2a`\3\2\2\2b\17\3\2\2\2ce\5*\26\2dc\3\2\2\2eh\3\2"+
		"\2\2fd\3\2\2\2fg\3\2\2\2gi\3\2\2\2hf\3\2\2\2ij\7\32\2\2jk\7\b\2\2kl\5"+
		"\"\22\2l\21\3\2\2\2mo\5*\26\2nm\3\2\2\2or\3\2\2\2pn\3\2\2\2pq\3\2\2\2"+
		"qs\3\2\2\2rp\3\2\2\2st\7\32\2\2tx\7\t\2\2uw\5\26\f\2vu\3\2\2\2wz\3\2\2"+
		"\2xv\3\2\2\2xy\3\2\2\2y{\3\2\2\2zx\3\2\2\2{}\7\n\2\2|~\5\24\13\2}|\3\2"+
		"\2\2}~\3\2\2\2~\23\3\2\2\2\177\u0080\7\b\2\2\u0080\u0081\5\"\22\2\u0081"+
		"\25\3\2\2\2\u0082\u0083\7\32\2\2\u0083\u0084\7\b\2\2\u0084\u0086\5\"\22"+
		"\2\u0085\u0087\7\13\2\2\u0086\u0085\3\2\2\2\u0086\u0087\3\2\2\2\u0087"+
		"\27\3\2\2\2\u0088\u008a\5*\26\2\u0089\u0088\3\2\2\2\u008a\u008d\3\2\2"+
		"\2\u008b\u0089\3\2\2\2\u008b\u008c\3\2\2\2\u008c\u008e\3\2\2\2\u008d\u008b"+
		"\3\2\2\2\u008e\u008f\7\f\2\2\u008f\u0090\7\32\2\2\u0090\u0094\7\t\2\2"+
		"\u0091\u0093\5\26\f\2\u0092\u0091\3\2\2\2\u0093\u0096\3\2\2\2\u0094\u0092"+
		"\3\2\2\2\u0094\u0095\3\2\2\2\u0095\u0097\3\2\2\2\u0096\u0094\3\2\2\2\u0097"+
		"\u0098\7\n\2\2\u0098\31\3\2\2\2\u0099\u009b\5*\26\2\u009a\u0099\3\2\2"+
		"\2\u009b\u009e\3\2\2\2\u009c\u009a\3\2\2\2\u009c\u009d\3\2\2\2\u009d\u009f"+
		"\3\2\2\2\u009e\u009c\3\2\2\2\u009f\u00a0\7\r\2\2\u00a0\u00a1\7\32\2\2"+
		"\u00a1\u00a5\7\6\2\2\u00a2\u00a4\5\34\17\2\u00a3\u00a2\3\2\2\2\u00a4\u00a7"+
		"\3\2\2\2\u00a5\u00a3\3\2\2\2\u00a5\u00a6\3\2\2\2\u00a6\u00a8\3\2\2\2\u00a7"+
		"\u00a5\3\2\2\2\u00a8\u00a9\7\7\2\2\u00a9\33\3\2\2\2\u00aa\u00ac\5*\26"+
		"\2\u00ab\u00aa\3\2\2\2\u00ac\u00af\3\2\2\2\u00ad\u00ab\3\2\2\2\u00ad\u00ae"+
		"\3\2\2\2\u00ae\u00b0\3\2\2\2\u00af\u00ad\3\2\2\2\u00b0\u00b1\7\32\2\2"+
		"\u00b1\u00b2\7\b\2\2\u00b2\u00b3\5\"\22\2\u00b3\35\3\2\2\2\u00b4\u00b6"+
		"\5*\26\2\u00b5\u00b4\3\2\2\2\u00b6\u00b9\3\2\2\2\u00b7\u00b5\3\2\2\2\u00b7"+
		"\u00b8\3\2\2\2\u00b8\u00ba\3\2\2\2\u00b9\u00b7\3\2\2\2\u00ba\u00bb\7\16"+
		"\2\2\u00bb\u00bc\7\32\2\2\u00bc\u00c0\7\6\2\2\u00bd\u00bf\5 \21\2\u00be"+
		"\u00bd\3\2\2\2\u00bf\u00c2\3\2\2\2\u00c0\u00be\3\2\2\2\u00c0\u00c1\3\2"+
		"\2\2\u00c1\u00c3\3\2\2\2\u00c2\u00c0\3\2\2\2\u00c3\u00c4\7\7\2\2\u00c4"+
		"\37\3\2\2\2\u00c5\u00c7\5*\26\2\u00c6\u00c5\3\2\2\2\u00c7\u00ca\3\2\2"+
		"\2\u00c8\u00c6\3\2\2\2\u00c8\u00c9\3\2\2\2\u00c9\u00cb\3\2\2\2\u00ca\u00c8"+
		"\3\2\2\2\u00cb\u00ce\7\32\2\2\u00cc\u00cd\7\17\2\2\u00cd\u00cf\7\27\2"+
		"\2\u00ce\u00cc\3\2\2\2\u00ce\u00cf\3\2\2\2\u00cf\u00d1\3\2\2\2\u00d0\u00d2"+
		"\7\13\2\2\u00d1\u00d0\3\2\2\2\u00d1\u00d2\3\2\2\2\u00d2!\3\2\2\2\u00d3"+
		"\u00d6\5&\24\2\u00d4\u00d6\5(\25\2\u00d5\u00d3\3\2\2\2\u00d5\u00d4\3\2"+
		"\2\2\u00d6\u00d8\3\2\2\2\u00d7\u00d9\5$\23\2\u00d8\u00d7\3\2\2\2\u00d8"+
		"\u00d9\3\2\2\2\u00d9#\3\2\2\2\u00da\u00db\7\20\2\2\u00db\u00dc\7\21\2"+
		"\2\u00dc%\3\2\2\2\u00dd\u00e2\7\22\2\2\u00de\u00e2\7\23\2\2\u00df\u00e2"+
		"\7\24\2\2\u00e0\u00e2\7\25\2\2\u00e1\u00dd\3\2\2\2\u00e1\u00de\3\2\2\2"+
		"\u00e1\u00df\3\2\2\2\u00e1\u00e0\3\2\2\2\u00e2\'\3\2\2\2\u00e3\u00e4\7"+
		"\32\2\2\u00e4)\3\2\2\2\u00e5\u00e8\7\35\2\2\u00e6\u00e8\7\34\2\2\u00e7"+
		"\u00e5\3\2\2\2\u00e7\u00e6\3\2\2\2\u00e8+\3\2\2\2\34\60\67=KPYafpx}\u0086"+
		"\u008b\u0094\u009c\u00a5\u00ad\u00b7\u00c0\u00c8\u00ce\u00d1\u00d5\u00d8"+
		"\u00e1\u00e7";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}