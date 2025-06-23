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
		"", "'int'", "'='", "';'", "'['", "']'", "','",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "WS", "NAME", "DIGIT",
	}
	staticData.RuleNames = []string{
		"prog", "stmt", "decl", "assg", "var", "expr", "val", "index", "list",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 9, 83, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7, 4,
		2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 5, 0, 20, 8, 0, 10,
		0, 12, 0, 23, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 3, 1, 29, 8, 1, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 3, 2, 43, 8,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 3, 5, 59, 8, 5, 1, 6, 1, 6, 3, 6, 63, 8, 6, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 74, 8, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 3, 8, 81, 8, 8, 1, 8, 0, 0, 9, 0, 2, 4, 6, 8, 10, 12, 14,
		16, 0, 0, 80, 0, 21, 1, 0, 0, 0, 2, 28, 1, 0, 0, 0, 4, 42, 1, 0, 0, 0,
		6, 44, 1, 0, 0, 0, 8, 50, 1, 0, 0, 0, 10, 58, 1, 0, 0, 0, 12, 62, 1, 0,
		0, 0, 14, 73, 1, 0, 0, 0, 16, 80, 1, 0, 0, 0, 18, 20, 3, 2, 1, 0, 19, 18,
		1, 0, 0, 0, 20, 23, 1, 0, 0, 0, 21, 19, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0,
		22, 24, 1, 0, 0, 0, 23, 21, 1, 0, 0, 0, 24, 25, 5, 0, 0, 1, 25, 1, 1, 0,
		0, 0, 26, 29, 3, 4, 2, 0, 27, 29, 3, 6, 3, 0, 28, 26, 1, 0, 0, 0, 28, 27,
		1, 0, 0, 0, 29, 3, 1, 0, 0, 0, 30, 31, 5, 1, 0, 0, 31, 32, 3, 8, 4, 0,
		32, 33, 5, 2, 0, 0, 33, 34, 3, 10, 5, 0, 34, 35, 5, 3, 0, 0, 35, 43, 1,
		0, 0, 0, 36, 37, 5, 1, 0, 0, 37, 38, 3, 8, 4, 0, 38, 39, 5, 2, 0, 0, 39,
		40, 5, 9, 0, 0, 40, 41, 5, 3, 0, 0, 41, 43, 1, 0, 0, 0, 42, 30, 1, 0, 0,
		0, 42, 36, 1, 0, 0, 0, 43, 5, 1, 0, 0, 0, 44, 45, 3, 8, 4, 0, 45, 46, 5,
		2, 0, 0, 46, 47, 3, 8, 4, 0, 47, 48, 3, 14, 7, 0, 48, 49, 5, 3, 0, 0, 49,
		7, 1, 0, 0, 0, 50, 51, 5, 8, 0, 0, 51, 9, 1, 0, 0, 0, 52, 53, 5, 4, 0,
		0, 53, 59, 5, 5, 0, 0, 54, 55, 5, 4, 0, 0, 55, 56, 3, 16, 8, 0, 56, 57,
		5, 5, 0, 0, 57, 59, 1, 0, 0, 0, 58, 52, 1, 0, 0, 0, 58, 54, 1, 0, 0, 0,
		59, 11, 1, 0, 0, 0, 60, 63, 5, 9, 0, 0, 61, 63, 3, 10, 5, 0, 62, 60, 1,
		0, 0, 0, 62, 61, 1, 0, 0, 0, 63, 13, 1, 0, 0, 0, 64, 65, 5, 4, 0, 0, 65,
		66, 3, 12, 6, 0, 66, 67, 5, 5, 0, 0, 67, 68, 3, 14, 7, 0, 68, 74, 1, 0,
		0, 0, 69, 70, 5, 4, 0, 0, 70, 71, 3, 12, 6, 0, 71, 72, 5, 5, 0, 0, 72,
		74, 1, 0, 0, 0, 73, 64, 1, 0, 0, 0, 73, 69, 1, 0, 0, 0, 74, 15, 1, 0, 0,
		0, 75, 76, 3, 12, 6, 0, 76, 77, 5, 6, 0, 0, 77, 78, 3, 16, 8, 0, 78, 81,
		1, 0, 0, 0, 79, 81, 3, 12, 6, 0, 80, 75, 1, 0, 0, 0, 80, 79, 1, 0, 0, 0,
		81, 17, 1, 0, 0, 0, 7, 21, 28, 42, 58, 62, 73, 80,
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
	GrammarParserEOF   = antlr.TokenEOF
	GrammarParserT__0  = 1
	GrammarParserT__1  = 2
	GrammarParserT__2  = 3
	GrammarParserT__3  = 4
	GrammarParserT__4  = 5
	GrammarParserT__5  = 6
	GrammarParserWS    = 7
	GrammarParserNAME  = 8
	GrammarParserDIGIT = 9
)

