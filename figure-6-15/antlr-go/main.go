package main

import (
    "fmt"
    "os"
    "strings"

	"figure_6_15/parser" 
	"github.com/antlr4-go/antlr/v4"
)

type Attr struct {
    Str string
    Obj any
}

type Variable struct {
    Name       string
    Expression any
    Type       string
}

type Translator struct {
    parser.BaseGrammarVisitor
    Variables map[string]Variable
    Data      string
    Text      string
}

func NewTranslator() *Translator {
    return &Translator{
		Variables: make(map[string]Variable),
        Data: ".global _start\n\n.data\n",
        Text: "\n.text\n_start:\n",
    }
}


func (t *Translator) VisitProgram(ctx *parser.ProgramContext) interface{} {
    for _, stmt := range ctx.AllStmt() {
		stmt.Accept(t)
    }
    t.Text += "\tmov x8, #93\n\tsvc #0\n"
    return Attr{
        Str: t.Data + t.Text,
        Obj: nil,
    }
}

func (t *Translator) VisitStmtDeclaration(ctx *parser.StmtDeclarationContext) interface{} {
	return ctx.Decl().Accept(t)
}

func (t *Translator) VisitStmtAssigment(ctx *parser.StmtAssigmentContext) interface{} {
	return ctx.Assg().Accept(t)
}

func (t *Translator) VisitDeclExpr(ctx *parser.DeclExprContext) interface{} {
    name := ctx.Var_().GetText()
    typ := ctx.GetType_().GetText()
	exprIface := ctx.Expr().Accept(t)
	exprAttr := exprIface.(Attr) 
    expr := exprAttr.Obj.([]any)

    t.Variables[name] = Variable{name, expr, typ}

    dimensions := getDimensions(expr)
	
    flat := flattenArray(expr)

    t.Data += name + ": .dword "
	nums := strings.Join(flat, ", ")
	if nums == "" {
		t.Data += "0\n"
	} else {
		t.Data += nums + "\n"
	}

	if len(dimensions) == 1 {
		t.Data += name + "_cols: .dword " + fmt.Sprint(dimensions[0]) + "\n"
	} else if len(dimensions) == 2 {
		t.Data += name + "_rows: .dword " + fmt.Sprint(dimensions[0]) + "\n"
		t.Data += name + "_cols: .dword " + fmt.Sprint(dimensions[1]) + "\n"
	} else if len(dimensions) == 3 {
		t.Data += name + "_face: .dword " + fmt.Sprint(dimensions[0]) + "\n"
		t.Data += name + "_rows: .dword " + fmt.Sprint(dimensions[1]) + "\n"
		t.Data += name + "_cols: .dword " + fmt.Sprint(dimensions[2]) + "\n"
	}

    return Attr{}
}

func (t *Translator) VisitAssgArray(ctx *parser.AssgArrayContext) interface{} {
	dest := ctx.Var_(0).GetText()
	src := ctx.Var_(1).GetText()
	indexIface := ctx.Index().Accept(t)
	indices := indexIface.([]string)

	srcVar := t.Variables[src]
	dimensions := getDimensions(srcVar.Expression)

	if len(dimensions) == 1 {
		t.Text += fmt.Sprintf("\tmov x0, #%s\n", indices[0])
		t.Text += "\tmov x1, #8\n"
		t.Text += "\tmul x0, x0, x1\n"
		t.Text += fmt.Sprintf("\tldr x1, =%s\n", src)
		t.Text += "\tadd x2, x1, x0\n"
		t.Text += fmt.Sprintf("\tldr x0, =%s\n", dest)
		t.Text += "\tldr x1, [x2]\n"
		t.Text += "\tstr x1, [x0]\n\n"

	} else if len(dimensions) == 2 {
		t.Text += fmt.Sprintf("\tmov x0, #%s\n", indices[0])
		t.Text += "\tmov x1, #8\n"
		t.Text += "\tmul x0, x0, x1\n"
		t.Text += fmt.Sprintf("\tmov x1, #%s\n", indices[1])
		t.Text += "\tmov x2, #8\n"
		t.Text += "\tmul x1, x1, x2\n"
		t.Text += fmt.Sprintf("\tldr x2, =%s_cols\n", src)
        t.Text += "\tldr x2, [x2]\n"
		t.Text += "\tmul x3, x0, x2\n"
		t.Text += "\tadd x3, x3, x1\n"
		t.Text += fmt.Sprintf("\tldr x4, =%s\n", src)
		t.Text += "\tadd x4, x4, x3\n"
		t.Text += fmt.Sprintf("\tldr x0, =%s\n", dest)
		t.Text += "\tldr x1, [x4]\n"
		t.Text += "\tstr x1, [x0]\n\n"

	} else if len(dimensions) == 3 {
		t.Text += fmt.Sprintf("\tmov x0, #%s\n", indices[0])
		t.Text += "\tmov x1, #8\n"
		t.Text += "\tmul x0, x0, x1\n"
		t.Text += fmt.Sprintf("\tmov x1, #%s\n", indices[1])
		t.Text += "\tmov x2, #8\n"
		t.Text += "\tmul x1, x1, x2\n"
		t.Text += fmt.Sprintf("\tmov x2, #%s\n", indices[2])
		t.Text += "\tmov x3, #8\n"
		t.Text += "\tmul x2, x2, x3\n"
		t.Text += fmt.Sprintf("\tldr x3, =%s_rows\n", src)
        t.Text += "\tldr x3, [x3]\n"
		t.Text += "\tmul x4, x3, x0\n"
		t.Text += "\tadd x4, x4, x1\n"
		t.Text += fmt.Sprintf("\tldr x5, =%s_cols\n", src)
        t.Text += "\tldr x5, [x5]\n"
		t.Text += "\tmul x6, x5, x4\n"
		t.Text += "\tadd x6, x6, x2\n"
		t.Text += fmt.Sprintf("\tldr x0, =%s\n", src)
		t.Text += "\tadd x0, x0, x6\n"
		t.Text += fmt.Sprintf("\tldr x1, =%s\n", dest)
		t.Text += "\tldr x2, [x0]\n"
		t.Text += "\tstr x2, [x1]\n\n"
	}

	return Attr{}
}

