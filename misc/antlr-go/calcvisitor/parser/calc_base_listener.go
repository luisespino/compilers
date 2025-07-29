// Code generated from Calc.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Calc

import "github.com/antlr4-go/antlr/v4"

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

// EnterOp is called when production Op is entered.
func (s *BaseCalcListener) EnterOp(ctx *OpContext) {}

// ExitOp is called when production Op is exited.
func (s *BaseCalcListener) ExitOp(ctx *OpContext) {}

// EnterDigit is called when production Digit is entered.
func (s *BaseCalcListener) EnterDigit(ctx *DigitContext) {}

// ExitDigit is called when production Digit is exited.
func (s *BaseCalcListener) ExitDigit(ctx *DigitContext) {}

// EnterParen is called when production Paren is entered.
func (s *BaseCalcListener) EnterParen(ctx *ParenContext) {}

// ExitParen is called when production Paren is exited.
func (s *BaseCalcListener) ExitParen(ctx *ParenContext) {}
