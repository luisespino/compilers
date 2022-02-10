// Code generated from Calc.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseCalcVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCalcVisitor) VisitOp(ctx *OpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcVisitor) VisitDigit(ctx *DigitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCalcVisitor) VisitParen(ctx *ParenContext) interface{} {
	return v.VisitChildren(ctx)
}
