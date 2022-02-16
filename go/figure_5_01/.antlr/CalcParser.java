// Generated from c:\Users\lufes\Documents\GitHub\compilers\go\figure_5_1\Calc.g4 by ANTLR 4.8
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
		MUL=1, ADD=2, LB=3, RB=4, DIGIT=5, WS=6;
	public static final int
		RULE_l = 0, RULE_e = 1, RULE_t = 2, RULE_f = 3;
	private static String[] makeRuleNames() {
		return new String[] {
			"l", "e", "t", "f"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'*'", "'+'", "'('", "')'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "MUL", "ADD", "LB", "RB", "DIGIT", "WS"
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
			setState(8);
			e(0);
			setState(9);
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
	public static class PassTContext extends EContext {
		public TContext t() {
			return getRuleContext(TContext.class,0);
		}
		public PassTContext(EContext ctx) { copyFrom(ctx); }
	}
	public static class SumContext extends EContext {
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode ADD() { return getToken(CalcParser.ADD, 0); }
		public TContext t() {
			return getRuleContext(TContext.class,0);
		}
		public SumContext(EContext ctx) { copyFrom(ctx); }
	}

	public final EContext e() throws RecognitionException {
		return e(0);
	}

	private EContext e(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		EContext _localctx = new EContext(_ctx, _parentState);
		EContext _prevctx = _localctx;
		int _startState = 2;
		enterRecursionRule(_localctx, 2, RULE_e, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			_localctx = new PassTContext(_localctx);
			_ctx = _localctx;
			_prevctx = _localctx;

			setState(12);
			t(0);
			}
			_ctx.stop = _input.LT(-1);
			setState(19);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new SumContext(new EContext(_parentctx, _parentState));
					pushNewRecursionContext(_localctx, _startState, RULE_e);
					setState(14);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(15);
					match(ADD);
					setState(16);
					t(0);
					}
					} 
				}
				setState(21);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
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
	public static class PassFContext extends TContext {
		public FContext f() {
			return getRuleContext(FContext.class,0);
		}
		public PassFContext(TContext ctx) { copyFrom(ctx); }
	}
	public static class MulContext extends TContext {
		public TContext t() {
			return getRuleContext(TContext.class,0);
		}
		public TerminalNode MUL() { return getToken(CalcParser.MUL, 0); }
		public FContext f() {
			return getRuleContext(FContext.class,0);
		}
		public MulContext(TContext ctx) { copyFrom(ctx); }
	}

	public final TContext t() throws RecognitionException {
		return t(0);
	}

	private TContext t(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		TContext _localctx = new TContext(_ctx, _parentState);
		TContext _prevctx = _localctx;
		int _startState = 4;
		enterRecursionRule(_localctx, 4, RULE_t, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			_localctx = new PassFContext(_localctx);
			_ctx = _localctx;
			_prevctx = _localctx;

			setState(23);
			f();
			}
			_ctx.stop = _input.LT(-1);
			setState(30);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new MulContext(new TContext(_parentctx, _parentState));
					pushNewRecursionContext(_localctx, _startState, RULE_t);
					setState(25);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(26);
					match(MUL);
					setState(27);
					f();
					}
					} 
				}
				setState(32);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,1,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
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
	public static class PassEContext extends FContext {
		public TerminalNode LB() { return getToken(CalcParser.LB, 0); }
		public EContext e() {
			return getRuleContext(EContext.class,0);
		}
		public TerminalNode RB() { return getToken(CalcParser.RB, 0); }
		public PassEContext(FContext ctx) { copyFrom(ctx); }
	}
	public static class DigitContext extends FContext {
		public TerminalNode DIGIT() { return getToken(CalcParser.DIGIT, 0); }
		public DigitContext(FContext ctx) { copyFrom(ctx); }
	}

	public final FContext f() throws RecognitionException {
		FContext _localctx = new FContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_f);
		try {
			setState(38);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LB:
				_localctx = new PassEContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(33);
				match(LB);
				setState(34);
				e(0);
				setState(35);
				match(RB);
				}
				break;
			case DIGIT:
				_localctx = new DigitContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(37);
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 1:
			return e_sempred((EContext)_localctx, predIndex);
		case 2:
			return t_sempred((TContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean e_sempred(EContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean t_sempred(TContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\b+\4\2\t\2\4\3\t"+
		"\3\4\4\t\4\4\5\t\5\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3\3\7\3\24\n\3\f\3"+
		"\16\3\27\13\3\3\4\3\4\3\4\3\4\3\4\3\4\7\4\37\n\4\f\4\16\4\"\13\4\3\5\3"+
		"\5\3\5\3\5\3\5\5\5)\n\5\3\5\2\4\4\6\6\2\4\6\b\2\2\2)\2\n\3\2\2\2\4\r\3"+
		"\2\2\2\6\30\3\2\2\2\b(\3\2\2\2\n\13\5\4\3\2\13\f\7\2\2\3\f\3\3\2\2\2\r"+
		"\16\b\3\1\2\16\17\5\6\4\2\17\25\3\2\2\2\20\21\f\4\2\2\21\22\7\4\2\2\22"+
		"\24\5\6\4\2\23\20\3\2\2\2\24\27\3\2\2\2\25\23\3\2\2\2\25\26\3\2\2\2\26"+
		"\5\3\2\2\2\27\25\3\2\2\2\30\31\b\4\1\2\31\32\5\b\5\2\32 \3\2\2\2\33\34"+
		"\f\4\2\2\34\35\7\3\2\2\35\37\5\b\5\2\36\33\3\2\2\2\37\"\3\2\2\2 \36\3"+
		"\2\2\2 !\3\2\2\2!\7\3\2\2\2\" \3\2\2\2#$\7\5\2\2$%\5\4\3\2%&\7\6\2\2&"+
		")\3\2\2\2\')\7\7\2\2(#\3\2\2\2(\'\3\2\2\2)\t\3\2\2\2\5\25 (";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}