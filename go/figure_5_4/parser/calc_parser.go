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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 8, 45, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 3,
	2, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 26,
	10, 4, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 36, 10, 6,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 43, 10, 7, 3, 7, 2, 2, 8, 2, 4, 6,
	8, 10, 12, 2, 2, 2, 41, 2, 14, 3, 2, 2, 2, 4, 17, 3, 2, 2, 2, 6, 25, 3,
	2, 2, 2, 8, 27, 3, 2, 2, 2, 10, 35, 3, 2, 2, 2, 12, 42, 3, 2, 2, 2, 14,
	15, 5, 4, 3, 2, 15, 16, 7, 2, 2, 3, 16, 3, 3, 2, 2, 2, 17, 18, 5, 8, 5,
	2, 18, 19, 5, 6, 4, 2, 19, 5, 3, 2, 2, 2, 20, 21, 7, 3, 2, 2, 21, 22, 5,
	8, 5, 2, 22, 23, 5, 6, 4, 2, 23, 26, 3, 2, 2, 2, 24, 26, 3, 2, 2, 2, 25,
	20, 3, 2, 2, 2, 25, 24, 3, 2, 2, 2, 26, 7, 3, 2, 2, 2, 27, 28, 5, 12, 7,
	2, 28, 29, 5, 10, 6, 2, 29, 9, 3, 2, 2, 2, 30, 31, 7, 4, 2, 2, 31, 32,
	5, 12, 7, 2, 32, 33, 5, 10, 6, 2, 33, 36, 3, 2, 2, 2, 34, 36, 3, 2, 2,
	2, 35, 30, 3, 2, 2, 2, 35, 34, 3, 2, 2, 2, 36, 11, 3, 2, 2, 2, 37, 38,
	7, 5, 2, 2, 38, 39, 5, 4, 3, 2, 39, 40, 7, 6, 2, 2, 40, 43, 3, 2, 2, 2,
	41, 43, 7, 7, 2, 2, 42, 37, 3, 2, 2, 2, 42, 41, 3, 2, 2, 2, 43, 13, 3,
	2, 2, 2, 5, 25, 35, 42,
}
var literalNames = []string{
	"", "'+'", "'*'", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "", "", "", "DIGIT", "WS",
}

var ruleNames = []string{
	"l", "e", "ep", "t", "tp", "f",
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
	CalcParserT__0  = 1
	CalcParserT__1  = 2
	CalcParserT__2  = 3
	CalcParserT__3  = 4
	CalcParserDIGIT = 5
	CalcParserWS    = 6
)

// CalcParser rules.
const (
	CalcParserRULE_l  = 0
	CalcParserRULE_e  = 1
	CalcParserRULE_ep = 2
	CalcParserRULE_t  = 3
	CalcParserRULE_tp = 4
	CalcParserRULE_f  = 5
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
		p.SetState(12)
		p.E()
	}
	{
		p.SetState(13)
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

type ETContext struct {
	*EContext
}

func NewETContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ETContext {
	var p = new(ETContext)

	p.EContext = NewEmptyEContext()
	p.parser = parser
	p.CopyFrom(ctx.(*EContext))

	return p
}

func (s *ETContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ETContext) T() ITContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITContext)
}

func (s *ETContext) Ep() IEpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEpContext)
}

func (s *ETContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterET(s)
	}
}

func (s *ETContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitET(s)
	}
}

func (p *CalcParser) E() (localctx IEContext) {
	this := p
	_ = this

	localctx = NewEContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, CalcParserRULE_e)

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

	localctx = NewETContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(15)
		p.T()
	}
	{
		p.SetState(16)
		p.Ep()
	}

	return localctx
}

// IEpContext is an interface to support dynamic dispatch.
type IEpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEpContext differentiates from other interfaces.
	IsEpContext()
}

type EpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEpContext() *EpContext {
	var p = new(EpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_ep
	return p
}

func (*EpContext) IsEpContext() {}

func NewEpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EpContext {
	var p = new(EpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_ep

	return p
}

func (s *EpContext) GetParser() antlr.Parser { return s.parser }

func (s *EpContext) CopyFrom(ctx *EpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *EpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SumContext struct {
	*EpContext
}

func NewSumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SumContext {
	var p = new(SumContext)

	p.EpContext = NewEmptyEpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*EpContext))

	return p
}

func (s *SumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SumContext) T() ITContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITContext)
}

