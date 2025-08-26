// Code generated from Calc.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CalcListener is a complete listener for a parse tree produced by CalcParser.
type CalcListener interface {
	antlr.ParseTreeListener

	// EnterL is called when entering the l production.
	EnterL(c *LContext)

	// EnterET is called when entering the ET production.
	EnterET(c *ETContext)

	// EnterSum is called when entering the Sum production.
	EnterSum(c *SumContext)

	// EnterEpsSum is called when entering the EpsSum production.
	EnterEpsSum(c *EpsSumContext)

	// EnterTF is called when entering the TF production.
	EnterTF(c *TFContext)

	// EnterMul is called when entering the Mul production.
	EnterMul(c *MulContext)

	// EnterEpsMul is called when entering the EpsMul production.
	EnterEpsMul(c *EpsMulContext)

	// EnterBrace is called when entering the Brace production.
	EnterBrace(c *BraceContext)

	// EnterDigit is called when entering the Digit production.
	EnterDigit(c *DigitContext)

	// ExitL is called when exiting the l production.
	ExitL(c *LContext)

	// ExitET is called when exiting the ET production.
	ExitET(c *ETContext)

	// ExitSum is called when exiting the Sum production.
	ExitSum(c *SumContext)

	// ExitEpsSum is called when exiting the EpsSum production.
	ExitEpsSum(c *EpsSumContext)

	// ExitTF is called when exiting the TF production.
	ExitTF(c *TFContext)

	// ExitMul is called when exiting the Mul production.
	ExitMul(c *MulContext)

	// ExitEpsMul is called when exiting the EpsMul production.
	ExitEpsMul(c *EpsMulContext)

	// ExitBrace is called when exiting the Brace production.
	ExitBrace(c *BraceContext)

	// ExitDigit is called when exiting the Digit production.
	ExitDigit(c *DigitContext)
}
