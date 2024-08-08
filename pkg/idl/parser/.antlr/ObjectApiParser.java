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
		T__24=25, T__25=26, T__26=27, WHITESPACE=28, INTEGER=29, HEX=30, IDENTIFIER=31, 
		VERSION=32, DOCLINE=33, TAGLINE=34, COMMENT=35, DOT=36, LETTER=37, DIGIT=38, 
		UNDERSCORE=39, SEMICOLON=40;
	public static final int
		RULE_documentRule = 0, RULE_headerRule = 1, RULE_moduleRule = 2, RULE_importRule = 3, 
		RULE_declarationsRule = 4, RULE_externRule = 5, RULE_interfaceRule = 6, 
		RULE_interfaceMembersRule = 7, RULE_propertyRule = 8, RULE_operationRule = 9, 
		RULE_operationReturnRule = 10, RULE_operationParamRule = 11, RULE_signalRule = 12, 
		RULE_structRule = 13, RULE_structFieldRule = 14, RULE_enumRule = 15, RULE_enumMemberRule = 16, 
		RULE_schemaRule = 17, RULE_arrayRule = 18, RULE_primitiveSchema = 19, 
		RULE_symbolSchema = 20, RULE_metaRule = 21;
	private static String[] makeRuleNames() {
		return new String[] {
			"documentRule", "headerRule", "moduleRule", "importRule", "declarationsRule", 
			"externRule", "interfaceRule", "interfaceMembersRule", "propertyRule", 
			"operationRule", "operationReturnRule", "operationParamRule", "signalRule", 
			"structRule", "structFieldRule", "enumRule", "enumMemberRule", "schemaRule", 
			"arrayRule", "primitiveSchema", "symbolSchema", "metaRule"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'module'", "'import'", "'extern'", "'interface'", "'extends'", 
			"'{'", "'}'", "'readonly'", "':'", "'('", "')'", "','", "'signal'", "'struct'", 
			"'enum'", "'='", "'['", "']'", "'bool'", "'int'", "'int32'", "'int64'", 
			"'float'", "'float32'", "'float64'", "'string'", "'bytes'", null, null, 
			null, null, null, null, null, null, "'.'", null, null, "'_'", "';'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, "WHITESPACE", "INTEGER", "HEX", "IDENTIFIER", 
			"VERSION", "DOCLINE", "TAGLINE", "COMMENT", "DOT", "LETTER", "DIGIT", 
			"UNDERSCORE", "SEMICOLON"
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
	}

	public final DocumentRuleContext documentRule() throws RecognitionException {
		DocumentRuleContext _localctx = new DocumentRuleContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_documentRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(44);
			headerRule();
			setState(48);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 25769852952L) != 0)) {
				{
				{
				setState(45);
				declarationsRule();
				}
				}
				setState(50);
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
	}

	public final HeaderRuleContext headerRule() throws RecognitionException {
		HeaderRuleContext _localctx = new HeaderRuleContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_headerRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(51);
			moduleRule();
			setState(55);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__1) {
				{
				{
				setState(52);
				importRule();
				}
				}
				setState(57);
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
		public List<MetaRuleContext> metaRule() {
			return getRuleContexts(MetaRuleContext.class);
		}
		public MetaRuleContext metaRule(int i) {
			return getRuleContext(MetaRuleContext.class,i);
		}
		public TerminalNode SEMICOLON() { return getToken(ObjectApiParser.SEMICOLON, 0); }
		public TerminalNode VERSION() { return getToken(ObjectApiParser.VERSION, 0); }
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
			setState(61);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOCLINE || _la==TAGLINE) {
				{
				{
				setState(58);
				metaRule();
				}
				}
				setState(63);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(64);
			match(T__0);
			setState(65);
			((ModuleRuleContext)_localctx).name = match(IDENTIFIER);
			setState(67);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==VERSION) {
				{
				setState(66);
				((ModuleRuleContext)_localctx).version = match(VERSION);
				}
			}

			setState(70);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMICOLON) {
				{
				setState(69);
				match(SEMICOLON);
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
	public static class ImportRuleContext extends ParserRuleContext {
		public Token name;
		public Token version;
		public TerminalNode IDENTIFIER() { return getToken(ObjectApiParser.IDENTIFIER, 0); }
		public TerminalNode SEMICOLON() { return getToken(ObjectApiParser.SEMICOLON, 0); }
		public TerminalNode VERSION() { return getToken(ObjectApiParser.VERSION, 0); }
		public ImportRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_importRule; }
	}

	public final ImportRuleContext importRule() throws RecognitionException {
		ImportRuleContext _localctx = new ImportRuleContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_importRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(72);
			match(T__1);
			setState(73);
			((ImportRuleContext)_localctx).name = match(IDENTIFIER);
			setState(75);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==VERSION) {
				{
				setState(74);
				((ImportRuleContext)_localctx).version = match(VERSION);
				}
			}

			setState(78);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMICOLON) {
				{
				setState(77);
				match(SEMICOLON);
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
	public static class DeclarationsRuleContext extends ParserRuleContext {
		public ExternRuleContext externRule() {
			return getRuleContext(ExternRuleContext.class,0);
		}
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
			setState(84);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(80);
				externRule();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(81);
				interfaceRule();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(82);
				structRule();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(83);
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
	public static class ExternRuleContext extends ParserRuleContext {
		public Token name;
		public TerminalNode IDENTIFIER() { return getToken(ObjectApiParser.IDENTIFIER, 0); }
		public List<MetaRuleContext> metaRule() {
			return getRuleContexts(MetaRuleContext.class);
		}
		public MetaRuleContext metaRule(int i) {
			return getRuleContext(MetaRuleContext.class,i);
		}
		public TerminalNode SEMICOLON() { return getToken(ObjectApiParser.SEMICOLON, 0); }
		public ExternRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_externRule; }
	}

	public final ExternRuleContext externRule() throws RecognitionException {
		ExternRuleContext _localctx = new ExternRuleContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_externRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(89);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOCLINE || _la==TAGLINE) {
				{
				{
				setState(86);
				metaRule();
				}
				}
				setState(91);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(92);
			match(T__2);
			setState(93);
			((ExternRuleContext)_localctx).name = match(IDENTIFIER);
			setState(95);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMICOLON) {
				{
				setState(94);
				match(SEMICOLON);
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
	public static class InterfaceRuleContext extends ParserRuleContext {
		public Token name;
		public Token extends_;
		public List<TerminalNode> IDENTIFIER() { return getTokens(ObjectApiParser.IDENTIFIER); }
		public TerminalNode IDENTIFIER(int i) {
			return getToken(ObjectApiParser.IDENTIFIER, i);
		}
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
		enterRule(_localctx, 12, RULE_interfaceRule);
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
			match(T__3);
			setState(104);
			((InterfaceRuleContext)_localctx).name = match(IDENTIFIER);
			setState(107);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__4) {
				{
				setState(105);
				match(T__4);
				setState(106);
				((InterfaceRuleContext)_localctx).extends_ = match(IDENTIFIER);
				}
			}

			setState(109);
			match(T__5);
			setState(113);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 27917295872L) != 0)) {
				{
				{
				setState(110);
				interfaceMembersRule();
				}
				}
				setState(115);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(116);
			match(T__6);
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
	}

	public final InterfaceMembersRuleContext interfaceMembersRule() throws RecognitionException {
		InterfaceMembersRuleContext _localctx = new InterfaceMembersRuleContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_interfaceMembersRule);
		try {
			setState(121);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(118);
				propertyRule();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(119);
				operationRule();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(120);
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
		public TerminalNode SEMICOLON() { return getToken(ObjectApiParser.SEMICOLON, 0); }
		public PropertyRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_propertyRule; }
	}

	public final PropertyRuleContext propertyRule() throws RecognitionException {
		PropertyRuleContext _localctx = new PropertyRuleContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_propertyRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(126);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOCLINE || _la==TAGLINE) {
				{
				{
				setState(123);
				metaRule();
				}
				}
				setState(128);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(130);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__7) {
				{
				setState(129);
				((PropertyRuleContext)_localctx).readonly = match(T__7);
				}
			}

			setState(132);
			((PropertyRuleContext)_localctx).name = match(IDENTIFIER);
			setState(133);
			match(T__8);
			setState(134);
			((PropertyRuleContext)_localctx).schema = schemaRule();
			setState(136);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMICOLON) {
				{
				setState(135);
				match(SEMICOLON);
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
		public TerminalNode SEMICOLON() { return getToken(ObjectApiParser.SEMICOLON, 0); }
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
		enterRule(_localctx, 18, RULE_operationRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(141);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOCLINE || _la==TAGLINE) {
				{
				{
				setState(138);
				metaRule();
				}
				}
				setState(143);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(144);
			((OperationRuleContext)_localctx).name = match(IDENTIFIER);
			setState(145);
			match(T__9);
			setState(149);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENTIFIER) {
				{
				{
				setState(146);
				((OperationRuleContext)_localctx).params = operationParamRule();
				}
				}
				setState(151);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(152);
			match(T__10);
			setState(154);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__8) {
				{
				setState(153);
				operationReturnRule();
				}
			}

			setState(157);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMICOLON) {
				{
				setState(156);
				match(SEMICOLON);
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
	}

	public final OperationReturnRuleContext operationReturnRule() throws RecognitionException {
		OperationReturnRuleContext _localctx = new OperationReturnRuleContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_operationReturnRule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(159);
			match(T__8);
			setState(160);
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
	}

	public final OperationParamRuleContext operationParamRule() throws RecognitionException {
		OperationParamRuleContext _localctx = new OperationParamRuleContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_operationParamRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(162);
			((OperationParamRuleContext)_localctx).name = match(IDENTIFIER);
			setState(163);
			match(T__8);
			setState(164);
			((OperationParamRuleContext)_localctx).schema = schemaRule();
			setState(166);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__11) {
				{
				setState(165);
				match(T__11);
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
		public TerminalNode SEMICOLON() { return getToken(ObjectApiParser.SEMICOLON, 0); }
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
		enterRule(_localctx, 24, RULE_signalRule);
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
			match(T__12);
			setState(175);
			((SignalRuleContext)_localctx).name = match(IDENTIFIER);
			setState(176);
			match(T__9);
			setState(180);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENTIFIER) {
				{
				{
				setState(177);
				((SignalRuleContext)_localctx).params = operationParamRule();
				}
				}
				setState(182);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(183);
			match(T__10);
			setState(185);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMICOLON) {
				{
				setState(184);
				match(SEMICOLON);
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
		enterRule(_localctx, 26, RULE_structRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(190);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOCLINE || _la==TAGLINE) {
				{
				{
				setState(187);
				metaRule();
				}
				}
				setState(192);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(193);
			match(T__13);
			setState(194);
			((StructRuleContext)_localctx).name = match(IDENTIFIER);
			setState(195);
			match(T__5);
			setState(199);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 27917287680L) != 0)) {
				{
				{
				setState(196);
				structFieldRule();
				}
				}
				setState(201);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(202);
			match(T__6);
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
		public TerminalNode SEMICOLON() { return getToken(ObjectApiParser.SEMICOLON, 0); }
		public StructFieldRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structFieldRule; }
	}

	public final StructFieldRuleContext structFieldRule() throws RecognitionException {
		StructFieldRuleContext _localctx = new StructFieldRuleContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_structFieldRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(207);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOCLINE || _la==TAGLINE) {
				{
				{
				setState(204);
				metaRule();
				}
				}
				setState(209);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(211);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__7) {
				{
				setState(210);
				((StructFieldRuleContext)_localctx).readonly = match(T__7);
				}
			}

			setState(213);
			((StructFieldRuleContext)_localctx).name = match(IDENTIFIER);
			setState(214);
			match(T__8);
			setState(215);
			((StructFieldRuleContext)_localctx).schema = schemaRule();
			setState(217);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMICOLON) {
				{
				setState(216);
				match(SEMICOLON);
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
		enterRule(_localctx, 30, RULE_enumRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(222);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOCLINE || _la==TAGLINE) {
				{
				{
				setState(219);
				metaRule();
				}
				}
				setState(224);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(225);
			match(T__14);
			setState(226);
			((EnumRuleContext)_localctx).name = match(IDENTIFIER);
			setState(227);
			match(T__5);
			setState(231);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 27917287424L) != 0)) {
				{
				{
				setState(228);
				enumMemberRule();
				}
				}
				setState(233);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(234);
			match(T__6);
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
	}

	public final EnumMemberRuleContext enumMemberRule() throws RecognitionException {
		EnumMemberRuleContext _localctx = new EnumMemberRuleContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_enumMemberRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(239);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==DOCLINE || _la==TAGLINE) {
				{
				{
				setState(236);
				metaRule();
				}
				}
				setState(241);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(242);
			((EnumMemberRuleContext)_localctx).name = match(IDENTIFIER);
			setState(245);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__15) {
				{
				setState(243);
				match(T__15);
				setState(244);
				((EnumMemberRuleContext)_localctx).value = match(INTEGER);
				}
			}

			setState(248);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__11) {
				{
				setState(247);
				match(T__11);
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
	}

	public final SchemaRuleContext schemaRule() throws RecognitionException {
		SchemaRuleContext _localctx = new SchemaRuleContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_schemaRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(252);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__18:
			case T__19:
			case T__20:
			case T__21:
			case T__22:
			case T__23:
			case T__24:
			case T__25:
			case T__26:
				{
				setState(250);
				primitiveSchema();
				}
				break;
			case IDENTIFIER:
				{
				setState(251);
				symbolSchema();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(255);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__16) {
				{
				setState(254);
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
	}

	public final ArrayRuleContext arrayRule() throws RecognitionException {
		ArrayRuleContext _localctx = new ArrayRuleContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_arrayRule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(257);
			match(T__16);
			setState(258);
			match(T__17);
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
	}

	public final PrimitiveSchemaContext primitiveSchema() throws RecognitionException {
		PrimitiveSchemaContext _localctx = new PrimitiveSchemaContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_primitiveSchema);
		try {
			setState(269);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__18:
				enterOuterAlt(_localctx, 1);
				{
				setState(260);
				((PrimitiveSchemaContext)_localctx).name = match(T__18);
				}
				break;
			case T__19:
				enterOuterAlt(_localctx, 2);
				{
				setState(261);
				((PrimitiveSchemaContext)_localctx).name = match(T__19);
				}
				break;
			case T__20:
				enterOuterAlt(_localctx, 3);
				{
				setState(262);
				((PrimitiveSchemaContext)_localctx).name = match(T__20);
				}
				break;
			case T__21:
				enterOuterAlt(_localctx, 4);
				{
				setState(263);
				((PrimitiveSchemaContext)_localctx).name = match(T__21);
				}
				break;
			case T__22:
				enterOuterAlt(_localctx, 5);
				{
				setState(264);
				((PrimitiveSchemaContext)_localctx).name = match(T__22);
				}
				break;
			case T__23:
				enterOuterAlt(_localctx, 6);
				{
				setState(265);
				((PrimitiveSchemaContext)_localctx).name = match(T__23);
				}
				break;
			case T__24:
				enterOuterAlt(_localctx, 7);
				{
				setState(266);
				((PrimitiveSchemaContext)_localctx).name = match(T__24);
				}
				break;
			case T__25:
				enterOuterAlt(_localctx, 8);
				{
				setState(267);
				((PrimitiveSchemaContext)_localctx).name = match(T__25);
				}
				break;
			case T__26:
				enterOuterAlt(_localctx, 9);
				{
				setState(268);
				((PrimitiveSchemaContext)_localctx).name = match(T__26);
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
	}

	public final SymbolSchemaContext symbolSchema() throws RecognitionException {
		SymbolSchemaContext _localctx = new SymbolSchemaContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_symbolSchema);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(271);
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
	}

	public final MetaRuleContext metaRule() throws RecognitionException {
		MetaRuleContext _localctx = new MetaRuleContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_metaRule);
		try {
			setState(275);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TAGLINE:
				enterOuterAlt(_localctx, 1);
				{
				setState(273);
				((MetaRuleContext)_localctx).tagLine = match(TAGLINE);
				}
				break;
			case DOCLINE:
				enterOuterAlt(_localctx, 2);
				{
				setState(274);
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
		"\u0004\u0001(\u0116\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0001\u0000\u0001\u0000\u0005\u0000/\b\u0000\n\u0000\f\u00002\t\u0000"+
		"\u0001\u0001\u0001\u0001\u0005\u00016\b\u0001\n\u0001\f\u00019\t\u0001"+
		"\u0001\u0002\u0005\u0002<\b\u0002\n\u0002\f\u0002?\t\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0003\u0002D\b\u0002\u0001\u0002\u0003\u0002"+
		"G\b\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0003\u0003L\b\u0003\u0001"+
		"\u0003\u0003\u0003O\b\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0003\u0004U\b\u0004\u0001\u0005\u0005\u0005X\b\u0005\n\u0005\f"+
		"\u0005[\t\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0003\u0005`\b\u0005"+
		"\u0001\u0006\u0005\u0006c\b\u0006\n\u0006\f\u0006f\t\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006l\b\u0006\u0001\u0006"+
		"\u0001\u0006\u0005\u0006p\b\u0006\n\u0006\f\u0006s\t\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0003\u0007z\b\u0007"+
		"\u0001\b\u0005\b}\b\b\n\b\f\b\u0080\t\b\u0001\b\u0003\b\u0083\b\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0003\b\u0089\b\b\u0001\t\u0005\t\u008c\b\t"+
		"\n\t\f\t\u008f\t\t\u0001\t\u0001\t\u0001\t\u0005\t\u0094\b\t\n\t\f\t\u0097"+
		"\t\t\u0001\t\u0001\t\u0003\t\u009b\b\t\u0001\t\u0003\t\u009e\b\t\u0001"+
		"\n\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0003"+
		"\u000b\u00a7\b\u000b\u0001\f\u0005\f\u00aa\b\f\n\f\f\f\u00ad\t\f\u0001"+
		"\f\u0001\f\u0001\f\u0001\f\u0005\f\u00b3\b\f\n\f\f\f\u00b6\t\f\u0001\f"+
		"\u0001\f\u0003\f\u00ba\b\f\u0001\r\u0005\r\u00bd\b\r\n\r\f\r\u00c0\t\r"+
		"\u0001\r\u0001\r\u0001\r\u0001\r\u0005\r\u00c6\b\r\n\r\f\r\u00c9\t\r\u0001"+
		"\r\u0001\r\u0001\u000e\u0005\u000e\u00ce\b\u000e\n\u000e\f\u000e\u00d1"+
		"\t\u000e\u0001\u000e\u0003\u000e\u00d4\b\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0003\u000e\u00da\b\u000e\u0001\u000f\u0005\u000f"+
		"\u00dd\b\u000f\n\u000f\f\u000f\u00e0\t\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0005\u000f\u00e6\b\u000f\n\u000f\f\u000f\u00e9\t\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u0010\u0005\u0010\u00ee\b\u0010\n\u0010"+
		"\f\u0010\u00f1\t\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0003\u0010"+
		"\u00f6\b\u0010\u0001\u0010\u0003\u0010\u00f9\b\u0010\u0001\u0011\u0001"+
		"\u0011\u0003\u0011\u00fd\b\u0011\u0001\u0011\u0003\u0011\u0100\b\u0011"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0013\u0001\u0013\u0001\u0013"+
		"\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013"+
		"\u0003\u0013\u010e\b\u0013\u0001\u0014\u0001\u0014\u0001\u0015\u0001\u0015"+
		"\u0003\u0015\u0114\b\u0015\u0001\u0015\u0000\u0000\u0016\u0000\u0002\u0004"+
		"\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \""+
		"$&(*\u0000\u0000\u0130\u0000,\u0001\u0000\u0000\u0000\u00023\u0001\u0000"+
		"\u0000\u0000\u0004=\u0001\u0000\u0000\u0000\u0006H\u0001\u0000\u0000\u0000"+
		"\bT\u0001\u0000\u0000\u0000\nY\u0001\u0000\u0000\u0000\fd\u0001\u0000"+
		"\u0000\u0000\u000ey\u0001\u0000\u0000\u0000\u0010~\u0001\u0000\u0000\u0000"+
		"\u0012\u008d\u0001\u0000\u0000\u0000\u0014\u009f\u0001\u0000\u0000\u0000"+
		"\u0016\u00a2\u0001\u0000\u0000\u0000\u0018\u00ab\u0001\u0000\u0000\u0000"+
		"\u001a\u00be\u0001\u0000\u0000\u0000\u001c\u00cf\u0001\u0000\u0000\u0000"+
		"\u001e\u00de\u0001\u0000\u0000\u0000 \u00ef\u0001\u0000\u0000\u0000\""+
		"\u00fc\u0001\u0000\u0000\u0000$\u0101\u0001\u0000\u0000\u0000&\u010d\u0001"+
		"\u0000\u0000\u0000(\u010f\u0001\u0000\u0000\u0000*\u0113\u0001\u0000\u0000"+
		"\u0000,0\u0003\u0002\u0001\u0000-/\u0003\b\u0004\u0000.-\u0001\u0000\u0000"+
		"\u0000/2\u0001\u0000\u0000\u00000.\u0001\u0000\u0000\u000001\u0001\u0000"+
		"\u0000\u00001\u0001\u0001\u0000\u0000\u000020\u0001\u0000\u0000\u0000"+
		"37\u0003\u0004\u0002\u000046\u0003\u0006\u0003\u000054\u0001\u0000\u0000"+
		"\u000069\u0001\u0000\u0000\u000075\u0001\u0000\u0000\u000078\u0001\u0000"+
		"\u0000\u00008\u0003\u0001\u0000\u0000\u000097\u0001\u0000\u0000\u0000"+
		":<\u0003*\u0015\u0000;:\u0001\u0000\u0000\u0000<?\u0001\u0000\u0000\u0000"+
		"=;\u0001\u0000\u0000\u0000=>\u0001\u0000\u0000\u0000>@\u0001\u0000\u0000"+
		"\u0000?=\u0001\u0000\u0000\u0000@A\u0005\u0001\u0000\u0000AC\u0005\u001f"+
		"\u0000\u0000BD\u0005 \u0000\u0000CB\u0001\u0000\u0000\u0000CD\u0001\u0000"+
		"\u0000\u0000DF\u0001\u0000\u0000\u0000EG\u0005(\u0000\u0000FE\u0001\u0000"+
		"\u0000\u0000FG\u0001\u0000\u0000\u0000G\u0005\u0001\u0000\u0000\u0000"+
		"HI\u0005\u0002\u0000\u0000IK\u0005\u001f\u0000\u0000JL\u0005 \u0000\u0000"+
		"KJ\u0001\u0000\u0000\u0000KL\u0001\u0000\u0000\u0000LN\u0001\u0000\u0000"+
		"\u0000MO\u0005(\u0000\u0000NM\u0001\u0000\u0000\u0000NO\u0001\u0000\u0000"+
		"\u0000O\u0007\u0001\u0000\u0000\u0000PU\u0003\n\u0005\u0000QU\u0003\f"+
		"\u0006\u0000RU\u0003\u001a\r\u0000SU\u0003\u001e\u000f\u0000TP\u0001\u0000"+
		"\u0000\u0000TQ\u0001\u0000\u0000\u0000TR\u0001\u0000\u0000\u0000TS\u0001"+
		"\u0000\u0000\u0000U\t\u0001\u0000\u0000\u0000VX\u0003*\u0015\u0000WV\u0001"+
		"\u0000\u0000\u0000X[\u0001\u0000\u0000\u0000YW\u0001\u0000\u0000\u0000"+
		"YZ\u0001\u0000\u0000\u0000Z\\\u0001\u0000\u0000\u0000[Y\u0001\u0000\u0000"+
		"\u0000\\]\u0005\u0003\u0000\u0000]_\u0005\u001f\u0000\u0000^`\u0005(\u0000"+
		"\u0000_^\u0001\u0000\u0000\u0000_`\u0001\u0000\u0000\u0000`\u000b\u0001"+
		"\u0000\u0000\u0000ac\u0003*\u0015\u0000ba\u0001\u0000\u0000\u0000cf\u0001"+
		"\u0000\u0000\u0000db\u0001\u0000\u0000\u0000de\u0001\u0000\u0000\u0000"+
		"eg\u0001\u0000\u0000\u0000fd\u0001\u0000\u0000\u0000gh\u0005\u0004\u0000"+
		"\u0000hk\u0005\u001f\u0000\u0000ij\u0005\u0005\u0000\u0000jl\u0005\u001f"+
		"\u0000\u0000ki\u0001\u0000\u0000\u0000kl\u0001\u0000\u0000\u0000lm\u0001"+
		"\u0000\u0000\u0000mq\u0005\u0006\u0000\u0000np\u0003\u000e\u0007\u0000"+
		"on\u0001\u0000\u0000\u0000ps\u0001\u0000\u0000\u0000qo\u0001\u0000\u0000"+
		"\u0000qr\u0001\u0000\u0000\u0000rt\u0001\u0000\u0000\u0000sq\u0001\u0000"+
		"\u0000\u0000tu\u0005\u0007\u0000\u0000u\r\u0001\u0000\u0000\u0000vz\u0003"+
		"\u0010\b\u0000wz\u0003\u0012\t\u0000xz\u0003\u0018\f\u0000yv\u0001\u0000"+
		"\u0000\u0000yw\u0001\u0000\u0000\u0000yx\u0001\u0000\u0000\u0000z\u000f"+
		"\u0001\u0000\u0000\u0000{}\u0003*\u0015\u0000|{\u0001\u0000\u0000\u0000"+
		"}\u0080\u0001\u0000\u0000\u0000~|\u0001\u0000\u0000\u0000~\u007f\u0001"+
		"\u0000\u0000\u0000\u007f\u0082\u0001\u0000\u0000\u0000\u0080~\u0001\u0000"+
		"\u0000\u0000\u0081\u0083\u0005\b\u0000\u0000\u0082\u0081\u0001\u0000\u0000"+
		"\u0000\u0082\u0083\u0001\u0000\u0000\u0000\u0083\u0084\u0001\u0000\u0000"+
		"\u0000\u0084\u0085\u0005\u001f\u0000\u0000\u0085\u0086\u0005\t\u0000\u0000"+
		"\u0086\u0088\u0003\"\u0011\u0000\u0087\u0089\u0005(\u0000\u0000\u0088"+
		"\u0087\u0001\u0000\u0000\u0000\u0088\u0089\u0001\u0000\u0000\u0000\u0089"+
		"\u0011\u0001\u0000\u0000\u0000\u008a\u008c\u0003*\u0015\u0000\u008b\u008a"+
		"\u0001\u0000\u0000\u0000\u008c\u008f\u0001\u0000\u0000\u0000\u008d\u008b"+
		"\u0001\u0000\u0000\u0000\u008d\u008e\u0001\u0000\u0000\u0000\u008e\u0090"+
		"\u0001\u0000\u0000\u0000\u008f\u008d\u0001\u0000\u0000\u0000\u0090\u0091"+
		"\u0005\u001f\u0000\u0000\u0091\u0095\u0005\n\u0000\u0000\u0092\u0094\u0003"+
		"\u0016\u000b\u0000\u0093\u0092\u0001\u0000\u0000\u0000\u0094\u0097\u0001"+
		"\u0000\u0000\u0000\u0095\u0093\u0001\u0000\u0000\u0000\u0095\u0096\u0001"+
		"\u0000\u0000\u0000\u0096\u0098\u0001\u0000\u0000\u0000\u0097\u0095\u0001"+
		"\u0000\u0000\u0000\u0098\u009a\u0005\u000b\u0000\u0000\u0099\u009b\u0003"+
		"\u0014\n\u0000\u009a\u0099\u0001\u0000\u0000\u0000\u009a\u009b\u0001\u0000"+
		"\u0000\u0000\u009b\u009d\u0001\u0000\u0000\u0000\u009c\u009e\u0005(\u0000"+
		"\u0000\u009d\u009c\u0001\u0000\u0000\u0000\u009d\u009e\u0001\u0000\u0000"+
		"\u0000\u009e\u0013\u0001\u0000\u0000\u0000\u009f\u00a0\u0005\t\u0000\u0000"+
		"\u00a0\u00a1\u0003\"\u0011\u0000\u00a1\u0015\u0001\u0000\u0000\u0000\u00a2"+
		"\u00a3\u0005\u001f\u0000\u0000\u00a3\u00a4\u0005\t\u0000\u0000\u00a4\u00a6"+
		"\u0003\"\u0011\u0000\u00a5\u00a7\u0005\f\u0000\u0000\u00a6\u00a5\u0001"+
		"\u0000\u0000\u0000\u00a6\u00a7\u0001\u0000\u0000\u0000\u00a7\u0017\u0001"+
		"\u0000\u0000\u0000\u00a8\u00aa\u0003*\u0015\u0000\u00a9\u00a8\u0001\u0000"+
		"\u0000\u0000\u00aa\u00ad\u0001\u0000\u0000\u0000\u00ab\u00a9\u0001\u0000"+
		"\u0000\u0000\u00ab\u00ac\u0001\u0000\u0000\u0000\u00ac\u00ae\u0001\u0000"+
		"\u0000\u0000\u00ad\u00ab\u0001\u0000\u0000\u0000\u00ae\u00af\u0005\r\u0000"+
		"\u0000\u00af\u00b0\u0005\u001f\u0000\u0000\u00b0\u00b4\u0005\n\u0000\u0000"+
		"\u00b1\u00b3\u0003\u0016\u000b\u0000\u00b2\u00b1\u0001\u0000\u0000\u0000"+
		"\u00b3\u00b6\u0001\u0000\u0000\u0000\u00b4\u00b2\u0001\u0000\u0000\u0000"+
		"\u00b4\u00b5\u0001\u0000\u0000\u0000\u00b5\u00b7\u0001\u0000\u0000\u0000"+
		"\u00b6\u00b4\u0001\u0000\u0000\u0000\u00b7\u00b9\u0005\u000b\u0000\u0000"+
		"\u00b8\u00ba\u0005(\u0000\u0000\u00b9\u00b8\u0001\u0000\u0000\u0000\u00b9"+
		"\u00ba\u0001\u0000\u0000\u0000\u00ba\u0019\u0001\u0000\u0000\u0000\u00bb"+
		"\u00bd\u0003*\u0015\u0000\u00bc\u00bb\u0001\u0000\u0000\u0000\u00bd\u00c0"+
		"\u0001\u0000\u0000\u0000\u00be\u00bc\u0001\u0000\u0000\u0000\u00be\u00bf"+
		"\u0001\u0000\u0000\u0000\u00bf\u00c1\u0001\u0000\u0000\u0000\u00c0\u00be"+
		"\u0001\u0000\u0000\u0000\u00c1\u00c2\u0005\u000e\u0000\u0000\u00c2\u00c3"+
		"\u0005\u001f\u0000\u0000\u00c3\u00c7\u0005\u0006\u0000\u0000\u00c4\u00c6"+
		"\u0003\u001c\u000e\u0000\u00c5\u00c4\u0001\u0000\u0000\u0000\u00c6\u00c9"+
		"\u0001\u0000\u0000\u0000\u00c7\u00c5\u0001\u0000\u0000\u0000\u00c7\u00c8"+
		"\u0001\u0000\u0000\u0000\u00c8\u00ca\u0001\u0000\u0000\u0000\u00c9\u00c7"+
		"\u0001\u0000\u0000\u0000\u00ca\u00cb\u0005\u0007\u0000\u0000\u00cb\u001b"+
		"\u0001\u0000\u0000\u0000\u00cc\u00ce\u0003*\u0015\u0000\u00cd\u00cc\u0001"+
		"\u0000\u0000\u0000\u00ce\u00d1\u0001\u0000\u0000\u0000\u00cf\u00cd\u0001"+
		"\u0000\u0000\u0000\u00cf\u00d0\u0001\u0000\u0000\u0000\u00d0\u00d3\u0001"+
		"\u0000\u0000\u0000\u00d1\u00cf\u0001\u0000\u0000\u0000\u00d2\u00d4\u0005"+
		"\b\u0000\u0000\u00d3\u00d2\u0001\u0000\u0000\u0000\u00d3\u00d4\u0001\u0000"+
		"\u0000\u0000\u00d4\u00d5\u0001\u0000\u0000\u0000\u00d5\u00d6\u0005\u001f"+
		"\u0000\u0000\u00d6\u00d7\u0005\t\u0000\u0000\u00d7\u00d9\u0003\"\u0011"+
		"\u0000\u00d8\u00da\u0005(\u0000\u0000\u00d9\u00d8\u0001\u0000\u0000\u0000"+
		"\u00d9\u00da\u0001\u0000\u0000\u0000\u00da\u001d\u0001\u0000\u0000\u0000"+
		"\u00db\u00dd\u0003*\u0015\u0000\u00dc\u00db\u0001\u0000\u0000\u0000\u00dd"+
		"\u00e0\u0001\u0000\u0000\u0000\u00de\u00dc\u0001\u0000\u0000\u0000\u00de"+
		"\u00df\u0001\u0000\u0000\u0000\u00df\u00e1\u0001\u0000\u0000\u0000\u00e0"+
		"\u00de\u0001\u0000\u0000\u0000\u00e1\u00e2\u0005\u000f\u0000\u0000\u00e2"+
		"\u00e3\u0005\u001f\u0000\u0000\u00e3\u00e7\u0005\u0006\u0000\u0000\u00e4"+
		"\u00e6\u0003 \u0010\u0000\u00e5\u00e4\u0001\u0000\u0000\u0000\u00e6\u00e9"+
		"\u0001\u0000\u0000\u0000\u00e7\u00e5\u0001\u0000\u0000\u0000\u00e7\u00e8"+
		"\u0001\u0000\u0000\u0000\u00e8\u00ea\u0001\u0000\u0000\u0000\u00e9\u00e7"+
		"\u0001\u0000\u0000\u0000\u00ea\u00eb\u0005\u0007\u0000\u0000\u00eb\u001f"+
		"\u0001\u0000\u0000\u0000\u00ec\u00ee\u0003*\u0015\u0000\u00ed\u00ec\u0001"+
		"\u0000\u0000\u0000\u00ee\u00f1\u0001\u0000\u0000\u0000\u00ef\u00ed\u0001"+
		"\u0000\u0000\u0000\u00ef\u00f0\u0001\u0000\u0000\u0000\u00f0\u00f2\u0001"+
		"\u0000\u0000\u0000\u00f1\u00ef\u0001\u0000\u0000\u0000\u00f2\u00f5\u0005"+
		"\u001f\u0000\u0000\u00f3\u00f4\u0005\u0010\u0000\u0000\u00f4\u00f6\u0005"+
		"\u001d\u0000\u0000\u00f5\u00f3\u0001\u0000\u0000\u0000\u00f5\u00f6\u0001"+
		"\u0000\u0000\u0000\u00f6\u00f8\u0001\u0000\u0000\u0000\u00f7\u00f9\u0005"+
		"\f\u0000\u0000\u00f8\u00f7\u0001\u0000\u0000\u0000\u00f8\u00f9\u0001\u0000"+
		"\u0000\u0000\u00f9!\u0001\u0000\u0000\u0000\u00fa\u00fd\u0003&\u0013\u0000"+
		"\u00fb\u00fd\u0003(\u0014\u0000\u00fc\u00fa\u0001\u0000\u0000\u0000\u00fc"+
		"\u00fb\u0001\u0000\u0000\u0000\u00fd\u00ff\u0001\u0000\u0000\u0000\u00fe"+
		"\u0100\u0003$\u0012\u0000\u00ff\u00fe\u0001\u0000\u0000\u0000\u00ff\u0100"+
		"\u0001\u0000\u0000\u0000\u0100#\u0001\u0000\u0000\u0000\u0101\u0102\u0005"+
		"\u0011\u0000\u0000\u0102\u0103\u0005\u0012\u0000\u0000\u0103%\u0001\u0000"+
		"\u0000\u0000\u0104\u010e\u0005\u0013\u0000\u0000\u0105\u010e\u0005\u0014"+
		"\u0000\u0000\u0106\u010e\u0005\u0015\u0000\u0000\u0107\u010e\u0005\u0016"+
		"\u0000\u0000\u0108\u010e\u0005\u0017\u0000\u0000\u0109\u010e\u0005\u0018"+
		"\u0000\u0000\u010a\u010e\u0005\u0019\u0000\u0000\u010b\u010e\u0005\u001a"+
		"\u0000\u0000\u010c\u010e\u0005\u001b\u0000\u0000\u010d\u0104\u0001\u0000"+
		"\u0000\u0000\u010d\u0105\u0001\u0000\u0000\u0000\u010d\u0106\u0001\u0000"+
		"\u0000\u0000\u010d\u0107\u0001\u0000\u0000\u0000\u010d\u0108\u0001\u0000"+
		"\u0000\u0000\u010d\u0109\u0001\u0000\u0000\u0000\u010d\u010a\u0001\u0000"+
		"\u0000\u0000\u010d\u010b\u0001\u0000\u0000\u0000\u010d\u010c\u0001\u0000"+
		"\u0000\u0000\u010e\'\u0001\u0000\u0000\u0000\u010f\u0110\u0005\u001f\u0000"+
		"\u0000\u0110)\u0001\u0000\u0000\u0000\u0111\u0114\u0005\"\u0000\u0000"+
		"\u0112\u0114\u0005!\u0000\u0000\u0113\u0111\u0001\u0000\u0000\u0000\u0113"+
		"\u0112\u0001\u0000\u0000\u0000\u0114+\u0001\u0000\u0000\u0000\'07=CFK"+
		"NTY_dkqy~\u0082\u0088\u008d\u0095\u009a\u009d\u00a6\u00ab\u00b4\u00b9"+
		"\u00be\u00c7\u00cf\u00d3\u00d9\u00de\u00e7\u00ef\u00f5\u00f8\u00fc\u00ff"+
		"\u010d\u0113";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}