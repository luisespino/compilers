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

// EnterPassT is called when production PassT is entered.
func (s *BaseCalcListener) EnterPassT(ctx *PassTContext) {}

// ExitPassT is called when production PassT is exited.
func (s *BaseCalcListener) ExitPassT(ctx *PassTContext) {}

// EnterSum is called when production Sum is entered.
func (s *BaseCalcListener) EnterSum(ctx *SumContext) {}

// ExitSum is called when production Sum is exited.
func (s *BaseCalcListener) ExitSum(ctx *SumContext) {}

// EnterPassF is called when production PassF is entered.
func (s *BaseCalcListener) EnterPassF(ctx *PassFContext) {}

// ExitPassF is called when production PassF is exited.
func (s *BaseCalcListener) ExitPassF(ctx *PassFContext) {}

// EnterMul is called when production Mul is entered.
func (s *BaseCalcListener) EnterMul(ctx *MulContext) {}

// ExitMul is called when production Mul is exited.
func (s *BaseCalcListener) ExitMul(ctx *MulContext) {}

// EnterPassE is called when production PassE is entered.
func (s *BaseCalcListener) EnterPassE(ctx *PassEContext) {}

// ExitPassE is called when production PassE is exited.
func (s *BaseCalcListener) ExitPassE(ctx *PassEContext) {}

// EnterDigit is called when production Digit is entered.
func (s *BaseCalcListener) EnterDigit(ctx *DigitContext) {}

// ExitDigit is called when production Digit is exited.
func (s *BaseCalcListener) ExitDigit(ctx *DigitContext) {}
