// Code generated from Grammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Grammar

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by GrammarParser.
type GrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GrammarParser#Program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by GrammarParser#StmtDeclaration.
	VisitStmtDeclaration(ctx *StmtDeclarationContext) interface{}

	// Visit a parse tree produced by GrammarParser#StmtAssigment.
	VisitStmtAssigment(ctx *StmtAssigmentContext) interface{}

	// Visit a parse tree produced by GrammarParser#DeclExpr.
	VisitDeclExpr(ctx *DeclExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#DeclVal.
	VisitDeclVal(ctx *DeclValContext) interface{}

	// Visit a parse tree produced by GrammarParser#AssgArray.
	VisitAssgArray(ctx *AssgArrayContext) interface{}

	// Visit a parse tree produced by GrammarParser#VarName.
	VisitVarName(ctx *VarNameContext) interface{}

	// Visit a parse tree produced by GrammarParser#ExprEmpty.
	VisitExprEmpty(ctx *ExprEmptyContext) interface{}

	// Visit a parse tree produced by GrammarParser#ExprValues.
	VisitExprValues(ctx *ExprValuesContext) interface{}

	// Visit a parse tree produced by GrammarParser#ValDigit.
	VisitValDigit(ctx *ValDigitContext) interface{}

	// Visit a parse tree produced by GrammarParser#ValExpr.
	VisitValExpr(ctx *ValExprContext) interface{}

	// Visit a parse tree produced by GrammarParser#IndexMany.
	VisitIndexMany(ctx *IndexManyContext) interface{}

	// Visit a parse tree produced by GrammarParser#IndexOne.
	VisitIndexOne(ctx *IndexOneContext) interface{}

	// Visit a parse tree produced by GrammarParser#Values.
	VisitValues(ctx *ValuesContext) interface{}

	// Visit a parse tree produced by GrammarParser#Value.
	VisitValue(ctx *ValueContext) interface{}
}
