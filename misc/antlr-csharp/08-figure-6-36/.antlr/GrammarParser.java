// Generated from /home/luisespino/Documents/GitHub/compilers/antlr-csharp/08-figure-6-36/Grammar.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class GrammarParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, NAME=17, 
		DIGIT=18, WS=19, EQ=20, NEQ=21, LE=22, LT=23, GE=24, GT=25, T=26, F=27;
	public static final int
		RULE_prog = 0, RULE_stmt = 1, RULE_decl = 2, RULE_var = 3, RULE_expr = 4, 
		RULE_term = 5, RULE_fact = 6, RULE_num = 7, RULE_assg = 8, RULE_if = 9, 
		RULE_while = 10, RULE_cond = 11, RULE_bool = 12, RULE_or = 13, RULE_and = 14, 
		RULE_not = 15, RULE_rel = 16, RULE_oprel = 17;
	private static String[] makeRuleNames() {
		return new String[] {
			"prog", "stmt", "decl", "var", "expr", "term", "fact", "num", "assg", 
			"if", "while", "cond", "bool", "or", "and", "not", "rel", "oprel"
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

	@Override
	public String getGrammarFileName() { return "Grammar.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public GrammarParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(GrammarParser.EOF, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public ProgContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prog; }
	}

	public final ProgContext prog() throws RecognitionException {
		ProgContext _localctx = new ProgContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_prog);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(39);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 136194L) != 0)) {
				{
				{
				setState(36);
				stmt();
				}
				}
				setState(41);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(42);
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

	@SuppressWarnings("CheckReturnValue")
	public static class StmtContext extends ParserRuleContext {
		public DeclContext decl() {
			return getRuleContext(DeclContext.class,0);
		}
		public AssgContext assg() {
			return getRuleContext(AssgContext.class,0);
		}
		public IfContext if_() {
			return getRuleContext(IfContext.class,0);
		}
		public WhileContext while_() {
			return getRuleContext(WhileContext.class,0);
		}
		public StmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stmt; }
	}

	public final StmtContext stmt() throws RecognitionException {
		StmtContext _localctx = new StmtContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_stmt);
		try {
			setState(48);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
				enterOuterAlt(_localctx, 1);
				{
				setState(44);
				decl();
				}
				break;
			case NAME:
				enterOuterAlt(_localctx, 2);
				{
				setState(45);
				assg();
				}
				break;
			case T__9:
				enterOuterAlt(_localctx, 3);
				{
				setState(46);
				if_();
				}
				break;
			case T__11:
				enterOuterAlt(_localctx, 4);
				{
				setState(47);
				while_();
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
	public static class DeclContext extends ParserRuleContext {
		public Token type;
		public VarContext var() {
			return getRuleContext(VarContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public DeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_decl; }
	}

	public final DeclContext decl() throws RecognitionException {
		DeclContext _localctx = new DeclContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_decl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(50);
			((DeclContext)_localctx).type = match(T__0);
			setState(51);
			var();
			setState(52);
			match(T__1);
			setState(53);
			expr(0);
			setState(54);
			match(T__2);
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
	public static class VarContext extends ParserRuleContext {
		public TerminalNode NAME() { return getToken(GrammarParser.NAME, 0); }
		public VarContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_var; }
	}

	public final VarContext var() throws RecognitionException {
		VarContext _localctx = new VarContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_var);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(56);
			match(NAME);
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
	public static class ExprContext extends ParserRuleContext {
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	 
		public ExprContext() { }
		public void copyFrom(ExprContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExprSumResContext extends ExprContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TermContext term() {
			return getRuleContext(TermContext.class,0);
		}
		public ExprSumResContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExprTermContext extends ExprContext {
		public TermContext term() {
			return getRuleContext(TermContext.class,0);
		}
		public ExprTermContext(ExprContext ctx) { copyFrom(ctx); }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 8;
		enterRecursionRule(_localctx, 8, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			_localctx = new ExprTermContext(_localctx);
			_ctx = _localctx;
			_prevctx = _localctx;

			setState(59);
			term(0);
			}
			_ctx.stop = _input.LT(-1);
			setState(66);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ExprSumResContext(new ExprContext(_parentctx, _parentState));
					pushNewRecursionContext(_localctx, _startState, RULE_expr);
					setState(61);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(62);
					_la = _input.LA(1);
					if ( !(_la==T__3 || _la==T__4) ) {
					_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					setState(63);
					term(0);
					}
					} 
				}
				setState(68);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
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

	@SuppressWarnings("CheckReturnValue")
	public static class TermContext extends ParserRuleContext {
		public TermContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_term; }
	 
		public TermContext() { }
		public void copyFrom(TermContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExprFactContext extends TermContext {
		public FactContext fact() {
			return getRuleContext(FactContext.class,0);
		}
		public ExprFactContext(TermContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExprMulDivContext extends TermContext {
		public TermContext term() {
			return getRuleContext(TermContext.class,0);
		}
		public FactContext fact() {
			return getRuleContext(FactContext.class,0);
		}
		public ExprMulDivContext(TermContext ctx) { copyFrom(ctx); }
	}

	public final TermContext term() throws RecognitionException {
		return term(0);
	}

	private TermContext term(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		TermContext _localctx = new TermContext(_ctx, _parentState);
		TermContext _prevctx = _localctx;
		int _startState = 10;
		enterRecursionRule(_localctx, 10, RULE_term, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			_localctx = new ExprFactContext(_localctx);
			_ctx = _localctx;
			_prevctx = _localctx;

			setState(70);
			fact();
			}
			_ctx.stop = _input.LT(-1);
			setState(77);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ExprMulDivContext(new TermContext(_parentctx, _parentState));
					pushNewRecursionContext(_localctx, _startState, RULE_term);
					setState(72);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(73);
					_la = _input.LA(1);
					if ( !(_la==T__5 || _la==T__6) ) {
					_errHandler.recoverInline(this);
					}
					else {
						if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
						_errHandler.reportMatch(this);
						consume();
					}
					setState(74);
					fact();
					}
					} 
				}
				setState(79);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,3,_ctx);
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

	@SuppressWarnings("CheckReturnValue")
	public static class FactContext extends ParserRuleContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public VarContext var() {
			return getRuleContext(VarContext.class,0);
		}
		public NumContext num() {
			return getRuleContext(NumContext.class,0);
		}
		public FactContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_fact; }
	}

	public final FactContext fact() throws RecognitionException {
		FactContext _localctx = new FactContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_fact);
		try {
			setState(86);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__7:
				enterOuterAlt(_localctx, 1);
				{
				setState(80);
				match(T__7);
				setState(81);
				expr(0);
				setState(82);
				match(T__8);
				}
				break;
			case NAME:
				enterOuterAlt(_localctx, 2);
				{
				setState(84);
				var();
				}
				break;
			case DIGIT:
				enterOuterAlt(_localctx, 3);
				{
				setState(85);
				num();
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
	public static class NumContext extends ParserRuleContext {
		public TerminalNode DIGIT() { return getToken(GrammarParser.DIGIT, 0); }
		public NumContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_num; }
	}

	public final NumContext num() throws RecognitionException {
		NumContext _localctx = new NumContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_num);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(88);
			match(DIGIT);
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
	public static class AssgContext extends ParserRuleContext {
		public VarContext var() {
			return getRuleContext(VarContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public AssgContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assg; }
	}

	public final AssgContext assg() throws RecognitionException {
		AssgContext _localctx = new AssgContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_assg);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(90);
			var();
			setState(91);
			match(T__1);
			setState(92);
			expr(0);
			setState(93);
			match(T__2);
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
	public static class IfContext extends ParserRuleContext {
		public CondContext cond() {
			return getRuleContext(CondContext.class,0);
		}
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public IfContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if; }
	}

	public final IfContext if_() throws RecognitionException {
		IfContext _localctx = new IfContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_if);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(95);
			match(T__9);
			setState(96);
			cond();
			setState(100);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 136194L) != 0)) {
				{
				{
				setState(97);
				stmt();
				}
				}
				setState(102);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(103);
			match(T__10);
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
	public static class WhileContext extends ParserRuleContext {
		public CondContext cond() {
			return getRuleContext(CondContext.class,0);
		}
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public WhileContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_while; }
	}

	public final WhileContext while_() throws RecognitionException {
		WhileContext _localctx = new WhileContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_while);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(105);
			match(T__11);
			setState(106);
			cond();
			setState(110);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 136194L) != 0)) {
				{
				{
				setState(107);
				stmt();
				}
				}
				setState(112);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(113);
			match(T__12);
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
	public static class CondContext extends ParserRuleContext {
		public BoolContext bool() {
			return getRuleContext(BoolContext.class,0);
		}
		public CondContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cond; }
	}

	public final CondContext cond() throws RecognitionException {
		CondContext _localctx = new CondContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_cond);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(115);
			match(T__7);
			setState(116);
			bool();
			setState(117);
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
	public static class BoolContext extends ParserRuleContext {
		public OrContext or() {
			return getRuleContext(OrContext.class,0);
		}
		public BoolContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_bool; }
	}

	public final BoolContext bool() throws RecognitionException {
		BoolContext _localctx = new BoolContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_bool);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(119);
			or();
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
	public static class OrContext extends ParserRuleContext {
		public OrContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_or; }
	 
		public OrContext() { }
		public void copyFrom(OrContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class OrExprContext extends OrContext {
		public AndContext and() {
			return getRuleContext(AndContext.class,0);
		}
		public OrContext or() {
			return getRuleContext(OrContext.class,0);
		}
		public OrExprContext(OrContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class OrPassContext extends OrContext {
		public AndContext and() {
			return getRuleContext(AndContext.class,0);
		}
		public OrPassContext(OrContext ctx) { copyFrom(ctx); }
	}

	public final OrContext or() throws RecognitionException {
		OrContext _localctx = new OrContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_or);
		try {
			setState(126);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				_localctx = new OrExprContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(121);
				and();
				setState(122);
				match(T__13);
				setState(123);
				or();
				}
				break;
			case 2:
				_localctx = new OrPassContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(125);
				and();
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
	public static class AndContext extends ParserRuleContext {
		public AndContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_and; }
	 
		public AndContext() { }
		public void copyFrom(AndContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AndExprContext extends AndContext {
		public NotContext not() {
			return getRuleContext(NotContext.class,0);
		}
		public AndContext and() {
			return getRuleContext(AndContext.class,0);
		}
		public AndExprContext(AndContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AndPassContext extends AndContext {
		public NotContext not() {
			return getRuleContext(NotContext.class,0);
		}
		public AndPassContext(AndContext ctx) { copyFrom(ctx); }
	}

	public final AndContext and() throws RecognitionException {
		AndContext _localctx = new AndContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_and);
		try {
			setState(133);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				_localctx = new AndExprContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(128);
				not();
				setState(129);
				match(T__14);
				setState(130);
				and();
				}
				break;
			case 2:
				_localctx = new AndPassContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(132);
				not();
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
	public static class NotContext extends ParserRuleContext {
		public NotContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_not; }
	 
		public NotContext() { }
		public void copyFrom(NotContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class NotExprContext extends NotContext {
		public NotContext not() {
			return getRuleContext(NotContext.class,0);
		}
		public NotExprContext(NotContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class NotPassContext extends NotContext {
		public RelContext rel() {
			return getRuleContext(RelContext.class,0);
		}
		public NotPassContext(NotContext ctx) { copyFrom(ctx); }
	}

	public final NotContext not() throws RecognitionException {
		NotContext _localctx = new NotContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_not);
		try {
			setState(138);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__15:
				_localctx = new NotExprContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(135);
				match(T__15);
				setState(136);
				not();
				}
				break;
			case T__7:
			case NAME:
			case DIGIT:
				_localctx = new NotPassContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(137);
				rel();
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
	public static class RelContext extends ParserRuleContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public OprelContext oprel() {
			return getRuleContext(OprelContext.class,0);
		}
		public RelContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rel; }
	}

	public final RelContext rel() throws RecognitionException {
		RelContext _localctx = new RelContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_rel);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(140);
			expr(0);
			setState(141);
			oprel();
			setState(142);
			expr(0);
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
	public static class OprelContext extends ParserRuleContext {
		public TerminalNode EQ() { return getToken(GrammarParser.EQ, 0); }
		public TerminalNode NEQ() { return getToken(GrammarParser.NEQ, 0); }
		public TerminalNode LE() { return getToken(GrammarParser.LE, 0); }
		public TerminalNode LT() { return getToken(GrammarParser.LT, 0); }
		public TerminalNode GE() { return getToken(GrammarParser.GE, 0); }
		public TerminalNode GT() { return getToken(GrammarParser.GT, 0); }
		public OprelContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_oprel; }
	}

	public final OprelContext oprel() throws RecognitionException {
		OprelContext _localctx = new OprelContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_oprel);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(144);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 66060288L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 4:
			return expr_sempred((ExprContext)_localctx, predIndex);
		case 5:
			return term_sempred((TermContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean term_sempred(TermContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001\u001b\u0093\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001"+
		"\u0002\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004"+
		"\u0002\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007"+
		"\u0002\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b"+
		"\u0002\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007"+
		"\u000f\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0001\u0000\u0005"+
		"\u0000&\b\u0000\n\u0000\f\u0000)\t\u0000\u0001\u0000\u0001\u0000\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0003\u00011\b\u0001\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0005\u0004A\b\u0004\n\u0004\f\u0004D\t\u0004\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0005"+
		"\u0005L\b\u0005\n\u0005\f\u0005O\t\u0005\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006W\b\u0006\u0001"+
		"\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\t\u0001"+
		"\t\u0001\t\u0005\tc\b\t\n\t\f\tf\t\t\u0001\t\u0001\t\u0001\n\u0001\n\u0001"+
		"\n\u0005\nm\b\n\n\n\f\np\t\n\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0001\f\u0001\f\u0001\r\u0001\r\u0001\r\u0001\r\u0001"+
		"\r\u0003\r\u007f\b\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0003\u000e\u0086\b\u000e\u0001\u000f\u0001\u000f\u0001\u000f\u0003"+
		"\u000f\u008b\b\u000f\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0011\u0001\u0011\u0001\u0011\u0000\u0002\b\n\u0012\u0000\u0002\u0004"+
		"\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \""+
		"\u0000\u0003\u0001\u0000\u0004\u0005\u0001\u0000\u0006\u0007\u0001\u0000"+
		"\u0014\u0019\u008d\u0000\'\u0001\u0000\u0000\u0000\u00020\u0001\u0000"+
		"\u0000\u0000\u00042\u0001\u0000\u0000\u0000\u00068\u0001\u0000\u0000\u0000"+
		"\b:\u0001\u0000\u0000\u0000\nE\u0001\u0000\u0000\u0000\fV\u0001\u0000"+
		"\u0000\u0000\u000eX\u0001\u0000\u0000\u0000\u0010Z\u0001\u0000\u0000\u0000"+
		"\u0012_\u0001\u0000\u0000\u0000\u0014i\u0001\u0000\u0000\u0000\u0016s"+
		"\u0001\u0000\u0000\u0000\u0018w\u0001\u0000\u0000\u0000\u001a~\u0001\u0000"+
		"\u0000\u0000\u001c\u0085\u0001\u0000\u0000\u0000\u001e\u008a\u0001\u0000"+
		"\u0000\u0000 \u008c\u0001\u0000\u0000\u0000\"\u0090\u0001\u0000\u0000"+
		"\u0000$&\u0003\u0002\u0001\u0000%$\u0001\u0000\u0000\u0000&)\u0001\u0000"+
		"\u0000\u0000\'%\u0001\u0000\u0000\u0000\'(\u0001\u0000\u0000\u0000(*\u0001"+
		"\u0000\u0000\u0000)\'\u0001\u0000\u0000\u0000*+\u0005\u0000\u0000\u0001"+
		"+\u0001\u0001\u0000\u0000\u0000,1\u0003\u0004\u0002\u0000-1\u0003\u0010"+
		"\b\u0000.1\u0003\u0012\t\u0000/1\u0003\u0014\n\u00000,\u0001\u0000\u0000"+
		"\u00000-\u0001\u0000\u0000\u00000.\u0001\u0000\u0000\u00000/\u0001\u0000"+
		"\u0000\u00001\u0003\u0001\u0000\u0000\u000023\u0005\u0001\u0000\u0000"+
		"34\u0003\u0006\u0003\u000045\u0005\u0002\u0000\u000056\u0003\b\u0004\u0000"+
		"67\u0005\u0003\u0000\u00007\u0005\u0001\u0000\u0000\u000089\u0005\u0011"+
		"\u0000\u00009\u0007\u0001\u0000\u0000\u0000:;\u0006\u0004\uffff\uffff"+
		"\u0000;<\u0003\n\u0005\u0000<B\u0001\u0000\u0000\u0000=>\n\u0002\u0000"+
		"\u0000>?\u0007\u0000\u0000\u0000?A\u0003\n\u0005\u0000@=\u0001\u0000\u0000"+
		"\u0000AD\u0001\u0000\u0000\u0000B@\u0001\u0000\u0000\u0000BC\u0001\u0000"+
		"\u0000\u0000C\t\u0001\u0000\u0000\u0000DB\u0001\u0000\u0000\u0000EF\u0006"+
		"\u0005\uffff\uffff\u0000FG\u0003\f\u0006\u0000GM\u0001\u0000\u0000\u0000"+
		"HI\n\u0002\u0000\u0000IJ\u0007\u0001\u0000\u0000JL\u0003\f\u0006\u0000"+
		"KH\u0001\u0000\u0000\u0000LO\u0001\u0000\u0000\u0000MK\u0001\u0000\u0000"+
		"\u0000MN\u0001\u0000\u0000\u0000N\u000b\u0001\u0000\u0000\u0000OM\u0001"+
		"\u0000\u0000\u0000PQ\u0005\b\u0000\u0000QR\u0003\b\u0004\u0000RS\u0005"+
		"\t\u0000\u0000SW\u0001\u0000\u0000\u0000TW\u0003\u0006\u0003\u0000UW\u0003"+
		"\u000e\u0007\u0000VP\u0001\u0000\u0000\u0000VT\u0001\u0000\u0000\u0000"+
		"VU\u0001\u0000\u0000\u0000W\r\u0001\u0000\u0000\u0000XY\u0005\u0012\u0000"+
		"\u0000Y\u000f\u0001\u0000\u0000\u0000Z[\u0003\u0006\u0003\u0000[\\\u0005"+
		"\u0002\u0000\u0000\\]\u0003\b\u0004\u0000]^\u0005\u0003\u0000\u0000^\u0011"+
		"\u0001\u0000\u0000\u0000_`\u0005\n\u0000\u0000`d\u0003\u0016\u000b\u0000"+
		"ac\u0003\u0002\u0001\u0000ba\u0001\u0000\u0000\u0000cf\u0001\u0000\u0000"+
		"\u0000db\u0001\u0000\u0000\u0000de\u0001\u0000\u0000\u0000eg\u0001\u0000"+
		"\u0000\u0000fd\u0001\u0000\u0000\u0000gh\u0005\u000b\u0000\u0000h\u0013"+
		"\u0001\u0000\u0000\u0000ij\u0005\f\u0000\u0000jn\u0003\u0016\u000b\u0000"+
		"km\u0003\u0002\u0001\u0000lk\u0001\u0000\u0000\u0000mp\u0001\u0000\u0000"+
		"\u0000nl\u0001\u0000\u0000\u0000no\u0001\u0000\u0000\u0000oq\u0001\u0000"+
		"\u0000\u0000pn\u0001\u0000\u0000\u0000qr\u0005\r\u0000\u0000r\u0015\u0001"+
		"\u0000\u0000\u0000st\u0005\b\u0000\u0000tu\u0003\u0018\f\u0000uv\u0005"+
		"\t\u0000\u0000v\u0017\u0001\u0000\u0000\u0000wx\u0003\u001a\r\u0000x\u0019"+
		"\u0001\u0000\u0000\u0000yz\u0003\u001c\u000e\u0000z{\u0005\u000e\u0000"+
		"\u0000{|\u0003\u001a\r\u0000|\u007f\u0001\u0000\u0000\u0000}\u007f\u0003"+
		"\u001c\u000e\u0000~y\u0001\u0000\u0000\u0000~}\u0001\u0000\u0000\u0000"+
		"\u007f\u001b\u0001\u0000\u0000\u0000\u0080\u0081\u0003\u001e\u000f\u0000"+
		"\u0081\u0082\u0005\u000f\u0000\u0000\u0082\u0083\u0003\u001c\u000e\u0000"+
		"\u0083\u0086\u0001\u0000\u0000\u0000\u0084\u0086\u0003\u001e\u000f\u0000"+
		"\u0085\u0080\u0001\u0000\u0000\u0000\u0085\u0084\u0001\u0000\u0000\u0000"+
		"\u0086\u001d\u0001\u0000\u0000\u0000\u0087\u0088\u0005\u0010\u0000\u0000"+
		"\u0088\u008b\u0003\u001e\u000f\u0000\u0089\u008b\u0003 \u0010\u0000\u008a"+
		"\u0087\u0001\u0000\u0000\u0000\u008a\u0089\u0001\u0000\u0000\u0000\u008b"+
		"\u001f\u0001\u0000\u0000\u0000\u008c\u008d\u0003\b\u0004\u0000\u008d\u008e"+
		"\u0003\"\u0011\u0000\u008e\u008f\u0003\b\u0004\u0000\u008f!\u0001\u0000"+
		"\u0000\u0000\u0090\u0091\u0007\u0002\u0000\u0000\u0091#\u0001\u0000\u0000"+
		"\u0000\n\'0BMVdn~\u0085\u008a";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}