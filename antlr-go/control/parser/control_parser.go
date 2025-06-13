// Code generated from Control.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // Control

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

type ControlParser struct {
	*antlr.BaseParser
}

var ControlParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func controlParserInit() {
	staticData := &ControlParserStaticData
	staticData.LiteralNames = []string{
		"", "'='", "'println'", "'('", "')'", "'print'", "'if'", "'{'", "'}'",
		"'while'", "'!'", "'*'", "'/'", "'+'", "'-'", "'>='", "'>'", "'<='",
		"'<'", "'=='", "'!='", "'&&'", "'||'", "'true'", "'false'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "", "INT", "ID", "STRING", "WS",
	}
	staticData.RuleNames = []string{
		"prog", "block", "stmt", "assignstmt", "printlnstmt", "printstmt", "ifstmt",
		"whilestmt", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 28, 104, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1, 0, 1,
		1, 5, 1, 23, 8, 1, 10, 1, 12, 1, 26, 9, 1, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		3, 2, 33, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1,
		6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 76, 8, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8,
		1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 5, 8, 99, 8, 8, 10, 8, 12, 8,
		102, 9, 8, 1, 8, 0, 1, 16, 9, 0, 2, 4, 6, 8, 10, 12, 14, 16, 0, 6, 1, 0,
		23, 24, 1, 0, 11, 12, 1, 0, 13, 14, 1, 0, 15, 16, 1, 0, 17, 18, 1, 0, 19,
		20, 111, 0, 18, 1, 0, 0, 0, 2, 24, 1, 0, 0, 0, 4, 32, 1, 0, 0, 0, 6, 34,
		1, 0, 0, 0, 8, 38, 1, 0, 0, 0, 10, 43, 1, 0, 0, 0, 12, 48, 1, 0, 0, 0,
		14, 56, 1, 0, 0, 0, 16, 75, 1, 0, 0, 0, 18, 19, 3, 2, 1, 0, 19, 20, 5,
		0, 0, 1, 20, 1, 1, 0, 0, 0, 21, 23, 3, 4, 2, 0, 22, 21, 1, 0, 0, 0, 23,
		26, 1, 0, 0, 0, 24, 22, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 3, 1, 0, 0,
		0, 26, 24, 1, 0, 0, 0, 27, 33, 3, 6, 3, 0, 28, 33, 3, 8, 4, 0, 29, 33,
		3, 10, 5, 0, 30, 33, 3, 12, 6, 0, 31, 33, 3, 14, 7, 0, 32, 27, 1, 0, 0,
		0, 32, 28, 1, 0, 0, 0, 32, 29, 1, 0, 0, 0, 32, 30, 1, 0, 0, 0, 32, 31,
		1, 0, 0, 0, 33, 5, 1, 0, 0, 0, 34, 35, 5, 26, 0, 0, 35, 36, 5, 1, 0, 0,
		36, 37, 3, 16, 8, 0, 37, 7, 1, 0, 0, 0, 38, 39, 5, 2, 0, 0, 39, 40, 5,
		3, 0, 0, 40, 41, 3, 16, 8, 0, 41, 42, 5, 4, 0, 0, 42, 9, 1, 0, 0, 0, 43,
		44, 5, 5, 0, 0, 44, 45, 5, 3, 0, 0, 45, 46, 3, 16, 8, 0, 46, 47, 5, 4,
		0, 0, 47, 11, 1, 0, 0, 0, 48, 49, 5, 6, 0, 0, 49, 50, 5, 3, 0, 0, 50, 51,
		3, 16, 8, 0, 51, 52, 5, 4, 0, 0, 52, 53, 5, 7, 0, 0, 53, 54, 3, 2, 1, 0,
		54, 55, 5, 8, 0, 0, 55, 13, 1, 0, 0, 0, 56, 57, 5, 9, 0, 0, 57, 58, 5,
		3, 0, 0, 58, 59, 3, 16, 8, 0, 59, 60, 5, 4, 0, 0, 60, 61, 5, 7, 0, 0, 61,
		62, 3, 2, 1, 0, 62, 63, 5, 8, 0, 0, 63, 15, 1, 0, 0, 0, 64, 65, 6, 8, -1,
		0, 65, 66, 5, 10, 0, 0, 66, 76, 3, 16, 8, 13, 67, 68, 5, 3, 0, 0, 68, 69,
		3, 16, 8, 0, 69, 70, 5, 4, 0, 0, 70, 76, 1, 0, 0, 0, 71, 76, 5, 25, 0,
		0, 72, 76, 5, 26, 0, 0, 73, 76, 5, 27, 0, 0, 74, 76, 7, 0, 0, 0, 75, 64,
		1, 0, 0, 0, 75, 67, 1, 0, 0, 0, 75, 71, 1, 0, 0, 0, 75, 72, 1, 0, 0, 0,
		75, 73, 1, 0, 0, 0, 75, 74, 1, 0, 0, 0, 76, 100, 1, 0, 0, 0, 77, 78, 10,
		12, 0, 0, 78, 79, 7, 1, 0, 0, 79, 99, 3, 16, 8, 13, 80, 81, 10, 11, 0,
		0, 81, 82, 7, 2, 0, 0, 82, 99, 3, 16, 8, 12, 83, 84, 10, 10, 0, 0, 84,
		85, 7, 3, 0, 0, 85, 99, 3, 16, 8, 11, 86, 87, 10, 9, 0, 0, 87, 88, 7, 4,
		0, 0, 88, 99, 3, 16, 8, 10, 89, 90, 10, 8, 0, 0, 90, 91, 7, 5, 0, 0, 91,
		99, 3, 16, 8, 9, 92, 93, 10, 7, 0, 0, 93, 94, 5, 21, 0, 0, 94, 99, 3, 16,
		8, 8, 95, 96, 10, 6, 0, 0, 96, 97, 5, 22, 0, 0, 97, 99, 3, 16, 8, 7, 98,
		77, 1, 0, 0, 0, 98, 80, 1, 0, 0, 0, 98, 83, 1, 0, 0, 0, 98, 86, 1, 0, 0,
		0, 98, 89, 1, 0, 0, 0, 98, 92, 1, 0, 0, 0, 98, 95, 1, 0, 0, 0, 99, 102,
		1, 0, 0, 0, 100, 98, 1, 0, 0, 0, 100, 101, 1, 0, 0, 0, 101, 17, 1, 0, 0,
		0, 102, 100, 1, 0, 0, 0, 5, 24, 32, 75, 98, 100,
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

// ControlParserInit initializes any static state used to implement ControlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewControlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ControlParserInit() {
	staticData := &ControlParserStaticData
	staticData.once.Do(controlParserInit)
}

// NewControlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewControlParser(input antlr.TokenStream) *ControlParser {
	ControlParserInit()
	this := new(ControlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ControlParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Control.g4"

	return this
}

// ControlParser tokens.
const (
	ControlParserEOF    = antlr.TokenEOF
	ControlParserT__0   = 1
	ControlParserT__1   = 2
	ControlParserT__2   = 3
	ControlParserT__3   = 4
	ControlParserT__4   = 5
	ControlParserT__5   = 6
	ControlParserT__6   = 7
	ControlParserT__7   = 8
	ControlParserT__8   = 9
	ControlParserT__9   = 10
	ControlParserT__10  = 11
	ControlParserT__11  = 12
	ControlParserT__12  = 13
	ControlParserT__13  = 14
	ControlParserT__14  = 15
	ControlParserT__15  = 16
	ControlParserT__16  = 17
	ControlParserT__17  = 18
	ControlParserT__18  = 19
	ControlParserT__19  = 20
	ControlParserT__20  = 21
	ControlParserT__21  = 22
	ControlParserT__22  = 23
	ControlParserT__23  = 24
	ControlParserINT    = 25
	ControlParserID     = 26
	ControlParserSTRING = 27
	ControlParserWS     = 28
)

// ControlParser rules.
const (
	ControlParserRULE_prog        = 0
	ControlParserRULE_block       = 1
	ControlParserRULE_stmt        = 2
	ControlParserRULE_assignstmt  = 3
	ControlParserRULE_printlnstmt = 4
	ControlParserRULE_printstmt   = 5
	ControlParserRULE_ifstmt      = 6
	ControlParserRULE_whilestmt   = 7
	ControlParserRULE_expr        = 8
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Block() IBlockContext
	EOF() antlr.TerminalNode

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
	p.RuleIndex = ControlParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(ControlParserEOF, 0)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ControlParserRULE_prog)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(18)
		p.Block()
	}
	{
		p.SetState(19)
		p.Match(ControlParserEOF)
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

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStmt() []IStmtContext
	Stmt(i int) IStmtContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) AllStmt() []IStmtContext {
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

func (s *BlockContext) Stmt(i int) IStmtContext {
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

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ControlParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(24)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&67109476) != 0 {
		{
			p.SetState(21)
			p.Stmt()
		}

		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
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

	// Getter signatures
	Assignstmt() IAssignstmtContext
	Printlnstmt() IPrintlnstmtContext
	Printstmt() IPrintstmtContext
	Ifstmt() IIfstmtContext
	Whilestmt() IWhilestmtContext

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
	p.RuleIndex = ControlParserRULE_stmt
	return p
}

func InitEmptyStmtContext(p *StmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_stmt
}

func (*StmtContext) IsStmtContext() {}

func NewStmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StmtContext {
	var p = new(StmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_stmt

	return p
}

func (s *StmtContext) GetParser() antlr.Parser { return s.parser }

func (s *StmtContext) Assignstmt() IAssignstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignstmtContext)
}

