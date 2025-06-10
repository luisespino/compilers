// Code generated from Calc.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Calc

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type CalcParser struct {
	*antlr.BaseParser
}

var CalcParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func calcParserInit() {
	staticData := &CalcParserStaticData
	staticData.LiteralNames = []string{
		"", "'+'", "'-'", "'*'", "'/'", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "ID", "NUM", "WS",
	}
	staticData.RuleNames = []string{
		"e", "t", "f",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 9, 37, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 1, 0, 1, 0, 1, 0,
		1, 0, 1, 0, 5, 0, 13, 8, 0, 10, 0, 12, 0, 16, 9, 0, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 5, 1, 24, 8, 1, 10, 1, 12, 1, 27, 9, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 3, 2, 35, 8, 2, 1, 2, 0, 2, 0, 2, 3, 0, 2, 4, 0, 2, 1,
		0, 1, 2, 1, 0, 3, 4, 37, 0, 6, 1, 0, 0, 0, 2, 17, 1, 0, 0, 0, 4, 34, 1,
		0, 0, 0, 6, 7, 6, 0, -1, 0, 7, 8, 3, 2, 1, 0, 8, 14, 1, 0, 0, 0, 9, 10,
		10, 2, 0, 0, 10, 11, 7, 0, 0, 0, 11, 13, 3, 2, 1, 0, 12, 9, 1, 0, 0, 0,
		13, 16, 1, 0, 0, 0, 14, 12, 1, 0, 0, 0, 14, 15, 1, 0, 0, 0, 15, 1, 1, 0,
		0, 0, 16, 14, 1, 0, 0, 0, 17, 18, 6, 1, -1, 0, 18, 19, 3, 4, 2, 0, 19,
		25, 1, 0, 0, 0, 20, 21, 10, 2, 0, 0, 21, 22, 7, 1, 0, 0, 22, 24, 3, 4,
		2, 0, 23, 20, 1, 0, 0, 0, 24, 27, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 25, 26,
		1, 0, 0, 0, 26, 3, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 28, 29, 5, 5, 0, 0,
		29, 30, 3, 0, 0, 0, 30, 31, 5, 6, 0, 0, 31, 35, 1, 0, 0, 0, 32, 35, 5,
		7, 0, 0, 33, 35, 5, 8, 0, 0, 34, 28, 1, 0, 0, 0, 34, 32, 1, 0, 0, 0, 34,
		33, 1, 0, 0, 0, 35, 5, 1, 0, 0, 0, 3, 14, 25, 34,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// CalcParserInit initializes any static state used to implement CalcParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewCalcParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func CalcParserInit() {
	staticData := &CalcParserStaticData
	staticData.once.Do(calcParserInit)
}

// NewCalcParser produces a new parser instance for the optional input antlr.TokenStream.
func NewCalcParser(input antlr.TokenStream) *CalcParser {
	CalcParserInit()
	this := new(CalcParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &CalcParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEContext() *EContext {
	var p = new(EContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_e
	return p
}

func InitEmptyEContext(p *EContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_e
}

func (*EContext) IsEContext() {}

func NewEContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EContext {
	var p = new(EContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_e

	return p
}

func (s *EContext) GetParser() antlr.Parser { return s.parser }

func (s *EContext) CopyAll(ctx *EContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *EContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SumresContext struct {
	EContext
	op antlr.Token
}

func NewSumresContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SumresContext {
	var p = new(SumresContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *SumresContext) GetOp() antlr.Token { return s.op }

func (s *SumresContext) SetOp(v antlr.Token) { s.op = v }

func (s *SumresContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SumresContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *SumresContext) T() ITContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	EContext
}

func NewPassTContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PassTContext {
	var p = new(PassTContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *PassTContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassTContext) T() ITContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewEContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IEContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, CalcParserRULE_e, _p)
	var _la int

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
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
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
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
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
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 0, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTContext() *TContext {
	var p = new(TContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_t
	return p
}

func InitEmptyTContext(p *TContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_t
}

func (*TContext) IsTContext() {}

func NewTContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TContext {
	var p = new(TContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_t

	return p
}

func (s *TContext) GetParser() antlr.Parser { return s.parser }

func (s *TContext) CopyAll(ctx *TContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PassFContext struct {
	TContext
}

func NewPassFContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PassFContext {
	var p = new(PassFContext)

	InitEmptyTContext(&p.TContext)
	p.parser = parser
	p.CopyAll(ctx.(*TContext))

	return p
}

func (s *PassFContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassFContext) F() IFContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	TContext
	op antlr.Token
}

func NewMuldivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MuldivContext {
	var p = new(MuldivContext)

	InitEmptyTContext(&p.TContext)
	p.parser = parser
	p.CopyAll(ctx.(*TContext))

	return p
}

func (s *MuldivContext) GetOp() antlr.Token { return s.op }

func (s *MuldivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MuldivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MuldivContext) T() ITContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITContext)
}

func (s *MuldivContext) F() IFContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewTContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ITContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, CalcParserRULE_t, _p)
	var _la int

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
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
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
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
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
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFContext() *FContext {
	var p = new(FContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_f
	return p
}

func InitEmptyFContext(p *FContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = CalcParserRULE_f
}

func (*FContext) IsFContext() {}

func NewFContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FContext {
	var p = new(FContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_f

	return p
}

func (s *FContext) GetParser() antlr.Parser { return s.parser }

func (s *FContext) CopyAll(ctx *FContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *FContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PassEContext struct {
	FContext
}

func NewPassEContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PassEContext {
	var p = new(PassEContext)

	InitEmptyFContext(&p.FContext)
	p.parser = parser
	p.CopyAll(ctx.(*FContext))

	return p
}

func (s *PassEContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassEContext) E() IEContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

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
	FContext
}

func NewNumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumContext {
	var p = new(NumContext)

	InitEmptyFContext(&p.FContext)
	p.parser = parser
	p.CopyAll(ctx.(*FContext))

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
	FContext
}

func NewIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdContext {
	var p = new(IdContext)

	InitEmptyFContext(&p.FContext)
	p.parser = parser
	p.CopyAll(ctx.(*FContext))

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
	localctx = NewFContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CalcParserRULE_f)
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case CalcParserT__4:
		localctx = NewPassEContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.Match(CalcParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(29)
			p.e(0)
		}
		{
			p.SetState(30)
			p.Match(CalcParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CalcParserID:
		localctx = NewIdContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(32)
			p.Match(CalcParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case CalcParserNUM:
		localctx = NewNumContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(33)
			p.Match(CalcParserNUM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
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
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *CalcParser) T_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
