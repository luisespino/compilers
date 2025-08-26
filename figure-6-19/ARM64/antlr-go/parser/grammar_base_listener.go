// Code generated from Grammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Grammar

import "github.com/antlr4-go/antlr/v4"

// BaseGrammarListener is a complete listener for a parse tree produced by GrammarParser.
type BaseGrammarListener struct{}

var _ GrammarListener = &BaseGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production Program is entered.
func (s *BaseGrammarListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production Program is exited.
func (s *BaseGrammarListener) ExitProgram(ctx *ProgramContext) {}

// EnterAssign is called when production Assign is entered.
func (s *BaseGrammarListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production Assign is exited.
func (s *BaseGrammarListener) ExitAssign(ctx *AssignContext) {}

// EnterPassE is called when production PassE is entered.
func (s *BaseGrammarListener) EnterPassE(ctx *PassEContext) {}

// ExitPassE is called when production PassE is exited.
func (s *BaseGrammarListener) ExitPassE(ctx *PassEContext) {}

// EnterSumres is called when production Sumres is entered.
func (s *BaseGrammarListener) EnterSumres(ctx *SumresContext) {}

// ExitSumres is called when production Sumres is exited.
func (s *BaseGrammarListener) ExitSumres(ctx *SumresContext) {}

// EnterMuldiv is called when production Muldiv is entered.
func (s *BaseGrammarListener) EnterMuldiv(ctx *MuldivContext) {}

// ExitMuldiv is called when production Muldiv is exited.
func (s *BaseGrammarListener) ExitMuldiv(ctx *MuldivContext) {}

// EnterNum is called when production Num is entered.
func (s *BaseGrammarListener) EnterNum(ctx *NumContext) {}

// ExitNum is called when production Num is exited.
func (s *BaseGrammarListener) ExitNum(ctx *NumContext) {}

// EnterId is called when production Id is entered.
func (s *BaseGrammarListener) EnterId(ctx *IdContext) {}

// ExitId is called when production Id is exited.
func (s *BaseGrammarListener) ExitId(ctx *IdContext) {}
