// Code generated from Calc.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type CalcLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var CalcLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func calclexerLexerInit() {
	staticData := &CalcLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'+'", "'-'", "'*'", "'/'", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "ID", "NUM", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "ID", "NUM", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 9, 50, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 5,
		6, 34, 8, 6, 10, 6, 12, 6, 37, 9, 6, 1, 7, 4, 7, 40, 8, 7, 11, 7, 12, 7,
		41, 1, 8, 4, 8, 45, 8, 8, 11, 8, 12, 8, 46, 1, 8, 1, 8, 0, 0, 9, 1, 1,
		3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 1, 0, 4, 3, 0, 65,
		90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 48, 57,
		3, 0, 9, 10, 13, 13, 32, 32, 52, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 1, 19, 1, 0, 0, 0,
		3, 21, 1, 0, 0, 0, 5, 23, 1, 0, 0, 0, 7, 25, 1, 0, 0, 0, 9, 27, 1, 0, 0,
		0, 11, 29, 1, 0, 0, 0, 13, 31, 1, 0, 0, 0, 15, 39, 1, 0, 0, 0, 17, 44,
		1, 0, 0, 0, 19, 20, 5, 43, 0, 0, 20, 2, 1, 0, 0, 0, 21, 22, 5, 45, 0, 0,
		22, 4, 1, 0, 0, 0, 23, 24, 5, 42, 0, 0, 24, 6, 1, 0, 0, 0, 25, 26, 5, 47,
		0, 0, 26, 8, 1, 0, 0, 0, 27, 28, 5, 40, 0, 0, 28, 10, 1, 0, 0, 0, 29, 30,
		5, 41, 0, 0, 30, 12, 1, 0, 0, 0, 31, 35, 7, 0, 0, 0, 32, 34, 7, 1, 0, 0,
		33, 32, 1, 0, 0, 0, 34, 37, 1, 0, 0, 0, 35, 33, 1, 0, 0, 0, 35, 36, 1,
		0, 0, 0, 36, 14, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 38, 40, 7, 2, 0, 0, 39,
		38, 1, 0, 0, 0, 40, 41, 1, 0, 0, 0, 41, 39, 1, 0, 0, 0, 41, 42, 1, 0, 0,
		0, 42, 16, 1, 0, 0, 0, 43, 45, 7, 3, 0, 0, 44, 43, 1, 0, 0, 0, 45, 46,
		1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 48, 1, 0, 0, 0,
		48, 49, 6, 8, 0, 0, 49, 18, 1, 0, 0, 0, 4, 0, 35, 41, 46, 1, 6, 0, 0,
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

// CalcLexerInit initializes any static state used to implement CalcLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewCalcLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func CalcLexerInit() {
	staticData := &CalcLexerLexerStaticData
	staticData.once.Do(calclexerLexerInit)
}

// NewCalcLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCalcLexer(input antlr.CharStream) *CalcLexer {
	CalcLexerInit()
	l := new(CalcLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &CalcLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
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
