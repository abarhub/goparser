// Code generated from Calc.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser

import (
	"fmt"
	"sync"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
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

var calclexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	channelNames           []string
	modeNames              []string
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func calclexerLexerInit() {
	staticData := &calclexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'('", "')'", "'*'", "'/'", "'+'", "'-'", "", "", "','",
	}
	staticData.symbolicNames = []string{
		"", "", "", "MUL", "DIV", "ADD", "SUB", "NUMBER", "WHITESPACE", "COMMA",
	}
	staticData.ruleNames = []string{
		"T__0", "T__1", "MUL", "DIV", "ADD", "SUB", "NUMBER", "WHITESPACE",
		"COMMA",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 9, 45, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 1, 0, 1, 0, 1,
		1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 4, 6, 33,
		8, 6, 11, 6, 12, 6, 34, 1, 7, 4, 7, 38, 8, 7, 11, 7, 12, 7, 39, 1, 7, 1,
		7, 1, 8, 1, 8, 0, 0, 9, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15,
		8, 17, 9, 1, 0, 2, 1, 0, 48, 57, 3, 0, 9, 10, 13, 13, 32, 32, 46, 0, 1,
		1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9,
		1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0,
		17, 1, 0, 0, 0, 1, 19, 1, 0, 0, 0, 3, 21, 1, 0, 0, 0, 5, 23, 1, 0, 0, 0,
		7, 25, 1, 0, 0, 0, 9, 27, 1, 0, 0, 0, 11, 29, 1, 0, 0, 0, 13, 32, 1, 0,
		0, 0, 15, 37, 1, 0, 0, 0, 17, 43, 1, 0, 0, 0, 19, 20, 5, 40, 0, 0, 20,
		2, 1, 0, 0, 0, 21, 22, 5, 41, 0, 0, 22, 4, 1, 0, 0, 0, 23, 24, 5, 42, 0,
		0, 24, 6, 1, 0, 0, 0, 25, 26, 5, 47, 0, 0, 26, 8, 1, 0, 0, 0, 27, 28, 5,
		43, 0, 0, 28, 10, 1, 0, 0, 0, 29, 30, 5, 45, 0, 0, 30, 12, 1, 0, 0, 0,
		31, 33, 7, 0, 0, 0, 32, 31, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 32, 1,
		0, 0, 0, 34, 35, 1, 0, 0, 0, 35, 14, 1, 0, 0, 0, 36, 38, 7, 1, 0, 0, 37,
		36, 1, 0, 0, 0, 38, 39, 1, 0, 0, 0, 39, 37, 1, 0, 0, 0, 39, 40, 1, 0, 0,
		0, 40, 41, 1, 0, 0, 0, 41, 42, 6, 7, 0, 0, 42, 16, 1, 0, 0, 0, 43, 44,
		5, 44, 0, 0, 44, 18, 1, 0, 0, 0, 3, 0, 34, 39, 1, 6, 0, 0,
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
	staticData := &calclexerLexerStaticData
	staticData.once.Do(calclexerLexerInit)
}

// NewCalcLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewCalcLexer(input antlr.CharStream) *CalcLexer {
	CalcLexerInit()
	l := new(CalcLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &calclexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Calc.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CalcLexer tokens.
const (
	CalcLexerT__0       = 1
	CalcLexerT__1       = 2
	CalcLexerMUL        = 3
	CalcLexerDIV        = 4
	CalcLexerADD        = 5
	CalcLexerSUB        = 6
	CalcLexerNUMBER     = 7
	CalcLexerWHITESPACE = 8
	CalcLexerCOMMA      = 9
)
