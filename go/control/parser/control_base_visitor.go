// Code generated from Control.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Control

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseControlVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseControlVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitBlock(ctx *BlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitAssignstmt(ctx *AssignstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitPrintlnstmt(ctx *PrintlnstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitPrintstmt(ctx *PrintstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitIfstmt(ctx *IfstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitWhilestmt(ctx *WhilestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitBoolExpr(ctx *BoolExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitIdExpr(ctx *IdExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitParExpr(ctx *ParExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitStrExpr(ctx *StrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitIntExpr(ctx *IntExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseControlVisitor) VisitOpExpr(ctx *OpExprContext) interface{} {
	return v.VisitChildren(ctx)
}
