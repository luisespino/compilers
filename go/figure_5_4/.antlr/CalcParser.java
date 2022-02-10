// Generated from c:\Users\lufes\Documents\GitHub\compilers\go\figure_5_4\Calc.g4 by ANTLR 4.8
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class CalcParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, DIGIT=5, WS=6;
	public static final int
		RULE_l = 0, RULE_e = 1, RULE_ep = 2, RULE_t = 3, RULE_tp = 4, RULE_f = 5;
	private static String[] makeRuleNames() {
		return new String[] {
			"l", "e", "ep", "t", "tp", "f"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'+'", "'*'", "'('", "')'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, "DIGIT", "WS"
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
	public String getGrammarFileName() { return "Calc.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public CalcParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class LContext extends ParserRuleContext {
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode EOF() { return getToken(CalcParser.EOF, 0); }
		public LContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_l; }
	}

	public final LContext l() throws RecognitionException {
		LContext _localctx = new LContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_l);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(12);
			e();
			setState(13);
			match(EOF);
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

	public static class EContext extends ParserRuleContext {
		public EContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_e; }
	 
		public EContext() { }
		public void copyFrom(EContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class ETContext extends EContext {
		public TContext t() {
			return getRuleContext(TContext.class,0);
		}
		public EpContext ep() {
			return getRuleContext(EpContext.class,0);
		}
		public ETContext(EContext ctx) { copyFrom(ctx); }
	}

	public final EContext e() throws RecognitionException {
		EContext _localctx = new EContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_e);
		try {
			_localctx = new ETContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(15);
			t();
			setState(16);
			ep();
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

	public static class EpContext extends ParserRuleContext {
		public EpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ep; }
	 
		public EpContext() { }
		public void copyFrom(EpContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class SumContext extends EpContext {
		public TContext t() {
			return getRuleContext(TContext.class,0);
		}
		public EpContext ep() {
			return getRuleContext(EpContext.class,0);
		}
		public SumContext(EpContext ctx) { copyFrom(ctx); }
	}
	public static class EpsSumContext extends EpContext {
		public EpsSumContext(EpContext ctx) { copyFrom(ctx); }
	}

	public final EpContext ep() throws RecognitionException {
		EpContext _localctx = new EpContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_ep);
		try {
			setState(23);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
				_localctx = new SumContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(18);
				match(T__0);
				setState(19);
				t();
				setState(20);
				ep();
				}
				break;
			case EOF:
			case T__3:
				_localctx = new EpsSumContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
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

	public static class TContext extends ParserRuleContext {
		public TContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_t; }
	 
		public TContext() { }
		public void copyFrom(TContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class TFContext extends TContext {
		public FContext f() {
			return getRuleContext(FContext.class,0);
		}
		public TpContext tp() {
			return getRuleContext(TpContext.class,0);
		}
		public TFContext(TContext ctx) { copyFrom(ctx); }
	}

	public final TContext t() throws RecognitionException {
		TContext _localctx = new TContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_t);
		try {
			_localctx = new TFContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(25);
			f();
			setState(26);
			tp();
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

	public static class TpContext extends ParserRuleContext {
		public TpContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tp; }
	 
		public TpContext() { }
		public void copyFrom(TpContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class EpsMulContext extends TpContext {
		public EpsMulContext(TpContext ctx) { copyFrom(ctx); }
	}
	public static class MulContext extends TpContext {
		public FContext f() {
			return getRuleContext(FContext.class,0);
		}
		public TpContext tp() {
			return getRuleContext(TpContext.class,0);
		}
		public MulContext(TpContext ctx) { copyFrom(ctx); }
	}

	public final TpContext tp() throws RecognitionException {
		TpContext _localctx = new TpContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_tp);
		try {
			setState(33);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__1:
				_localctx = new MulContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(28);
				match(T__1);
				setState(29);
				f();
				setState(30);
				tp();
				}
				break;
			case EOF:
			case T__0:
			case T__3:
				_localctx = new EpsMulContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
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

	public static class FContext extends ParserRuleContext {
		public FContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_f; }
	 
		public FContext() { }
		public void copyFrom(FContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class DigitContext extends FContext {
		public TerminalNode DIGIT() { return getToken(CalcParser.DIGIT, 0); }
		public DigitContext(FContext ctx) { copyFrom(ctx); }
	}
	public static class BraceContext extends FContext {
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public BraceContext(FContext ctx) { copyFrom(ctx); }
	}

	public final FContext f() throws RecognitionException {
		FContext _localctx = new FContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_f);
		try {
			setState(40);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__2:
				_localctx = new BraceContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(35);
				match(T__2);
				setState(36);
				e();
				setState(37);
				match(T__3);
				}
				break;
			case DIGIT:
				_localctx = new DigitContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(39);
				match(DIGIT);
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
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\b-\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\3\2\3\2\3\2\3\3\3\3\3\3\3\4\3\4\3\4"+
		"\3\4\3\4\5\4\32\n\4\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\5\6$\n\6\3\7\3\7\3"+
		"\7\3\7\3\7\5\7+\n\7\3\7\2\2\b\2\4\6\b\n\f\2\2\2)\2\16\3\2\2\2\4\21\3\2"+
		"\2\2\6\31\3\2\2\2\b\33\3\2\2\2\n#\3\2\2\2\f*\3\2\2\2\16\17\5\4\3\2\17"+
		"\20\7\2\2\3\20\3\3\2\2\2\21\22\5\b\5\2\22\23\5\6\4\2\23\5\3\2\2\2\24\25"+
		"\7\3\2\2\25\26\5\b\5\2\26\27\5\6\4\2\27\32\3\2\2\2\30\32\3\2\2\2\31\24"+
		"\3\2\2\2\31\30\3\2\2\2\32\7\3\2\2\2\33\34\5\f\7\2\34\35\5\n\6\2\35\t\3"+
		"\2\2\2\36\37\7\4\2\2\37 \5\f\7\2 !\5\n\6\2!$\3\2\2\2\"$\3\2\2\2#\36\3"+
		"\2\2\2#\"\3\2\2\2$\13\3\2\2\2%&\7\5\2\2&\'\5\4\3\2\'(\7\6\2\2(+\3\2\2"+
		"\2)+\7\7\2\2*%\3\2\2\2*)\3\2\2\2+\r\3\2\2\2\5\31#*";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}