// GrammarParser rules.
const (
	GrammarParserRULE_prog  = 0
	GrammarParserRULE_stmt  = 1
	GrammarParserRULE_decl  = 2
	GrammarParserRULE_assg  = 3
	GrammarParserRULE_var   = 4
	GrammarParserRULE_expr  = 5
	GrammarParserRULE_val   = 6
	GrammarParserRULE_index = 7
	GrammarParserRULE_list  = 8
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) CopyAll(ctx *ProgContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ProgramContext struct {
	ProgContext
}

func NewProgramContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ProgramContext {
	var p = new(ProgramContext)

	InitEmptyProgContext(&p.ProgContext)
	p.parser = parser
	p.CopyAll(ctx.(*ProgContext))

	return p
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(GrammarParserEOF, 0)
}

func (s *ProgramContext) AllStmt() []IStmtContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStmtContext); ok {
			len++
		}
	}

	tst := make([]IStmtContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStmtContext); ok {
			tst[i] = t.(IStmtContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Stmt(i int) IStmtContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStmtContext); ok {
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

	return t.(IStmtContext)
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

func (p *GrammarParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GrammarParserRULE_prog)
	var _la int

	localctx = NewProgramContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == GrammarParserT__0 || _la == GrammarParserNAME {
		{
			p.SetState(18)
			p.Stmt()
		}

		p.SetState(23)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(24)
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

// IStmtContext is an interface to support dynamic dispatch.
type IStmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStmtContext differentiates from other interfaces.
	IsStmtContext()
}

type StmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStmtContext() *StmtContext {
	var p = new(StmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) CopyAll(ctx *StmtContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StmtAssigmentContext struct {
	StmtContext
}

func NewStmtAssigmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtAssigmentContext {
	var p = new(StmtAssigmentContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *StmtAssigmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtAssigmentContext) Assg() IAssgContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssgContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssgContext)
}

func (s *StmtAssigmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterStmtAssigment(s)
	}
}

func (s *StmtAssigmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitStmtAssigment(s)
	}
}

func (s *StmtAssigmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitStmtAssigment(s)

	default:
		return t.VisitChildren(s)
	}
}

type StmtDeclarationContext struct {
	StmtContext
}

func NewStmtDeclarationContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StmtDeclarationContext {
	var p = new(StmtDeclarationContext)

	InitEmptyStmtContext(&p.StmtContext)
	p.parser = parser
	p.CopyAll(ctx.(*StmtContext))

	return p
}

func (s *StmtDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtDeclarationContext) Decl() IDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclContext)
}

func (s *StmtDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterStmtDeclaration(s)
	}
}

func (s *StmtDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitStmtDeclaration(s)
	}
}

func (s *StmtDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitStmtDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GrammarParserRULE_stmt)
	p.SetState(28)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserT__0:
		localctx = NewStmtDeclarationContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(26)
			p.Decl()
		}

	case GrammarParserNAME:
		localctx = NewStmtAssigmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(27)
			p.Assg()
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

// IDeclContext is an interface to support dynamic dispatch.
type IDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsDeclContext differentiates from other interfaces.
	IsDeclContext()
}

type DeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclContext() *DeclContext {
	var p = new(DeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_decl
	return p
}

func InitEmptyDeclContext(p *DeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_decl
}

func (*DeclContext) IsDeclContext() {}

func NewDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclContext {
	var p = new(DeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_decl

	return p
}

func (s *DeclContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclContext) CopyAll(ctx *DeclContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *DeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DeclValContext struct {
	DeclContext
	type_ antlr.Token
}

func NewDeclValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclValContext {
	var p = new(DeclValContext)

	InitEmptyDeclContext(&p.DeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclContext))

	return p
}

func (s *DeclValContext) GetType_() antlr.Token { return s.type_ }

func (s *DeclValContext) SetType_(v antlr.Token) { s.type_ = v }

func (s *DeclValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclValContext) Var_() IVarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarContext)
}

func (s *DeclValContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(GrammarParserDIGIT, 0)
}

func (s *DeclValContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterDeclVal(s)
	}
}

func (s *DeclValContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitDeclVal(s)
	}
}

func (s *DeclValContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitDeclVal(s)

	default:
		return t.VisitChildren(s)
	}
}

type DeclExprContext struct {
	DeclContext
	type_ antlr.Token
}

func NewDeclExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DeclExprContext {
	var p = new(DeclExprContext)

	InitEmptyDeclContext(&p.DeclContext)
	p.parser = parser
	p.CopyAll(ctx.(*DeclContext))

	return p
}

func (s *DeclExprContext) GetType_() antlr.Token { return s.type_ }

func (s *DeclExprContext) SetType_(v antlr.Token) { s.type_ = v }

func (s *DeclExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclExprContext) Var_() IVarContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarContext)
}

func (s *DeclExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *DeclExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterDeclExpr(s)
	}
}

func (s *DeclExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitDeclExpr(s)
	}
}

func (s *DeclExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitDeclExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Decl() (localctx IDeclContext) {
	localctx = NewDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GrammarParserRULE_decl)
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDeclExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(30)

			var _m = p.Match(GrammarParserT__0)

			localctx.(*DeclExprContext).type_ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(31)
			p.Var_()
		}
		{
			p.SetState(32)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(33)
			p.Expr()
		}
		{
			p.SetState(34)
			p.Match(GrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewDeclValContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(36)

			var _m = p.Match(GrammarParserT__0)

			localctx.(*DeclValContext).type_ = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(37)
			p.Var_()
		}
		{
			p.SetState(38)
			p.Match(GrammarParserT__1)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.Match(GrammarParserDIGIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(40)
			p.Match(GrammarParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// IAssgContext is an interface to support dynamic dispatch.
type IAssgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsAssgContext differentiates from other interfaces.
	IsAssgContext()
}

type AssgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssgContext() *AssgContext {
	var p = new(AssgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_assg
	return p
}

func InitEmptyAssgContext(p *AssgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_assg
}

func (*AssgContext) IsAssgContext() {}

func NewAssgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssgContext {
	var p = new(AssgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_assg

	return p
}

func (s *AssgContext) GetParser() antlr.Parser { return s.parser }

func (s *AssgContext) CopyAll(ctx *AssgContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *AssgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AssgArrayContext struct {
	AssgContext
}

func NewAssgArrayContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssgArrayContext {
	var p = new(AssgArrayContext)

	InitEmptyAssgContext(&p.AssgContext)
	p.parser = parser
	p.CopyAll(ctx.(*AssgContext))

	return p
}

func (s *AssgArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssgArrayContext) AllVar_() []IVarContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarContext); ok {
			len++
		}
	}

	tst := make([]IVarContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarContext); ok {
			tst[i] = t.(IVarContext)
			i++
		}
	}

	return tst
}

func (s *AssgArrayContext) Var_(i int) IVarContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarContext); ok {
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

	return t.(IVarContext)
}

func (s *AssgArrayContext) Index() IIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *AssgArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterAssgArray(s)
	}
}

func (s *AssgArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitAssgArray(s)
	}
}

