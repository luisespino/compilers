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
		"", "'+'", "'*'", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "DIGIT", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "DIGIT", "WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 6, 33, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		4, 4, 4, 23, 8, 4, 11, 4, 12, 4, 24, 1, 5, 4, 5, 28, 8, 5, 11, 5, 12, 5,
		29, 1, 5, 1, 5, 0, 0, 6, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 1, 0, 2,
		1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 34, 0, 1, 1, 0, 0, 0, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 1, 13, 1, 0, 0, 0, 3, 15, 1, 0, 0, 0, 5, 17, 1, 0, 0, 0, 7,
		19, 1, 0, 0, 0, 9, 22, 1, 0, 0, 0, 11, 27, 1, 0, 0, 0, 13, 14, 5, 43, 0,
		0, 14, 2, 1, 0, 0, 0, 15, 16, 5, 42, 0, 0, 16, 4, 1, 0, 0, 0, 17, 18, 5,
		40, 0, 0, 18, 6, 1, 0, 0, 0, 19, 20, 5, 41, 0, 0, 20, 8, 1, 0, 0, 0, 21,
		23, 7, 0, 0, 0, 22, 21, 1, 0, 0, 0, 23, 24, 1, 0, 0, 0, 24, 22, 1, 0, 0,
		0, 24, 25, 1, 0, 0, 0, 25, 10, 1, 0, 0, 0, 26, 28, 7, 1, 0, 0, 27, 26,
		1, 0, 0, 0, 28, 29, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0,
		30, 31, 1, 0, 0, 0, 31, 32, 6, 5, 0, 0, 32, 12, 1, 0, 0, 0, 3, 0, 24, 29,
		1, 6, 0, 0,
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
	CalcLexerT__0  = 1
	CalcLexerT__1  = 2
	CalcLexerT__2  = 3
	CalcLexerT__3  = 4
	CalcLexerDIGIT = 5
	CalcLexerWS    = 6
)
