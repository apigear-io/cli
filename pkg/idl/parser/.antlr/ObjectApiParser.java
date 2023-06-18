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
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__2) | (1L << T__11) | (1L << T__12) | (1L << DOCLINE) | (1L << TAGLINE))) != 0)) {
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
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__5) | (1L << T__10) | (1L << IDENTIFIER) | (1L << DOCLINE) | (1L << TAGLINE))) != 0)) {
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
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << IDENTIFIER) | (1L << DOCLINE) | (1L << TAGLINE))) != 0)) {
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
			setState(177);
			((StructFieldRuleContext)_localctx).name = match(IDENTIFIER);
			setState(178);
			match(T__6);
			setState(179);
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
			setState(184);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOCLINE || _la==TAGLINE) {
				{
				{
				setState(181);
				metaRule();
				}
				}
				setState(186);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(187);
			match(T__12);
			setState(188);
			((EnumRuleContext)_localctx).name = match(IDENTIFIER);
			setState(189);
			match(T__3);
			setState(193);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << IDENTIFIER) | (1L << DOCLINE) | (1L << TAGLINE))) != 0)) {
				{
				{
				setState(190);
				enumMemberRule();
				}
				}
				setState(195);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(196);
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
			setState(201);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOCLINE || _la==TAGLINE) {
				{
				{
				setState(198);
				metaRule();
				}
				}
				setState(203);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(204);
			((EnumMemberRuleContext)_localctx).name = match(IDENTIFIER);
			setState(207);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__13) {
				{
				setState(205);
				match(T__13);
				setState(206);
				((EnumMemberRuleContext)_localctx).value = match(INTEGER);
				}
			}

			setState(210);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__9) {
				{
				setState(209);
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
			setState(214);
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
				setState(212);
				primitiveSchema();
				}
				break;
			case IDENTIFIER:
				{
				setState(213);
				symbolSchema();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(217);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__14) {
				{
				setState(216);
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
			setState(219);
			match(T__14);
			setState(220);
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
			setState(230);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__16:
				enterOuterAlt(_localctx, 1);
				{
				setState(222);
				((PrimitiveSchemaContext)_localctx).name = match(T__16);
				}
				break;
			case T__17:
				enterOuterAlt(_localctx, 2);
				{
				setState(223);
				((PrimitiveSchemaContext)_localctx).name = match(T__17);
				}
				break;
			case T__18:
				enterOuterAlt(_localctx, 3);
				{
				setState(224);
				((PrimitiveSchemaContext)_localctx).name = match(T__18);
				}
				break;
			case T__19:
				enterOuterAlt(_localctx, 4);
				{
				setState(225);
				((PrimitiveSchemaContext)_localctx).name = match(T__19);
				}
				break;
			case T__20:
				enterOuterAlt(_localctx, 5);
				{
				setState(226);
				((PrimitiveSchemaContext)_localctx).name = match(T__20);
				}
				break;
			case T__21:
				enterOuterAlt(_localctx, 6);
				{
				setState(227);
				((PrimitiveSchemaContext)_localctx).name = match(T__21);
				}
				break;
			case T__22:
				enterOuterAlt(_localctx, 7);
				{
				setState(228);
				((PrimitiveSchemaContext)_localctx).name = match(T__22);
				}
				break;
			case T__23:
				enterOuterAlt(_localctx, 8);
				{
				setState(229);
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
			setState(232);
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
			setState(236);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TAGLINE:
				enterOuterAlt(_localctx, 1);
				{
				setState(234);
				((MetaRuleContext)_localctx).tagLine = match(TAGLINE);
				}
				break;
			case DOCLINE:
				enterOuterAlt(_localctx, 2);
				{
				setState(235);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3#\u00f1\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\3\2\3\2\7\2/\n\2\f\2\16\2\62"+
		"\13\2\3\3\3\3\7\3\66\n\3\f\3\16\39\13\3\3\4\7\4<\n\4\f\4\16\4?\13\4\3"+
		"\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\6\3\6\3\6\5\6L\n\6\3\7\7\7O\n\7\f\7\16"+
		"\7R\13\7\3\7\3\7\3\7\3\7\7\7X\n\7\f\7\16\7[\13\7\3\7\3\7\3\b\3\b\3\b\5"+
		"\bb\n\b\3\t\7\te\n\t\f\t\16\th\13\t\3\t\5\tk\n\t\3\t\3\t\3\t\3\t\3\n\7"+
		"\nr\n\n\f\n\16\nu\13\n\3\n\3\n\3\n\7\nz\n\n\f\n\16\n}\13\n\3\n\3\n\5\n"+
		"\u0081\n\n\3\13\3\13\3\13\3\f\3\f\3\f\3\f\5\f\u008a\n\f\3\r\7\r\u008d"+
		"\n\r\f\r\16\r\u0090\13\r\3\r\3\r\3\r\3\r\7\r\u0096\n\r\f\r\16\r\u0099"+
		"\13\r\3\r\3\r\3\16\7\16\u009e\n\16\f\16\16\16\u00a1\13\16\3\16\3\16\3"+
		"\16\3\16\7\16\u00a7\n\16\f\16\16\16\u00aa\13\16\3\16\3\16\3\17\7\17\u00af"+
		"\n\17\f\17\16\17\u00b2\13\17\3\17\3\17\3\17\3\17\3\20\7\20\u00b9\n\20"+
		"\f\20\16\20\u00bc\13\20\3\20\3\20\3\20\3\20\7\20\u00c2\n\20\f\20\16\20"+
		"\u00c5\13\20\3\20\3\20\3\21\7\21\u00ca\n\21\f\21\16\21\u00cd\13\21\3\21"+
		"\3\21\3\21\5\21\u00d2\n\21\3\21\5\21\u00d5\n\21\3\22\3\22\5\22\u00d9\n"+
		"\22\3\22\5\22\u00dc\n\22\3\23\3\23\3\23\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\5\24\u00e9\n\24\3\25\3\25\3\26\3\26\5\26\u00ef\n\26\3\26\2"+
		"\2\27\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*\2\2\2\u00fe\2,\3\2"+
		"\2\2\4\63\3\2\2\2\6=\3\2\2\2\bD\3\2\2\2\nK\3\2\2\2\fP\3\2\2\2\16a\3\2"+
		"\2\2\20f\3\2\2\2\22s\3\2\2\2\24\u0082\3\2\2\2\26\u0085\3\2\2\2\30\u008e"+
		"\3\2\2\2\32\u009f\3\2\2\2\34\u00b0\3\2\2\2\36\u00ba\3\2\2\2 \u00cb\3\2"+
		"\2\2\"\u00d8\3\2\2\2$\u00dd\3\2\2\2&\u00e8\3\2\2\2(\u00ea\3\2\2\2*\u00ee"+
		"\3\2\2\2,\60\5\4\3\2-/\5\n\6\2.-\3\2\2\2/\62\3\2\2\2\60.\3\2\2\2\60\61"+
		"\3\2\2\2\61\3\3\2\2\2\62\60\3\2\2\2\63\67\5\6\4\2\64\66\5\b\5\2\65\64"+
		"\3\2\2\2\669\3\2\2\2\67\65\3\2\2\2\678\3\2\2\28\5\3\2\2\29\67\3\2\2\2"+
		":<\5*\26\2;:\3\2\2\2<?\3\2\2\2=;\3\2\2\2=>\3\2\2\2>@\3\2\2\2?=\3\2\2\2"+
		"@A\7\3\2\2AB\7\37\2\2BC\7 \2\2C\7\3\2\2\2DE\7\4\2\2EF\7\37\2\2FG\7 \2"+
		"\2G\t\3\2\2\2HL\5\f\7\2IL\5\32\16\2JL\5\36\20\2KH\3\2\2\2KI\3\2\2\2KJ"+
		"\3\2\2\2L\13\3\2\2\2MO\5*\26\2NM\3\2\2\2OR\3\2\2\2PN\3\2\2\2PQ\3\2\2\2"+
		"QS\3\2\2\2RP\3\2\2\2ST\7\5\2\2TU\7\37\2\2UY\7\6\2\2VX\5\16\b\2WV\3\2\2"+
		"\2X[\3\2\2\2YW\3\2\2\2YZ\3\2\2\2Z\\\3\2\2\2[Y\3\2\2\2\\]\7\7\2\2]\r\3"+
		"\2\2\2^b\5\20\t\2_b\5\22\n\2`b\5\30\r\2a^\3\2\2\2a_\3\2\2\2a`\3\2\2\2"+
		"b\17\3\2\2\2ce\5*\26\2dc\3\2\2\2eh\3\2\2\2fd\3\2\2\2fg\3\2\2\2gj\3\2\2"+
		"\2hf\3\2\2\2ik\7\b\2\2ji\3\2\2\2jk\3\2\2\2kl\3\2\2\2lm\7\37\2\2mn\7\t"+
		"\2\2no\5\"\22\2o\21\3\2\2\2pr\5*\26\2qp\3\2\2\2ru\3\2\2\2sq\3\2\2\2st"+
		"\3\2\2\2tv\3\2\2\2us\3\2\2\2vw\7\37\2\2w{\7\n\2\2xz\5\26\f\2yx\3\2\2\2"+
		"z}\3\2\2\2{y\3\2\2\2{|\3\2\2\2|~\3\2\2\2}{\3\2\2\2~\u0080\7\13\2\2\177"+
		"\u0081\5\24\13\2\u0080\177\3\2\2\2\u0080\u0081\3\2\2\2\u0081\23\3\2\2"+
		"\2\u0082\u0083\7\t\2\2\u0083\u0084\5\"\22\2\u0084\25\3\2\2\2\u0085\u0086"+
		"\7\37\2\2\u0086\u0087\7\t\2\2\u0087\u0089\5\"\22\2\u0088\u008a\7\f\2\2"+
		"\u0089\u0088\3\2\2\2\u0089\u008a\3\2\2\2\u008a\27\3\2\2\2\u008b\u008d"+
		"\5*\26\2\u008c\u008b\3\2\2\2\u008d\u0090\3\2\2\2\u008e\u008c\3\2\2\2\u008e"+
		"\u008f\3\2\2\2\u008f\u0091\3\2\2\2\u0090\u008e\3\2\2\2\u0091\u0092\7\r"+
		"\2\2\u0092\u0093\7\37\2\2\u0093\u0097\7\n\2\2\u0094\u0096\5\26\f\2\u0095"+
		"\u0094\3\2\2\2\u0096\u0099\3\2\2\2\u0097\u0095\3\2\2\2\u0097\u0098\3\2"+
		"\2\2\u0098\u009a\3\2\2\2\u0099\u0097\3\2\2\2\u009a\u009b\7\13\2\2\u009b"+
		"\31\3\2\2\2\u009c\u009e\5*\26\2\u009d\u009c\3\2\2\2\u009e\u00a1\3\2\2"+
		"\2\u009f\u009d\3\2\2\2\u009f\u00a0\3\2\2\2\u00a0\u00a2\3\2\2\2\u00a1\u009f"+
		"\3\2\2\2\u00a2\u00a3\7\16\2\2\u00a3\u00a4\7\37\2\2\u00a4\u00a8\7\6\2\2"+
		"\u00a5\u00a7\5\34\17\2\u00a6\u00a5\3\2\2\2\u00a7\u00aa\3\2\2\2\u00a8\u00a6"+
		"\3\2\2\2\u00a8\u00a9\3\2\2\2\u00a9\u00ab\3\2\2\2\u00aa\u00a8\3\2\2\2\u00ab"+
		"\u00ac\7\7\2\2\u00ac\33\3\2\2\2\u00ad\u00af\5*\26\2\u00ae\u00ad\3\2\2"+
		"\2\u00af\u00b2\3\2\2\2\u00b0\u00ae\3\2\2\2\u00b0\u00b1\3\2\2\2\u00b1\u00b3"+
		"\3\2\2\2\u00b2\u00b0\3\2\2\2\u00b3\u00b4\7\37\2\2\u00b4\u00b5\7\t\2\2"+
		"\u00b5\u00b6\5\"\22\2\u00b6\35\3\2\2\2\u00b7\u00b9\5*\26\2\u00b8\u00b7"+
		"\3\2\2\2\u00b9\u00bc\3\2\2\2\u00ba\u00b8\3\2\2\2\u00ba\u00bb\3\2\2\2\u00bb"+
		"\u00bd\3\2\2\2\u00bc\u00ba\3\2\2\2\u00bd\u00be\7\17\2\2\u00be\u00bf\7"+
		"\37\2\2\u00bf\u00c3\7\6\2\2\u00c0\u00c2\5 \21\2\u00c1\u00c0\3\2\2\2\u00c2"+
		"\u00c5\3\2\2\2\u00c3\u00c1\3\2\2\2\u00c3\u00c4\3\2\2\2\u00c4\u00c6\3\2"+
		"\2\2\u00c5\u00c3\3\2\2\2\u00c6\u00c7\7\7\2\2\u00c7\37\3\2\2\2\u00c8\u00ca"+
		"\5*\26\2\u00c9\u00c8\3\2\2\2\u00ca\u00cd\3\2\2\2\u00cb\u00c9\3\2\2\2\u00cb"+
		"\u00cc\3\2\2\2\u00cc\u00ce\3\2\2\2\u00cd\u00cb\3\2\2\2\u00ce\u00d1\7\37"+
		"\2\2\u00cf\u00d0\7\20\2\2\u00d0\u00d2\7\34\2\2\u00d1\u00cf\3\2\2\2\u00d1"+
		"\u00d2\3\2\2\2\u00d2\u00d4\3\2\2\2\u00d3\u00d5\7\f\2\2\u00d4\u00d3\3\2"+
		"\2\2\u00d4\u00d5\3\2\2\2\u00d5!\3\2\2\2\u00d6\u00d9\5&\24\2\u00d7\u00d9"+
		"\5(\25\2\u00d8\u00d6\3\2\2\2\u00d8\u00d7\3\2\2\2\u00d9\u00db\3\2\2\2\u00da"+
		"\u00dc\5$\23\2\u00db\u00da\3\2\2\2\u00db\u00dc\3\2\2\2\u00dc#\3\2\2\2"+
		"\u00dd\u00de\7\21\2\2\u00de\u00df\7\22\2\2\u00df%\3\2\2\2\u00e0\u00e9"+
		"\7\23\2\2\u00e1\u00e9\7\24\2\2\u00e2\u00e9\7\25\2\2\u00e3\u00e9\7\26\2"+
		"\2\u00e4\u00e9\7\27\2\2\u00e5\u00e9\7\30\2\2\u00e6\u00e9\7\31\2\2\u00e7"+
		"\u00e9\7\32\2\2\u00e8\u00e0\3\2\2\2\u00e8\u00e1\3\2\2\2\u00e8\u00e2\3"+
		"\2\2\2\u00e8\u00e3\3\2\2\2\u00e8\u00e4\3\2\2\2\u00e8\u00e5\3\2\2\2\u00e8"+
		"\u00e6\3\2\2\2\u00e8\u00e7\3\2\2\2\u00e9\'\3\2\2\2\u00ea\u00eb\7\37\2"+
		"\2\u00eb)\3\2\2\2\u00ec\u00ef\7\"\2\2\u00ed\u00ef\7!\2\2\u00ee\u00ec\3"+
		"\2\2\2\u00ee\u00ed\3\2\2\2\u00ef+\3\2\2\2\35\60\67=KPYafjs{\u0080\u0089"+
		"\u008e\u0097\u009f\u00a8\u00b0\u00ba\u00c3\u00cb\u00d1\u00d4\u00d8\u00db"+
		"\u00e8\u00ee";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}