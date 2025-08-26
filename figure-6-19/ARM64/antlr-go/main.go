package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"figure_6_19/parser" 
	"github.com/antlr4-go/antlr/v4"
)

type Attr struct {
	addrNUM string
	addrID  string
	code    string
}

type SymbolTable struct {
	ids map[string]struct{}
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		ids: make(map[string]struct{}),
	}
}

func (st *SymbolTable) Insert(name string) {
	st.ids[name] = struct{}{}
}

func (st *SymbolTable) Exists(name string) bool {
	_, exists := st.ids[name]
	return exists
}

type Translator struct {
	parser.BaseGrammarVisitor
	tid int
	data string
	code string
	symbolTable *SymbolTable
}

func NewTranslator() *Translator {
    return &Translator{
        tid:  0,
        data: ".global _start\n\n.data\n",
        code: "\n.text\n_start:\n",
		symbolTable: NewSymbolTable(),
    }
}

func isNUM(attr Attr) bool {
    return attr.addrNUM != "" && attr.addrID == ""
}

func (t *Translator) VisitProgram(ctx *parser.ProgramContext) interface{} {
	for _, s := range ctx.AllS() {
		s.Accept(t)
	}
	t.code += "\tmov X8, #93\n\tsvc #0\n";
	return Attr{ code : t.data + t.code };
}

func (t *Translator) VisitAssign(ctx *parser.AssignContext) interface{} {
	id := ctx.ID().GetText()
	
	exprIface := ctx.E().Accept(t)
	expr := exprIface.(Attr) 
	
	if isNUM(expr) {
		if t.symbolTable.Exists(id) {
			t.code += "\tldr X0, =" + id + "\n"
			t.code += "\tmov X1, #" + expr.addrNUM + "\n"
			t.code += "\tstr X1, [X0]\n\n"
		} else {
			t.data += id + ": .word " + expr.addrNUM + "\n"
			t.symbolTable.Insert(id)
		}
	} else {
		if t.symbolTable.Exists(id) {
			t.code += "\tldr X0, =" + id + "\n"
			t.code += "\tldr X1, =" + expr.addrID + "\n"
			t.code += "\tldr X2, [X1]\n"
			t.code += "\tstr X2, [X0]\n\n"
		} else {
			t.data += id + ": .word 0\n"
			t.code += "\tldr X0, =" + id + "\n"
			t.code += "\tldr X1, =" + expr.addrID + "\n"
			t.code += "\tldr X2, [X1]\n"
			t.code += "\tstr X2, [X0]\n\n"
			t.symbolTable.Insert(id)
		}
	}

	return Attr{}
}

func (t *Translator) VisitSumres(ctx *parser.SumresContext) interface{} {
	leftIface := ctx.E(0).Accept(t)
	left := leftIface.(Attr) 
	rightIface := ctx.E(1).Accept(t)
	right := rightIface.(Attr) 
	
	t.tid++
	addr := fmt.Sprintf("t%d", t.tid)
	t.data += addr + ": .word 0\n"

	t.code += "\tldr X0, =" + addr + "\n"

	op := ""
	if ctx.GetOp().GetText() == "+" {
		op = "add"
	} else {
		op = "sub"
	}

	if isNUM(left) {
		t.code += "\tldr X1, [X0]\n"
		if isNUM(right) {
			t.code += "\tadd X1, X1, #" + left.addrNUM + "\n"
			t.code += "\t" + op + " X1, X1, #" + right.addrNUM + "\n"
			t.code += "\tstr X1, [X0]\n"
		} else {
			t.code += "\tadd X1, X1, #" + left.addrNUM + "\n"
			t.code += "\tldr X2, =" + right.addrID + "\n"
			t.code += "\tldr X3, [X2]\n"
			t.code += "\t" + op + " X1, X1, X3\n"
			t.code += "\tstr X1, [X0]\n"
		}
	} else {
		if isNUM(right) {
			t.code += "\tldr X1, =" + left.addrID + "\n"
			t.code += "\tldr X2, [X1]\n"
			t.code += "\t" + op + " X2, X2, #" + right.addrNUM + "\n"
			t.code += "\tstr X2, [X0]\n"
		} else {
			t.code += "\tldr X1, =" + left.addrID + "\n"
			t.code += "\tldr X2, [X1]\n"
			t.code += "\tldr X3, =" + right.addrID + "\n"
			t.code += "\tldr X4, [X3]\n"
			t.code += "\t" + op + " X2, X2, X4\n"
			t.code += "\tstr X2, [X0]\n"
		}
	}

	t.code += "\n"

	return Attr{addrID: addr}
}

