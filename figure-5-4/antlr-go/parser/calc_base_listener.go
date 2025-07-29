// Code generated from Calc.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Calc

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCalcListener is a complete listener for a parse tree produced by CalcParser.
type BaseCalcListener struct{}

var _ CalcListener = &BaseCalcListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCalcListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCalcListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCalcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCalcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterL is called when production l is entered.
func (s *BaseCalcListener) EnterL(ctx *LContext) {}

// ExitL is called when production l is exited.
func (s *BaseCalcListener) ExitL(ctx *LContext) {}

// EnterET is called when production ET is entered.
func (s *BaseCalcListener) EnterET(ctx *ETContext) {}

// ExitET is called when production ET is exited.
func (s *BaseCalcListener) ExitET(ctx *ETContext) {}

// EnterSum is called when production Sum is entered.
func (s *BaseCalcListener) EnterSum(ctx *SumContext) {}

// ExitSum is called when production Sum is exited.
func (s *BaseCalcListener) ExitSum(ctx *SumContext) {}

// EnterEpsSum is called when production EpsSum is entered.
func (s *BaseCalcListener) EnterEpsSum(ctx *EpsSumContext) {}

// ExitEpsSum is called when production EpsSum is exited.
func (s *BaseCalcListener) ExitEpsSum(ctx *EpsSumContext) {}

// EnterTF is called when production TF is entered.
func (s *BaseCalcListener) EnterTF(ctx *TFContext) {}

// ExitTF is called when production TF is exited.
func (s *BaseCalcListener) ExitTF(ctx *TFContext) {}

// EnterMul is called when production Mul is entered.
func (s *BaseCalcListener) EnterMul(ctx *MulContext) {}

// ExitMul is called when production Mul is exited.
func (s *BaseCalcListener) ExitMul(ctx *MulContext) {}

// EnterEpsMul is called when production EpsMul is entered.
func (s *BaseCalcListener) EnterEpsMul(ctx *EpsMulContext) {}

// ExitEpsMul is called when production EpsMul is exited.
func (s *BaseCalcListener) ExitEpsMul(ctx *EpsMulContext) {}

// EnterBrace is called when production Brace is entered.
func (s *BaseCalcListener) EnterBrace(ctx *BraceContext) {}

// ExitBrace is called when production Brace is exited.
func (s *BaseCalcListener) ExitBrace(ctx *BraceContext) {}

// EnterDigit is called when production Digit is entered.
func (s *BaseCalcListener) EnterDigit(ctx *DigitContext) {}

// ExitDigit is called when production Digit is exited.
func (s *BaseCalcListener) ExitDigit(ctx *DigitContext) {}
