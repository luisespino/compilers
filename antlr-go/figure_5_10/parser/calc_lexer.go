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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 11, 52, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 7, 8, 36, 10, 8,
	12, 8, 14, 8, 39, 11, 8, 3, 9, 6, 9, 42, 10, 9, 13, 9, 14, 9, 43, 3, 10,
	6, 10, 47, 10, 10, 13, 10, 14, 10, 48, 3, 10, 3, 10, 2, 2, 11, 3, 3, 5,
	4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 3, 2, 6, 5, 2, 67,
	92, 97, 97, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59,
	5, 2, 11, 12, 15, 15, 34, 34, 2, 54, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2,
	2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2,
	2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 3, 21, 3, 2,
	2, 2, 5, 23, 3, 2, 2, 2, 7, 25, 3, 2, 2, 2, 9, 27, 3, 2, 2, 2, 11, 29,
	3, 2, 2, 2, 13, 31, 3, 2, 2, 2, 15, 33, 3, 2, 2, 2, 17, 41, 3, 2, 2, 2,
	19, 46, 3, 2, 2, 2, 21, 22, 7, 45, 2, 2, 22, 4, 3, 2, 2, 2, 23, 24, 7,
	47, 2, 2, 24, 6, 3, 2, 2, 2, 25, 26, 7, 44, 2, 2, 26, 8, 3, 2, 2, 2, 27,
	28, 7, 49, 2, 2, 28, 10, 3, 2, 2, 2, 29, 30, 7, 42, 2, 2, 30, 12, 3, 2,
	2, 2, 31, 32, 7, 43, 2, 2, 32, 14, 3, 2, 2, 2, 33, 37, 9, 2, 2, 2, 34,
	36, 9, 3, 2, 2, 35, 34, 3, 2, 2, 2, 36, 39, 3, 2, 2, 2, 37, 35, 3, 2, 2,
	2, 37, 38, 3, 2, 2, 2, 38, 16, 3, 2, 2, 2, 39, 37, 3, 2, 2, 2, 40, 42,
	9, 4, 2, 2, 41, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2,
	43, 44, 3, 2, 2, 2, 44, 18, 3, 2, 2, 2, 45, 47, 9, 5, 2, 2, 46, 45, 3,
	2, 2, 2, 47, 48, 3, 2, 2, 2, 48, 46, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49,
	50, 3, 2, 2, 2, 50, 51, 8, 10, 2, 2, 51, 20, 3, 2, 2, 2, 6, 2, 37, 43,
	48, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'+'", "'-'", "'*'", "'/'", "'('", "')'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "ID", "NUM", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "ID", "NUM", "WS",
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
	CalcLexerT__0 = 1
	CalcLexerT__1 = 2
	CalcLexerT__2 = 3
	CalcLexerT__3 = 4
	CalcLexerT__4 = 5
	CalcLexerT__5 = 6
	CalcLexerID   = 7
	CalcLexerNUM  = 8
	CalcLexerWS   = 9
)
