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
		"", "", "", "", "'*'", "'/'", "'+'", "'-'", "'%'", "'('", "')'", "'{'",
		"'}'", "';'", "'=='", "'!='", "'>'", "'>='", "'<'", "'<='", "'='", "','",
		"'&&'", "'||'", "'!'", "", "'void'", "'int'", "'string'", "'boolean'",
		"'true'", "'false'",
	}
	staticData.symbolicNames = []string{
		"", "WHITESPACE", "LINE_COMMENT", "COMMENT", "MUL", "DIV", "ADD", "SUB",
		"MOD", "PARENTHESIS_OPEN", "PARENTHESIS_CLOSE", "CURLY_BRACKET_OPEN",
		"CURLY_BRACKET_CLOSE", "SEMICOLON", "EQUALS_TEST", "NOT_EQUALS_TEST",
		"GREATER_THAN", "GREATER_OR_EQUALS", "LESS_THAN", "LESS_OR_EQUALS",
		"EQUALS", "COMMA", "AND_TEST", "OR_TEST", "NOT_TEST", "NUMBER", "TYPE_VOID",
		"TYPE_INT", "TYPE_STRING", "TYPE_BOOLEAN", "VALUE_TRUE", "VALUE_FALSE",
		"IDENT", "STRING",
	}
	staticData.ruleNames = []string{
		"WHITESPACE", "LINE_COMMENT", "COMMENT", "MUL", "DIV", "ADD", "SUB",
		"MOD", "PARENTHESIS_OPEN", "PARENTHESIS_CLOSE", "CURLY_BRACKET_OPEN",
		"CURLY_BRACKET_CLOSE", "SEMICOLON", "EQUALS_TEST", "NOT_EQUALS_TEST",
		"GREATER_THAN", "GREATER_OR_EQUALS", "LESS_THAN", "LESS_OR_EQUALS",
		"EQUALS", "COMMA", "AND_TEST", "OR_TEST", "NOT_TEST", "NUMBER", "TYPE_VOID",
		"TYPE_INT", "TYPE_STRING", "TYPE_BOOLEAN", "VALUE_TRUE", "VALUE_FALSE",
		"IDENT", "STRING",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 33, 208, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 1, 0, 4, 0, 69, 8, 0, 11, 0, 12, 0, 70, 1, 0,
		1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 79, 8, 1, 10, 1, 12, 1, 82, 9, 1, 1,
		1, 3, 1, 85, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2,
		95, 8, 2, 10, 2, 12, 2, 98, 9, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1,
		3, 1, 4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1,
		9, 1, 10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14,
		1, 14, 1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1,
		18, 1, 18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22,
		1, 22, 1, 23, 1, 23, 1, 24, 4, 24, 154, 8, 24, 11, 24, 12, 24, 155, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 30, 1, 31, 1, 31, 5, 31, 195, 8, 31, 10, 31, 12, 31, 198,
		9, 31, 1, 32, 1, 32, 5, 32, 202, 8, 32, 10, 32, 12, 32, 205, 9, 32, 1,
		32, 1, 32, 2, 80, 96, 0, 33, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7,
		15, 8, 17, 9, 19, 10, 21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33,
		17, 35, 18, 37, 19, 39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51,
		26, 53, 27, 55, 28, 57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 1, 0, 5, 3,
		0, 9, 10, 13, 13, 32, 32, 1, 0, 48, 57, 2, 0, 65, 90, 97, 122, 4, 0, 48,
		57, 65, 90, 95, 95, 97, 122, 1, 0, 34, 34, 214, 0, 1, 1, 0, 0, 0, 0, 3,
		1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11,
		1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0,
		19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0, 25, 1, 0, 0, 0,
		0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0,
		0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0,
		0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0, 0, 0, 0, 49, 1,
		0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1, 0, 0, 0, 0, 57,
		1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63, 1, 0, 0, 0, 0,
		65, 1, 0, 0, 0, 1, 68, 1, 0, 0, 0, 3, 74, 1, 0, 0, 0, 5, 90, 1, 0, 0, 0,
		7, 104, 1, 0, 0, 0, 9, 106, 1, 0, 0, 0, 11, 108, 1, 0, 0, 0, 13, 110, 1,
		0, 0, 0, 15, 112, 1, 0, 0, 0, 17, 114, 1, 0, 0, 0, 19, 116, 1, 0, 0, 0,
		21, 118, 1, 0, 0, 0, 23, 120, 1, 0, 0, 0, 25, 122, 1, 0, 0, 0, 27, 124,
		1, 0, 0, 0, 29, 127, 1, 0, 0, 0, 31, 130, 1, 0, 0, 0, 33, 132, 1, 0, 0,
		0, 35, 135, 1, 0, 0, 0, 37, 137, 1, 0, 0, 0, 39, 140, 1, 0, 0, 0, 41, 142,
		1, 0, 0, 0, 43, 144, 1, 0, 0, 0, 45, 147, 1, 0, 0, 0, 47, 150, 1, 0, 0,
		0, 49, 153, 1, 0, 0, 0, 51, 157, 1, 0, 0, 0, 53, 162, 1, 0, 0, 0, 55, 166,
		1, 0, 0, 0, 57, 173, 1, 0, 0, 0, 59, 181, 1, 0, 0, 0, 61, 186, 1, 0, 0,
		0, 63, 192, 1, 0, 0, 0, 65, 199, 1, 0, 0, 0, 67, 69, 7, 0, 0, 0, 68, 67,
		1, 0, 0, 0, 69, 70, 1, 0, 0, 0, 70, 68, 1, 0, 0, 0, 70, 71, 1, 0, 0, 0,
		71, 72, 1, 0, 0, 0, 72, 73, 6, 0, 0, 0, 73, 2, 1, 0, 0, 0, 74, 75, 5, 47,
		0, 0, 75, 76, 5, 47, 0, 0, 76, 80, 1, 0, 0, 0, 77, 79, 9, 0, 0, 0, 78,
		77, 1, 0, 0, 0, 79, 82, 1, 0, 0, 0, 80, 81, 1, 0, 0, 0, 80, 78, 1, 0, 0,
		0, 81, 84, 1, 0, 0, 0, 82, 80, 1, 0, 0, 0, 83, 85, 5, 13, 0, 0, 84, 83,
		1, 0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 86, 1, 0, 0, 0, 86, 87, 5, 10, 0, 0,
		87, 88, 1, 0, 0, 0, 88, 89, 6, 1, 0, 0, 89, 4, 1, 0, 0, 0, 90, 91, 5, 47,
		0, 0, 91, 92, 5, 42, 0, 0, 92, 96, 1, 0, 0, 0, 93, 95, 9, 0, 0, 0, 94,
		93, 1, 0, 0, 0, 95, 98, 1, 0, 0, 0, 96, 97, 1, 0, 0, 0, 96, 94, 1, 0, 0,
		0, 97, 99, 1, 0, 0, 0, 98, 96, 1, 0, 0, 0, 99, 100, 5, 42, 0, 0, 100, 101,
		5, 47, 0, 0, 101, 102, 1, 0, 0, 0, 102, 103, 6, 2, 0, 0, 103, 6, 1, 0,
		0, 0, 104, 105, 5, 42, 0, 0, 105, 8, 1, 0, 0, 0, 106, 107, 5, 47, 0, 0,
		107, 10, 1, 0, 0, 0, 108, 109, 5, 43, 0, 0, 109, 12, 1, 0, 0, 0, 110, 111,
		5, 45, 0, 0, 111, 14, 1, 0, 0, 0, 112, 113, 5, 37, 0, 0, 113, 16, 1, 0,
		0, 0, 114, 115, 5, 40, 0, 0, 115, 18, 1, 0, 0, 0, 116, 117, 5, 41, 0, 0,
		117, 20, 1, 0, 0, 0, 118, 119, 5, 123, 0, 0, 119, 22, 1, 0, 0, 0, 120,
		121, 5, 125, 0, 0, 121, 24, 1, 0, 0, 0, 122, 123, 5, 59, 0, 0, 123, 26,
		1, 0, 0, 0, 124, 125, 5, 61, 0, 0, 125, 126, 5, 61, 0, 0, 126, 28, 1, 0,
		0, 0, 127, 128, 5, 33, 0, 0, 128, 129, 5, 61, 0, 0, 129, 30, 1, 0, 0, 0,
		130, 131, 5, 62, 0, 0, 131, 32, 1, 0, 0, 0, 132, 133, 5, 62, 0, 0, 133,
		134, 5, 61, 0, 0, 134, 34, 1, 0, 0, 0, 135, 136, 5, 60, 0, 0, 136, 36,
		1, 0, 0, 0, 137, 138, 5, 60, 0, 0, 138, 139, 5, 61, 0, 0, 139, 38, 1, 0,
		0, 0, 140, 141, 5, 61, 0, 0, 141, 40, 1, 0, 0, 0, 142, 143, 5, 44, 0, 0,
		143, 42, 1, 0, 0, 0, 144, 145, 5, 38, 0, 0, 145, 146, 5, 38, 0, 0, 146,
		44, 1, 0, 0, 0, 147, 148, 5, 124, 0, 0, 148, 149, 5, 124, 0, 0, 149, 46,
		1, 0, 0, 0, 150, 151, 5, 33, 0, 0, 151, 48, 1, 0, 0, 0, 152, 154, 7, 1,
		0, 0, 153, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0, 155, 153, 1, 0, 0, 0,
		155, 156, 1, 0, 0, 0, 156, 50, 1, 0, 0, 0, 157, 158, 5, 118, 0, 0, 158,
		159, 5, 111, 0, 0, 159, 160, 5, 105, 0, 0, 160, 161, 5, 100, 0, 0, 161,
		52, 1, 0, 0, 0, 162, 163, 5, 105, 0, 0, 163, 164, 5, 110, 0, 0, 164, 165,
		5, 116, 0, 0, 165, 54, 1, 0, 0, 0, 166, 167, 5, 115, 0, 0, 167, 168, 5,
		116, 0, 0, 168, 169, 5, 114, 0, 0, 169, 170, 5, 105, 0, 0, 170, 171, 5,
		110, 0, 0, 171, 172, 5, 103, 0, 0, 172, 56, 1, 0, 0, 0, 173, 174, 5, 98,
		0, 0, 174, 175, 5, 111, 0, 0, 175, 176, 5, 111, 0, 0, 176, 177, 5, 108,
		0, 0, 177, 178, 5, 101, 0, 0, 178, 179, 5, 97, 0, 0, 179, 180, 5, 110,
		0, 0, 180, 58, 1, 0, 0, 0, 181, 182, 5, 116, 0, 0, 182, 183, 5, 114, 0,
		0, 183, 184, 5, 117, 0, 0, 184, 185, 5, 101, 0, 0, 185, 60, 1, 0, 0, 0,
		186, 187, 5, 102, 0, 0, 187, 188, 5, 97, 0, 0, 188, 189, 5, 108, 0, 0,
		189, 190, 5, 115, 0, 0, 190, 191, 5, 101, 0, 0, 191, 62, 1, 0, 0, 0, 192,
		196, 7, 2, 0, 0, 193, 195, 7, 3, 0, 0, 194, 193, 1, 0, 0, 0, 195, 198,
		1, 0, 0, 0, 196, 194, 1, 0, 0, 0, 196, 197, 1, 0, 0, 0, 197, 64, 1, 0,
		0, 0, 198, 196, 1, 0, 0, 0, 199, 203, 5, 34, 0, 0, 200, 202, 8, 4, 0, 0,
		201, 200, 1, 0, 0, 0, 202, 205, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0, 203,
		204, 1, 0, 0, 0, 204, 206, 1, 0, 0, 0, 205, 203, 1, 0, 0, 0, 206, 207,
		5, 34, 0, 0, 207, 66, 1, 0, 0, 0, 8, 0, 70, 80, 84, 96, 155, 196, 203,
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
	TinylangLexerLINE_COMMENT        = 2
	TinylangLexerCOMMENT             = 3
	TinylangLexerMUL                 = 4
	TinylangLexerDIV                 = 5
	TinylangLexerADD                 = 6
	TinylangLexerSUB                 = 7
	TinylangLexerMOD                 = 8
	TinylangLexerPARENTHESIS_OPEN    = 9
	TinylangLexerPARENTHESIS_CLOSE   = 10
	TinylangLexerCURLY_BRACKET_OPEN  = 11
	TinylangLexerCURLY_BRACKET_CLOSE = 12
	TinylangLexerSEMICOLON           = 13
	TinylangLexerEQUALS_TEST         = 14
	TinylangLexerNOT_EQUALS_TEST     = 15
	TinylangLexerGREATER_THAN        = 16
	TinylangLexerGREATER_OR_EQUALS   = 17
	TinylangLexerLESS_THAN           = 18
	TinylangLexerLESS_OR_EQUALS      = 19
	TinylangLexerEQUALS              = 20
	TinylangLexerCOMMA               = 21
	TinylangLexerAND_TEST            = 22
	TinylangLexerOR_TEST             = 23
	TinylangLexerNOT_TEST            = 24
	TinylangLexerNUMBER              = 25
	TinylangLexerTYPE_VOID           = 26
	TinylangLexerTYPE_INT            = 27
	TinylangLexerTYPE_STRING         = 28
	TinylangLexerTYPE_BOOLEAN        = 29
	TinylangLexerVALUE_TRUE          = 30
	TinylangLexerVALUE_FALSE         = 31
	TinylangLexerIDENT               = 32
	TinylangLexerSTRING              = 33
)