func (s *SumContext) Ep() IEpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEpContext)
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

type EpsSumContext struct {
	*EpContext
}

func NewEpsSumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EpsSumContext {
	var p = new(EpsSumContext)

	p.EpContext = NewEmptyEpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*EpContext))

	return p
}

func (s *EpsSumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EpsSumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterEpsSum(s)
	}
}

func (s *EpsSumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitEpsSum(s)
	}
}

func (p *CalcParser) Ep() (localctx IEpContext) {
	this := p
	_ = this

	localctx = NewEpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, CalcParserRULE_ep)

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

	p.SetState(23)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalcParserT__0:
		localctx = NewSumContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(18)
			p.Match(CalcParserT__0)
		}
		{
			p.SetState(19)
			p.T()
		}
		{
			p.SetState(20)
			p.Ep()
		}

	case CalcParserEOF, CalcParserT__3:
		localctx = NewEpsSumContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

type TFContext struct {
	*TContext
}

func NewTFContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TFContext {
	var p = new(TFContext)

	p.TContext = NewEmptyTContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TContext))

	return p
}

func (s *TFContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TFContext) F() IFContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFContext)
}

func (s *TFContext) Tp() ITpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITpContext)
}

func (s *TFContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterTF(s)
	}
}

func (s *TFContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitTF(s)
	}
}

func (p *CalcParser) T() (localctx ITContext) {
	this := p
	_ = this

	localctx = NewTContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, CalcParserRULE_t)

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

	localctx = NewTFContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.F()
	}
	{
		p.SetState(26)
		p.Tp()
	}

	return localctx
}

// ITpContext is an interface to support dynamic dispatch.
type ITpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTpContext differentiates from other interfaces.
	IsTpContext()
}

type TpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTpContext() *TpContext {
	var p = new(TpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_tp
	return p
}

func (*TpContext) IsTpContext() {}

func NewTpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TpContext {
	var p = new(TpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_tp

	return p
}

func (s *TpContext) GetParser() antlr.Parser { return s.parser }

func (s *TpContext) CopyFrom(ctx *TpContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type EpsMulContext struct {
	*TpContext
}

func NewEpsMulContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *EpsMulContext {
	var p = new(EpsMulContext)

	p.TpContext = NewEmptyTpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TpContext))

	return p
}

func (s *EpsMulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EpsMulContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterEpsMul(s)
	}
}

func (s *EpsMulContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitEpsMul(s)
	}
}

type MulContext struct {
	*TpContext
}

func NewMulContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulContext {
	var p = new(MulContext)

	p.TpContext = NewEmptyTpContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TpContext))

	return p
}

func (s *MulContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulContext) F() IFContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFContext)
}

func (s *MulContext) Tp() ITpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITpContext)
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

func (p *CalcParser) Tp() (localctx ITpContext) {
	this := p
	_ = this

	localctx = NewTpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, CalcParserRULE_tp)

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

	p.SetState(33)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalcParserT__1:
		localctx = NewMulContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(28)
			p.Match(CalcParserT__1)
		}
		{
			p.SetState(29)
			p.F()
		}
		{
			p.SetState(30)
			p.Tp()
		}

	case CalcParserEOF, CalcParserT__0, CalcParserT__3:
		localctx = NewEpsMulContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

type BraceContext struct {
	*FContext
}

func NewBraceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BraceContext {
	var p = new(BraceContext)

	p.FContext = NewEmptyFContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FContext))

	return p
}

func (s *BraceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BraceContext) E() IEContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *BraceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.EnterBrace(s)
	}
}

func (s *BraceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(CalcListener); ok {
		listenerT.ExitBrace(s)
	}
}

func (p *CalcParser) F() (localctx IFContext) {
	this := p
	_ = this

	localctx = NewFContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, CalcParserRULE_f)

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

	p.SetState(40)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalcParserT__2:
		localctx = NewBraceContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(35)
			p.Match(CalcParserT__2)
		}
		{
			p.SetState(36)
			p.E()
		}
		{
			p.SetState(37)
			p.Match(CalcParserT__3)
		}

	case CalcParserDIGIT:
		localctx = NewDigitContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(39)
			p.Match(CalcParserDIGIT)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}