func (t *Translator) VisitIndexMany(ctx *parser.IndexManyContext) interface{} {
	val := ctx.Val().GetText()
	rest := ctx.Index().Accept(t).([]string)
	return append([]string{val}, rest...)
}

func (t *Translator) VisitIndexOne(ctx *parser.IndexOneContext) interface{} {
	return []string{ctx.Val().GetText()}
}

func (t *Translator) VisitIndex(ctx *parser.IndexContext) interface{} {
	return ctx.Accept(t)
}


func (t *Translator) VisitDeclVal(ctx *parser.DeclValContext) interface{} {
    name := ctx.Var_().GetText()
    typ := ctx.GetType_().GetText()
	expr := ctx.DIGIT().GetText()


    t.Variables[name] = Variable{name, expr, typ}

	t.Data += name + ": .dword "+expr+"\n"


    return Attr{}
}

func getDimensions(obj any) []int {
    arr, ok := obj.([]any)
    if !ok {
        return nil
    }
    dims := []int{len(arr)}
    if len(arr) > 0 {
        innerDims := getDimensions(arr[0])
        dims = append(dims, innerDims...)
    }
    return dims
}

func flattenArray(obj any) []string {
    arr, ok := obj.([]any)
    if !ok {
        return nil
    }
    var res []string
    for _, v := range arr {
        switch vv := v.(type) {
        case string:
            res = append(res, vv)
        case []any:
            res = append(res, flattenArray(vv)...)
        }
    }
    if len(res) == 0 {
        return []string{"0"}
    }
    return res
}

func (t *Translator) VisitVarName(ctx *parser.VarNameContext) interface{} {
	fmt.Println(ctx.NAME().GetText())
    return Attr{Str: ctx.NAME().GetText()}
}

func (t *Translator) VisitExprEmpty(ctx *parser.ExprEmptyContext) interface{} {
    return Attr{Obj: []any{}}
}

func (t *Translator) VisitExprValues(ctx *parser.ExprValuesContext) interface{} {
	listIface := ctx.List().Accept(t)
	listAttr := listIface.(Attr) 
    list := listAttr.Obj.([]any)
    return Attr{Str: "", Obj: list}
}

func (t *Translator) VisitValues(ctx *parser.ValuesContext) interface{} {
	valIface := ctx.Val().Accept(t)
	valAttr := valIface.(Attr) 
	listIface := ctx.List().Accept(t)
	listAttr := listIface.(Attr) 	

    var value []any
    if valAttr.Str != "" {
        value = []any{valAttr.Str}
    } else {
        value = []any{valAttr.Obj}
    }

    combined := append(value, listAttr.Obj.([]any)...)

    return Attr{Str: "", Obj: combined}
}


func (t *Translator) VisitValue(ctx *parser.ValueContext) interface{} {
	valIface := ctx.Val().Accept(t)
	valAttr := valIface.(Attr) 
    if valAttr.Str != "" {
        return Attr{Obj: []any{valAttr.Str}}
    }
    return Attr{Obj: []any{valAttr.Obj}}
}

func (t *Translator) VisitValDigit(ctx *parser.ValDigitContext) interface{} {
    return Attr{Str: ctx.DIGIT().GetText()}
}

func (t *Translator) VisitValExpr(ctx *parser.ValExprContext) interface{} {
	exprIface := ctx.Expr().Accept(t)
	exprAttr := exprIface.(Attr) 
    expr := exprAttr.Obj.([]any)
    return Attr{Str: "", Obj: expr}
}


func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: go run main.go <inputfile>")
        return
    }
    input, err := os.ReadFile(os.Args[1])
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
	
    is := antlr.NewInputStream(string(input))
    lexer := parser.NewGrammarLexer(is)
    stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
    p := parser.NewGrammarParser(stream)

    tree := p.Prog()
    translator := NewTranslator()
    res := tree.Accept(translator).(Attr)
    fmt.Println(res.Str)
}
