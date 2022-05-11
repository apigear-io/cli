// Generated from /Users/jryannel/work/apigear-go/apigear-cli-go/pkg/idl/parser/ObjectApi.g4 by ANTLR 4.8
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
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, WHITESPACE=20, INTEGER=21, HEX=22, TYPE_IDENTIFIER=23, 
		IDENTIFIER=24, VERSION=25;
	public static final int
		RULE_documentRule = 0, RULE_headerRule = 1, RULE_moduleRule = 2, RULE_importRule = 3, 
		RULE_declarationsRule = 4, RULE_interfaceRule = 5, RULE_interfaceMembersRule = 6, 
		RULE_propertyRule = 7, RULE_methodRule = 8, RULE_outputRule = 9, RULE_inputRule = 10, 
		RULE_signalRule = 11, RULE_structRule = 12, RULE_structFieldRule = 13, 
		RULE_enumRule = 14, RULE_enumMemberRule = 15, RULE_schemaRule = 16, RULE_arrayRule = 17, 
		RULE_primitiveSchema = 18, RULE_symbolSchema = 19;
	private static String[] makeRuleNames() {
		return new String[] {
			"documentRule", "headerRule", "moduleRule", "importRule", "declarationsRule", 
			"interfaceRule", "interfaceMembersRule", "propertyRule", "methodRule", 
			"outputRule", "inputRule", "signalRule", "structRule", "structFieldRule", 
			"enumRule", "enumMemberRule", "schemaRule", "arrayRule", "primitiveSchema", 
			"symbolSchema"
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
			"HEX", "TYPE_IDENTIFIER", "IDENTIFIER", "VERSION"
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
			setState(40);
			headerRule();
			setState(44);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__2) | (1L << T__10) | (1L << T__11))) != 0)) {
				{
				{
				setState(41);
				declarationsRule();
				}
				}
				setState(46);
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
			setState(47);
			moduleRule();
			setState(51);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__1) {
				{
				{
				setState(48);
				importRule();
				}
				}
				setState(53);
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
		public ModuleRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_moduleRule; }
	}

	public final ModuleRuleContext moduleRule() throws RecognitionException {
		ModuleRuleContext _localctx = new ModuleRuleContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_moduleRule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(54);
			match(T__0);
			setState(55);
			((ModuleRuleContext)_localctx).name = match(IDENTIFIER);
			setState(56);
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
			setState(58);
			match(T__1);
			setState(59);
			((ImportRuleContext)_localctx).name = match(IDENTIFIER);
			setState(60);
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
			setState(65);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__2:
				enterOuterAlt(_localctx, 1);
				{
				setState(62);
				interfaceRule();
				}
				break;
			case T__10:
				enterOuterAlt(_localctx, 2);
				{
				setState(63);
				structRule();
				}
				break;
			case T__11:
				enterOuterAlt(_localctx, 3);
				{
				setState(64);
				enumRule();
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

	public static class InterfaceRuleContext extends ParserRuleContext {
		public Token name;
		public TerminalNode IDENTIFIER() { return getToken(ObjectApiParser.IDENTIFIER, 0); }
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
			setState(67);
			match(T__2);
			setState(68);
			((InterfaceRuleContext)_localctx).name = match(IDENTIFIER);
			setState(69);
			match(T__3);
			setState(73);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__9 || _la==IDENTIFIER) {
				{
				{
				setState(70);
				interfaceMembersRule();
				}
				}
				setState(75);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(76);
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
		public MethodRuleContext methodRule() {
			return getRuleContext(MethodRuleContext.class,0);
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
			setState(81);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(78);
				propertyRule();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(79);
				methodRule();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(80);
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
		public PropertyRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_propertyRule; }
	}

	public final PropertyRuleContext propertyRule() throws RecognitionException {
		PropertyRuleContext _localctx = new PropertyRuleContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_propertyRule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(83);
			((PropertyRuleContext)_localctx).name = match(IDENTIFIER);
			setState(84);
			match(T__5);
			setState(85);
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

	public static class MethodRuleContext extends ParserRuleContext {
		public Token name;
		public InputRuleContext inputs;
		public TerminalNode IDENTIFIER() { return getToken(ObjectApiParser.IDENTIFIER, 0); }
		public OutputRuleContext outputRule() {
			return getRuleContext(OutputRuleContext.class,0);
		}
		public List<InputRuleContext> inputRule() {
			return getRuleContexts(InputRuleContext.class);
		}
		public InputRuleContext inputRule(int i) {
			return getRuleContext(InputRuleContext.class,i);
		}
		public MethodRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_methodRule; }
	}

	public final MethodRuleContext methodRule() throws RecognitionException {
		MethodRuleContext _localctx = new MethodRuleContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_methodRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(87);
			((MethodRuleContext)_localctx).name = match(IDENTIFIER);
			setState(88);
			match(T__6);
			setState(92);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENTIFIER) {
				{
				{
				setState(89);
				((MethodRuleContext)_localctx).inputs = inputRule();
				}
				}
				setState(94);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(95);
			match(T__7);
			setState(97);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__5) {
				{
				setState(96);
				outputRule();
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

	public static class OutputRuleContext extends ParserRuleContext {
		public SchemaRuleContext schema;
		public SchemaRuleContext schemaRule() {
			return getRuleContext(SchemaRuleContext.class,0);
		}
		public OutputRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_outputRule; }
	}

	public final OutputRuleContext outputRule() throws RecognitionException {
		OutputRuleContext _localctx = new OutputRuleContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_outputRule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(99);
			match(T__5);
			setState(100);
			((OutputRuleContext)_localctx).schema = schemaRule();
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

	public static class InputRuleContext extends ParserRuleContext {
		public Token name;
		public SchemaRuleContext schema;
		public TerminalNode IDENTIFIER() { return getToken(ObjectApiParser.IDENTIFIER, 0); }
		public SchemaRuleContext schemaRule() {
			return getRuleContext(SchemaRuleContext.class,0);
		}
		public InputRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_inputRule; }
	}

	public final InputRuleContext inputRule() throws RecognitionException {
		InputRuleContext _localctx = new InputRuleContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_inputRule);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(102);
			((InputRuleContext)_localctx).name = match(IDENTIFIER);
			setState(103);
			match(T__5);
			setState(104);
			((InputRuleContext)_localctx).schema = schemaRule();
			setState(106);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__8) {
				{
				setState(105);
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
		public InputRuleContext inputs;
		public TerminalNode IDENTIFIER() { return getToken(ObjectApiParser.IDENTIFIER, 0); }
		public List<InputRuleContext> inputRule() {
			return getRuleContexts(InputRuleContext.class);
		}
		public InputRuleContext inputRule(int i) {
			return getRuleContext(InputRuleContext.class,i);
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
			setState(108);
			match(T__9);
			setState(109);
			((SignalRuleContext)_localctx).name = match(IDENTIFIER);
			setState(110);
			match(T__6);
			setState(114);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENTIFIER) {
				{
				{
				setState(111);
				((SignalRuleContext)_localctx).inputs = inputRule();
				}
				}
				setState(116);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(117);
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
			setState(119);
			match(T__10);
			setState(120);
			((StructRuleContext)_localctx).name = match(IDENTIFIER);
			setState(121);
			match(T__3);
			setState(125);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENTIFIER) {
				{
				{
				setState(122);
				structFieldRule();
				}
				}
				setState(127);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(128);
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
		public StructFieldRuleContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structFieldRule; }
	}

	public final StructFieldRuleContext structFieldRule() throws RecognitionException {
		StructFieldRuleContext _localctx = new StructFieldRuleContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_structFieldRule);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(130);
			((StructFieldRuleContext)_localctx).name = match(IDENTIFIER);
			setState(131);
			match(T__5);
			setState(132);
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
			setState(134);
			match(T__11);
			setState(135);
			((EnumRuleContext)_localctx).name = match(IDENTIFIER);
			setState(136);
			match(T__3);
			setState(140);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==IDENTIFIER) {
				{
				{
				setState(137);
				enumMemberRule();
				}
				}
				setState(142);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(143);
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
			setState(145);
			((EnumMemberRuleContext)_localctx).name = match(IDENTIFIER);
			setState(148);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__12) {
				{
				setState(146);
				match(T__12);
				setState(147);
				((EnumMemberRuleContext)_localctx).value = match(INTEGER);
				}
			}

			setState(151);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__8) {
				{
				setState(150);
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
			setState(155);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__15:
			case T__16:
			case T__17:
			case T__18:
				{
				setState(153);
				primitiveSchema();
				}
				break;
			case IDENTIFIER:
				{
				setState(154);
				symbolSchema();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(158);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__13) {
				{
				setState(157);
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
			setState(160);
			match(T__13);
			setState(161);
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
			setState(167);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__15:
				enterOuterAlt(_localctx, 1);
				{
				setState(163);
				((PrimitiveSchemaContext)_localctx).name = match(T__15);
				}
				break;
			case T__16:
				enterOuterAlt(_localctx, 2);
				{
				setState(164);
				((PrimitiveSchemaContext)_localctx).name = match(T__16);
				}
				break;
			case T__17:
				enterOuterAlt(_localctx, 3);
				{
				setState(165);
				((PrimitiveSchemaContext)_localctx).name = match(T__17);
				}
				break;
			case T__18:
				enterOuterAlt(_localctx, 4);
				{
				setState(166);
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
			setState(169);
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

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\33\u00ae\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\3\2\3\2\7\2-\n\2\f\2\16\2\60\13\2\3\3\3"+
		"\3\7\3\64\n\3\f\3\16\3\67\13\3\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\6\3\6"+
		"\3\6\5\6D\n\6\3\7\3\7\3\7\3\7\7\7J\n\7\f\7\16\7M\13\7\3\7\3\7\3\b\3\b"+
		"\3\b\5\bT\n\b\3\t\3\t\3\t\3\t\3\n\3\n\3\n\7\n]\n\n\f\n\16\n`\13\n\3\n"+
		"\3\n\5\nd\n\n\3\13\3\13\3\13\3\f\3\f\3\f\3\f\5\fm\n\f\3\r\3\r\3\r\3\r"+
		"\7\rs\n\r\f\r\16\rv\13\r\3\r\3\r\3\16\3\16\3\16\3\16\7\16~\n\16\f\16\16"+
		"\16\u0081\13\16\3\16\3\16\3\17\3\17\3\17\3\17\3\20\3\20\3\20\3\20\7\20"+
		"\u008d\n\20\f\20\16\20\u0090\13\20\3\20\3\20\3\21\3\21\3\21\5\21\u0097"+
		"\n\21\3\21\5\21\u009a\n\21\3\22\3\22\5\22\u009e\n\22\3\22\5\22\u00a1\n"+
		"\22\3\23\3\23\3\23\3\24\3\24\3\24\3\24\5\24\u00aa\n\24\3\25\3\25\3\25"+
		"\2\2\26\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(\2\2\2\u00ad\2*\3"+
		"\2\2\2\4\61\3\2\2\2\68\3\2\2\2\b<\3\2\2\2\nC\3\2\2\2\fE\3\2\2\2\16S\3"+
		"\2\2\2\20U\3\2\2\2\22Y\3\2\2\2\24e\3\2\2\2\26h\3\2\2\2\30n\3\2\2\2\32"+
		"y\3\2\2\2\34\u0084\3\2\2\2\36\u0088\3\2\2\2 \u0093\3\2\2\2\"\u009d\3\2"+
		"\2\2$\u00a2\3\2\2\2&\u00a9\3\2\2\2(\u00ab\3\2\2\2*.\5\4\3\2+-\5\n\6\2"+
		",+\3\2\2\2-\60\3\2\2\2.,\3\2\2\2./\3\2\2\2/\3\3\2\2\2\60.\3\2\2\2\61\65"+
		"\5\6\4\2\62\64\5\b\5\2\63\62\3\2\2\2\64\67\3\2\2\2\65\63\3\2\2\2\65\66"+
		"\3\2\2\2\66\5\3\2\2\2\67\65\3\2\2\289\7\3\2\29:\7\32\2\2:;\7\33\2\2;\7"+
		"\3\2\2\2<=\7\4\2\2=>\7\32\2\2>?\7\33\2\2?\t\3\2\2\2@D\5\f\7\2AD\5\32\16"+
		"\2BD\5\36\20\2C@\3\2\2\2CA\3\2\2\2CB\3\2\2\2D\13\3\2\2\2EF\7\5\2\2FG\7"+
		"\32\2\2GK\7\6\2\2HJ\5\16\b\2IH\3\2\2\2JM\3\2\2\2KI\3\2\2\2KL\3\2\2\2L"+
		"N\3\2\2\2MK\3\2\2\2NO\7\7\2\2O\r\3\2\2\2PT\5\20\t\2QT\5\22\n\2RT\5\30"+
		"\r\2SP\3\2\2\2SQ\3\2\2\2SR\3\2\2\2T\17\3\2\2\2UV\7\32\2\2VW\7\b\2\2WX"+
		"\5\"\22\2X\21\3\2\2\2YZ\7\32\2\2Z^\7\t\2\2[]\5\26\f\2\\[\3\2\2\2]`\3\2"+
		"\2\2^\\\3\2\2\2^_\3\2\2\2_a\3\2\2\2`^\3\2\2\2ac\7\n\2\2bd\5\24\13\2cb"+
		"\3\2\2\2cd\3\2\2\2d\23\3\2\2\2ef\7\b\2\2fg\5\"\22\2g\25\3\2\2\2hi\7\32"+
		"\2\2ij\7\b\2\2jl\5\"\22\2km\7\13\2\2lk\3\2\2\2lm\3\2\2\2m\27\3\2\2\2n"+
		"o\7\f\2\2op\7\32\2\2pt\7\t\2\2qs\5\26\f\2rq\3\2\2\2sv\3\2\2\2tr\3\2\2"+
		"\2tu\3\2\2\2uw\3\2\2\2vt\3\2\2\2wx\7\n\2\2x\31\3\2\2\2yz\7\r\2\2z{\7\32"+
		"\2\2{\177\7\6\2\2|~\5\34\17\2}|\3\2\2\2~\u0081\3\2\2\2\177}\3\2\2\2\177"+
		"\u0080\3\2\2\2\u0080\u0082\3\2\2\2\u0081\177\3\2\2\2\u0082\u0083\7\7\2"+
		"\2\u0083\33\3\2\2\2\u0084\u0085\7\32\2\2\u0085\u0086\7\b\2\2\u0086\u0087"+
		"\5\"\22\2\u0087\35\3\2\2\2\u0088\u0089\7\16\2\2\u0089\u008a\7\32\2\2\u008a"+
		"\u008e\7\6\2\2\u008b\u008d\5 \21\2\u008c\u008b\3\2\2\2\u008d\u0090\3\2"+
		"\2\2\u008e\u008c\3\2\2\2\u008e\u008f\3\2\2\2\u008f\u0091\3\2\2\2\u0090"+
		"\u008e\3\2\2\2\u0091\u0092\7\7\2\2\u0092\37\3\2\2\2\u0093\u0096\7\32\2"+
		"\2\u0094\u0095\7\17\2\2\u0095\u0097\7\27\2\2\u0096\u0094\3\2\2\2\u0096"+
		"\u0097\3\2\2\2\u0097\u0099\3\2\2\2\u0098\u009a\7\13\2\2\u0099\u0098\3"+
		"\2\2\2\u0099\u009a\3\2\2\2\u009a!\3\2\2\2\u009b\u009e\5&\24\2\u009c\u009e"+
		"\5(\25\2\u009d\u009b\3\2\2\2\u009d\u009c\3\2\2\2\u009e\u00a0\3\2\2\2\u009f"+
		"\u00a1\5$\23\2\u00a0\u009f\3\2\2\2\u00a0\u00a1\3\2\2\2\u00a1#\3\2\2\2"+
		"\u00a2\u00a3\7\20\2\2\u00a3\u00a4\7\21\2\2\u00a4%\3\2\2\2\u00a5\u00aa"+
		"\7\22\2\2\u00a6\u00aa\7\23\2\2\u00a7\u00aa\7\24\2\2\u00a8\u00aa\7\25\2"+
		"\2\u00a9\u00a5\3\2\2\2\u00a9\u00a6\3\2\2\2\u00a9\u00a7\3\2\2\2\u00a9\u00a8"+
		"\3\2\2\2\u00aa\'\3\2\2\2\u00ab\u00ac\7\32\2\2\u00ac)\3\2\2\2\22.\65CK"+
		"S^clt\177\u008e\u0096\u0099\u009d\u00a0\u00a9";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}