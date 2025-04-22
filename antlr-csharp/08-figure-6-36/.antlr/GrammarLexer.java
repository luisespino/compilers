// Generated from /home/luisespino/Documents/GitHub/compilers/antlr-csharp/08-figure-6-36/Grammar.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue", "this-escape"})
public class GrammarLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, NAME=17, 
		DIGIT=18, WS=19, EQ=20, NEQ=21, LE=22, LT=23, GE=24, GT=25, T=26, F=27;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
			"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "NAME", 
			"DIGIT", "WS", "EQ", "NEQ", "LE", "LT", "GE", "GT", "T", "F"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'int'", "'='", "';'", "'+'", "'-'", "'*'", "'/'", "'('", "')'", 
			"'if'", "'endif'", "'while'", "'endwhile'", "'||'", "'&&'", "'!'", null, 
			null, null, "'=='", "'!='", "'<='", "'<'", "'>='", "'>'", "'true'", "'false'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, "NAME", "DIGIT", "WS", "EQ", "NEQ", "LE", 
			"LT", "GE", "GT", "T", "F"
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


	public GrammarLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Grammar.g4"; }

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
		"\u0004\u0000\u001b\u0099\u0006\uffff\uffff\u0002\u0000\u0007\u0000\u0002"+
		"\u0001\u0007\u0001\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002"+
		"\u0004\u0007\u0004\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002"+
		"\u0007\u0007\u0007\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002"+
		"\u000b\u0007\u000b\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e"+
		"\u0002\u000f\u0007\u000f\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011"+
		"\u0002\u0012\u0007\u0012\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014"+
		"\u0002\u0015\u0007\u0015\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017"+
		"\u0002\u0018\u0007\u0018\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a"+
		"\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001"+
		"\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004"+
		"\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007"+
		"\u0001\b\u0001\b\u0001\t\u0001\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001"+
		"\f\u0001\f\u0001\f\u0001\r\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000f\u0001\u000f\u0001\u0010\u0001\u0010\u0005\u0010n\b"+
		"\u0010\n\u0010\f\u0010q\t\u0010\u0001\u0011\u0004\u0011t\b\u0011\u000b"+
		"\u0011\f\u0011u\u0001\u0012\u0004\u0012y\b\u0012\u000b\u0012\f\u0012z"+
		"\u0001\u0012\u0001\u0012\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0016"+
		"\u0001\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0018\u0001\u0018"+
		"\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u001a"+
		"\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0000\u0000"+
		"\u001b\u0001\u0001\u0003\u0002\u0005\u0003\u0007\u0004\t\u0005\u000b\u0006"+
		"\r\u0007\u000f\b\u0011\t\u0013\n\u0015\u000b\u0017\f\u0019\r\u001b\u000e"+
		"\u001d\u000f\u001f\u0010!\u0011#\u0012%\u0013\'\u0014)\u0015+\u0016-\u0017"+
		"/\u00181\u00193\u001a5\u001b\u0001\u0000\u0004\u0003\u0000AZ__az\u0004"+
		"\u000009AZ__az\u0001\u000009\u0003\u0000\t\n\f\r  \u009b\u0000\u0001\u0001"+
		"\u0000\u0000\u0000\u0000\u0003\u0001\u0000\u0000\u0000\u0000\u0005\u0001"+
		"\u0000\u0000\u0000\u0000\u0007\u0001\u0000\u0000\u0000\u0000\t\u0001\u0000"+
		"\u0000\u0000\u0000\u000b\u0001\u0000\u0000\u0000\u0000\r\u0001\u0000\u0000"+
		"\u0000\u0000\u000f\u0001\u0000\u0000\u0000\u0000\u0011\u0001\u0000\u0000"+
		"\u0000\u0000\u0013\u0001\u0000\u0000\u0000\u0000\u0015\u0001\u0000\u0000"+
		"\u0000\u0000\u0017\u0001\u0000\u0000\u0000\u0000\u0019\u0001\u0000\u0000"+
		"\u0000\u0000\u001b\u0001\u0000\u0000\u0000\u0000\u001d\u0001\u0000\u0000"+
		"\u0000\u0000\u001f\u0001\u0000\u0000\u0000\u0000!\u0001\u0000\u0000\u0000"+
		"\u0000#\u0001\u0000\u0000\u0000\u0000%\u0001\u0000\u0000\u0000\u0000\'"+
		"\u0001\u0000\u0000\u0000\u0000)\u0001\u0000\u0000\u0000\u0000+\u0001\u0000"+
		"\u0000\u0000\u0000-\u0001\u0000\u0000\u0000\u0000/\u0001\u0000\u0000\u0000"+
		"\u00001\u0001\u0000\u0000\u0000\u00003\u0001\u0000\u0000\u0000\u00005"+
		"\u0001\u0000\u0000\u0000\u00017\u0001\u0000\u0000\u0000\u0003;\u0001\u0000"+
		"\u0000\u0000\u0005=\u0001\u0000\u0000\u0000\u0007?\u0001\u0000\u0000\u0000"+
		"\tA\u0001\u0000\u0000\u0000\u000bC\u0001\u0000\u0000\u0000\rE\u0001\u0000"+
		"\u0000\u0000\u000fG\u0001\u0000\u0000\u0000\u0011I\u0001\u0000\u0000\u0000"+
		"\u0013K\u0001\u0000\u0000\u0000\u0015N\u0001\u0000\u0000\u0000\u0017T"+
		"\u0001\u0000\u0000\u0000\u0019Z\u0001\u0000\u0000\u0000\u001bc\u0001\u0000"+
		"\u0000\u0000\u001df\u0001\u0000\u0000\u0000\u001fi\u0001\u0000\u0000\u0000"+
		"!k\u0001\u0000\u0000\u0000#s\u0001\u0000\u0000\u0000%x\u0001\u0000\u0000"+
		"\u0000\'~\u0001\u0000\u0000\u0000)\u0081\u0001\u0000\u0000\u0000+\u0084"+
		"\u0001\u0000\u0000\u0000-\u0087\u0001\u0000\u0000\u0000/\u0089\u0001\u0000"+
		"\u0000\u00001\u008c\u0001\u0000\u0000\u00003\u008e\u0001\u0000\u0000\u0000"+
		"5\u0093\u0001\u0000\u0000\u000078\u0005i\u0000\u000089\u0005n\u0000\u0000"+
		"9:\u0005t\u0000\u0000:\u0002\u0001\u0000\u0000\u0000;<\u0005=\u0000\u0000"+
		"<\u0004\u0001\u0000\u0000\u0000=>\u0005;\u0000\u0000>\u0006\u0001\u0000"+
		"\u0000\u0000?@\u0005+\u0000\u0000@\b\u0001\u0000\u0000\u0000AB\u0005-"+
		"\u0000\u0000B\n\u0001\u0000\u0000\u0000CD\u0005*\u0000\u0000D\f\u0001"+
		"\u0000\u0000\u0000EF\u0005/\u0000\u0000F\u000e\u0001\u0000\u0000\u0000"+
		"GH\u0005(\u0000\u0000H\u0010\u0001\u0000\u0000\u0000IJ\u0005)\u0000\u0000"+
		"J\u0012\u0001\u0000\u0000\u0000KL\u0005i\u0000\u0000LM\u0005f\u0000\u0000"+
		"M\u0014\u0001\u0000\u0000\u0000NO\u0005e\u0000\u0000OP\u0005n\u0000\u0000"+
		"PQ\u0005d\u0000\u0000QR\u0005i\u0000\u0000RS\u0005f\u0000\u0000S\u0016"+
		"\u0001\u0000\u0000\u0000TU\u0005w\u0000\u0000UV\u0005h\u0000\u0000VW\u0005"+
		"i\u0000\u0000WX\u0005l\u0000\u0000XY\u0005e\u0000\u0000Y\u0018\u0001\u0000"+
		"\u0000\u0000Z[\u0005e\u0000\u0000[\\\u0005n\u0000\u0000\\]\u0005d\u0000"+
		"\u0000]^\u0005w\u0000\u0000^_\u0005h\u0000\u0000_`\u0005i\u0000\u0000"+
		"`a\u0005l\u0000\u0000ab\u0005e\u0000\u0000b\u001a\u0001\u0000\u0000\u0000"+
		"cd\u0005|\u0000\u0000de\u0005|\u0000\u0000e\u001c\u0001\u0000\u0000\u0000"+
		"fg\u0005&\u0000\u0000gh\u0005&\u0000\u0000h\u001e\u0001\u0000\u0000\u0000"+
		"ij\u0005!\u0000\u0000j \u0001\u0000\u0000\u0000ko\u0007\u0000\u0000\u0000"+
		"ln\u0007\u0001\u0000\u0000ml\u0001\u0000\u0000\u0000nq\u0001\u0000\u0000"+
		"\u0000om\u0001\u0000\u0000\u0000op\u0001\u0000\u0000\u0000p\"\u0001\u0000"+
		"\u0000\u0000qo\u0001\u0000\u0000\u0000rt\u0007\u0002\u0000\u0000sr\u0001"+
		"\u0000\u0000\u0000tu\u0001\u0000\u0000\u0000us\u0001\u0000\u0000\u0000"+
		"uv\u0001\u0000\u0000\u0000v$\u0001\u0000\u0000\u0000wy\u0007\u0003\u0000"+
		"\u0000xw\u0001\u0000\u0000\u0000yz\u0001\u0000\u0000\u0000zx\u0001\u0000"+
		"\u0000\u0000z{\u0001\u0000\u0000\u0000{|\u0001\u0000\u0000\u0000|}\u0006"+
		"\u0012\u0000\u0000}&\u0001\u0000\u0000\u0000~\u007f\u0005=\u0000\u0000"+
		"\u007f\u0080\u0005=\u0000\u0000\u0080(\u0001\u0000\u0000\u0000\u0081\u0082"+
		"\u0005!\u0000\u0000\u0082\u0083\u0005=\u0000\u0000\u0083*\u0001\u0000"+
		"\u0000\u0000\u0084\u0085\u0005<\u0000\u0000\u0085\u0086\u0005=\u0000\u0000"+
		"\u0086,\u0001\u0000\u0000\u0000\u0087\u0088\u0005<\u0000\u0000\u0088."+
		"\u0001\u0000\u0000\u0000\u0089\u008a\u0005>\u0000\u0000\u008a\u008b\u0005"+
		"=\u0000\u0000\u008b0\u0001\u0000\u0000\u0000\u008c\u008d\u0005>\u0000"+
		"\u0000\u008d2\u0001\u0000\u0000\u0000\u008e\u008f\u0005t\u0000\u0000\u008f"+
		"\u0090\u0005r\u0000\u0000\u0090\u0091\u0005u\u0000\u0000\u0091\u0092\u0005"+
		"e\u0000\u0000\u00924\u0001\u0000\u0000\u0000\u0093\u0094\u0005f\u0000"+
		"\u0000\u0094\u0095\u0005a\u0000\u0000\u0095\u0096\u0005l\u0000\u0000\u0096"+
		"\u0097\u0005s\u0000\u0000\u0097\u0098\u0005e\u0000\u0000\u00986\u0001"+
		"\u0000\u0000\u0000\u0004\u0000ouz\u0001\u0006\u0000\u0000";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}