// Code generated from Calc.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Calc

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 8, 43, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 3, 2, 3, 2, 3, 2, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 20, 10, 3, 12, 3, 14, 3, 23, 11, 3, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 7, 4, 31, 10, 4, 12, 4, 14, 4, 34, 11,
	4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 41, 10, 5, 3, 5, 2, 4, 4, 6, 6,
	2, 4, 6, 8, 2, 2, 2, 41, 2, 10, 3, 2, 2, 2, 4, 13, 3, 2, 2, 2, 6, 24, 3,
	2, 2, 2, 8, 40, 3, 2, 2, 2, 10, 11, 5, 4, 3, 2, 11, 12, 7, 2, 2, 3, 12,
	3, 3, 2, 2, 2, 13, 14, 8, 3, 1, 2, 14, 15, 5, 6, 4, 2, 15, 21, 3, 2, 2,
	2, 16, 17, 12, 4, 2, 2, 17, 18, 7, 4, 2, 2, 18, 20, 5, 6, 4, 2, 19, 16,
	3, 2, 2, 2, 20, 23, 3, 2, 2, 2, 21, 19, 3, 2, 2, 2, 21, 22, 3, 2, 2, 2,
	22, 5, 3, 2, 2, 2, 23, 21, 3, 2, 2, 2, 24, 25, 8, 4, 1, 2, 25, 26, 5, 8,
	5, 2, 26, 32, 3, 2, 2, 2, 27, 28, 12, 4, 2, 2, 28, 29, 7, 3, 2, 2, 29,
	31, 5, 8, 5, 2, 30, 27, 3, 2, 2, 2, 31, 34, 3, 2, 2, 2, 32, 30, 3, 2, 2,
	2, 32, 33, 3, 2, 2, 2, 33, 7, 3, 2, 2, 2, 34, 32, 3, 2, 2, 2, 35, 36, 7,
	5, 2, 2, 36, 37, 5, 4, 3, 2, 37, 38, 7, 6, 2, 2, 38, 41, 3, 2, 2, 2, 39,
	41, 7, 7, 2, 2, 40, 35, 3, 2, 2, 2, 40, 39, 3, 2, 2, 2, 41, 9, 3, 2, 2,
	2, 5, 21, 32, 40,
}
var literalNames = []string{
	"", "'*'", "'+'", "'('", "')'",
}
var symbolicNames = []string{
	"", "MUL", "ADD", "LB", "RB", "DIGIT", "WS",
}

var ruleNames = []string{
	"l", "e", "t", "f",
}

type CalcParser struct {
	*antlr.BaseParser
}

// NewCalcParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *CalcParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewCalcParser(input antlr.TokenStream) *CalcParser {
	this := new(CalcParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Calc.g4"

	return this
}

// CalcParser tokens.
const (
	CalcParserEOF   = antlr.TokenEOF
	CalcParserMUL   = 1
	CalcParserADD   = 2
	CalcParserLB    = 3
	CalcParserRB    = 4
	CalcParserDIGIT = 5
	CalcParserWS    = 6
)

// CalcParser rules.
const (
	CalcParserRULE_l = 0
	CalcParserRULE_e = 1
	CalcParserRULE_t = 2
	CalcParserRULE_f = 3
)

// ILContext is an interface to support dynamic dispatch.
type ILContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLContext differentiates from other interfaces.
	IsLContext()
}

type LContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLContext() *LContext {
	var p = new(LContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_l
	return p
}

func (*LContext) IsLContext() {}

func NewLContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LContext {
	var p = new(LContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_l

	return p
}

func (s *LContext) GetParser() antlr.Parser { return s.parser }

func (s *LContext) E() IEContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *LContext) EOF() antlr.TerminalNode {
	return s.GetToken(CalcParserEOF, 0)
}

func (s *LContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterL(s)
	}
}

func (s *LContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitL(s)
	}
}

func (p *CalcParser) L() (localctx ILContext) {
	this := p
	_ = this

	localctx = NewLContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, CalcParserRULE_l)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(8)
		p.e(0)
	}
	{
		p.SetState(9)
		p.Match(CalcParserEOF)
	}

	return localctx
}

// IEContext is an interface to support dynamic dispatch.
type IEContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEContext differentiates from other interfaces.
	IsEContext()
}

type EContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEContext() *EContext {
	var p = new(EContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_e
	return p
}

func (*EContext) IsEContext() {}

func NewEContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EContext {
	var p = new(EContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_e

	return p
}

func (s *EContext) GetParser() antlr.Parser { return s.parser }

func (s *EContext) CopyFrom(ctx *EContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *EContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PassTContext struct {
	*EContext
}

func NewPassTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PassTContext {
	var p = new(PassTContext)

	p.EContext = NewEmptyEContext()
	p.parser = parser
	p.CopyFrom(ctx.(*EContext))

	return p
}

func (s *PassTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassTContext) T() ITContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITContext)
}

func (s *PassTContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterPassT(s)
	}
}

func (s *PassTContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitPassT(s)
	}
}

type SumContext struct {
	*EContext
}

func NewSumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SumContext {
	var p = new(SumContext)

	p.EContext = NewEmptyEContext()
	p.parser = parser
	p.CopyFrom(ctx.(*EContext))

	return p
}

func (s *SumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SumContext) E() IEContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *SumContext) ADD() antlr.TerminalNode {
	return s.GetToken(CalcParserADD, 0)
}

func (s *SumContext) T() ITContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITContext)
}

func (s *SumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterSum(s)
	}
}

func (s *SumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitSum(s)
	}
}

func (p *CalcParser) E() (localctx IEContext) {
	return p.e(0)
}

func (p *CalcParser) e(_p int) (localctx IEContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewEContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IEContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, CalcParserRULE_e, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewPassTContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(12)
		p.t(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(19)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewSumContext(p, NewEContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_e)
			p.SetState(14)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(15)
				p.Match(CalcParserADD)
			}
			{
				p.SetState(16)
				p.t(0)
			}

		}
		p.SetState(21)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}

	return localctx
}

// ITContext is an interface to support dynamic dispatch.
type ITContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTContext differentiates from other interfaces.
	IsTContext()
}

type TContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTContext() *TContext {
	var p = new(TContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_t
	return p
}

func (*TContext) IsTContext() {}

func NewTContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TContext {
	var p = new(TContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_t

	return p
}

func (s *TContext) GetParser() antlr.Parser { return s.parser }

func (s *TContext) CopyFrom(ctx *TContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PassFContext struct {
	*TContext
}

func NewPassFContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PassFContext {
	var p = new(PassFContext)

	p.TContext = NewEmptyTContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TContext))

	return p
}

func (s *PassFContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassFContext) F() IFContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFContext)
}

func (s *PassFContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterPassF(s)
	}
}

func (s *PassFContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitPassF(s)
	}
}

type MulContext struct {
	*TContext
}

func NewMulContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulContext {
	var p = new(MulContext)

	p.TContext = NewEmptyTContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TContext))

	return p
}

func (s *MulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulContext) T() ITContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITContext)
}

func (s *MulContext) MUL() antlr.TerminalNode {
	return s.GetToken(CalcParserMUL, 0)
}

func (s *MulContext) F() IFContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFContext)
}

func (s *MulContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterMul(s)
	}
}

func (s *MulContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitMul(s)
	}
}

func (p *CalcParser) T() (localctx ITContext) {
	return p.t(0)
}

func (p *CalcParser) t(_p int) (localctx ITContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewTContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, CalcParserRULE_t, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	localctx = NewPassFContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(23)
		p.F()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(30)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMulContext(p, NewTContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_t)
			p.SetState(25)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(26)
				p.Match(CalcParserMUL)
			}
			{
				p.SetState(27)
				p.F()
			}

		}
		p.SetState(32)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}

	return localctx
}

// IFContext is an interface to support dynamic dispatch.
type IFContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFContext differentiates from other interfaces.
	IsFContext()
}

type FContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFContext() *FContext {
	var p = new(FContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_f
	return p
}

func (*FContext) IsFContext() {}

func NewFContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FContext {
	var p = new(FContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_f

	return p
}

func (s *FContext) GetParser() antlr.Parser { return s.parser }

func (s *FContext) CopyFrom(ctx *FContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PassEContext struct {
	*FContext
}

func NewPassEContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PassEContext {
	var p = new(PassEContext)

	p.FContext = NewEmptyFContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FContext))

	return p
}

func (s *PassEContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassEContext) LB() antlr.TerminalNode {
	return s.GetToken(CalcParserLB, 0)
}

func (s *PassEContext) E() IEContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *PassEContext) RB() antlr.TerminalNode {
	return s.GetToken(CalcParserRB, 0)
}

func (s *PassEContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterPassE(s)
	}
}

func (s *PassEContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitPassE(s)
	}
}

type DigitContext struct {
	*FContext
}

func NewDigitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DigitContext {
	var p = new(DigitContext)

	p.FContext = NewEmptyFContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FContext))

	return p
}

func (s *DigitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DigitContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(CalcParserDIGIT, 0)
}

func (s *DigitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterDigit(s)
	}
}

func (s *DigitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitDigit(s)
	}
}

func (p *CalcParser) F() (localctx IFContext) {
	this := p
	_ = this

	localctx = NewFContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CalcParserRULE_f)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(38)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalcParserLB:
		localctx = NewPassEContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(33)
			p.Match(CalcParserLB)
		}
		{
			p.SetState(34)
			p.e(0)
		}
		{
			p.SetState(35)
			p.Match(CalcParserRB)
		}

	case CalcParserDIGIT:
		localctx = NewDigitContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(37)
			p.Match(CalcParserDIGIT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *CalcParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *EContext = nil
		if localctx != nil {
			t = localctx.(*EContext)
		}
		return p.E_Sempred(t, predIndex)

	case 2:
		var t *TContext = nil
		if localctx != nil {
			t = localctx.(*TContext)
		}
		return p.T_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CalcParser) E_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CalcParser) T_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
