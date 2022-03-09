package main

import (
	"fmt"
	"strconv"

	"parser"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var id int = 0
var dot string = "graph{" + "\n"

type Node struct {
	id    int
	label string
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (t Tree) inorder(tmp *Node) {
	if tmp != nil {
		t.inorder(tmp.left)
		fmt.Print(tmp.label, " ")
		t.inorder(tmp.right)
	}
}

func (t Tree) postorder(tmp *Node) {
	if tmp != nil {
		t.postorder(tmp.left)
		t.postorder(tmp.right)
		fmt.Print(tmp.label, " ")
		dot += strconv.Itoa(tmp.id) + "[label=\"" + tmp.label + "\"];\n"
		if tmp.left != nil {
			dot += strconv.Itoa(tmp.id) + "--" + strconv.Itoa(tmp.left.id) + ";\n"
		}
		if tmp.right != nil {
			dot += strconv.Itoa(tmp.id) + "--" + strconv.Itoa(tmp.right.id) + ";\n"
		}
	}
}

type calcListener struct {
	*parser.BaseCalcListener
	stack []*Node
}

func (l *calcListener) push(i *Node) {
	l.stack = append(l.stack, i)
}

func (l *calcListener) pop() *Node {
	if len(l.stack) < 1 {
		panic("empty stack")
	}
	result := l.stack[len(l.stack)-1]
	l.stack = l.stack[:len(l.stack)-1]
	return result
}

func (l *calcListener) ExitSumres(c *parser.SumresContext) {
	right, left := l.pop(), l.pop()
	id += 1
	l.push(&Node{
		id:    id,
		label: c.GetOp().GetText(),
		right: right,
		left:  left})
}

func (l *calcListener) ExitMuldiv(c *parser.MuldivContext) {
	right, left := l.pop(), l.pop()
	id += 1
	l.push(&Node{
		id:    id,
		label: c.GetOp().GetText(),
		right: right,
		left:  left})
}

func (l *calcListener) ExitId(c *parser.IdContext) {
	id += 1
	l.push(&Node{id: id, label: c.GetText()})
}

func (l *calcListener) ExitNum(c *parser.NumContext) {
	id += 1
	l.push(&Node{id: id, label: c.GetText()})
}

func main() {
	is := antlr.NewInputStream("3*b+4")
	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewCalcParser(stream)
	var listener calcListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.E())
	tree := &Tree{root: listener.pop()}
	fmt.Println("Inorder:")
	tree.inorder(tree.root)
	fmt.Println("\n\nPostorder:")
	tree.postorder(tree.root)
	fmt.Println("\n\nDot:")
	fmt.Println(dot + "}")
}
