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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 10, 24, 4,
	2, 9, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 11, 10, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 19, 10, 2, 12, 2, 14, 2, 22, 11, 2, 3,
	2, 2, 3, 2, 3, 2, 2, 4, 3, 2, 3, 4, 3, 2, 5, 6, 2, 25, 2, 10, 3, 2, 2,
	2, 4, 5, 8, 2, 1, 2, 5, 11, 7, 9, 2, 2, 6, 7, 7, 7, 2, 2, 7, 8, 5, 2, 2,
	2, 8, 9, 7, 8, 2, 2, 9, 11, 3, 2, 2, 2, 10, 4, 3, 2, 2, 2, 10, 6, 3, 2,
	2, 2, 11, 20, 3, 2, 2, 2, 12, 13, 12, 6, 2, 2, 13, 14, 9, 2, 2, 2, 14,
	19, 5, 2, 2, 7, 15, 16, 12, 5, 2, 2, 16, 17, 9, 3, 2, 2, 17, 19, 5, 2,
	2, 6, 18, 12, 3, 2, 2, 2, 18, 15, 3, 2, 2, 2, 19, 22, 3, 2, 2, 2, 20, 18,
	3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 21, 3, 3, 2, 2, 2, 22, 20, 3, 2, 2, 2,
	5, 10, 18, 20,
}
var literalNames = []string{
	"", "'*'", "'/'", "'+'", "'-'", "'('", "')'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "INT", "WS",
}

var ruleNames = []string{
	"expr",
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
	CalcParserINT  = 7
	CalcParserWS   = 8
)

// CalcParserRULE_expr is the CalcParser rule.
const CalcParserRULE_expr = 0

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = CalcParserRULE_expr
	return p
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = CalcParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyFrom(ctx *ExprContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type OpContext struct {
	*ExprContext
	left  IExprContext
	op    antlr.Token
	right IExprContext
}

func NewOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *OpContext {
	var p = new(OpContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *OpContext) GetOp() antlr.Token { return s.op }

func (s *OpContext) SetOp(v antlr.Token) { s.op = v }

func (s *OpContext) GetLeft() IExprContext { return s.left }

func (s *OpContext) GetRight() IExprContext { return s.right }

func (s *OpContext) SetLeft(v IExprContext) { s.left = v }

func (s *OpContext) SetRight(v IExprContext) { s.right = v }

func (s *OpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OpContext) AllExpr() []IExprContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExprContext)(nil)).Elem())
	var tst = make([]IExprContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExprContext)
		}
	}

	return tst
}

func (s *OpContext) Expr(i int) IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *OpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcVisitor:
		return t.VisitOp(s)

	default:
		return t.VisitChildren(s)
	}
}

type DigitContext struct {
	*ExprContext
	digit antlr.Token
}

func NewDigitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DigitContext {
	var p = new(DigitContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *DigitContext) GetDigit() antlr.Token { return s.digit }

func (s *DigitContext) SetDigit(v antlr.Token) { s.digit = v }

func (s *DigitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DigitContext) INT() antlr.TerminalNode {
	return s.GetToken(CalcParserINT, 0)
}

func (s *DigitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcVisitor:
		return t.VisitDigit(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenContext struct {
	*ExprContext
}

func NewParenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenContext {
	var p = new(ParenContext)

	p.ExprContext = NewEmptyExprContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExprContext))

	return p
}

func (s *ParenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenContext) Expr() IExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case CalcVisitor:
		return t.VisitParen(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *CalcParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *CalcParser) expr(_p int) (localctx IExprContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, CalcParserRULE_expr, _p)
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
	p.SetState(8)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case CalcParserINT:
		localctx = NewDigitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(3)

			var _m = p.Match(CalcParserINT)

			localctx.(*DigitContext).digit = _m
		}

	case CalcParserT__4:
		localctx = NewParenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(4)
			p.Match(CalcParserT__4)
		}
		{
			p.SetState(5)
			p.expr(0)
		}
		{
			p.SetState(6)
			p.Match(CalcParserT__5)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(16)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expr)
				p.SetState(10)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(11)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CalcParserT__0 || _la == CalcParserT__1) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(12)

					var _x = p.expr(5)

					localctx.(*OpContext).right = _x
				}

			case 2:
				localctx = NewOpContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*OpContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, CalcParserRULE_expr)
				p.SetState(13)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(14)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*OpContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == CalcParserT__2 || _la == CalcParserT__3) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*OpContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(15)

					var _x = p.expr(4)

					localctx.(*OpContext).right = _x
				}

			}

		}
		p.SetState(20)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
	}

	return localctx
}

func (p *CalcParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *CalcParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
