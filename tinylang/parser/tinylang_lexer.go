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
		"", "", "'*'", "'/'", "'+'", "'-'", "'('", "')'", "'{'", "'}'", "';'",
		"'=='", "'>'", "'>='", "'<'", "'<='", "'='", "','", "", "'void'", "'int'",
		"'string'", "'boolean'", "'true'", "'false'",
	}
	staticData.symbolicNames = []string{
		"", "WHITESPACE", "MUL", "DIV", "ADD", "SUB", "PARENTHESIS_OPEN", "PARENTHESIS_CLOSE",
		"CURLY_BRACKET_OPEN", "CURLY_BRACKET_CLOSE", "SEMICOLON", "EQUALS_TEST",
		"GREATER_THAN", "GREATER_OR_EQUALS", "LESS_THAN", "LESS_OR_EQUALS",
		"EQUALS", "COMMA", "NUMBER", "TYPE_VOID", "TYPE_INT", "TYPE_STRING",
		"TYPE_BOOLEAN", "VALUE_TRUE", "VALUE_FALSE", "IDENT", "STRING",
	}
	staticData.ruleNames = []string{
		"WHITESPACE", "MUL", "DIV", "ADD", "SUB", "PARENTHESIS_OPEN", "PARENTHESIS_CLOSE",
		"CURLY_BRACKET_OPEN", "CURLY_BRACKET_CLOSE", "SEMICOLON", "EQUALS_TEST",
		"GREATER_THAN", "GREATER_OR_EQUALS", "LESS_THAN", "LESS_OR_EQUALS",
		"EQUALS", "COMMA", "NUMBER", "TYPE_VOID", "TYPE_INT", "TYPE_STRING",
		"TYPE_BOOLEAN", "VALUE_TRUE", "VALUE_FALSE", "IDENT", "STRING",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 26, 151, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		1, 0, 4, 0, 55, 8, 0, 11, 0, 12, 0, 56, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1,
		8, 1, 9, 1, 9, 1, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12,
		1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 4,
		17, 97, 8, 17, 11, 17, 12, 17, 98, 1, 18, 1, 18, 1, 18, 1, 18, 1, 18, 1,
		19, 1, 19, 1, 19, 1, 19, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20, 1, 20,
		1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1,
		22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24,
		5, 24, 138, 8, 24, 10, 24, 12, 24, 141, 9, 24, 1, 25, 1, 25, 5, 25, 145,
		8, 25, 10, 25, 12, 25, 148, 9, 25, 1, 25, 1, 25, 0, 0, 26, 1, 1, 3, 2,
		5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25,
		13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19, 39, 20, 41, 21, 43,
		22, 45, 23, 47, 24, 49, 25, 51, 26, 1, 0, 5, 3, 0, 9, 10, 13, 13, 32, 32,
		1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97,
		122, 1, 0, 34, 34, 154, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0,
		0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1,
		0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21,
		1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0,
		29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0,
		0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0,
		0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0,
		0, 0, 1, 54, 1, 0, 0, 0, 3, 60, 1, 0, 0, 0, 5, 62, 1, 0, 0, 0, 7, 64, 1,
		0, 0, 0, 9, 66, 1, 0, 0, 0, 11, 68, 1, 0, 0, 0, 13, 70, 1, 0, 0, 0, 15,
		72, 1, 0, 0, 0, 17, 74, 1, 0, 0, 0, 19, 76, 1, 0, 0, 0, 21, 78, 1, 0, 0,
		0, 23, 81, 1, 0, 0, 0, 25, 83, 1, 0, 0, 0, 27, 86, 1, 0, 0, 0, 29, 88,
		1, 0, 0, 0, 31, 91, 1, 0, 0, 0, 33, 93, 1, 0, 0, 0, 35, 96, 1, 0, 0, 0,
		37, 100, 1, 0, 0, 0, 39, 105, 1, 0, 0, 0, 41, 109, 1, 0, 0, 0, 43, 116,
		1, 0, 0, 0, 45, 124, 1, 0, 0, 0, 47, 129, 1, 0, 0, 0, 49, 135, 1, 0, 0,
		0, 51, 142, 1, 0, 0, 0, 53, 55, 7, 0, 0, 0, 54, 53, 1, 0, 0, 0, 55, 56,
		1, 0, 0, 0, 56, 54, 1, 0, 0, 0, 56, 57, 1, 0, 0, 0, 57, 58, 1, 0, 0, 0,
		58, 59, 6, 0, 0, 0, 59, 2, 1, 0, 0, 0, 60, 61, 5, 42, 0, 0, 61, 4, 1, 0,
		0, 0, 62, 63, 5, 47, 0, 0, 63, 6, 1, 0, 0, 0, 64, 65, 5, 43, 0, 0, 65,
		8, 1, 0, 0, 0, 66, 67, 5, 45, 0, 0, 67, 10, 1, 0, 0, 0, 68, 69, 5, 40,
		0, 0, 69, 12, 1, 0, 0, 0, 70, 71, 5, 41, 0, 0, 71, 14, 1, 0, 0, 0, 72,
		73, 5, 123, 0, 0, 73, 16, 1, 0, 0, 0, 74, 75, 5, 125, 0, 0, 75, 18, 1,
		0, 0, 0, 76, 77, 5, 59, 0, 0, 77, 20, 1, 0, 0, 0, 78, 79, 5, 61, 0, 0,
		79, 80, 5, 61, 0, 0, 80, 22, 1, 0, 0, 0, 81, 82, 5, 62, 0, 0, 82, 24, 1,
		0, 0, 0, 83, 84, 5, 62, 0, 0, 84, 85, 5, 61, 0, 0, 85, 26, 1, 0, 0, 0,
		86, 87, 5, 60, 0, 0, 87, 28, 1, 0, 0, 0, 88, 89, 5, 60, 0, 0, 89, 90, 5,
		61, 0, 0, 90, 30, 1, 0, 0, 0, 91, 92, 5, 61, 0, 0, 92, 32, 1, 0, 0, 0,
		93, 94, 5, 44, 0, 0, 94, 34, 1, 0, 0, 0, 95, 97, 7, 1, 0, 0, 96, 95, 1,
		0, 0, 0, 97, 98, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 98, 99, 1, 0, 0, 0, 99,
		36, 1, 0, 0, 0, 100, 101, 5, 118, 0, 0, 101, 102, 5, 111, 0, 0, 102, 103,
		5, 105, 0, 0, 103, 104, 5, 100, 0, 0, 104, 38, 1, 0, 0, 0, 105, 106, 5,
		105, 0, 0, 106, 107, 5, 110, 0, 0, 107, 108, 5, 116, 0, 0, 108, 40, 1,
		0, 0, 0, 109, 110, 5, 115, 0, 0, 110, 111, 5, 116, 0, 0, 111, 112, 5, 114,
		0, 0, 112, 113, 5, 105, 0, 0, 113, 114, 5, 110, 0, 0, 114, 115, 5, 103,
		0, 0, 115, 42, 1, 0, 0, 0, 116, 117, 5, 98, 0, 0, 117, 118, 5, 111, 0,
		0, 118, 119, 5, 111, 0, 0, 119, 120, 5, 108, 0, 0, 120, 121, 5, 101, 0,
		0, 121, 122, 5, 97, 0, 0, 122, 123, 5, 110, 0, 0, 123, 44, 1, 0, 0, 0,
		124, 125, 5, 116, 0, 0, 125, 126, 5, 114, 0, 0, 126, 127, 5, 117, 0, 0,
		127, 128, 5, 101, 0, 0, 128, 46, 1, 0, 0, 0, 129, 130, 5, 102, 0, 0, 130,
		131, 5, 97, 0, 0, 131, 132, 5, 108, 0, 0, 132, 133, 5, 115, 0, 0, 133,
		134, 5, 101, 0, 0, 134, 48, 1, 0, 0, 0, 135, 139, 7, 2, 0, 0, 136, 138,
		7, 3, 0, 0, 137, 136, 1, 0, 0, 0, 138, 141, 1, 0, 0, 0, 139, 137, 1, 0,
		0, 0, 139, 140, 1, 0, 0, 0, 140, 50, 1, 0, 0, 0, 141, 139, 1, 0, 0, 0,
		142, 146, 5, 34, 0, 0, 143, 145, 8, 4, 0, 0, 144, 143, 1, 0, 0, 0, 145,
		148, 1, 0, 0, 0, 146, 144, 1, 0, 0, 0, 146, 147, 1, 0, 0, 0, 147, 149,
		1, 0, 0, 0, 148, 146, 1, 0, 0, 0, 149, 150, 5, 34, 0, 0, 150, 52, 1, 0,
		0, 0, 5, 0, 56, 98, 139, 146, 1, 6, 0, 0,
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
	TinylangLexerWHITESPACE          = 1
	TinylangLexerMUL                 = 2
	TinylangLexerDIV                 = 3
	TinylangLexerADD                 = 4
	TinylangLexerSUB                 = 5
	TinylangLexerPARENTHESIS_OPEN    = 6
	TinylangLexerPARENTHESIS_CLOSE   = 7
	TinylangLexerCURLY_BRACKET_OPEN  = 8
	TinylangLexerCURLY_BRACKET_CLOSE = 9
	TinylangLexerSEMICOLON           = 10
	TinylangLexerEQUALS_TEST         = 11
	TinylangLexerGREATER_THAN        = 12
	TinylangLexerGREATER_OR_EQUALS   = 13
	TinylangLexerLESS_THAN           = 14
	TinylangLexerLESS_OR_EQUALS      = 15
	TinylangLexerEQUALS              = 16
	TinylangLexerCOMMA               = 17
	TinylangLexerNUMBER              = 18
	TinylangLexerTYPE_VOID           = 19
	TinylangLexerTYPE_INT            = 20
	TinylangLexerTYPE_STRING         = 21
	TinylangLexerTYPE_BOOLEAN        = 22
	TinylangLexerVALUE_TRUE          = 23
	TinylangLexerVALUE_FALSE         = 24
	TinylangLexerIDENT               = 25
	TinylangLexerSTRING              = 26
)
