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
		"", "", "'*'", "'/'", "'+'", "'-'", "'%'", "'('", "')'", "'{'", "'}'",
		"';'", "'=='", "'>'", "'>='", "'<'", "'<='", "'='", "','", "'&&'", "'||'",
		"", "'void'", "'int'", "'string'", "'boolean'", "'true'", "'false'",
	}
	staticData.symbolicNames = []string{
		"", "WHITESPACE", "MUL", "DIV", "ADD", "SUB", "MOD", "PARENTHESIS_OPEN",
		"PARENTHESIS_CLOSE", "CURLY_BRACKET_OPEN", "CURLY_BRACKET_CLOSE", "SEMICOLON",
		"EQUALS_TEST", "GREATER_THAN", "GREATER_OR_EQUALS", "LESS_THAN", "LESS_OR_EQUALS",
		"EQUALS", "COMMA", "AND_TEST", "OR_TEST", "NUMBER", "TYPE_VOID", "TYPE_INT",
		"TYPE_STRING", "TYPE_BOOLEAN", "VALUE_TRUE", "VALUE_FALSE", "IDENT",
		"STRING",
	}
	staticData.ruleNames = []string{
		"WHITESPACE", "MUL", "DIV", "ADD", "SUB", "MOD", "PARENTHESIS_OPEN",
		"PARENTHESIS_CLOSE", "CURLY_BRACKET_OPEN", "CURLY_BRACKET_CLOSE", "SEMICOLON",
		"EQUALS_TEST", "GREATER_THAN", "GREATER_OR_EQUALS", "LESS_THAN", "LESS_OR_EQUALS",
		"EQUALS", "COMMA", "AND_TEST", "OR_TEST", "NUMBER", "TYPE_VOID", "TYPE_INT",
		"TYPE_STRING", "TYPE_BOOLEAN", "VALUE_TRUE", "VALUE_FALSE", "IDENT",
		"STRING",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 29, 165, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 1, 0, 4, 0, 61, 8, 0, 11, 0,
		12, 0, 62, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4,
		1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1,
		15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 19,
		1, 19, 1, 19, 1, 20, 4, 20, 111, 8, 20, 11, 20, 12, 20, 112, 1, 21, 1,
		21, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22, 1, 22, 1, 23, 1, 23, 1, 23,
		1, 23, 1, 23, 1, 23, 1, 23, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1, 24, 1,
		24, 1, 24, 1, 25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26,
		1, 26, 1, 26, 1, 27, 1, 27, 5, 27, 152, 8, 27, 10, 27, 12, 27, 155, 9,
		27, 1, 28, 1, 28, 5, 28, 159, 8, 28, 10, 28, 12, 28, 162, 9, 28, 1, 28,
		1, 28, 0, 0, 29, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17,
		9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35,
		18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53,
		27, 55, 28, 57, 29, 1, 0, 5, 3, 0, 9, 10, 13, 13, 32, 32, 1, 0, 48, 57,
		2, 0, 65, 90, 97, 122, 4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 1, 0, 34,
		34, 168, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1,
		0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15,
		1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0,
		23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0,
		0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0,
		0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1,
		0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57, 1, 0, 0, 0, 1, 60, 1, 0, 0, 0, 3, 66,
		1, 0, 0, 0, 5, 68, 1, 0, 0, 0, 7, 70, 1, 0, 0, 0, 9, 72, 1, 0, 0, 0, 11,
		74, 1, 0, 0, 0, 13, 76, 1, 0, 0, 0, 15, 78, 1, 0, 0, 0, 17, 80, 1, 0, 0,
		0, 19, 82, 1, 0, 0, 0, 21, 84, 1, 0, 0, 0, 23, 86, 1, 0, 0, 0, 25, 89,
		1, 0, 0, 0, 27, 91, 1, 0, 0, 0, 29, 94, 1, 0, 0, 0, 31, 96, 1, 0, 0, 0,
		33, 99, 1, 0, 0, 0, 35, 101, 1, 0, 0, 0, 37, 103, 1, 0, 0, 0, 39, 106,
		1, 0, 0, 0, 41, 110, 1, 0, 0, 0, 43, 114, 1, 0, 0, 0, 45, 119, 1, 0, 0,
		0, 47, 123, 1, 0, 0, 0, 49, 130, 1, 0, 0, 0, 51, 138, 1, 0, 0, 0, 53, 143,
		1, 0, 0, 0, 55, 149, 1, 0, 0, 0, 57, 156, 1, 0, 0, 0, 59, 61, 7, 0, 0,
		0, 60, 59, 1, 0, 0, 0, 61, 62, 1, 0, 0, 0, 62, 60, 1, 0, 0, 0, 62, 63,
		1, 0, 0, 0, 63, 64, 1, 0, 0, 0, 64, 65, 6, 0, 0, 0, 65, 2, 1, 0, 0, 0,
		66, 67, 5, 42, 0, 0, 67, 4, 1, 0, 0, 0, 68, 69, 5, 47, 0, 0, 69, 6, 1,
		0, 0, 0, 70, 71, 5, 43, 0, 0, 71, 8, 1, 0, 0, 0, 72, 73, 5, 45, 0, 0, 73,
		10, 1, 0, 0, 0, 74, 75, 5, 37, 0, 0, 75, 12, 1, 0, 0, 0, 76, 77, 5, 40,
		0, 0, 77, 14, 1, 0, 0, 0, 78, 79, 5, 41, 0, 0, 79, 16, 1, 0, 0, 0, 80,
		81, 5, 123, 0, 0, 81, 18, 1, 0, 0, 0, 82, 83, 5, 125, 0, 0, 83, 20, 1,
		0, 0, 0, 84, 85, 5, 59, 0, 0, 85, 22, 1, 0, 0, 0, 86, 87, 5, 61, 0, 0,
		87, 88, 5, 61, 0, 0, 88, 24, 1, 0, 0, 0, 89, 90, 5, 62, 0, 0, 90, 26, 1,
		0, 0, 0, 91, 92, 5, 62, 0, 0, 92, 93, 5, 61, 0, 0, 93, 28, 1, 0, 0, 0,
		94, 95, 5, 60, 0, 0, 95, 30, 1, 0, 0, 0, 96, 97, 5, 60, 0, 0, 97, 98, 5,
		61, 0, 0, 98, 32, 1, 0, 0, 0, 99, 100, 5, 61, 0, 0, 100, 34, 1, 0, 0, 0,
		101, 102, 5, 44, 0, 0, 102, 36, 1, 0, 0, 0, 103, 104, 5, 38, 0, 0, 104,
		105, 5, 38, 0, 0, 105, 38, 1, 0, 0, 0, 106, 107, 5, 124, 0, 0, 107, 108,
		5, 124, 0, 0, 108, 40, 1, 0, 0, 0, 109, 111, 7, 1, 0, 0, 110, 109, 1, 0,
		0, 0, 111, 112, 1, 0, 0, 0, 112, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0,
		113, 42, 1, 0, 0, 0, 114, 115, 5, 118, 0, 0, 115, 116, 5, 111, 0, 0, 116,
		117, 5, 105, 0, 0, 117, 118, 5, 100, 0, 0, 118, 44, 1, 0, 0, 0, 119, 120,
		5, 105, 0, 0, 120, 121, 5, 110, 0, 0, 121, 122, 5, 116, 0, 0, 122, 46,
		1, 0, 0, 0, 123, 124, 5, 115, 0, 0, 124, 125, 5, 116, 0, 0, 125, 126, 5,
		114, 0, 0, 126, 127, 5, 105, 0, 0, 127, 128, 5, 110, 0, 0, 128, 129, 5,
		103, 0, 0, 129, 48, 1, 0, 0, 0, 130, 131, 5, 98, 0, 0, 131, 132, 5, 111,
		0, 0, 132, 133, 5, 111, 0, 0, 133, 134, 5, 108, 0, 0, 134, 135, 5, 101,
		0, 0, 135, 136, 5, 97, 0, 0, 136, 137, 5, 110, 0, 0, 137, 50, 1, 0, 0,
		0, 138, 139, 5, 116, 0, 0, 139, 140, 5, 114, 0, 0, 140, 141, 5, 117, 0,
		0, 141, 142, 5, 101, 0, 0, 142, 52, 1, 0, 0, 0, 143, 144, 5, 102, 0, 0,
		144, 145, 5, 97, 0, 0, 145, 146, 5, 108, 0, 0, 146, 147, 5, 115, 0, 0,
		147, 148, 5, 101, 0, 0, 148, 54, 1, 0, 0, 0, 149, 153, 7, 2, 0, 0, 150,
		152, 7, 3, 0, 0, 151, 150, 1, 0, 0, 0, 152, 155, 1, 0, 0, 0, 153, 151,
		1, 0, 0, 0, 153, 154, 1, 0, 0, 0, 154, 56, 1, 0, 0, 0, 155, 153, 1, 0,
		0, 0, 156, 160, 5, 34, 0, 0, 157, 159, 8, 4, 0, 0, 158, 157, 1, 0, 0, 0,
		159, 162, 1, 0, 0, 0, 160, 158, 1, 0, 0, 0, 160, 161, 1, 0, 0, 0, 161,
		163, 1, 0, 0, 0, 162, 160, 1, 0, 0, 0, 163, 164, 5, 34, 0, 0, 164, 58,
		1, 0, 0, 0, 5, 0, 62, 112, 153, 160, 1, 6, 0, 0,
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
	TinylangLexerMOD                 = 6
	TinylangLexerPARENTHESIS_OPEN    = 7
	TinylangLexerPARENTHESIS_CLOSE   = 8
	TinylangLexerCURLY_BRACKET_OPEN  = 9
	TinylangLexerCURLY_BRACKET_CLOSE = 10
	TinylangLexerSEMICOLON           = 11
	TinylangLexerEQUALS_TEST         = 12
	TinylangLexerGREATER_THAN        = 13
	TinylangLexerGREATER_OR_EQUALS   = 14
	TinylangLexerLESS_THAN           = 15
	TinylangLexerLESS_OR_EQUALS      = 16
	TinylangLexerEQUALS              = 17
	TinylangLexerCOMMA               = 18
	TinylangLexerAND_TEST            = 19
	TinylangLexerOR_TEST             = 20
	TinylangLexerNUMBER              = 21
	TinylangLexerTYPE_VOID           = 22
	TinylangLexerTYPE_INT            = 23
	TinylangLexerTYPE_STRING         = 24
	TinylangLexerTYPE_BOOLEAN        = 25
	TinylangLexerVALUE_TRUE          = 26
	TinylangLexerVALUE_FALSE         = 27
	TinylangLexerIDENT               = 28
	TinylangLexerSTRING              = 29
)
