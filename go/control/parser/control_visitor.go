// Code generated from Control.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Control

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ControlParser.
type ControlVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ControlParser#prog.
	VisitProg(ctx *ProgContext) interface{}

	// Visit a parse tree produced by ControlParser#block.
	VisitBlock(ctx *BlockContext) interface{}

	// Visit a parse tree produced by ControlParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by ControlParser#assignstmt.
	VisitAssignstmt(ctx *AssignstmtContext) interface{}

	// Visit a parse tree produced by ControlParser#printlnstmt.
	VisitPrintlnstmt(ctx *PrintlnstmtContext) interface{}

	// Visit a parse tree produced by ControlParser#printstmt.
	VisitPrintstmt(ctx *PrintstmtContext) interface{}

	// Visit a parse tree produced by ControlParser#ifstmt.
	VisitIfstmt(ctx *IfstmtContext) interface{}

	// Visit a parse tree produced by ControlParser#whilestmt.
	VisitWhilestmt(ctx *WhilestmtContext) interface{}

	// Visit a parse tree produced by ControlParser#BoolExpr.
	VisitBoolExpr(ctx *BoolExprContext) interface{}

	// Visit a parse tree produced by ControlParser#IdExpr.
	VisitIdExpr(ctx *IdExprContext) interface{}

	// Visit a parse tree produced by ControlParser#ParExpr.
	VisitParExpr(ctx *ParExprContext) interface{}

	// Visit a parse tree produced by ControlParser#StrExpr.
	VisitStrExpr(ctx *StrExprContext) interface{}

	// Visit a parse tree produced by ControlParser#NotExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by ControlParser#IntExpr.
	VisitIntExpr(ctx *IntExprContext) interface{}

	// Visit a parse tree produced by ControlParser#OpExpr.
	VisitOpExpr(ctx *OpExprContext) interface{}
}
