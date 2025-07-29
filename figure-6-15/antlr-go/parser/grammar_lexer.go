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
		"", "'int'", "'='", "';'", "'['", "']'", "','",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "WS", "NAME", "DIGIT",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "WS", "NAME", "DIGIT",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 9, 50, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1,
		6, 4, 6, 35, 8, 6, 11, 6, 12, 6, 36, 1, 6, 1, 6, 1, 7, 4, 7, 42, 8, 7,
		11, 7, 12, 7, 43, 1, 8, 4, 8, 47, 8, 8, 11, 8, 12, 8, 48, 0, 0, 9, 1, 1,
		3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 1, 0, 3, 3, 0, 9, 10,
		13, 13, 32, 32, 2, 0, 65, 90, 97, 122, 1, 0, 48, 57, 52, 0, 1, 1, 0, 0,
		0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0,
		0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0,
		0, 0, 1, 19, 1, 0, 0, 0, 3, 23, 1, 0, 0, 0, 5, 25, 1, 0, 0, 0, 7, 27, 1,
		0, 0, 0, 9, 29, 1, 0, 0, 0, 11, 31, 1, 0, 0, 0, 13, 34, 1, 0, 0, 0, 15,
		41, 1, 0, 0, 0, 17, 46, 1, 0, 0, 0, 19, 20, 5, 105, 0, 0, 20, 21, 5, 110,
		0, 0, 21, 22, 5, 116, 0, 0, 22, 2, 1, 0, 0, 0, 23, 24, 5, 61, 0, 0, 24,
		4, 1, 0, 0, 0, 25, 26, 5, 59, 0, 0, 26, 6, 1, 0, 0, 0, 27, 28, 5, 91, 0,
		0, 28, 8, 1, 0, 0, 0, 29, 30, 5, 93, 0, 0, 30, 10, 1, 0, 0, 0, 31, 32,
		5, 44, 0, 0, 32, 12, 1, 0, 0, 0, 33, 35, 7, 0, 0, 0, 34, 33, 1, 0, 0, 0,
		35, 36, 1, 0, 0, 0, 36, 34, 1, 0, 0, 0, 36, 37, 1, 0, 0, 0, 37, 38, 1,
		0, 0, 0, 38, 39, 6, 6, 0, 0, 39, 14, 1, 0, 0, 0, 40, 42, 7, 1, 0, 0, 41,
		40, 1, 0, 0, 0, 42, 43, 1, 0, 0, 0, 43, 41, 1, 0, 0, 0, 43, 44, 1, 0, 0,
		0, 44, 16, 1, 0, 0, 0, 45, 47, 7, 2, 0, 0, 46, 45, 1, 0, 0, 0, 47, 48,
		1, 0, 0, 0, 48, 46, 1, 0, 0, 0, 48, 49, 1, 0, 0, 0, 49, 18, 1, 0, 0, 0,
		4, 0, 36, 43, 48, 1, 6, 0, 0,
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
	GrammarLexerT__0  = 1
	GrammarLexerT__1  = 2
	GrammarLexerT__2  = 3
	GrammarLexerT__3  = 4
	GrammarLexerT__4  = 5
	GrammarLexerT__5  = 6
	GrammarLexerWS    = 7
	GrammarLexerNAME  = 8
	GrammarLexerDIGIT = 9
)
