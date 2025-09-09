// Generated from /home/luisespino/Documents/GitHub/compilers/misc/antlr-csharp/07-figure-6-15/Grammar.g4 by ANTLR 4.13.1
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
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, WS=7, NAME=8, DIGIT=9;
	public static final int
		RULE_prog = 0, RULE_stmt = 1, RULE_decl = 2, RULE_assg = 3, RULE_var = 4, 
		RULE_expr = 5, RULE_val = 6, RULE_index = 7, RULE_list = 8;
	private static String[] makeRuleNames() {
		return new String[] {
			"prog", "stmt", "decl", "assg", "var", "expr", "val", "index", "list"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'int'", "'='", "';'", "'['", "']'", "','"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, "WS", "NAME", "DIGIT"
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
		public ProgContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_prog; }
	 
		public ProgContext() { }
		public void copyFrom(ProgContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ProgramContext extends ProgContext {
		public TerminalNode EOF() { return getToken(GrammarParser.EOF, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public ProgramContext(ProgContext ctx) { copyFrom(ctx); }
	}

	public final ProgContext prog() throws RecognitionException {
		ProgContext _localctx = new ProgContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_prog);
		int _la;
		try {
			_localctx = new ProgramContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(21);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__0 || _la==NAME) {
				{
				{
				setState(18);
				stmt();
				}
				}
				setState(23);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(24);
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
		public StmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stmt; }
	 
		public StmtContext() { }
		public void copyFrom(StmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StmtAssigmentContext extends StmtContext {
		public AssgContext assg() {
			return getRuleContext(AssgContext.class,0);
		}
		public StmtAssigmentContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StmtDeclarationContext extends StmtContext {
		public DeclContext decl() {
			return getRuleContext(DeclContext.class,0);
		}
		public StmtDeclarationContext(StmtContext ctx) { copyFrom(ctx); }
	}

	public final StmtContext stmt() throws RecognitionException {
		StmtContext _localctx = new StmtContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_stmt);
		try {
			setState(28);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
				_localctx = new StmtDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(26);
				decl();
				}
				break;
			case NAME:
				_localctx = new StmtAssigmentContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(27);
				assg();
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
		public DeclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_decl; }
	 
		public DeclContext() { }
		public void copyFrom(DeclContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclValueContext extends DeclContext {
		public Token type;
		public VarContext var() {
			return getRuleContext(VarContext.class,0);
		}
		public ValContext val() {
			return getRuleContext(ValContext.class,0);
		}
		public DeclValueContext(DeclContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclExprContext extends DeclContext {
		public Token type;
		public VarContext var() {
			return getRuleContext(VarContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public DeclExprContext(DeclContext ctx) { copyFrom(ctx); }
	}

	public final DeclContext decl() throws RecognitionException {
		DeclContext _localctx = new DeclContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_decl);
		try {
			setState(42);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				_localctx = new DeclExprContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(30);
				((DeclExprContext)_localctx).type = match(T__0);
				setState(31);
				var();
				setState(32);
				match(T__1);
				setState(33);
				expr();
				setState(34);
				match(T__2);
				}
				break;
			case 2:
				_localctx = new DeclValueContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(36);
				((DeclValueContext)_localctx).type = match(T__0);
				setState(37);
				var();
				setState(38);
				match(T__1);
				setState(39);
				val();
				setState(40);
				match(T__2);
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
	public static class AssgContext extends ParserRuleContext {
		public AssgContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assg; }
	 
		public AssgContext() { }
		public void copyFrom(AssgContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssgArrayContext extends AssgContext {
		public List<VarContext> var() {
			return getRuleContexts(VarContext.class);
		}
		public VarContext var(int i) {
			return getRuleContext(VarContext.class,i);
		}
		public IndexContext index() {
			return getRuleContext(IndexContext.class,0);
		}
		public AssgArrayContext(AssgContext ctx) { copyFrom(ctx); }
	}

	public final AssgContext assg() throws RecognitionException {
		AssgContext _localctx = new AssgContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_assg);
		try {
			_localctx = new AssgArrayContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(44);
			var();
			setState(45);
			match(T__1);
			setState(46);
			var();
			setState(47);
			index();
			setState(48);
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
		public VarContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_var; }
	 
		public VarContext() { }
		public void copyFrom(VarContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VarNameContext extends VarContext {
		public TerminalNode NAME() { return getToken(GrammarParser.NAME, 0); }
		public VarNameContext(VarContext ctx) { copyFrom(ctx); }
	}

	public final VarContext var() throws RecognitionException {
		VarContext _localctx = new VarContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_var);
		try {
			_localctx = new VarNameContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(50);
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
	public static class ExprValuesContext extends ExprContext {
		public ListContext list() {
			return getRuleContext(ListContext.class,0);
		}
		public ExprValuesContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExprEmptyContext extends ExprContext {
		public ExprEmptyContext(ExprContext ctx) { copyFrom(ctx); }
	}

	public final ExprContext expr() throws RecognitionException {
		ExprContext _localctx = new ExprContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_expr);
		try {
			setState(58);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				_localctx = new ExprEmptyContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(52);
				match(T__3);
				setState(53);
				match(T__4);
				}
				break;
			case 2:
				_localctx = new ExprValuesContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(54);
				match(T__3);
				setState(55);
				list();
				setState(56);
				match(T__4);
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
	public static class ValContext extends ParserRuleContext {
		public ValContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_val; }
	 
		public ValContext() { }
		public void copyFrom(ValContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValDigitContext extends ValContext {
		public TerminalNode DIGIT() { return getToken(GrammarParser.DIGIT, 0); }
		public ValDigitContext(ValContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValExprContext extends ValContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ValExprContext(ValContext ctx) { copyFrom(ctx); }
	}

	public final ValContext val() throws RecognitionException {
		ValContext _localctx = new ValContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_val);
		try {
			setState(62);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case DIGIT:
				_localctx = new ValDigitContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(60);
				match(DIGIT);
				}
				break;
			case T__3:
				_localctx = new ValExprContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(61);
				expr();
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
	public static class IndexContext extends ParserRuleContext {
		public IndexContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_index; }
	 
		public IndexContext() { }
		public void copyFrom(IndexContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IndexOneContext extends IndexContext {
		public ValContext val() {
			return getRuleContext(ValContext.class,0);
		}
		public IndexOneContext(IndexContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IndexManyContext extends IndexContext {
		public ValContext val() {
			return getRuleContext(ValContext.class,0);
		}
		public IndexContext index() {
			return getRuleContext(IndexContext.class,0);
		}
		public IndexManyContext(IndexContext ctx) { copyFrom(ctx); }
	}

	public final IndexContext index() throws RecognitionException {
		IndexContext _localctx = new IndexContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_index);
		try {
			setState(73);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				_localctx = new IndexManyContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(64);
				match(T__3);
				setState(65);
				val();
				setState(66);
				match(T__4);
				setState(67);
				index();
				}
				break;
			case 2:
				_localctx = new IndexOneContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(69);
				match(T__3);
				setState(70);
				val();
				setState(71);
				match(T__4);
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
	public static class ListContext extends ParserRuleContext {
		public ListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_list; }
	 
		public ListContext() { }
		public void copyFrom(ListContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValuesContext extends ListContext {
		public ValContext val() {
			return getRuleContext(ValContext.class,0);
		}
		public ListContext list() {
			return getRuleContext(ListContext.class,0);
		}
		public ValuesContext(ListContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValueContext extends ListContext {
		public ValContext val() {
			return getRuleContext(ValContext.class,0);
		}
		public ValueContext(ListContext ctx) { copyFrom(ctx); }
	}

	public final ListContext list() throws RecognitionException {
		ListContext _localctx = new ListContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_list);
		try {
			setState(80);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				_localctx = new ValuesContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(75);
				val();
				setState(76);
				match(T__5);
				setState(77);
				list();
				}
				break;
			case 2:
				_localctx = new ValueContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(79);
				val();
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

	public static final String _serializedATN =
		"\u0004\u0001\tS\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0001\u0000\u0005\u0000\u0014\b\u0000\n\u0000\f\u0000\u0017"+
		"\t\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0003\u0001\u001d"+
		"\b\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0003\u0002+\b\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0003\u0005;\b"+
		"\u0005\u0001\u0006\u0001\u0006\u0003\u0006?\b\u0006\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0003\u0007J\b\u0007\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0003\bQ\b\b\u0001\b\u0000\u0000\t\u0000\u0002\u0004\u0006"+
		"\b\n\f\u000e\u0010\u0000\u0000P\u0000\u0015\u0001\u0000\u0000\u0000\u0002"+
		"\u001c\u0001\u0000\u0000\u0000\u0004*\u0001\u0000\u0000\u0000\u0006,\u0001"+
		"\u0000\u0000\u0000\b2\u0001\u0000\u0000\u0000\n:\u0001\u0000\u0000\u0000"+
		"\f>\u0001\u0000\u0000\u0000\u000eI\u0001\u0000\u0000\u0000\u0010P\u0001"+
		"\u0000\u0000\u0000\u0012\u0014\u0003\u0002\u0001\u0000\u0013\u0012\u0001"+
		"\u0000\u0000\u0000\u0014\u0017\u0001\u0000\u0000\u0000\u0015\u0013\u0001"+
		"\u0000\u0000\u0000\u0015\u0016\u0001\u0000\u0000\u0000\u0016\u0018\u0001"+
		"\u0000\u0000\u0000\u0017\u0015\u0001\u0000\u0000\u0000\u0018\u0019\u0005"+
		"\u0000\u0000\u0001\u0019\u0001\u0001\u0000\u0000\u0000\u001a\u001d\u0003"+
		"\u0004\u0002\u0000\u001b\u001d\u0003\u0006\u0003\u0000\u001c\u001a\u0001"+
		"\u0000\u0000\u0000\u001c\u001b\u0001\u0000\u0000\u0000\u001d\u0003\u0001"+
		"\u0000\u0000\u0000\u001e\u001f\u0005\u0001\u0000\u0000\u001f \u0003\b"+
		"\u0004\u0000 !\u0005\u0002\u0000\u0000!\"\u0003\n\u0005\u0000\"#\u0005"+
		"\u0003\u0000\u0000#+\u0001\u0000\u0000\u0000$%\u0005\u0001\u0000\u0000"+
		"%&\u0003\b\u0004\u0000&\'\u0005\u0002\u0000\u0000\'(\u0003\f\u0006\u0000"+
		"()\u0005\u0003\u0000\u0000)+\u0001\u0000\u0000\u0000*\u001e\u0001\u0000"+
		"\u0000\u0000*$\u0001\u0000\u0000\u0000+\u0005\u0001\u0000\u0000\u0000"+
		",-\u0003\b\u0004\u0000-.\u0005\u0002\u0000\u0000./\u0003\b\u0004\u0000"+
		"/0\u0003\u000e\u0007\u000001\u0005\u0003\u0000\u00001\u0007\u0001\u0000"+
		"\u0000\u000023\u0005\b\u0000\u00003\t\u0001\u0000\u0000\u000045\u0005"+
		"\u0004\u0000\u00005;\u0005\u0005\u0000\u000067\u0005\u0004\u0000\u0000"+
		"78\u0003\u0010\b\u000089\u0005\u0005\u0000\u00009;\u0001\u0000\u0000\u0000"+
		":4\u0001\u0000\u0000\u0000:6\u0001\u0000\u0000\u0000;\u000b\u0001\u0000"+
		"\u0000\u0000<?\u0005\t\u0000\u0000=?\u0003\n\u0005\u0000><\u0001\u0000"+
		"\u0000\u0000>=\u0001\u0000\u0000\u0000?\r\u0001\u0000\u0000\u0000@A\u0005"+
		"\u0004\u0000\u0000AB\u0003\f\u0006\u0000BC\u0005\u0005\u0000\u0000CD\u0003"+
		"\u000e\u0007\u0000DJ\u0001\u0000\u0000\u0000EF\u0005\u0004\u0000\u0000"+
		"FG\u0003\f\u0006\u0000GH\u0005\u0005\u0000\u0000HJ\u0001\u0000\u0000\u0000"+
		"I@\u0001\u0000\u0000\u0000IE\u0001\u0000\u0000\u0000J\u000f\u0001\u0000"+
		"\u0000\u0000KL\u0003\f\u0006\u0000LM\u0005\u0006\u0000\u0000MN\u0003\u0010"+
		"\b\u0000NQ\u0001\u0000\u0000\u0000OQ\u0003\f\u0006\u0000PK\u0001\u0000"+
		"\u0000\u0000PO\u0001\u0000\u0000\u0000Q\u0011\u0001\u0000\u0000\u0000"+
		"\u0007\u0015\u001c*:>IP";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}