// Code generated from Calc.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 8, 35, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 6, 6, 25, 10,
	6, 13, 6, 14, 6, 26, 3, 7, 6, 7, 30, 10, 7, 13, 7, 14, 7, 31, 3, 7, 3,
	7, 2, 2, 8, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 3, 2, 4, 3, 2, 50, 59,
	5, 2, 11, 12, 15, 15, 34, 34, 2, 36, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2,
	2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2,
	2, 3, 15, 3, 2, 2, 2, 5, 17, 3, 2, 2, 2, 7, 19, 3, 2, 2, 2, 9, 21, 3, 2,
	2, 2, 11, 24, 3, 2, 2, 2, 13, 29, 3, 2, 2, 2, 15, 16, 7, 44, 2, 2, 16,
	4, 3, 2, 2, 2, 17, 18, 7, 45, 2, 2, 18, 6, 3, 2, 2, 2, 19, 20, 7, 42, 2,
	2, 20, 8, 3, 2, 2, 2, 21, 22, 7, 43, 2, 2, 22, 10, 3, 2, 2, 2, 23, 25,
	9, 2, 2, 2, 24, 23, 3, 2, 2, 2, 25, 26, 3, 2, 2, 2, 26, 24, 3, 2, 2, 2,
	26, 27, 3, 2, 2, 2, 27, 12, 3, 2, 2, 2, 28, 30, 9, 3, 2, 2, 29, 28, 3,
	2, 2, 2, 30, 31, 3, 2, 2, 2, 31, 29, 3, 2, 2, 2, 31, 32, 3, 2, 2, 2, 32,
	33, 3, 2, 2, 2, 33, 34, 8, 7, 2, 2, 34, 14, 3, 2, 2, 2, 5, 2, 26, 31, 3,
	8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'*'", "'+'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "MUL", "ADD", "LB", "RB", "DIGIT", "WS",
}

var lexerRuleNames = []string{
	"MUL", "ADD", "LB", "RB", "DIGIT", "WS",
}

type CalcLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewCalcLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *CalcLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewCalcLexer(input antlr.CharStream) *CalcLexer {
	l := new(CalcLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Calc.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CalcLexer tokens.
const (
	CalcLexerMUL   = 1
	CalcLexerADD   = 2
	CalcLexerLB    = 3
	CalcLexerRB    = 4
	CalcLexerDIGIT = 5
	CalcLexerWS    = 6
)
