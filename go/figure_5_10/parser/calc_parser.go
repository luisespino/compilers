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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 11, 39, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7,
	2, 15, 10, 2, 12, 2, 14, 2, 18, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 7, 3, 26, 10, 3, 12, 3, 14, 3, 29, 11, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 5, 4, 37, 10, 4, 3, 4, 2, 4, 2, 4, 5, 2, 4, 6, 2, 4, 3, 2, 3,
	4, 3, 2, 5, 6, 2, 39, 2, 8, 3, 2, 2, 2, 4, 19, 3, 2, 2, 2, 6, 36, 3, 2,
	2, 2, 8, 9, 8, 2, 1, 2, 9, 10, 5, 4, 3, 2, 10, 16, 3, 2, 2, 2, 11, 12,
	12, 4, 2, 2, 12, 13, 9, 2, 2, 2, 13, 15, 5, 4, 3, 2, 14, 11, 3, 2, 2, 2,
	15, 18, 3, 2, 2, 2, 16, 14, 3, 2, 2, 2, 16, 17, 3, 2, 2, 2, 17, 3, 3, 2,
	2, 2, 18, 16, 3, 2, 2, 2, 19, 20, 8, 3, 1, 2, 20, 21, 5, 6, 4, 2, 21, 27,
	3, 2, 2, 2, 22, 23, 12, 4, 2, 2, 23, 24, 9, 3, 2, 2, 24, 26, 5, 6, 4, 2,
	25, 22, 3, 2, 2, 2, 26, 29, 3, 2, 2, 2, 27, 25, 3, 2, 2, 2, 27, 28, 3,
	2, 2, 2, 28, 5, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 30, 31, 7, 7, 2, 2, 31,
	32, 5, 2, 2, 2, 32, 33, 7, 8, 2, 2, 33, 37, 3, 2, 2, 2, 34, 37, 7, 9, 2,
	2, 35, 37, 7, 10, 2, 2, 36, 30, 3, 2, 2, 2, 36, 34, 3, 2, 2, 2, 36, 35,
	3, 2, 2, 2, 37, 7, 3, 2, 2, 2, 5, 16, 27, 36,
}
var literalNames = []string{
	"", "'+'", "'-'", "'*'", "'/'", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "ID", "NUM", "WS",
}

var ruleNames = []string{
	"e", "t", "f",
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
	CalcParserEOF  = antlr.TokenEOF
	CalcParserT__0 = 1
	CalcParserT__1 = 2
	CalcParserT__2 = 3
	CalcParserT__3 = 4
	CalcParserT__4 = 5
	CalcParserT__5 = 6
	CalcParserID   = 7
	CalcParserNUM  = 8
	CalcParserWS   = 9
)

// CalcParser rules.
const (
	CalcParserRULE_e = 0
	CalcParserRULE_t = 1
	CalcParserRULE_f = 2
)

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

type SumresContext struct {
	*EContext
	op antlr.Token
}

func NewSumresContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SumresContext {
	var p = new(SumresContext)

	p.EContext = NewEmptyEContext()
	p.parser = parser
	p.CopyFrom(ctx.(*EContext))

	return p
}

func (s *SumresContext) GetOp() antlr.Token { return s.op }

func (s *SumresContext) SetOp(v antlr.Token) { s.op = v }

func (s *SumresContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SumresContext) E() IEContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *SumresContext) T() ITContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITContext)
}

func (s *SumresContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterSumres(s)
	}
}

func (s *SumresContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitSumres(s)
	}
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
	_startState := 0
	p.EnterRecursionRule(localctx, 0, CalcParserRULE_e, _p)
	var _la int

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
		p.SetState(7)
		p.t(0)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(14)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewSumresContext(p, NewEContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_e)
			p.SetState(9)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(10)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*SumresContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == CalcParserT__0 || _la == CalcParserT__1) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*SumresContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(11)
				p.t(0)
			}

		}
		p.SetState(16)
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

type MuldivContext struct {
	*TContext
	op antlr.Token
}

func NewMuldivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MuldivContext {
	var p = new(MuldivContext)

	p.TContext = NewEmptyTContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TContext))

	return p
}

func (s *MuldivContext) GetOp() antlr.Token { return s.op }

func (s *MuldivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MuldivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MuldivContext) T() ITContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITContext)
}

func (s *MuldivContext) F() IFContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFContext)
}

func (s *MuldivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterMuldiv(s)
	}
}

func (s *MuldivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitMuldiv(s)
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
	_startState := 2
	p.EnterRecursionRule(localctx, 2, CalcParserRULE_t, _p)
	var _la int

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
		p.SetState(18)
		p.F()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMuldivContext(p, NewTContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_t)
			p.SetState(20)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(21)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*MuldivContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == CalcParserT__2 || _la == CalcParserT__3) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*MuldivContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(22)
				p.F()
			}

		}
		p.SetState(27)
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

func (s *PassEContext) E() IEContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEContext)
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

type NumContext struct {
	*FContext
}

func NewNumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumContext {
	var p = new(NumContext)

	p.FContext = NewEmptyFContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FContext))

	return p
}

func (s *NumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumContext) NUM() antlr.TerminalNode {
	return s.GetToken(CalcParserNUM, 0)
}

func (s *NumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterNum(s)
	}
}

func (s *NumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitNum(s)
	}
}

type IdContext struct {
	*FContext
}

func NewIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdContext {
	var p = new(IdContext)

	p.FContext = NewEmptyFContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FContext))

	return p
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ID() antlr.TerminalNode {
	return s.GetToken(CalcParserID, 0)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitId(s)
	}
}

func (p *CalcParser) F() (localctx IFContext) {
	this := p
	_ = this

	localctx = NewFContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CalcParserRULE_f)

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

	p.SetState(34)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalcParserT__4:
		localctx = NewPassEContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.Match(CalcParserT__4)
		}
		{
			p.SetState(29)
			p.e(0)
		}
		{
			p.SetState(30)
			p.Match(CalcParserT__5)
		}

	case CalcParserID:
		localctx = NewIdContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(32)
			p.Match(CalcParserID)
		}

	case CalcParserNUM:
		localctx = NewNumContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(33)
			p.Match(CalcParserNUM)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *CalcParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *EContext = nil
		if localctx != nil {
			t = localctx.(*EContext)
		}
		return p.E_Sempred(t, predIndex)

	case 1:
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
