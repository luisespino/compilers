// Code generated from Grammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Grammar

import "github.com/antlr4-go/antlr/v4"

// GrammarListener is a complete listener for a parse tree produced by GrammarParser.
type GrammarListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the Program production.
	EnterProgram(c *ProgramContext)

	// EnterStmtDeclaration is called when entering the StmtDeclaration production.
	EnterStmtDeclaration(c *StmtDeclarationContext)

	// EnterStmtAssigment is called when entering the StmtAssigment production.
	EnterStmtAssigment(c *StmtAssigmentContext)

	// EnterDeclExpr is called when entering the DeclExpr production.
	EnterDeclExpr(c *DeclExprContext)

	// EnterDeclVal is called when entering the DeclVal production.
	EnterDeclVal(c *DeclValContext)

	// EnterAssgArray is called when entering the AssgArray production.
	EnterAssgArray(c *AssgArrayContext)

	// EnterVarName is called when entering the VarName production.
	EnterVarName(c *VarNameContext)

	// EnterExprEmpty is called when entering the ExprEmpty production.
	EnterExprEmpty(c *ExprEmptyContext)

	// EnterExprValues is called when entering the ExprValues production.
	EnterExprValues(c *ExprValuesContext)

	// EnterValDigit is called when entering the ValDigit production.
	EnterValDigit(c *ValDigitContext)

	// EnterValExpr is called when entering the ValExpr production.
	EnterValExpr(c *ValExprContext)

	// EnterIndexMany is called when entering the IndexMany production.
	EnterIndexMany(c *IndexManyContext)

	// EnterIndexOne is called when entering the IndexOne production.
	EnterIndexOne(c *IndexOneContext)

	// EnterValues is called when entering the Values production.
	EnterValues(c *ValuesContext)

	// EnterValue is called when entering the Value production.
	EnterValue(c *ValueContext)

	// ExitProgram is called when exiting the Program production.
	ExitProgram(c *ProgramContext)

	// ExitStmtDeclaration is called when exiting the StmtDeclaration production.
	ExitStmtDeclaration(c *StmtDeclarationContext)

	// ExitStmtAssigment is called when exiting the StmtAssigment production.
	ExitStmtAssigment(c *StmtAssigmentContext)

	// ExitDeclExpr is called when exiting the DeclExpr production.
	ExitDeclExpr(c *DeclExprContext)

	// ExitDeclVal is called when exiting the DeclVal production.
	ExitDeclVal(c *DeclValContext)

	// ExitAssgArray is called when exiting the AssgArray production.
	ExitAssgArray(c *AssgArrayContext)

	// ExitVarName is called when exiting the VarName production.
	ExitVarName(c *VarNameContext)

	// ExitExprEmpty is called when exiting the ExprEmpty production.
	ExitExprEmpty(c *ExprEmptyContext)

	// ExitExprValues is called when exiting the ExprValues production.
	ExitExprValues(c *ExprValuesContext)

	// ExitValDigit is called when exiting the ValDigit production.
	ExitValDigit(c *ValDigitContext)

	// ExitValExpr is called when exiting the ValExpr production.
	ExitValExpr(c *ValExprContext)

	// ExitIndexMany is called when exiting the IndexMany production.
	ExitIndexMany(c *IndexManyContext)

	// ExitIndexOne is called when exiting the IndexOne production.
	ExitIndexOne(c *IndexOneContext)

	// ExitValues is called when exiting the Values production.
	ExitValues(c *ValuesContext)

	// ExitValue is called when exiting the Value production.
	ExitValue(c *ValueContext)
}
