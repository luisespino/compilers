// Code generated from Grammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Grammar

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by GrammarParser.
type GrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GrammarParser#Program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by GrammarParser#Assign.
	VisitAssign(ctx *AssignContext) interface{}

	// Visit a parse tree produced by GrammarParser#PassE.
	VisitPassE(ctx *PassEContext) interface{}

	// Visit a parse tree produced by GrammarParser#Sumres.
	VisitSumres(ctx *SumresContext) interface{}

	// Visit a parse tree produced by GrammarParser#Muldiv.
	VisitMuldiv(ctx *MuldivContext) interface{}

	// Visit a parse tree produced by GrammarParser#Num.
	VisitNum(ctx *NumContext) interface{}

	// Visit a parse tree produced by GrammarParser#Id.
	VisitId(ctx *IdContext) interface{}
}