func (s *StmtContext) Printlnstmt() IPrintlnstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintlnstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintlnstmtContext)
}

func (s *StmtContext) Printstmt() IPrintstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrintstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrintstmtContext)
}

func (s *StmtContext) Ifstmt() IIfstmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfstmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfstmtContext)
}

func (s *StmtContext) Whilestmt() IWhilestmtContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhilestmtContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhilestmtContext)
}

func (s *StmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterStmt(s)
	}
}

func (s *StmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitStmt(s)
	}
}

func (s *StmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitStmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Stmt() (localctx IStmtContext) {
	localctx = NewStmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ControlParserRULE_stmt)
	p.SetState(32)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ControlParserID:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(27)
			p.Assignstmt()
		}

	case ControlParserT__1:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(28)
			p.Printlnstmt()
		}

	case ControlParserT__4:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(29)
			p.Printstmt()
		}

	case ControlParserT__5:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(30)
			p.Ifstmt()
		}

	case ControlParserT__8:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(31)
			p.Whilestmt()
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

// IAssignstmtContext is an interface to support dynamic dispatch.
type IAssignstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	Expr() IExprContext

	// IsAssignstmtContext differentiates from other interfaces.
	IsAssignstmtContext()
}

type AssignstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignstmtContext() *AssignstmtContext {
	var p = new(AssignstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_assignstmt
	return p
}

func InitEmptyAssignstmtContext(p *AssignstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_assignstmt
}

func (*AssignstmtContext) IsAssignstmtContext() {}

func NewAssignstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignstmtContext {
	var p = new(AssignstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_assignstmt

	return p
}

func (s *AssignstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignstmtContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *AssignstmtContext) Expr() IExprContext {
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

func (s *AssignstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterAssignstmt(s)
	}
}

func (s *AssignstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitAssignstmt(s)
	}
}

func (s *AssignstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitAssignstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Assignstmt() (localctx IAssignstmtContext) {
	localctx = NewAssignstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ControlParserRULE_assignstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)
		p.Match(ControlParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(35)
		p.Match(ControlParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(36)
		p.expr(0)
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

// IPrintlnstmtContext is an interface to support dynamic dispatch.
type IPrintlnstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsPrintlnstmtContext differentiates from other interfaces.
	IsPrintlnstmtContext()
}

type PrintlnstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintlnstmtContext() *PrintlnstmtContext {
	var p = new(PrintlnstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_printlnstmt
	return p
}

func InitEmptyPrintlnstmtContext(p *PrintlnstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_printlnstmt
}

func (*PrintlnstmtContext) IsPrintlnstmtContext() {}

func NewPrintlnstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintlnstmtContext {
	var p = new(PrintlnstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_printlnstmt

	return p
}

func (s *PrintlnstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintlnstmtContext) Expr() IExprContext {
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

func (s *PrintlnstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintlnstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintlnstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterPrintlnstmt(s)
	}
}

func (s *PrintlnstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitPrintlnstmt(s)
	}
}

func (s *PrintlnstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitPrintlnstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Printlnstmt() (localctx IPrintlnstmtContext) {
	localctx = NewPrintlnstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ControlParserRULE_printlnstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(38)
		p.Match(ControlParserT__1)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(39)
		p.Match(ControlParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(40)
		p.expr(0)
	}
	{
		p.SetState(41)
		p.Match(ControlParserT__3)
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

// IPrintstmtContext is an interface to support dynamic dispatch.
type IPrintstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext

	// IsPrintstmtContext differentiates from other interfaces.
	IsPrintstmtContext()
}

type PrintstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrintstmtContext() *PrintstmtContext {
	var p = new(PrintstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_printstmt
	return p
}

func InitEmptyPrintstmtContext(p *PrintstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_printstmt
}

func (*PrintstmtContext) IsPrintstmtContext() {}

func NewPrintstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrintstmtContext {
	var p = new(PrintstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_printstmt

	return p
}

func (s *PrintstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *PrintstmtContext) Expr() IExprContext {
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

func (s *PrintstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrintstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterPrintstmt(s)
	}
}

func (s *PrintstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitPrintstmt(s)
	}
}

func (s *PrintstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitPrintstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Printstmt() (localctx IPrintstmtContext) {
	localctx = NewPrintstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ControlParserRULE_printstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(43)
		p.Match(ControlParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(44)
		p.Match(ControlParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.expr(0)
	}
	{
		p.SetState(46)
		p.Match(ControlParserT__3)
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

// IIfstmtContext is an interface to support dynamic dispatch.
type IIfstmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Block() IBlockContext

	// IsIfstmtContext differentiates from other interfaces.
	IsIfstmtContext()
}

type IfstmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfstmtContext() *IfstmtContext {
	var p = new(IfstmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_ifstmt
	return p
}

func InitEmptyIfstmtContext(p *IfstmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_ifstmt
}

func (*IfstmtContext) IsIfstmtContext() {}

func NewIfstmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfstmtContext {
	var p = new(IfstmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_ifstmt

	return p
}

func (s *IfstmtContext) GetParser() antlr.Parser { return s.parser }

func (s *IfstmtContext) Expr() IExprContext {
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

func (s *IfstmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfstmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfstmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfstmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterIfstmt(s)
	}
}

func (s *IfstmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitIfstmt(s)
	}
}

func (s *IfstmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitIfstmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Ifstmt() (localctx IIfstmtContext) {
	localctx = NewIfstmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ControlParserRULE_ifstmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.Match(ControlParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(49)
		p.Match(ControlParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.expr(0)
	}
	{
		p.SetState(51)
		p.Match(ControlParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(52)
		p.Match(ControlParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(53)
		p.Block()
	}
	{
		p.SetState(54)
		p.Match(ControlParserT__7)
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

// IWhilestmtContext is an interface to support dynamic dispatch.
type IWhilestmtContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	Block() IBlockContext

	// IsWhilestmtContext differentiates from other interfaces.
	IsWhilestmtContext()
}

type WhilestmtContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhilestmtContext() *WhilestmtContext {
	var p = new(WhilestmtContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_whilestmt
	return p
}

func InitEmptyWhilestmtContext(p *WhilestmtContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_whilestmt
}

func (*WhilestmtContext) IsWhilestmtContext() {}

func NewWhilestmtContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhilestmtContext {
	var p = new(WhilestmtContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_whilestmt

	return p
}

func (s *WhilestmtContext) GetParser() antlr.Parser { return s.parser }

func (s *WhilestmtContext) Expr() IExprContext {
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

func (s *WhilestmtContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhilestmtContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhilestmtContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhilestmtContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterWhilestmt(s)
	}
}

func (s *WhilestmtContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitWhilestmt(s)
	}
}

func (s *WhilestmtContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitWhilestmt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Whilestmt() (localctx IWhilestmtContext) {
	localctx = NewWhilestmtContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ControlParserRULE_whilestmt)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(56)
		p.Match(ControlParserT__8)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(57)
		p.Match(ControlParserT__2)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(58)
		p.expr(0)
	}
	{
		p.SetState(59)
		p.Match(ControlParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.Match(ControlParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)
		p.Block()
	}
	{
		p.SetState(62)
		p.Match(ControlParserT__7)
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
	p.RuleIndex = ControlParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ControlParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ControlParserRULE_expr

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

type BoolExprContext struct {
	ExprContext
}

func NewBoolExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolExprContext {
	var p = new(BoolExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *BoolExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterBoolExpr(s)
	}
}

func (s *BoolExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitBoolExpr(s)
	}
}

func (s *BoolExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitBoolExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdExprContext struct {
	ExprContext
}

func NewIdExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdExprContext {
	var p = new(IdExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdExprContext) ID() antlr.TerminalNode {
	return s.GetToken(ControlParserID, 0)
}

func (s *IdExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterIdExpr(s)
	}
}

func (s *IdExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitIdExpr(s)
	}
}

func (s *IdExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitIdExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParExprContext struct {
	ExprContext
}

func NewParExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParExprContext {
	var p = new(ParExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParExprContext) Expr() IExprContext {
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

func (s *ParExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterParExpr(s)
	}
}

func (s *ParExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitParExpr(s)
	}
}

func (s *ParExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitParExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type StrExprContext struct {
	ExprContext
}

func NewStrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StrExprContext {
	var p = new(StrExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *StrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StrExprContext) STRING() antlr.TerminalNode {
	return s.GetToken(ControlParserSTRING, 0)
}

func (s *StrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterStrExpr(s)
	}
}

func (s *StrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitStrExpr(s)
	}
}

func (s *StrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitStrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotExprContext struct {
	ExprContext
	right IExprContext
}

func NewNotExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotExprContext {
	var p = new(NotExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NotExprContext) GetRight() IExprContext { return s.right }

func (s *NotExprContext) SetRight(v IExprContext) { s.right = v }

func (s *NotExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotExprContext) Expr() IExprContext {
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

func (s *NotExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterNotExpr(s)
	}
}

func (s *NotExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitNotExpr(s)
	}
}

func (s *NotExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitNotExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntExprContext struct {
	ExprContext
}

func NewIntExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntExprContext {
	var p = new(IntExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IntExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntExprContext) INT() antlr.TerminalNode {
	return s.GetToken(ControlParserINT, 0)
}

func (s *IntExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterIntExpr(s)
	}
}

func (s *IntExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitIntExpr(s)
	}
}

func (s *IntExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitIntExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type OpExprContext struct {
	ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewOpExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpExprContext {
	var p = new(OpExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *OpExprContext) GetOp() antlr.Token { return s.op }

func (s *OpExprContext) SetOp(v antlr.Token) { s.op = v }

func (s *OpExprContext) GetLeft() IExprContext { return s.left }

func (s *OpExprContext) GetRight() IExprContext { return s.right }

func (s *OpExprContext) SetLeft(v IExprContext) { s.left = v }

func (s *OpExprContext) SetRight(v IExprContext) { s.right = v }

func (s *OpExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpExprContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *OpExprContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *OpExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.EnterOpExpr(s)
	}
}

func (s *OpExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ControlListener); ok {
		listenerT.ExitOpExpr(s)
	}
}

func (s *OpExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ControlVisitor:
		return t.VisitOpExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ControlParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *ControlParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 16
	p.EnterRecursionRule(localctx, 16, ControlParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ControlParserT__9:
		localctx = NewNotExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(65)
			p.Match(ControlParserT__9)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)

			var _x = p.expr(13)

			localctx.(*NotExprContext).right = _x
		}

	case ControlParserT__2:
		localctx = NewParExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(67)
			p.Match(ControlParserT__2)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(68)
			p.expr(0)
		}
		{
			p.SetState(69)
			p.Match(ControlParserT__3)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ControlParserINT:
		localctx = NewIntExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(71)
			p.Match(ControlParserINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ControlParserID:
		localctx = NewIdExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(72)
			p.Match(ControlParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ControlParserSTRING:
		localctx = NewStrExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(73)
			p.Match(ControlParserSTRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ControlParserT__22, ControlParserT__23:
		localctx = NewBoolExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(74)
			_la = p.GetTokenStream().LA(1)

			if !(_la == ControlParserT__22 || _la == ControlParserT__23) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(100)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(98)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ControlParserRULE_expr)
				p.SetState(77)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(78)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ControlParserT__10 || _la == ControlParserT__11) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(79)

					var _x = p.expr(13)

					localctx.(*OpExprContext).right = _x
				}

			case 2:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ControlParserRULE_expr)
				p.SetState(80)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(81)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ControlParserT__12 || _la == ControlParserT__13) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(82)

					var _x = p.expr(12)

					localctx.(*OpExprContext).right = _x
				}

			case 3:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ControlParserRULE_expr)
				p.SetState(83)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(84)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ControlParserT__14 || _la == ControlParserT__15) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(85)

					var _x = p.expr(11)

					localctx.(*OpExprContext).right = _x
				}

			case 4:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ControlParserRULE_expr)
				p.SetState(86)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(87)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ControlParserT__16 || _la == ControlParserT__17) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(88)

					var _x = p.expr(10)

					localctx.(*OpExprContext).right = _x
				}

			case 5:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ControlParserRULE_expr)
				p.SetState(89)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
					goto errorExit
				}
				{
					p.SetState(90)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpExprContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ControlParserT__18 || _la == ControlParserT__19) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpExprContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(91)

					var _x = p.expr(9)

					localctx.(*OpExprContext).right = _x
				}

			case 6:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ControlParserRULE_expr)
				p.SetState(92)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
					goto errorExit
				}
				{
					p.SetState(93)
					p.Match(ControlParserT__20)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(94)

					var _x = p.expr(8)

					localctx.(*OpExprContext).right = _x
				}

			case 7:
				localctx = NewOpExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpExprContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ControlParserRULE_expr)
				p.SetState(95)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
					goto errorExit
				}
				{
					p.SetState(96)
					p.Match(ControlParserT__21)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(97)

					var _x = p.expr(7)

					localctx.(*OpExprContext).right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(102)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
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

func (p *ControlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 8:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ControlParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 7)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 6)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
