// Generated from /home/luisespino/Documents/GitHub/compilers/antlr-csharp/02-figure-5-04/Grammar.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link GrammarParser}.
 */
public interface GrammarListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link GrammarParser#l}.
	 * @param ctx the parse tree
	 */
	void enterL(GrammarParser.LContext ctx);
	/**
	 * Exit a parse tree produced by {@link GrammarParser#l}.
	 * @param ctx the parse tree
	 */
	void exitL(GrammarParser.LContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ET}
	 * labeled alternative in {@link GrammarParser#e}.
	 * @param ctx the parse tree
	 */
	void enterET(GrammarParser.ETContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ET}
	 * labeled alternative in {@link GrammarParser#e}.
	 * @param ctx the parse tree
	 */
	void exitET(GrammarParser.ETContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Sum}
	 * labeled alternative in {@link GrammarParser#ep}.
	 * @param ctx the parse tree
	 */
	void enterSum(GrammarParser.SumContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Sum}
	 * labeled alternative in {@link GrammarParser#ep}.
	 * @param ctx the parse tree
	 */
	void exitSum(GrammarParser.SumContext ctx);
	/**
	 * Enter a parse tree produced by the {@code EpsSum}
	 * labeled alternative in {@link GrammarParser#ep}.
	 * @param ctx the parse tree
	 */
	void enterEpsSum(GrammarParser.EpsSumContext ctx);
	/**
	 * Exit a parse tree produced by the {@code EpsSum}
	 * labeled alternative in {@link GrammarParser#ep}.
	 * @param ctx the parse tree
	 */
	void exitEpsSum(GrammarParser.EpsSumContext ctx);
	/**
	 * Enter a parse tree produced by the {@code TF}
	 * labeled alternative in {@link GrammarParser#t}.
	 * @param ctx the parse tree
	 */
	void enterTF(GrammarParser.TFContext ctx);
	/**
	 * Exit a parse tree produced by the {@code TF}
	 * labeled alternative in {@link GrammarParser#t}.
	 * @param ctx the parse tree
	 */
	void exitTF(GrammarParser.TFContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Mul}
	 * labeled alternative in {@link GrammarParser#tp}.
	 * @param ctx the parse tree
	 */
	void enterMul(GrammarParser.MulContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Mul}
	 * labeled alternative in {@link GrammarParser#tp}.
	 * @param ctx the parse tree
	 */
	void exitMul(GrammarParser.MulContext ctx);
	/**
	 * Enter a parse tree produced by the {@code EpsMul}
	 * labeled alternative in {@link GrammarParser#tp}.
	 * @param ctx the parse tree
	 */
	void enterEpsMul(GrammarParser.EpsMulContext ctx);
	/**
	 * Exit a parse tree produced by the {@code EpsMul}
	 * labeled alternative in {@link GrammarParser#tp}.
	 * @param ctx the parse tree
	 */
	void exitEpsMul(GrammarParser.EpsMulContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Brace}
	 * labeled alternative in {@link GrammarParser#f}.
	 * @param ctx the parse tree
	 */
	void enterBrace(GrammarParser.BraceContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Brace}
	 * labeled alternative in {@link GrammarParser#f}.
	 * @param ctx the parse tree
	 */
	void exitBrace(GrammarParser.BraceContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Digit}
	 * labeled alternative in {@link GrammarParser#f}.
	 * @param ctx the parse tree
	 */
	void enterDigit(GrammarParser.DigitContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Digit}
	 * labeled alternative in {@link GrammarParser#f}.
	 * @param ctx the parse tree
	 */
	void exitDigit(GrammarParser.DigitContext ctx);
}