func (t *Translator) VisitMuldiv(ctx *parser.MuldivContext) interface{} {
	leftIface := ctx.E(0).Accept(t)
	left := leftIface.(Attr) 
	rightIface := ctx.E(1).Accept(t)
	right := rightIface.(Attr) 

	t.tid++
	addr := fmt.Sprintf("t%d", t.tid)
	t.data += addr + ": .word 0\n"

	t.code += "\tldr X0, =" + addr + "\n"

	op := ""
	if ctx.GetOp().GetText() == "*" {
		op = "mul"
	} else {
		op = "sdiv"
	}

	if isNUM(left) {
		t.code += "\tldr X1, [X0]\n"
		if isNUM(right) {
			t.code += "\tadd X1, X1, #" + left.addrNUM + "\n"
			t.code += "\tmov X2, #" + right.addrNUM + "\n"
			t.code += "\t" + op + " X1, X1, X2\n"
			t.code += "\tstr X1, [X0]\n"
		} else {
			t.code += "\tadd X1, X1, #" + left.addrNUM + "\n"
			t.code += "\tldr X2, =" + right.addrID + "\n"
			t.code += "\tldr X3, [X2]\n"
			t.code += "\t" + op + " X1, X1, X3\n"
			t.code += "\tstr X1, [X0]\n"
		}
	} else {
		if isNUM(right) {
			t.code += "\tldr X1, =" + left.addrID + "\n"
			t.code += "\tldr X2, [X1]\n"
			t.code += "\tmov X3, #" + right.addrNUM + "\n"
			t.code += "\t" + op + " X2, X2, X3\n"
			t.code += "\tstr X2, [X0]\n"
		} else {
			t.code += "\tldr X1, =" + left.addrID + "\n"
			t.code += "\tldr X2, [X1]\n"
			t.code += "\tldr X3, =" + right.addrID + "\n"
			t.code += "\tldr X4, [X3]\n"
			t.code += "\t" + op + " X2, X2, X4\n"
			t.code += "\tstr X2, [X0]\n"
		}
	}

	t.code += "\n"

	return Attr{addrID: addr}
}

func (t *Translator) VisitPassE(ctx *parser.PassEContext) interface{} {
	return ctx.E().Accept(t)
}

func (t *Translator) VisitId(ctx *parser.IdContext) interface{} {
	return Attr{addrID : ctx.ID().GetText()}
}

func (t *Translator) VisitNum(ctx *parser.NumContext) interface{} {
	return Attr{addrNUM : ctx.NUM().GetText()}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Ingrese instrucción o 'exit' para salir: ")
		input, err := reader.ReadString('\n')
		// test_input = a = 1  b = 2  c = a + b * 2   d = (c + 1) * 3
		if err != nil {
			fmt.Println("Error leyendo entrada:", err)
			continue
		}

		input = strings.TrimSpace(input)
		if strings.ToLower(input) == "exit" {
			fmt.Println("Exiting...")
			break
		}

		translator := NewTranslator()

		inputStream := antlr.NewInputStream(input)
		lexer := parser.NewGrammarLexer(inputStream)
		tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := parser.NewGrammarParser(tokens)

		tree := p.P()
		result := tree.Accept(translator).(Attr)
		fmt.Println(result.code)
	}
}
