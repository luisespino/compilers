// Code generated from Calc.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Calc

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by CalcParser.
type CalcVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CalcParser#Op.
	VisitOp(ctx *OpContext) interface{}

	// Visit a parse tree produced by CalcParser#Digit.
	VisitDigit(ctx *DigitContext) interface{}

	// Visit a parse tree produced by CalcParser#Paren.
	VisitParen(ctx *ParenContext) interface{}
}
