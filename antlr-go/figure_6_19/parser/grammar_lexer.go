// Code generated from Grammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

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

type GrammarLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var GrammarLexerLexerStaticData struct {
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

func grammarlexerLexerInit() {
	staticData := &GrammarLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'='", "'*'", "'/'", "'+'", "'-'", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "ID", "NUM", "WS",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "ID", "NUM",
		"WS",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 10, 54, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		6, 1, 6, 1, 7, 1, 7, 5, 7, 38, 8, 7, 10, 7, 12, 7, 41, 9, 7, 1, 8, 4, 8,
		44, 8, 8, 11, 8, 12, 8, 45, 1, 9, 4, 9, 49, 8, 9, 11, 9, 12, 9, 50, 1,
		9, 1, 9, 0, 0, 10, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17,
		9, 19, 10, 1, 0, 4, 3, 0, 65, 90, 95, 95, 97, 122, 4, 0, 48, 57, 65, 90,
		95, 95, 97, 122, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 56, 0, 1, 1,
		0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1,
		0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17,
		1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 1, 21, 1, 0, 0, 0, 3, 23, 1, 0, 0, 0, 5,
		25, 1, 0, 0, 0, 7, 27, 1, 0, 0, 0, 9, 29, 1, 0, 0, 0, 11, 31, 1, 0, 0,
		0, 13, 33, 1, 0, 0, 0, 15, 35, 1, 0, 0, 0, 17, 43, 1, 0, 0, 0, 19, 48,
		1, 0, 0, 0, 21, 22, 5, 61, 0, 0, 22, 2, 1, 0, 0, 0, 23, 24, 5, 42, 0, 0,
		24, 4, 1, 0, 0, 0, 25, 26, 5, 47, 0, 0, 26, 6, 1, 0, 0, 0, 27, 28, 5, 43,
		0, 0, 28, 8, 1, 0, 0, 0, 29, 30, 5, 45, 0, 0, 30, 10, 1, 0, 0, 0, 31, 32,
		5, 40, 0, 0, 32, 12, 1, 0, 0, 0, 33, 34, 5, 41, 0, 0, 34, 14, 1, 0, 0,
		0, 35, 39, 7, 0, 0, 0, 36, 38, 7, 1, 0, 0, 37, 36, 1, 0, 0, 0, 38, 41,
		1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 16, 1, 0, 0, 0,
		41, 39, 1, 0, 0, 0, 42, 44, 7, 2, 0, 0, 43, 42, 1, 0, 0, 0, 44, 45, 1,
		0, 0, 0, 45, 43, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 18, 1, 0, 0, 0, 47,
		49, 7, 3, 0, 0, 48, 47, 1, 0, 0, 0, 49, 50, 1, 0, 0, 0, 50, 48, 1, 0, 0,
		0, 50, 51, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 53, 6, 9, 0, 0, 53, 20,
		1, 0, 0, 0, 4, 0, 39, 45, 50, 1, 6, 0, 0,
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

// GrammarLexerInit initializes any static state used to implement GrammarLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewGrammarLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func GrammarLexerInit() {
	staticData := &GrammarLexerLexerStaticData
	staticData.once.Do(grammarlexerLexerInit)
}

// NewGrammarLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewGrammarLexer(input antlr.CharStream) *GrammarLexer {
	GrammarLexerInit()
	l := new(GrammarLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &GrammarLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "Grammar.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// GrammarLexer tokens.
const (
	GrammarLexerT__0 = 1
	GrammarLexerT__1 = 2
	GrammarLexerT__2 = 3
	GrammarLexerT__3 = 4
	GrammarLexerT__4 = 5
	GrammarLexerT__5 = 6
	GrammarLexerT__6 = 7
	GrammarLexerID   = 8
	GrammarLexerNUM  = 9
	GrammarLexerWS   = 10
)