func (s *AssgArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitAssgArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Assg() (localctx IAssgContext) {
	localctx = NewAssgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GrammarParserRULE_assg)
	localctx = NewAssgArrayContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Var_()
	}
	{
		p.SetState(45)
		p.Match(GrammarParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(46)
		p.Var_()
	}
	{
		p.SetState(47)
		p.Index()
	}
	{
		p.SetState(48)
		p.Match(GrammarParserT__2)
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

// IVarContext is an interface to support dynamic dispatch.
type IVarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsVarContext differentiates from other interfaces.
	IsVarContext()
}

type VarContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarContext() *VarContext {
	var p = new(VarContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_var
	return p
}

func InitEmptyVarContext(p *VarContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_var
}

func (*VarContext) IsVarContext() {}

func NewVarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarContext {
	var p = new(VarContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_var

	return p
}

func (s *VarContext) GetParser() antlr.Parser { return s.parser }

func (s *VarContext) CopyAll(ctx *VarContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *VarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VarNameContext struct {
	VarContext
}

func NewVarNameContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarNameContext {
	var p = new(VarNameContext)

	InitEmptyVarContext(&p.VarContext)
	p.parser = parser
	p.CopyAll(ctx.(*VarContext))

	return p
}

func (s *VarNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarNameContext) NAME() antlr.TerminalNode {
	return s.GetToken(GrammarParserNAME, 0)
}

func (s *VarNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterVarName(s)
	}
}

func (s *VarNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitVarName(s)
	}
}

func (s *VarNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitVarName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Var_() (localctx IVarContext) {
	localctx = NewVarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GrammarParserRULE_var)
	localctx = NewVarNameContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Match(GrammarParserNAME)
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

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExprValuesContext struct {
	ExprContext
}

func NewExprValuesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprValuesContext {
	var p = new(ExprValuesContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprValuesContext) List() IListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *ExprValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterExprValues(s)
	}
}

func (s *ExprValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitExprValues(s)
	}
}

func (s *ExprValuesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitExprValues(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprEmptyContext struct {
	ExprContext
}

func NewExprEmptyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprEmptyContext {
	var p = new(ExprEmptyContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ExprEmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprEmptyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterExprEmpty(s)
	}
}

func (s *ExprEmptyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitExprEmpty(s)
	}
}

func (s *ExprEmptyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitExprEmpty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GrammarParserRULE_expr)
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		localctx = NewExprEmptyContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(52)
			p.Match(GrammarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(53)
			p.Match(GrammarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		localctx = NewExprValuesContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)
			p.Match(GrammarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(55)
			p.List()
		}
		{
			p.SetState(56)
			p.Match(GrammarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// IValContext is an interface to support dynamic dispatch.
type IValContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsValContext differentiates from other interfaces.
	IsValContext()
}

type ValContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValContext() *ValContext {
	var p = new(ValContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_val
	return p
}

func InitEmptyValContext(p *ValContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_val
}

func (*ValContext) IsValContext() {}

func NewValContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValContext {
	var p = new(ValContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_val

	return p
}

func (s *ValContext) GetParser() antlr.Parser { return s.parser }

func (s *ValContext) CopyAll(ctx *ValContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ValDigitContext struct {
	ValContext
}

func NewValDigitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValDigitContext {
	var p = new(ValDigitContext)

	InitEmptyValContext(&p.ValContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValContext))

	return p
}

func (s *ValDigitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValDigitContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(GrammarParserDIGIT, 0)
}

func (s *ValDigitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterValDigit(s)
	}
}

func (s *ValDigitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitValDigit(s)
	}
}

func (s *ValDigitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitValDigit(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValExprContext struct {
	ValContext
}

func NewValExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValExprContext {
	var p = new(ValExprContext)

	InitEmptyValContext(&p.ValContext)
	p.parser = parser
	p.CopyAll(ctx.(*ValContext))

	return p
}

func (s *ValExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValExprContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ValExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterValExpr(s)
	}
}

func (s *ValExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitValExpr(s)
	}
}

func (s *ValExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitValExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Val() (localctx IValContext) {
	localctx = NewValContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GrammarParserRULE_val)
	p.SetState(62)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case GrammarParserDIGIT:
		localctx = NewValDigitContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(60)
			p.Match(GrammarParserDIGIT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case GrammarParserT__3:
		localctx = NewValExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(61)
			p.Expr()
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

// IIndexContext is an interface to support dynamic dispatch.
type IIndexContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsIndexContext differentiates from other interfaces.
	IsIndexContext()
}

type IndexContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIndexContext() *IndexContext {
	var p = new(IndexContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_index
	return p
}

func InitEmptyIndexContext(p *IndexContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_index
}

func (*IndexContext) IsIndexContext() {}

func NewIndexContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IndexContext {
	var p = new(IndexContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_index

	return p
}

func (s *IndexContext) GetParser() antlr.Parser { return s.parser }

func (s *IndexContext) CopyAll(ctx *IndexContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *IndexContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IndexOneContext struct {
	IndexContext
}

func NewIndexOneContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexOneContext {
	var p = new(IndexOneContext)

	InitEmptyIndexContext(&p.IndexContext)
	p.parser = parser
	p.CopyAll(ctx.(*IndexContext))

	return p
}

func (s *IndexOneContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexOneContext) Val() IValContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValContext)
}

func (s *IndexOneContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterIndexOne(s)
	}
}

func (s *IndexOneContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitIndexOne(s)
	}
}

func (s *IndexOneContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitIndexOne(s)

	default:
		return t.VisitChildren(s)
	}
}

type IndexManyContext struct {
	IndexContext
}

func NewIndexManyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexManyContext {
	var p = new(IndexManyContext)

	InitEmptyIndexContext(&p.IndexContext)
	p.parser = parser
	p.CopyAll(ctx.(*IndexContext))

	return p
}

func (s *IndexManyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexManyContext) Val() IValContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValContext)
}

func (s *IndexManyContext) Index() IIndexContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIndexContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIndexContext)
}

