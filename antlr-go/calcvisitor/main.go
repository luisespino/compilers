package main

import (
	"fmt"
	"calcvisitor/parser"
	"strconv"

	"github.com/antlr4-go/antlr/v4"
)

type Visitor struct {
	parser.CalcVisitor
	stackop string
}

func (v *Visitor) Visit(tree antlr.ParseTree) int64 {
	switch val := tree.(type) {
	case *parser.OpContext:
		return v.VisitOp(val)
	case *parser.DigitContext:
		return v.VisitDigit(val)
	case *parser.ParenContext:
		return v.VisitParen(val)
	default:
		panic("Unknown context")
	}
}

func (v *Visitor) VisitOp(ctx *parser.OpContext) int64 {
	l := v.Visit(ctx.GetLeft())
	r := v.Visit(ctx.GetRight())
	op := ctx.GetOp().GetText()
	v.stackop += "- Visit Op "+op+" ~ Two pop an apply "+op+"\n"
	switch op {
	case "+":
		return l + r
	case "-":
		return l - r
	case "*":
		return l * r
	case "/":
		return l / r
	}
	return 0
}

func (v *Visitor) VisitDigit(ctx *parser.DigitContext) int64 {
	v.stackop  += "- Visit Digit " + ctx.GetText() + " ~ Push " + ctx.GetText() + "\n"
	i1, _ := strconv.ParseInt(ctx.GetText(), 10, 64)
	return i1
}

func (v *Visitor) VisitParen(ctx *parser.ParenContext) int64 {
	tar := v.Visit(ctx.Expr())
	return tar
}

func main() {
	expression := "3*(5+4)"
	input := antlr.NewInputStream(expression)
	lexer := parser.NewCalcLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCalcParser(stream)
	p.BuildParseTrees = true
	tree := p.Expr()
	var visitor = Visitor{}
	visitor.stackop = "";
	var result = visitor.Visit(tree)
	fmt.Println("Visitor call stack:\n" + visitor.stackop)
	fmt.Println("Evaluation:", expression, "=", result)
}
