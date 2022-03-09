// Code generated from Calc.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CalcListener is a complete listener for a parse tree produced by CalcParser.
type CalcListener interface {
	antlr.ParseTreeListener

	// EnterSumres is called when entering the Sumres production.
	EnterSumres(c *SumresContext)

	// EnterPassT is called when entering the PassT production.
	EnterPassT(c *PassTContext)

	// EnterPassF is called when entering the PassF production.
	EnterPassF(c *PassFContext)

	// EnterMuldiv is called when entering the Muldiv production.
	EnterMuldiv(c *MuldivContext)

	// EnterPassE is called when entering the PassE production.
	EnterPassE(c *PassEContext)

	// EnterId is called when entering the Id production.
	EnterId(c *IdContext)

	// EnterNum is called when entering the Num production.
	EnterNum(c *NumContext)

	// ExitSumres is called when exiting the Sumres production.
	ExitSumres(c *SumresContext)

	// ExitPassT is called when exiting the PassT production.
	ExitPassT(c *PassTContext)

	// ExitPassF is called when exiting the PassF production.
	ExitPassF(c *PassFContext)

	// ExitMuldiv is called when exiting the Muldiv production.
	ExitMuldiv(c *MuldivContext)

	// ExitPassE is called when exiting the PassE production.
	ExitPassE(c *PassEContext)

	// ExitId is called when exiting the Id production.
	ExitId(c *IdContext)

	// ExitNum is called when exiting the Num production.
	ExitNum(c *NumContext)
}
