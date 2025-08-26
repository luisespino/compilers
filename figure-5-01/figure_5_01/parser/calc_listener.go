// Code generated from Calc.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CalcListener is a complete listener for a parse tree produced by CalcParser.
type CalcListener interface {
	antlr.ParseTreeListener

	// EnterL is called when entering the l production.
	EnterL(c *LContext)

	// EnterPassT is called when entering the PassT production.
	EnterPassT(c *PassTContext)

	// EnterSum is called when entering the Sum production.
	EnterSum(c *SumContext)

	// EnterPassF is called when entering the PassF production.
	EnterPassF(c *PassFContext)

	// EnterMul is called when entering the Mul production.
	EnterMul(c *MulContext)

	// EnterPassE is called when entering the PassE production.
	EnterPassE(c *PassEContext)

	// EnterDigit is called when entering the Digit production.
	EnterDigit(c *DigitContext)

	// ExitL is called when exiting the l production.
	ExitL(c *LContext)

	// ExitPassT is called when exiting the PassT production.
	ExitPassT(c *PassTContext)

	// ExitSum is called when exiting the Sum production.
	ExitSum(c *SumContext)

	// ExitPassF is called when exiting the PassF production.
	ExitPassF(c *PassFContext)

	// ExitMul is called when exiting the Mul production.
	ExitMul(c *MulContext)

	// ExitPassE is called when exiting the PassE production.
	ExitPassE(c *PassEContext)

	// ExitDigit is called when exiting the Digit production.
	ExitDigit(c *DigitContext)
}
