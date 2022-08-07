// Code generated from Tinylang.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

type TinylangLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var tinylanglexerLexerStaticData struct {
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

func tinylanglexerLexerInit() {
	staticData := &tinylanglexerLexerStaticData
	staticData.channelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.modeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.literalNames = []string{
		"", "'*'", "'/'", "'+'", "'-'", "", "", "','", "'void'", "'int'", "'string'",
		"", "'('", "')'", "'{'", "'}'", "';'", "'='",
	}
	staticData.symbolicNames = []string{
		"", "MUL", "DIV", "ADD", "SUB", "NUMBER", "WHITESPACE", "COMMA", "TYPE_VOID",
		"TYPE_INT", "TYPE_STRING", "IDENT", "PARENTHESIS_OPEN", "PARENTHESIS_CLOSE",
		"CURLY_BRACKET_OPEN", "CURLY_BRACKET_CLOSE", "SEMICOLON", "EQUALS",
	}
	staticData.ruleNames = []string{
		"MUL", "DIV", "ADD", "SUB", "NUMBER", "WHITESPACE", "COMMA", "TYPE_VOID",
		"TYPE_INT", "TYPE_STRING", "IDENT", "PARENTHESIS_OPEN", "PARENTHESIS_CLOSE",
		"CURLY_BRACKET_OPEN", "CURLY_BRACKET_CLOSE", "SEMICOLON", "EQUALS",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 17, 92, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		4, 4, 4, 45, 8, 4, 11, 4, 12, 4, 46, 1, 5, 4, 5, 50, 8, 5, 11, 5, 12, 5,
		51, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 5, 10,
		76, 8, 10, 10, 10, 12, 10, 79, 9, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13,
		1, 13, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 0, 0, 17, 1, 1, 3, 2,
		5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 1, 0, 4, 1, 0, 48, 57, 3, 0, 9, 10,
		13, 13, 32, 32, 2, 0, 65, 90, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97,
		122, 94, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 1, 35, 1, 0, 0, 0, 3, 37, 1, 0, 0,
		0, 5, 39, 1, 0, 0, 0, 7, 41, 1, 0, 0, 0, 9, 44, 1, 0, 0, 0, 11, 49, 1,
		0, 0, 0, 13, 55, 1, 0, 0, 0, 15, 57, 1, 0, 0, 0, 17, 62, 1, 0, 0, 0, 19,
		66, 1, 0, 0, 0, 21, 73, 1, 0, 0, 0, 23, 80, 1, 0, 0, 0, 25, 82, 1, 0, 0,
		0, 27, 84, 1, 0, 0, 0, 29, 86, 1, 0, 0, 0, 31, 88, 1, 0, 0, 0, 33, 90,
		1, 0, 0, 0, 35, 36, 5, 42, 0, 0, 36, 2, 1, 0, 0, 0, 37, 38, 5, 47, 0, 0,
		38, 4, 1, 0, 0, 0, 39, 40, 5, 43, 0, 0, 40, 6, 1, 0, 0, 0, 41, 42, 5, 45,
		0, 0, 42, 8, 1, 0, 0, 0, 43, 45, 7, 0, 0, 0, 44, 43, 1, 0, 0, 0, 45, 46,
		1, 0, 0, 0, 46, 44, 1, 0, 0, 0, 46, 47, 1, 0, 0, 0, 47, 10, 1, 0, 0, 0,
		48, 50, 7, 1, 0, 0, 49, 48, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 49, 1,
		0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 53, 1, 0, 0, 0, 53, 54, 6, 5, 0, 0, 54,
		12, 1, 0, 0, 0, 55, 56, 5, 44, 0, 0, 56, 14, 1, 0, 0, 0, 57, 58, 5, 118,
		0, 0, 58, 59, 5, 111, 0, 0, 59, 60, 5, 105, 0, 0, 60, 61, 5, 100, 0, 0,
		61, 16, 1, 0, 0, 0, 62, 63, 5, 105, 0, 0, 63, 64, 5, 110, 0, 0, 64, 65,
		5, 116, 0, 0, 65, 18, 1, 0, 0, 0, 66, 67, 5, 115, 0, 0, 67, 68, 5, 116,
		0, 0, 68, 69, 5, 114, 0, 0, 69, 70, 5, 105, 0, 0, 70, 71, 5, 110, 0, 0,
		71, 72, 5, 103, 0, 0, 72, 20, 1, 0, 0, 0, 73, 77, 7, 2, 0, 0, 74, 76, 7,
		3, 0, 0, 75, 74, 1, 0, 0, 0, 76, 79, 1, 0, 0, 0, 77, 75, 1, 0, 0, 0, 77,
		78, 1, 0, 0, 0, 78, 22, 1, 0, 0, 0, 79, 77, 1, 0, 0, 0, 80, 81, 5, 40,
		0, 0, 81, 24, 1, 0, 0, 0, 82, 83, 5, 41, 0, 0, 83, 26, 1, 0, 0, 0, 84,
		85, 5, 123, 0, 0, 85, 28, 1, 0, 0, 0, 86, 87, 5, 125, 0, 0, 87, 30, 1,
		0, 0, 0, 88, 89, 5, 59, 0, 0, 89, 32, 1, 0, 0, 0, 90, 91, 5, 61, 0, 0,
		91, 34, 1, 0, 0, 0, 4, 0, 46, 51, 77, 1, 6, 0, 0,
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

// TinylangLexerInit initializes any static state used to implement TinylangLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewTinylangLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func TinylangLexerInit() {
	staticData := &tinylanglexerLexerStaticData
	staticData.once.Do(tinylanglexerLexerInit)
}

// NewTinylangLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewTinylangLexer(input antlr.CharStream) *TinylangLexer {
	TinylangLexerInit()
	l := new(TinylangLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &tinylanglexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	l.channelNames = staticData.channelNames
	l.modeNames = staticData.modeNames
	l.RuleNames = staticData.ruleNames
	l.LiteralNames = staticData.literalNames
	l.SymbolicNames = staticData.symbolicNames
	l.GrammarFileName = "Tinylang.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// TinylangLexer tokens.
const (
	TinylangLexerMUL                 = 1
	TinylangLexerDIV                 = 2
	TinylangLexerADD                 = 3
	TinylangLexerSUB                 = 4
	TinylangLexerNUMBER              = 5
	TinylangLexerWHITESPACE          = 6
	TinylangLexerCOMMA               = 7
	TinylangLexerTYPE_VOID           = 8
	TinylangLexerTYPE_INT            = 9
	TinylangLexerTYPE_STRING         = 10
	TinylangLexerIDENT               = 11
	TinylangLexerPARENTHESIS_OPEN    = 12
	TinylangLexerPARENTHESIS_CLOSE   = 13
	TinylangLexerCURLY_BRACKET_OPEN  = 14
	TinylangLexerCURLY_BRACKET_CLOSE = 15
	TinylangLexerSEMICOLON           = 16
	TinylangLexerEQUALS              = 17
)