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

// EnterSumres is called when production Sumres is entered.
func (s *BaseCalcListener) EnterSumres(ctx *SumresContext) {}

// ExitSumres is called when production Sumres is exited.
func (s *BaseCalcListener) ExitSumres(ctx *SumresContext) {}

// EnterPassT is called when production PassT is entered.
func (s *BaseCalcListener) EnterPassT(ctx *PassTContext) {}

// ExitPassT is called when production PassT is exited.
func (s *BaseCalcListener) ExitPassT(ctx *PassTContext) {}

// EnterPassF is called when production PassF is entered.
func (s *BaseCalcListener) EnterPassF(ctx *PassFContext) {}

// ExitPassF is called when production PassF is exited.
func (s *BaseCalcListener) ExitPassF(ctx *PassFContext) {}

// EnterMuldiv is called when production Muldiv is entered.
func (s *BaseCalcListener) EnterMuldiv(ctx *MuldivContext) {}

// ExitMuldiv is called when production Muldiv is exited.
func (s *BaseCalcListener) ExitMuldiv(ctx *MuldivContext) {}

// EnterPassE is called when production PassE is entered.
func (s *BaseCalcListener) EnterPassE(ctx *PassEContext) {}

// ExitPassE is called when production PassE is exited.
func (s *BaseCalcListener) ExitPassE(ctx *PassEContext) {}

// EnterId is called when production Id is entered.
func (s *BaseCalcListener) EnterId(ctx *IdContext) {}

// ExitId is called when production Id is exited.
func (s *BaseCalcListener) ExitId(ctx *IdContext) {}

// EnterNum is called when production Num is entered.
func (s *BaseCalcListener) EnterNum(ctx *NumContext) {}

// ExitNum is called when production Num is exited.
func (s *BaseCalcListener) ExitNum(ctx *NumContext) {}
