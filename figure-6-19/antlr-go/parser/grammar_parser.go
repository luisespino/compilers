// Code generated from Grammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Grammar

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

type GrammarParser struct {
	*antlr.BaseParser
}

var GrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func grammarParserInit() {
	staticData := &GrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "'='", "'*'", "'/'", "'+'", "'-'", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "ID", "NUM", "WS",
	}
	staticData.RuleNames = []string{
		"p", "s", "e",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 10, 38, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 1, 0, 4, 0, 8, 8, 0,
		11, 0, 12, 0, 9, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 25, 8, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1,
		2, 5, 2, 33, 8, 2, 10, 2, 12, 2, 36, 9, 2, 1, 2, 0, 1, 4, 3, 0, 2, 4, 0,
		2, 1, 0, 2, 3, 1, 0, 4, 5, 39, 0, 7, 1, 0, 0, 0, 2, 13, 1, 0, 0, 0, 4,
		24, 1, 0, 0, 0, 6, 8, 3, 2, 1, 0, 7, 6, 1, 0, 0, 0, 8, 9, 1, 0, 0, 0, 9,
		7, 1, 0, 0, 0, 9, 10, 1, 0, 0, 0, 10, 11, 1, 0, 0, 0, 11, 12, 5, 0, 0,
		1, 12, 1, 1, 0, 0, 0, 13, 14, 5, 8, 0, 0, 14, 15, 5, 1, 0, 0, 15, 16, 3,
		4, 2, 0, 16, 3, 1, 0, 0, 0, 17, 18, 6, 2, -1, 0, 18, 19, 5, 6, 0, 0, 19,
		20, 3, 4, 2, 0, 20, 21, 5, 7, 0, 0, 21, 25, 1, 0, 0, 0, 22, 25, 5, 8, 0,
		0, 23, 25, 5, 9, 0, 0, 24, 17, 1, 0, 0, 0, 24, 22, 1, 0, 0, 0, 24, 23,
		1, 0, 0, 0, 25, 34, 1, 0, 0, 0, 26, 27, 10, 5, 0, 0, 27, 28, 7, 0, 0, 0,
		28, 33, 3, 4, 2, 6, 29, 30, 10, 4, 0, 0, 30, 31, 7, 1, 0, 0, 31, 33, 3,
		4, 2, 5, 32, 26, 1, 0, 0, 0, 32, 29, 1, 0, 0, 0, 33, 36, 1, 0, 0, 0, 34,
		32, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 5, 1, 0, 0, 0, 36, 34, 1, 0, 0,
		0, 4, 9, 24, 32, 34,
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

// GrammarParserInit initializes any static state used to implement GrammarParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewGrammarParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func GrammarParserInit() {
	staticData := &GrammarParserStaticData
	staticData.once.Do(grammarParserInit)
}

// NewGrammarParser produces a new parser instance for the optional input antlr.TokenStream.
func NewGrammarParser(input antlr.TokenStream) *GrammarParser {
	GrammarParserInit()
	this := new(GrammarParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &GrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Grammar.g4"

	return this
}

// GrammarParser tokens.
const (
	GrammarParserEOF  = antlr.TokenEOF
	GrammarParserT__0 = 1
	GrammarParserT__1 = 2
	GrammarParserT__2 = 3
	GrammarParserT__3 = 4
	GrammarParserT__4 = 5
	GrammarParserT__5 = 6
	GrammarParserT__6 = 7
	GrammarParserID   = 8
	GrammarParserNUM  = 9
	GrammarParserWS   = 10
)

// GrammarParser rules.
const (
	GrammarParserRULE_p = 0
	GrammarParserRULE_s = 1
	GrammarParserRULE_e = 2
)

// IPContext is an interface to support dynamic dispatch.
type IPContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsPContext differentiates from other interfaces.
	IsPContext()
}

type PContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPContext() *PContext {
	var p = new(PContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_p
	return p
}

func InitEmptyPContext(p *PContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_p
}

func (*PContext) IsPContext() {}

func NewPContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PContext {
	var p = new(PContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_p

	return p
}

func (s *PContext) GetParser() antlr.Parser { return s.parser }

func (s *PContext) CopyAll(ctx *PContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *PContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ProgramContext struct {
	PContext
}

func NewProgramContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ProgramContext {
	var p = new(ProgramContext)

	InitEmptyPContext(&p.PContext)
	p.parser = parser
	p.CopyAll(ctx.(*PContext))

	return p
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(GrammarParserEOF, 0)
}

func (s *ProgramContext) AllS() []ISContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISContext); ok {
			len++
		}
	}

	tst := make([]ISContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISContext); ok {
			tst[i] = t.(ISContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) S(i int) ISContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISContext)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) P() (localctx IPContext) {
	localctx = NewPContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GrammarParserRULE_p)
	var _la int

	localctx = NewProgramContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(7)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GrammarParserID {
		{
			p.SetState(6)
			p.S()
		}

		p.SetState(9)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(11)
		p.Match(GrammarParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
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
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISContext is an interface to support dynamic dispatch.
type ISContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsSContext differentiates from other interfaces.
	IsSContext()
}

type SContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySContext() *SContext {
	var p = new(SContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_s
	return p
}

func InitEmptySContext(p *SContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_s
}

func (*SContext) IsSContext() {}

func NewSContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SContext {
	var p = new(SContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_s

	return p
}

func (s *SContext) GetParser() antlr.Parser { return s.parser }

func (s *SContext) CopyAll(ctx *SContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *SContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssignContext struct {
	SContext
}

func NewAssignContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignContext {
	var p = new(AssignContext)

	InitEmptySContext(&p.SContext)
	p.parser = parser
	p.CopyAll(ctx.(*SContext))

	return p
}

func (s *AssignContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *AssignContext) E() IEContext {
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

func (s *AssignContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterAssign(s)
	}
}

func (s *AssignContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitAssign(s)
	}
}

func (s *AssignContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitAssign(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) S() (localctx ISContext) {
	localctx = NewSContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GrammarParserRULE_s)
	localctx = NewAssignContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(13)
		p.Match(GrammarParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(14)
		p.Match(GrammarParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(15)
		p.e(0)
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
	p.RuleIndex = GrammarParserRULE_e
	return p
}

func InitEmptyEContext(p *EContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_e
}

func (*EContext) IsEContext() {}

func NewEContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EContext {
	var p = new(EContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_e

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

type PassEContext struct {
	EContext
}

func NewPassEContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PassEContext {
	var p = new(PassEContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

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
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterPassE(s)
	}
}

func (s *PassEContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitPassE(s)
	}
}

func (s *PassEContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitPassE(s)

	default:
		return t.VisitChildren(s)
	}
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

func (s *SumresContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *SumresContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *SumresContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterSumres(s)
	}
}

func (s *SumresContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitSumres(s)
	}
}

func (s *SumresContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitSumres(s)

	default:
		return t.VisitChildren(s)
	}
}

type MuldivContext struct {
	EContext
	op antlr.Token
}

func NewMuldivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MuldivContext {
	var p = new(MuldivContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *MuldivContext) GetOp() antlr.Token { return s.op }

func (s *MuldivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MuldivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MuldivContext) AllE() []IEContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEContext); ok {
			len++
		}
	}

	tst := make([]IEContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEContext); ok {
			tst[i] = t.(IEContext)
			i++
		}
	}

	return tst
}

func (s *MuldivContext) E(i int) IEContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEContext)
}

func (s *MuldivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterMuldiv(s)
	}
}

func (s *MuldivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitMuldiv(s)
	}
}

func (s *MuldivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitMuldiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumContext struct {
	EContext
}

func NewNumContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumContext {
	var p = new(NumContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *NumContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumContext) NUM() antlr.TerminalNode {
	return s.GetToken(GrammarParserNUM, 0)
}

func (s *NumContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterNum(s)
	}
}

func (s *NumContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitNum(s)
	}
}

func (s *NumContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitNum(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdContext struct {
	EContext
}

func NewIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdContext {
	var p = new(IdContext)

	InitEmptyEContext(&p.EContext)
	p.parser = parser
	p.CopyAll(ctx.(*EContext))

	return p
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ID() antlr.TerminalNode {
	return s.GetToken(GrammarParserID, 0)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitId(s)
	}
}

func (s *IdContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitId(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) E() (localctx IEContext) {
	return p.e(0)
}

func (p *GrammarParser) e(_p int) (localctx IEContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewEContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IEContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 4
	p.EnterRecursionRule(localctx, 4, GrammarParserRULE_e, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserT__5:
		localctx = NewPassEContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(18)
			p.Match(GrammarParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(19)
			p.e(0)
		}
		{
			p.SetState(20)
			p.Match(GrammarParserT__6)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GrammarParserID:
		localctx = NewIdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(22)
			p.Match(GrammarParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GrammarParserNUM:
		localctx = NewNumContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(23)
			p.Match(GrammarParserNUM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(34)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(32)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMuldivContext(p, NewEContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_e)
				p.SetState(26)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(27)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MuldivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserT__1 || _la == GrammarParserT__2) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MuldivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(28)
					p.e(6)
				}

			case 2:
				localctx = NewSumresContext(p, NewEContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, GrammarParserRULE_e)
				p.SetState(29)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(30)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*SumresContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == GrammarParserT__3 || _la == GrammarParserT__4) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*SumresContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(31)
					p.e(5)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(36)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
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

func (p *GrammarParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *EContext = nil
		if localctx != nil {
			t = localctx.(*EContext)
		}
		return p.E_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GrammarParser) E_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
