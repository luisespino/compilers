// Code generated from Control.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Control

import "github.com/antlr4-go/antlr/v4"

// BaseControlListener is a complete listener for a parse tree produced by ControlParser.
type BaseControlListener struct{}

var _ ControlListener = &BaseControlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseControlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseControlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseControlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseControlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseControlListener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseControlListener) ExitProg(ctx *ProgContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseControlListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseControlListener) ExitBlock(ctx *BlockContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseControlListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseControlListener) ExitStmt(ctx *StmtContext) {}

// EnterAssignstmt is called when production assignstmt is entered.
func (s *BaseControlListener) EnterAssignstmt(ctx *AssignstmtContext) {}

// ExitAssignstmt is called when production assignstmt is exited.
func (s *BaseControlListener) ExitAssignstmt(ctx *AssignstmtContext) {}

// EnterPrintlnstmt is called when production printlnstmt is entered.
func (s *BaseControlListener) EnterPrintlnstmt(ctx *PrintlnstmtContext) {}

// ExitPrintlnstmt is called when production printlnstmt is exited.
func (s *BaseControlListener) ExitPrintlnstmt(ctx *PrintlnstmtContext) {}

// EnterPrintstmt is called when production printstmt is entered.
func (s *BaseControlListener) EnterPrintstmt(ctx *PrintstmtContext) {}

// ExitPrintstmt is called when production printstmt is exited.
func (s *BaseControlListener) ExitPrintstmt(ctx *PrintstmtContext) {}

// EnterIfstmt is called when production ifstmt is entered.
func (s *BaseControlListener) EnterIfstmt(ctx *IfstmtContext) {}

// ExitIfstmt is called when production ifstmt is exited.
func (s *BaseControlListener) ExitIfstmt(ctx *IfstmtContext) {}

// EnterWhilestmt is called when production whilestmt is entered.
func (s *BaseControlListener) EnterWhilestmt(ctx *WhilestmtContext) {}

// ExitWhilestmt is called when production whilestmt is exited.
func (s *BaseControlListener) ExitWhilestmt(ctx *WhilestmtContext) {}

// EnterBoolExpr is called when production BoolExpr is entered.
func (s *BaseControlListener) EnterBoolExpr(ctx *BoolExprContext) {}

// ExitBoolExpr is called when production BoolExpr is exited.
func (s *BaseControlListener) ExitBoolExpr(ctx *BoolExprContext) {}

// EnterIdExpr is called when production IdExpr is entered.
func (s *BaseControlListener) EnterIdExpr(ctx *IdExprContext) {}

// ExitIdExpr is called when production IdExpr is exited.
func (s *BaseControlListener) ExitIdExpr(ctx *IdExprContext) {}

// EnterParExpr is called when production ParExpr is entered.
func (s *BaseControlListener) EnterParExpr(ctx *ParExprContext) {}

// ExitParExpr is called when production ParExpr is exited.
func (s *BaseControlListener) ExitParExpr(ctx *ParExprContext) {}

// EnterStrExpr is called when production StrExpr is entered.
func (s *BaseControlListener) EnterStrExpr(ctx *StrExprContext) {}

// ExitStrExpr is called when production StrExpr is exited.
func (s *BaseControlListener) ExitStrExpr(ctx *StrExprContext) {}

// EnterNotExpr is called when production NotExpr is entered.
func (s *BaseControlListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production NotExpr is exited.
func (s *BaseControlListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterIntExpr is called when production IntExpr is entered.
func (s *BaseControlListener) EnterIntExpr(ctx *IntExprContext) {}

// ExitIntExpr is called when production IntExpr is exited.
func (s *BaseControlListener) ExitIntExpr(ctx *IntExprContext) {}

// EnterOpExpr is called when production OpExpr is entered.
func (s *BaseControlListener) EnterOpExpr(ctx *OpExprContext) {}

// ExitOpExpr is called when production OpExpr is exited.
func (s *BaseControlListener) ExitOpExpr(ctx *OpExprContext) {}
