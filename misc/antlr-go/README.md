# Example: Using ANTLR with Go on Arch Linux

This guide explains how to use ANTLR4 with Go on Arch Linux. The main requirements are Go and ANTLR4 (which includes `jdk-openjdk`).


## Install Go and ANTLR4 on Arch Linux

Update the system:

```
sudo pacman -Syu
```

Install dependencies:

```
sudo pacman -S antlr4 go
```

Check the installed version of ANTLR:
```
pacman -Q antlr4
```

You should see something like: antlr4 4.13.2-1

Check the installed version of Go:
```
go version
```

You should see something like: go version go1.24.4 linux/amd64

## Create and Run a New Project

Create a new project:
```
mkdir Calc
cd Calc
```

Initialize the Go module and add the ANTLR runtime dependency:
```
go mod init Calc
go get github.com/antlr4-go/antlr/v4
```

Create a file named [Calc.g4](https://github.com/luisespino/compilers/blob/main/antlr-go/Calc/Calc.g4) with the following content:
```
grammar Calc;

DIGIT: [0-9]+;
WS: [ \r\n\t]+ -> skip;

// Rules
l : e EOF;

e : e '+' t     # Sum
  | t           # PassT
  ;

t : t '*' f     # Mul
  | f           # PassF
  ;

f : '(' e ')'   # PassE
  | DIGIT       # Digit
  ;

```
This defines a simple arithmetic grammar for parsing expressions like 3*5+4.

Use the following command to generate the Go parser files inside a parser/ subdirectory:

```
antlr4 -Dlanguage=Go Calc.g4 -o parser
```

This will generate the following files in the parser/ folder:
- calc_base_listener.go
- calc_lexer.go
- calc_listener.go
- calc_parser.go

Create a [main.go](https://github.com/luisespino/compilers/blob/main/antlr-go/Calc/main.go) file with the following code (uses a listener-based approach):
```
package main

import (
	"fmt"
	"strconv"

	"Calc/parser"

	"github.com/antlr4-go/antlr/v4"
)

type calcListener struct {
	*parser.BaseCalcListener
	stack []int
}

func (l *calcListener) push(i int) {
	l.stack = append(l.stack, i)
}

func (l *calcListener) pop() int {
	if len(l.stack) < 1 {
		panic("empty stack")
	}
	result := l.stack[len(l.stack)-1]
	l.stack = l.stack[:len(l.stack)-1]
	return result
}

func (l *calcListener) ExitSum(c *parser.SumContext) {
	right, left := l.pop(), l.pop()
	l.push(left + right)
}

func (l *calcListener) ExitMul(c *parser.MulContext) {
	right, left := l.pop(), l.pop()
	l.push(left * right)
}

func (l *calcListener) ExitDigit(c *parser.DigitContext) {
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err.Error())
	}
	l.push(i)
}

func main() {

	// Setup the input
	is := antlr.NewInputStream("3*5+4")

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(stream)

	// Finally parse the expression (by walking the tree)
	var listener calcListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.L())

	fmt.Println(listener.pop())
}
```

Run the following command to clean up the go.mod and add missing imports:
```
go mod tidy
```

Run the project
```
go run main.go
````

Your calculator now successfully evaluates arithmetic expressions like 3*5+4, the expected output is 19.

![Alt text](https://github.com/luisespino/compilers/blob/main/antlr-go/Calc/run.png)
