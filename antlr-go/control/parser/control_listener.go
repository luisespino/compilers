// Code generated from Control.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Control

import "github.com/antlr4-go/antlr/v4"

// ControlListener is a complete listener for a parse tree produced by ControlParser.
type ControlListener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterAssignstmt is called when entering the assignstmt production.
	EnterAssignstmt(c *AssignstmtContext)

	// EnterPrintlnstmt is called when entering the printlnstmt production.
	EnterPrintlnstmt(c *PrintlnstmtContext)

	// EnterPrintstmt is called when entering the printstmt production.
	EnterPrintstmt(c *PrintstmtContext)

	// EnterIfstmt is called when entering the ifstmt production.
	EnterIfstmt(c *IfstmtContext)

	// EnterWhilestmt is called when entering the whilestmt production.
	EnterWhilestmt(c *WhilestmtContext)

	// EnterBoolExpr is called when entering the BoolExpr production.
	EnterBoolExpr(c *BoolExprContext)

	// EnterIdExpr is called when entering the IdExpr production.
	EnterIdExpr(c *IdExprContext)

	// EnterParExpr is called when entering the ParExpr production.
	EnterParExpr(c *ParExprContext)

	// EnterStrExpr is called when entering the StrExpr production.
	EnterStrExpr(c *StrExprContext)

	// EnterNotExpr is called when entering the NotExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterIntExpr is called when entering the IntExpr production.
	EnterIntExpr(c *IntExprContext)

	// EnterOpExpr is called when entering the OpExpr production.
	EnterOpExpr(c *OpExprContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitAssignstmt is called when exiting the assignstmt production.
	ExitAssignstmt(c *AssignstmtContext)

	// ExitPrintlnstmt is called when exiting the printlnstmt production.
	ExitPrintlnstmt(c *PrintlnstmtContext)

	// ExitPrintstmt is called when exiting the printstmt production.
	ExitPrintstmt(c *PrintstmtContext)

	// ExitIfstmt is called when exiting the ifstmt production.
	ExitIfstmt(c *IfstmtContext)

	// ExitWhilestmt is called when exiting the whilestmt production.
	ExitWhilestmt(c *WhilestmtContext)

	// ExitBoolExpr is called when exiting the BoolExpr production.
	ExitBoolExpr(c *BoolExprContext)

	// ExitIdExpr is called when exiting the IdExpr production.
	ExitIdExpr(c *IdExprContext)

	// ExitParExpr is called when exiting the ParExpr production.
	ExitParExpr(c *ParExprContext)

	// ExitStrExpr is called when exiting the StrExpr production.
	ExitStrExpr(c *StrExprContext)

	// ExitNotExpr is called when exiting the NotExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitIntExpr is called when exiting the IntExpr production.
	ExitIntExpr(c *IntExprContext)

	// ExitOpExpr is called when exiting the OpExpr production.
	ExitOpExpr(c *OpExprContext)
}