func (s *IndexManyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterIndexMany(s)
	}
}

func (s *IndexManyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitIndexMany(s)
	}
}

func (s *IndexManyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitIndexMany(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) Index() (localctx IIndexContext) {
	localctx = NewIndexContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GrammarParserRULE_index)
	p.SetState(73)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewIndexManyContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(64)
			p.Match(GrammarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(65)
			p.Val()
		}
		{
			p.SetState(66)
			p.Match(GrammarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(67)
			p.Index()
		}

	case 2:
		localctx = NewIndexOneContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(69)
			p.Match(GrammarParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(70)
			p.Val()
		}
		{
			p.SetState(71)
			p.Match(GrammarParserT__4)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
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

// IListContext is an interface to support dynamic dispatch.
type IListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsListContext differentiates from other interfaces.
	IsListContext()
}

type ListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListContext() *ListContext {
	var p = new(ListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_list
	return p
}

func InitEmptyListContext(p *ListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = GrammarParserRULE_list
}

func (*ListContext) IsListContext() {}

func NewListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListContext {
	var p = new(ListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = GrammarParserRULE_list

	return p
}

func (s *ListContext) GetParser() antlr.Parser { return s.parser }

func (s *ListContext) CopyAll(ctx *ListContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ValuesContext struct {
	ListContext
}

func NewValuesContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValuesContext {
	var p = new(ValuesContext)

	InitEmptyListContext(&p.ListContext)
	p.parser = parser
	p.CopyAll(ctx.(*ListContext))

	return p
}

func (s *ValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValuesContext) Val() IValContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValContext)
}

func (s *ValuesContext) List() IListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IListContext)
}

func (s *ValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterValues(s)
	}
}

func (s *ValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitValues(s)
	}
}

func (s *ValuesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitValues(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValueContext struct {
	ListContext
}

func NewValueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueContext {
	var p = new(ValueContext)

	InitEmptyListContext(&p.ListContext)
	p.parser = parser
	p.CopyAll(ctx.(*ListContext))

	return p
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) Val() IValContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValContext)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GrammarListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GrammarVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GrammarParser) List() (localctx IListContext) {
	localctx = NewListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GrammarParserRULE_list)
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		localctx = NewValuesContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(75)
			p.Val()
		}
		{
			p.SetState(76)
			p.Match(GrammarParserT__5)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(77)
			p.List()
		}

	case 2:
		localctx = NewValueContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
			p.Val()
		}

	case antlr.ATNInvalidAltNumber:
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
