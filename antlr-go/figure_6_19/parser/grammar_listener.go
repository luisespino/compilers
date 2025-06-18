// Code generated from Grammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Grammar

import "github.com/antlr4-go/antlr/v4"

// GrammarListener is a complete listener for a parse tree produced by GrammarParser.
type GrammarListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the Program production.
	EnterProgram(c *ProgramContext)

	// EnterAssign is called when entering the Assign production.
	EnterAssign(c *AssignContext)

	// EnterPassE is called when entering the PassE production.
	EnterPassE(c *PassEContext)

	// EnterSumres is called when entering the Sumres production.
	EnterSumres(c *SumresContext)

	// EnterMuldiv is called when entering the Muldiv production.
	EnterMuldiv(c *MuldivContext)

	// EnterNum is called when entering the Num production.
	EnterNum(c *NumContext)

	// EnterId is called when entering the Id production.
	EnterId(c *IdContext)

	// ExitProgram is called when exiting the Program production.
	ExitProgram(c *ProgramContext)

	// ExitAssign is called when exiting the Assign production.
	ExitAssign(c *AssignContext)

	// ExitPassE is called when exiting the PassE production.
	ExitPassE(c *PassEContext)

	// ExitSumres is called when exiting the Sumres production.
	ExitSumres(c *SumresContext)

	// ExitMuldiv is called when exiting the Muldiv production.
	ExitMuldiv(c *MuldivContext)

	// ExitNum is called when exiting the Num production.
	ExitNum(c *NumContext)

	// ExitId is called when exiting the Id production.
	ExitId(c *IdContext)
}
