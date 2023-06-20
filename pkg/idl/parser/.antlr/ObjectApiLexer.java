// Generated from /Users/jryannel/dev/apigear/cli/pkg/idl/parser/ObjectApi.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class ObjectApiLexer extends Lexer {
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
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
			"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16", 
			"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "WHITESPACE", 
			"INTEGER", "HEX", "TYPE_IDENTIFIER", "IDENTIFIER", "VERSION", "DOCLINE", 
			"TAGLINE", "COMMENT"
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


	public ObjectApiLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "ObjectApi.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2#\u0103\b\1\4\2\t"+
		"\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\3\2\3\2\3\2\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\3\3\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\7\3\7"+
		"\3\7\3\7\3\7\3\7\3\7\3\b\3\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\f\3\f"+
		"\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16\3\16\3\17"+
		"\3\17\3\20\3\20\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\23\3\23\3\23\3\23"+
		"\3\24\3\24\3\24\3\24\3\24\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\26\3\26"+
		"\3\26\3\26\3\26\3\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\30\3\30"+
		"\3\30\3\30\3\30\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\31\3\31\3\32"+
		"\6\32\u00bf\n\32\r\32\16\32\u00c0\3\32\3\32\3\33\5\33\u00c6\n\33\3\33"+
		"\6\33\u00c9\n\33\r\33\16\33\u00ca\3\34\3\34\3\34\3\34\6\34\u00d1\n\34"+
		"\r\34\16\34\u00d2\3\35\3\35\7\35\u00d7\n\35\f\35\16\35\u00da\13\35\3\35"+
		"\3\35\5\35\u00de\n\35\3\36\3\36\7\36\u00e2\n\36\f\36\16\36\u00e5\13\36"+
		"\3\37\3\37\3\37\3\37\3 \3 \3 \3 \7 \u00ef\n \f \16 \u00f2\13 \3!\3!\7"+
		"!\u00f6\n!\f!\16!\u00f9\13!\3\"\3\"\7\"\u00fd\n\"\f\"\16\"\u0100\13\""+
		"\3\"\3\"\2\2#\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16"+
		"\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34"+
		"\67\359\36;\37= ?!A\"C#\3\2\13\5\2\13\f\17\17\"\"\4\2--//\5\2\62;CHch"+
		"\4\2C\\aa\5\2\62;C\\aa\5\2C\\aac|\7\2\60\60\62;C\\aac|\3\2\62;\4\2\f\f"+
		"\17\17\2\u010c\2\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3"+
		"\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2"+
		"\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3"+
		"\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2"+
		"\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\2\67\3\2\2\2\2"+
		"9\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2\2\2A\3\2\2\2\2C\3\2\2\2\3E\3"+
		"\2\2\2\5L\3\2\2\2\7S\3\2\2\2\t]\3\2\2\2\13_\3\2\2\2\ra\3\2\2\2\17j\3\2"+
		"\2\2\21l\3\2\2\2\23n\3\2\2\2\25p\3\2\2\2\27r\3\2\2\2\31y\3\2\2\2\33\u0080"+
		"\3\2\2\2\35\u0085\3\2\2\2\37\u0087\3\2\2\2!\u0089\3\2\2\2#\u008b\3\2\2"+
		"\2%\u0090\3\2\2\2\'\u0094\3\2\2\2)\u009a\3\2\2\2+\u00a0\3\2\2\2-\u00a6"+
		"\3\2\2\2/\u00ae\3\2\2\2\61\u00b6\3\2\2\2\63\u00be\3\2\2\2\65\u00c5\3\2"+
		"\2\2\67\u00cc\3\2\2\29\u00d4\3\2\2\2;\u00df\3\2\2\2=\u00e6\3\2\2\2?\u00ea"+
		"\3\2\2\2A\u00f3\3\2\2\2C\u00fa\3\2\2\2EF\7o\2\2FG\7q\2\2GH\7f\2\2HI\7"+
		"w\2\2IJ\7n\2\2JK\7g\2\2K\4\3\2\2\2LM\7k\2\2MN\7o\2\2NO\7r\2\2OP\7q\2\2"+
		"PQ\7t\2\2QR\7v\2\2R\6\3\2\2\2ST\7k\2\2TU\7p\2\2UV\7v\2\2VW\7g\2\2WX\7"+
		"t\2\2XY\7h\2\2YZ\7c\2\2Z[\7e\2\2[\\\7g\2\2\\\b\3\2\2\2]^\7}\2\2^\n\3\2"+
		"\2\2_`\7\177\2\2`\f\3\2\2\2ab\7t\2\2bc\7g\2\2cd\7c\2\2de\7f\2\2ef\7q\2"+
		"\2fg\7p\2\2gh\7n\2\2hi\7{\2\2i\16\3\2\2\2jk\7<\2\2k\20\3\2\2\2lm\7*\2"+
		"\2m\22\3\2\2\2no\7+\2\2o\24\3\2\2\2pq\7.\2\2q\26\3\2\2\2rs\7u\2\2st\7"+
		"k\2\2tu\7i\2\2uv\7p\2\2vw\7c\2\2wx\7n\2\2x\30\3\2\2\2yz\7u\2\2z{\7v\2"+
		"\2{|\7t\2\2|}\7w\2\2}~\7e\2\2~\177\7v\2\2\177\32\3\2\2\2\u0080\u0081\7"+
		"g\2\2\u0081\u0082\7p\2\2\u0082\u0083\7w\2\2\u0083\u0084\7o\2\2\u0084\34"+
		"\3\2\2\2\u0085\u0086\7?\2\2\u0086\36\3\2\2\2\u0087\u0088\7]\2\2\u0088"+
		" \3\2\2\2\u0089\u008a\7_\2\2\u008a\"\3\2\2\2\u008b\u008c\7d\2\2\u008c"+
		"\u008d\7q\2\2\u008d\u008e\7q\2\2\u008e\u008f\7n\2\2\u008f$\3\2\2\2\u0090"+
		"\u0091\7k\2\2\u0091\u0092\7p\2\2\u0092\u0093\7v\2\2\u0093&\3\2\2\2\u0094"+
		"\u0095\7k\2\2\u0095\u0096\7p\2\2\u0096\u0097\7v\2\2\u0097\u0098\7\65\2"+
		"\2\u0098\u0099\7\64\2\2\u0099(\3\2\2\2\u009a\u009b\7k\2\2\u009b\u009c"+
		"\7p\2\2\u009c\u009d\7v\2\2\u009d\u009e\78\2\2\u009e\u009f\7\66\2\2\u009f"+
		"*\3\2\2\2\u00a0\u00a1\7h\2\2\u00a1\u00a2\7n\2\2\u00a2\u00a3\7q\2\2\u00a3"+
		"\u00a4\7c\2\2\u00a4\u00a5\7v\2\2\u00a5,\3\2\2\2\u00a6\u00a7\7h\2\2\u00a7"+
		"\u00a8\7n\2\2\u00a8\u00a9\7q\2\2\u00a9\u00aa\7c\2\2\u00aa\u00ab\7v\2\2"+
		"\u00ab\u00ac\7\65\2\2\u00ac\u00ad\7\64\2\2\u00ad.\3\2\2\2\u00ae\u00af"+
		"\7h\2\2\u00af\u00b0\7n\2\2\u00b0\u00b1\7q\2\2\u00b1\u00b2\7c\2\2\u00b2"+
		"\u00b3\7v\2\2\u00b3\u00b4\78\2\2\u00b4\u00b5\7\66\2\2\u00b5\60\3\2\2\2"+
		"\u00b6\u00b7\7u\2\2\u00b7\u00b8\7v\2\2\u00b8\u00b9\7t\2\2\u00b9\u00ba"+
		"\7k\2\2\u00ba\u00bb\7p\2\2\u00bb\u00bc\7i\2\2\u00bc\62\3\2\2\2\u00bd\u00bf"+
		"\t\2\2\2\u00be\u00bd\3\2\2\2\u00bf\u00c0\3\2\2\2\u00c0\u00be\3\2\2\2\u00c0"+
		"\u00c1\3\2\2\2\u00c1\u00c2\3\2\2\2\u00c2\u00c3\b\32\2\2\u00c3\64\3\2\2"+
		"\2\u00c4\u00c6\t\3\2\2\u00c5\u00c4\3\2\2\2\u00c5\u00c6\3\2\2\2\u00c6\u00c8"+
		"\3\2\2\2\u00c7\u00c9\4\62;\2\u00c8\u00c7\3\2\2\2\u00c9\u00ca\3\2\2\2\u00ca"+
		"\u00c8\3\2\2\2\u00ca\u00cb\3\2\2\2\u00cb\66\3\2\2\2\u00cc\u00cd\7\62\2"+
		"\2\u00cd\u00ce\7z\2\2\u00ce\u00d0\3\2\2\2\u00cf\u00d1\t\4\2\2\u00d0\u00cf"+
		"\3\2\2\2\u00d1\u00d2\3\2\2\2\u00d2\u00d0\3\2\2\2\u00d2\u00d3\3\2\2\2\u00d3"+
		"8\3\2\2\2\u00d4\u00d8\t\5\2\2\u00d5\u00d7\t\6\2\2\u00d6\u00d5\3\2\2\2"+
		"\u00d7\u00da\3\2\2\2\u00d8\u00d6\3\2\2\2\u00d8\u00d9\3\2\2\2\u00d9\u00dd"+
		"\3\2\2\2\u00da\u00d8\3\2\2\2\u00db\u00dc\7]\2\2\u00dc\u00de\7_\2\2\u00dd"+
		"\u00db\3\2\2\2\u00dd\u00de\3\2\2\2\u00de:\3\2\2\2\u00df\u00e3\t\7\2\2"+
		"\u00e0\u00e2\t\b\2\2\u00e1\u00e0\3\2\2\2\u00e2\u00e5\3\2\2\2\u00e3\u00e1"+
		"\3\2\2\2\u00e3\u00e4\3\2\2\2\u00e4<\3\2\2\2\u00e5\u00e3\3\2\2\2\u00e6"+
		"\u00e7\t\t\2\2\u00e7\u00e8\7\60\2\2\u00e8\u00e9\t\t\2\2\u00e9>\3\2\2\2"+
		"\u00ea\u00eb\7\61\2\2\u00eb\u00ec\7\61\2\2\u00ec\u00f0\3\2\2\2\u00ed\u00ef"+
		"\n\n\2\2\u00ee\u00ed\3\2\2\2\u00ef\u00f2\3\2\2\2\u00f0\u00ee\3\2\2\2\u00f0"+
		"\u00f1\3\2\2\2\u00f1@\3\2\2\2\u00f2\u00f0\3\2\2\2\u00f3\u00f7\7B\2\2\u00f4"+
		"\u00f6\n\n\2\2\u00f5\u00f4\3\2\2\2\u00f6\u00f9\3\2\2\2\u00f7\u00f5\3\2"+
		"\2\2\u00f7\u00f8\3\2\2\2\u00f8B\3\2\2\2\u00f9\u00f7\3\2\2\2\u00fa\u00fe"+
		"\7%\2\2\u00fb\u00fd\n\n\2\2\u00fc\u00fb\3\2\2\2\u00fd\u0100\3\2\2\2\u00fe"+
		"\u00fc\3\2\2\2\u00fe\u00ff\3\2\2\2\u00ff\u0101\3\2\2\2\u0100\u00fe\3\2"+
		"\2\2\u0101\u0102\b\"\2\2\u0102D\3\2\2\2\r\2\u00c0\u00c5\u00ca\u00d2\u00d8"+
		"\u00dd\u00e3\u00f0\u00f7\u00fe\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}