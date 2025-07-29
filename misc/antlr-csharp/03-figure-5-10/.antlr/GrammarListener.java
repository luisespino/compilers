// Generated from /home/luisespino/Documents/GitHub/compilers/antlr-csharp/03-figure-5-10/Grammar.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link GrammarParser}.
 */
public interface GrammarListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by the {@code Sumres}
	 * labeled alternative in {@link GrammarParser#e}.
	 * @param ctx the parse tree
	 */
	void enterSumres(GrammarParser.SumresContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Sumres}
	 * labeled alternative in {@link GrammarParser#e}.
	 * @param ctx the parse tree
	 */
	void exitSumres(GrammarParser.SumresContext ctx);
	/**
	 * Enter a parse tree produced by the {@code PassT}
	 * labeled alternative in {@link GrammarParser#e}.
	 * @param ctx the parse tree
	 */
	void enterPassT(GrammarParser.PassTContext ctx);
	/**
	 * Exit a parse tree produced by the {@code PassT}
	 * labeled alternative in {@link GrammarParser#e}.
	 * @param ctx the parse tree
	 */
	void exitPassT(GrammarParser.PassTContext ctx);
	/**
	 * Enter a parse tree produced by the {@code PassF}
	 * labeled alternative in {@link GrammarParser#t}.
	 * @param ctx the parse tree
	 */
	void enterPassF(GrammarParser.PassFContext ctx);
	/**
	 * Exit a parse tree produced by the {@code PassF}
	 * labeled alternative in {@link GrammarParser#t}.
	 * @param ctx the parse tree
	 */
	void exitPassF(GrammarParser.PassFContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Muldiv}
	 * labeled alternative in {@link GrammarParser#t}.
	 * @param ctx the parse tree
	 */
	void enterMuldiv(GrammarParser.MuldivContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Muldiv}
	 * labeled alternative in {@link GrammarParser#t}.
	 * @param ctx the parse tree
	 */
	void exitMuldiv(GrammarParser.MuldivContext ctx);
	/**
	 * Enter a parse tree produced by the {@code PassE}
	 * labeled alternative in {@link GrammarParser#f}.
	 * @param ctx the parse tree
	 */
	void enterPassE(GrammarParser.PassEContext ctx);
	/**
	 * Exit a parse tree produced by the {@code PassE}
	 * labeled alternative in {@link GrammarParser#f}.
	 * @param ctx the parse tree
	 */
	void exitPassE(GrammarParser.PassEContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Id}
	 * labeled alternative in {@link GrammarParser#f}.
	 * @param ctx the parse tree
	 */
	void enterId(GrammarParser.IdContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Id}
	 * labeled alternative in {@link GrammarParser#f}.
	 * @param ctx the parse tree
	 */
	void exitId(GrammarParser.IdContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Num}
	 * labeled alternative in {@link GrammarParser#f}.
	 * @param ctx the parse tree
	 */
	void enterNum(GrammarParser.NumContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Num}
	 * labeled alternative in {@link GrammarParser#f}.
	 * @param ctx the parse tree
	 */
	void exitNum(GrammarParser.NumContext ctx);
}