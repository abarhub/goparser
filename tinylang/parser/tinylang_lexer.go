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
		"", "STRING", "LINE_COMMENT", "COMMENT", "MUL", "DIV", "ADD", "SUB",
		"MOD", "PARENTHESIS_OPEN", "PARENTHESIS_CLOSE", "CURLY_BRACKET_OPEN",
		"CURLY_BRACKET_CLOSE", "SEMICOLON", "EQUALS_TEST", "NOT_EQUALS_TEST",
		"GREATER_THAN", "GREATER_OR_EQUALS", "LESS_THAN", "LESS_OR_EQUALS",
		"EQUALS", "COMMA", "AND_TEST", "OR_TEST", "NOT_TEST", "NUMBER", "TYPE_VOID",
		"TYPE_INT", "TYPE_STRING", "TYPE_BOOLEAN", "VALUE_TRUE", "VALUE_FALSE",
		"IDENT", "WHITESPACE", "NEWLINE",
	}
	staticData.ruleNames = []string{
		"STRING", "LINE_COMMENT", "COMMENT", "MUL", "DIV", "ADD", "SUB", "MOD",
		"PARENTHESIS_OPEN", "PARENTHESIS_CLOSE", "CURLY_BRACKET_OPEN", "CURLY_BRACKET_CLOSE",
		"SEMICOLON", "EQUALS_TEST", "NOT_EQUALS_TEST", "GREATER_THAN", "GREATER_OR_EQUALS",
		"LESS_THAN", "LESS_OR_EQUALS", "EQUALS", "COMMA", "AND_TEST", "OR_TEST",
		"NOT_TEST", "NUMBER", "TYPE_VOID", "TYPE_INT", "TYPE_STRING", "TYPE_BOOLEAN",
		"VALUE_TRUE", "VALUE_FALSE", "IDENT", "WHITESPACE", "NEWLINE",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 34, 214, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25,
		2, 26, 7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2,
		31, 7, 31, 2, 32, 7, 32, 2, 33, 7, 33, 1, 0, 1, 0, 5, 0, 72, 8, 0, 10,
		0, 12, 0, 75, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 5, 1, 83, 8, 1,
		10, 1, 12, 1, 86, 9, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 2, 1, 2, 5, 2, 94, 8,
		2, 10, 2, 12, 2, 97, 9, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1,
		4, 1, 4, 1, 5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 8, 1, 8, 1, 9, 1, 9, 1,
		10, 1, 10, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14,
		1, 14, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 17, 1, 17, 1, 18, 1, 18, 1,
		18, 1, 19, 1, 19, 1, 20, 1, 20, 1, 21, 1, 21, 1, 21, 1, 22, 1, 22, 1, 22,
		1, 23, 1, 23, 1, 24, 4, 24, 153, 8, 24, 11, 24, 12, 24, 154, 1, 25, 1,
		25, 1, 25, 1, 25, 1, 25, 1, 26, 1, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27,
		1, 27, 1, 27, 1, 27, 1, 27, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1, 28, 1,
		28, 1, 28, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1, 30,
		1, 30, 1, 30, 1, 31, 1, 31, 5, 31, 194, 8, 31, 10, 31, 12, 31, 197, 9,
		31, 1, 32, 4, 32, 200, 8, 32, 11, 32, 12, 32, 201, 1, 32, 1, 32, 1, 33,
		1, 33, 3, 33, 208, 8, 33, 1, 33, 3, 33, 211, 8, 33, 1, 33, 1, 33, 1, 95,
		0, 34, 1, 1, 3, 2, 5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10,
		21, 11, 23, 12, 25, 13, 27, 14, 29, 15, 31, 16, 33, 17, 35, 18, 37, 19,
		39, 20, 41, 21, 43, 22, 45, 23, 47, 24, 49, 25, 51, 26, 53, 27, 55, 28,
		57, 29, 59, 30, 61, 31, 63, 32, 65, 33, 67, 34, 1, 0, 6, 3, 0, 10, 10,
		13, 13, 34, 34, 2, 0, 10, 10, 13, 13, 1, 0, 48, 57, 2, 0, 65, 90, 97, 122,
		4, 0, 48, 57, 65, 90, 95, 95, 97, 122, 2, 0, 9, 9, 32, 32, 221, 0, 1, 1,
		0, 0, 0, 0, 3, 1, 0, 0, 0, 0, 5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1,
		0, 0, 0, 0, 11, 1, 0, 0, 0, 0, 13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17,
		1, 0, 0, 0, 0, 19, 1, 0, 0, 0, 0, 21, 1, 0, 0, 0, 0, 23, 1, 0, 0, 0, 0,
		25, 1, 0, 0, 0, 0, 27, 1, 0, 0, 0, 0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0,
		0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0, 0, 0, 37, 1, 0, 0, 0, 0, 39, 1, 0, 0,
		0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 0, 45, 1, 0, 0, 0, 0, 47, 1, 0,
		0, 0, 0, 49, 1, 0, 0, 0, 0, 51, 1, 0, 0, 0, 0, 53, 1, 0, 0, 0, 0, 55, 1,
		0, 0, 0, 0, 57, 1, 0, 0, 0, 0, 59, 1, 0, 0, 0, 0, 61, 1, 0, 0, 0, 0, 63,
		1, 0, 0, 0, 0, 65, 1, 0, 0, 0, 0, 67, 1, 0, 0, 0, 1, 69, 1, 0, 0, 0, 3,
		78, 1, 0, 0, 0, 5, 89, 1, 0, 0, 0, 7, 103, 1, 0, 0, 0, 9, 105, 1, 0, 0,
		0, 11, 107, 1, 0, 0, 0, 13, 109, 1, 0, 0, 0, 15, 111, 1, 0, 0, 0, 17, 113,
		1, 0, 0, 0, 19, 115, 1, 0, 0, 0, 21, 117, 1, 0, 0, 0, 23, 119, 1, 0, 0,
		0, 25, 121, 1, 0, 0, 0, 27, 123, 1, 0, 0, 0, 29, 126, 1, 0, 0, 0, 31, 129,
		1, 0, 0, 0, 33, 131, 1, 0, 0, 0, 35, 134, 1, 0, 0, 0, 37, 136, 1, 0, 0,
		0, 39, 139, 1, 0, 0, 0, 41, 141, 1, 0, 0, 0, 43, 143, 1, 0, 0, 0, 45, 146,
		1, 0, 0, 0, 47, 149, 1, 0, 0, 0, 49, 152, 1, 0, 0, 0, 51, 156, 1, 0, 0,
		0, 53, 161, 1, 0, 0, 0, 55, 165, 1, 0, 0, 0, 57, 172, 1, 0, 0, 0, 59, 180,
		1, 0, 0, 0, 61, 185, 1, 0, 0, 0, 63, 191, 1, 0, 0, 0, 65, 199, 1, 0, 0,
		0, 67, 210, 1, 0, 0, 0, 69, 73, 5, 34, 0, 0, 70, 72, 8, 0, 0, 0, 71, 70,
		1, 0, 0, 0, 72, 75, 1, 0, 0, 0, 73, 71, 1, 0, 0, 0, 73, 74, 1, 0, 0, 0,
		74, 76, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 76, 77, 5, 34, 0, 0, 77, 2, 1,
		0, 0, 0, 78, 79, 5, 47, 0, 0, 79, 80, 5, 47, 0, 0, 80, 84, 1, 0, 0, 0,
		81, 83, 8, 1, 0, 0, 82, 81, 1, 0, 0, 0, 83, 86, 1, 0, 0, 0, 84, 82, 1,
		0, 0, 0, 84, 85, 1, 0, 0, 0, 85, 87, 1, 0, 0, 0, 86, 84, 1, 0, 0, 0, 87,
		88, 6, 1, 0, 0, 88, 4, 1, 0, 0, 0, 89, 90, 5, 47, 0, 0, 90, 91, 5, 42,
		0, 0, 91, 95, 1, 0, 0, 0, 92, 94, 9, 0, 0, 0, 93, 92, 1, 0, 0, 0, 94, 97,
		1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 95, 93, 1, 0, 0, 0, 96, 98, 1, 0, 0, 0,
		97, 95, 1, 0, 0, 0, 98, 99, 5, 42, 0, 0, 99, 100, 5, 47, 0, 0, 100, 101,
		1, 0, 0, 0, 101, 102, 6, 2, 0, 0, 102, 6, 1, 0, 0, 0, 103, 104, 5, 42,
		0, 0, 104, 8, 1, 0, 0, 0, 105, 106, 5, 47, 0, 0, 106, 10, 1, 0, 0, 0, 107,
		108, 5, 43, 0, 0, 108, 12, 1, 0, 0, 0, 109, 110, 5, 45, 0, 0, 110, 14,
		1, 0, 0, 0, 111, 112, 5, 37, 0, 0, 112, 16, 1, 0, 0, 0, 113, 114, 5, 40,
		0, 0, 114, 18, 1, 0, 0, 0, 115, 116, 5, 41, 0, 0, 116, 20, 1, 0, 0, 0,
		117, 118, 5, 123, 0, 0, 118, 22, 1, 0, 0, 0, 119, 120, 5, 125, 0, 0, 120,
		24, 1, 0, 0, 0, 121, 122, 5, 59, 0, 0, 122, 26, 1, 0, 0, 0, 123, 124, 5,
		61, 0, 0, 124, 125, 5, 61, 0, 0, 125, 28, 1, 0, 0, 0, 126, 127, 5, 33,
		0, 0, 127, 128, 5, 61, 0, 0, 128, 30, 1, 0, 0, 0, 129, 130, 5, 62, 0, 0,
		130, 32, 1, 0, 0, 0, 131, 132, 5, 62, 0, 0, 132, 133, 5, 61, 0, 0, 133,
		34, 1, 0, 0, 0, 134, 135, 5, 60, 0, 0, 135, 36, 1, 0, 0, 0, 136, 137, 5,
		60, 0, 0, 137, 138, 5, 61, 0, 0, 138, 38, 1, 0, 0, 0, 139, 140, 5, 61,
		0, 0, 140, 40, 1, 0, 0, 0, 141, 142, 5, 44, 0, 0, 142, 42, 1, 0, 0, 0,
		143, 144, 5, 38, 0, 0, 144, 145, 5, 38, 0, 0, 145, 44, 1, 0, 0, 0, 146,
		147, 5, 124, 0, 0, 147, 148, 5, 124, 0, 0, 148, 46, 1, 0, 0, 0, 149, 150,
		5, 33, 0, 0, 150, 48, 1, 0, 0, 0, 151, 153, 7, 2, 0, 0, 152, 151, 1, 0,
		0, 0, 153, 154, 1, 0, 0, 0, 154, 152, 1, 0, 0, 0, 154, 155, 1, 0, 0, 0,
		155, 50, 1, 0, 0, 0, 156, 157, 5, 118, 0, 0, 157, 158, 5, 111, 0, 0, 158,
		159, 5, 105, 0, 0, 159, 160, 5, 100, 0, 0, 160, 52, 1, 0, 0, 0, 161, 162,
		5, 105, 0, 0, 162, 163, 5, 110, 0, 0, 163, 164, 5, 116, 0, 0, 164, 54,
		1, 0, 0, 0, 165, 166, 5, 115, 0, 0, 166, 167, 5, 116, 0, 0, 167, 168, 5,
		114, 0, 0, 168, 169, 5, 105, 0, 0, 169, 170, 5, 110, 0, 0, 170, 171, 5,
		103, 0, 0, 171, 56, 1, 0, 0, 0, 172, 173, 5, 98, 0, 0, 173, 174, 5, 111,
		0, 0, 174, 175, 5, 111, 0, 0, 175, 176, 5, 108, 0, 0, 176, 177, 5, 101,
		0, 0, 177, 178, 5, 97, 0, 0, 178, 179, 5, 110, 0, 0, 179, 58, 1, 0, 0,
		0, 180, 181, 5, 116, 0, 0, 181, 182, 5, 114, 0, 0, 182, 183, 5, 117, 0,
		0, 183, 184, 5, 101, 0, 0, 184, 60, 1, 0, 0, 0, 185, 186, 5, 102, 0, 0,
		186, 187, 5, 97, 0, 0, 187, 188, 5, 108, 0, 0, 188, 189, 5, 115, 0, 0,
		189, 190, 5, 101, 0, 0, 190, 62, 1, 0, 0, 0, 191, 195, 7, 3, 0, 0, 192,
		194, 7, 4, 0, 0, 193, 192, 1, 0, 0, 0, 194, 197, 1, 0, 0, 0, 195, 193,
		1, 0, 0, 0, 195, 196, 1, 0, 0, 0, 196, 64, 1, 0, 0, 0, 197, 195, 1, 0,
		0, 0, 198, 200, 7, 5, 0, 0, 199, 198, 1, 0, 0, 0, 200, 201, 1, 0, 0, 0,
		201, 199, 1, 0, 0, 0, 201, 202, 1, 0, 0, 0, 202, 203, 1, 0, 0, 0, 203,
		204, 6, 32, 0, 0, 204, 66, 1, 0, 0, 0, 205, 207, 5, 13, 0, 0, 206, 208,
		5, 10, 0, 0, 207, 206, 1, 0, 0, 0, 207, 208, 1, 0, 0, 0, 208, 211, 1, 0,
		0, 0, 209, 211, 5, 10, 0, 0, 210, 205, 1, 0, 0, 0, 210, 209, 1, 0, 0, 0,
		211, 212, 1, 0, 0, 0, 212, 213, 6, 33, 0, 0, 213, 68, 1, 0, 0, 0, 9, 0,
		73, 84, 95, 154, 195, 201, 207, 210, 1, 6, 0, 0,
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
	TinylangLexerSTRING              = 1
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
	TinylangLexerWHITESPACE          = 33
	TinylangLexerNEWLINE             = 34
)
