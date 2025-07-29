// Code generated from Grammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Grammar

import "github.com/antlr4-go/antlr/v4"

// BaseGrammarListener is a complete listener for a parse tree produced by GrammarParser.
type BaseGrammarListener struct{}

var _ GrammarListener = &BaseGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production Program is entered.
func (s *BaseGrammarListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production Program is exited.
func (s *BaseGrammarListener) ExitProgram(ctx *ProgramContext) {}

// EnterStmtDeclaration is called when production StmtDeclaration is entered.
func (s *BaseGrammarListener) EnterStmtDeclaration(ctx *StmtDeclarationContext) {}

// ExitStmtDeclaration is called when production StmtDeclaration is exited.
func (s *BaseGrammarListener) ExitStmtDeclaration(ctx *StmtDeclarationContext) {}

// EnterStmtAssigment is called when production StmtAssigment is entered.
func (s *BaseGrammarListener) EnterStmtAssigment(ctx *StmtAssigmentContext) {}

// ExitStmtAssigment is called when production StmtAssigment is exited.
func (s *BaseGrammarListener) ExitStmtAssigment(ctx *StmtAssigmentContext) {}

// EnterDeclExpr is called when production DeclExpr is entered.
func (s *BaseGrammarListener) EnterDeclExpr(ctx *DeclExprContext) {}

// ExitDeclExpr is called when production DeclExpr is exited.
func (s *BaseGrammarListener) ExitDeclExpr(ctx *DeclExprContext) {}

// EnterDeclVal is called when production DeclVal is entered.
func (s *BaseGrammarListener) EnterDeclVal(ctx *DeclValContext) {}

// ExitDeclVal is called when production DeclVal is exited.
func (s *BaseGrammarListener) ExitDeclVal(ctx *DeclValContext) {}

// EnterAssgArray is called when production AssgArray is entered.
func (s *BaseGrammarListener) EnterAssgArray(ctx *AssgArrayContext) {}

// ExitAssgArray is called when production AssgArray is exited.
func (s *BaseGrammarListener) ExitAssgArray(ctx *AssgArrayContext) {}

// EnterVarName is called when production VarName is entered.
func (s *BaseGrammarListener) EnterVarName(ctx *VarNameContext) {}

// ExitVarName is called when production VarName is exited.
func (s *BaseGrammarListener) ExitVarName(ctx *VarNameContext) {}

// EnterExprEmpty is called when production ExprEmpty is entered.
func (s *BaseGrammarListener) EnterExprEmpty(ctx *ExprEmptyContext) {}

// ExitExprEmpty is called when production ExprEmpty is exited.
func (s *BaseGrammarListener) ExitExprEmpty(ctx *ExprEmptyContext) {}

// EnterExprValues is called when production ExprValues is entered.
func (s *BaseGrammarListener) EnterExprValues(ctx *ExprValuesContext) {}

// ExitExprValues is called when production ExprValues is exited.
func (s *BaseGrammarListener) ExitExprValues(ctx *ExprValuesContext) {}

// EnterValDigit is called when production ValDigit is entered.
func (s *BaseGrammarListener) EnterValDigit(ctx *ValDigitContext) {}

// ExitValDigit is called when production ValDigit is exited.
func (s *BaseGrammarListener) ExitValDigit(ctx *ValDigitContext) {}

// EnterValExpr is called when production ValExpr is entered.
func (s *BaseGrammarListener) EnterValExpr(ctx *ValExprContext) {}

// ExitValExpr is called when production ValExpr is exited.
func (s *BaseGrammarListener) ExitValExpr(ctx *ValExprContext) {}

// EnterIndexMany is called when production IndexMany is entered.
func (s *BaseGrammarListener) EnterIndexMany(ctx *IndexManyContext) {}

// ExitIndexMany is called when production IndexMany is exited.
func (s *BaseGrammarListener) ExitIndexMany(ctx *IndexManyContext) {}

// EnterIndexOne is called when production IndexOne is entered.
func (s *BaseGrammarListener) EnterIndexOne(ctx *IndexOneContext) {}

// ExitIndexOne is called when production IndexOne is exited.
func (s *BaseGrammarListener) ExitIndexOne(ctx *IndexOneContext) {}

// EnterValues is called when production Values is entered.
func (s *BaseGrammarListener) EnterValues(ctx *ValuesContext) {}

// ExitValues is called when production Values is exited.
func (s *BaseGrammarListener) ExitValues(ctx *ValuesContext) {}

// EnterValue is called when production Value is entered.
func (s *BaseGrammarListener) EnterValue(ctx *ValueContext) {}

// ExitValue is called when production Value is exited.
func (s *BaseGrammarListener) ExitValue(ctx *ValueContext) {}
