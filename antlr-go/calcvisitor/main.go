package main

import (
	"fmt"
	"parser"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Visitor struct {
	parser.CalcVisitor
}

func (v *Visitor) Visit(tree antlr.ParseTree) float64 {
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

func (v *Visitor) VisitOp(ctx *parser.OpContext) float64 {
	l := v.Visit(ctx.GetLeft())
	r := v.Visit(ctx.GetRight())
	op := ctx.GetOp().GetText()
	fmt.Println("op", op)
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

func (v *Visitor) VisitDigit(ctx *parser.DigitContext) float64 {
	fmt.Println("digit", ctx.GetText())
	i1, _ := strconv.ParseFloat(ctx.GetText(), 64)
	return i1
}

func (v *Visitor) VisitParen(ctx *parser.ParenContext) float64 {
	fmt.Println("parent", ctx.GetText())
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
	var result = visitor.Visit(tree)
	fmt.Println(expression, "=", result)
}
