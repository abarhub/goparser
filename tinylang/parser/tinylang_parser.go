// Code generated from Tinylang.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Tinylang

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type TinylangParser struct {
	*antlr.BaseParser
}

var tinylangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tinylangParserInit() {
	staticData := &tinylangParserStaticData
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
		"start", "function", "instruction_list", "instruction", "type", "expression",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 29, 76, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 2, 1, 2, 1, 2, 5, 2, 26, 8, 2, 10, 2, 12, 2, 29, 9, 2, 1, 3, 1, 3,
		1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 39, 8, 3, 1, 4, 1, 4, 1, 4, 1,
		4, 3, 4, 45, 8, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5,
		1, 5, 3, 5, 57, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 5, 5, 71, 8, 5, 10, 5, 12, 5, 74, 9, 5, 1, 5, 0, 1,
		10, 6, 0, 2, 4, 6, 8, 10, 0, 4, 2, 0, 2, 3, 6, 6, 1, 0, 4, 5, 1, 0, 12,
		16, 1, 0, 19, 20, 83, 0, 12, 1, 0, 0, 0, 2, 14, 1, 0, 0, 0, 4, 27, 1, 0,
		0, 0, 6, 38, 1, 0, 0, 0, 8, 44, 1, 0, 0, 0, 10, 56, 1, 0, 0, 0, 12, 13,
		3, 2, 1, 0, 13, 1, 1, 0, 0, 0, 14, 15, 3, 8, 4, 0, 15, 16, 5, 28, 0, 0,
		16, 17, 5, 7, 0, 0, 17, 18, 5, 8, 0, 0, 18, 19, 5, 9, 0, 0, 19, 20, 3,
		4, 2, 0, 20, 21, 5, 10, 0, 0, 21, 3, 1, 0, 0, 0, 22, 23, 3, 6, 3, 0, 23,
		24, 5, 11, 0, 0, 24, 26, 1, 0, 0, 0, 25, 22, 1, 0, 0, 0, 26, 29, 1, 0,
		0, 0, 27, 25, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 5, 1, 0, 0, 0, 29, 27,
		1, 0, 0, 0, 30, 31, 5, 28, 0, 0, 31, 32, 5, 17, 0, 0, 32, 39, 3, 10, 5,
		0, 33, 34, 3, 8, 4, 0, 34, 35, 5, 28, 0, 0, 35, 36, 5, 17, 0, 0, 36, 37,
		3, 10, 5, 0, 37, 39, 1, 0, 0, 0, 38, 30, 1, 0, 0, 0, 38, 33, 1, 0, 0, 0,
		39, 7, 1, 0, 0, 0, 40, 45, 5, 22, 0, 0, 41, 45, 5, 23, 0, 0, 42, 45, 5,
		24, 0, 0, 43, 45, 5, 25, 0, 0, 44, 40, 1, 0, 0, 0, 44, 41, 1, 0, 0, 0,
		44, 42, 1, 0, 0, 0, 44, 43, 1, 0, 0, 0, 45, 9, 1, 0, 0, 0, 46, 47, 6, 5,
		-1, 0, 47, 57, 5, 21, 0, 0, 48, 57, 5, 26, 0, 0, 49, 57, 5, 27, 0, 0, 50,
		57, 5, 28, 0, 0, 51, 57, 5, 29, 0, 0, 52, 53, 5, 7, 0, 0, 53, 54, 3, 10,
		5, 0, 54, 55, 5, 8, 0, 0, 55, 57, 1, 0, 0, 0, 56, 46, 1, 0, 0, 0, 56, 48,
		1, 0, 0, 0, 56, 49, 1, 0, 0, 0, 56, 50, 1, 0, 0, 0, 56, 51, 1, 0, 0, 0,
		56, 52, 1, 0, 0, 0, 57, 72, 1, 0, 0, 0, 58, 59, 10, 10, 0, 0, 59, 60, 7,
		0, 0, 0, 60, 71, 3, 10, 5, 11, 61, 62, 10, 9, 0, 0, 62, 63, 7, 1, 0, 0,
		63, 71, 3, 10, 5, 10, 64, 65, 10, 8, 0, 0, 65, 66, 7, 2, 0, 0, 66, 71,
		3, 10, 5, 9, 67, 68, 10, 7, 0, 0, 68, 69, 7, 3, 0, 0, 69, 71, 3, 10, 5,
		8, 70, 58, 1, 0, 0, 0, 70, 61, 1, 0, 0, 0, 70, 64, 1, 0, 0, 0, 70, 67,
		1, 0, 0, 0, 71, 74, 1, 0, 0, 0, 72, 70, 1, 0, 0, 0, 72, 73, 1, 0, 0, 0,
		73, 11, 1, 0, 0, 0, 74, 72, 1, 0, 0, 0, 6, 27, 38, 44, 56, 70, 72,
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

// TinylangParserInit initializes any static state used to implement TinylangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewTinylangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func TinylangParserInit() {
	staticData := &tinylangParserStaticData
	staticData.once.Do(tinylangParserInit)
}

// NewTinylangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewTinylangParser(input antlr.TokenStream) *TinylangParser {
	TinylangParserInit()
	this := new(TinylangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &tinylangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Tinylang.g4"

	return this
}

// TinylangParser tokens.
const (
	TinylangParserEOF                 = antlr.TokenEOF
	TinylangParserWHITESPACE          = 1
	TinylangParserMUL                 = 2
	TinylangParserDIV                 = 3
	TinylangParserADD                 = 4
	TinylangParserSUB                 = 5
	TinylangParserMOD                 = 6
	TinylangParserPARENTHESIS_OPEN    = 7
	TinylangParserPARENTHESIS_CLOSE   = 8
	TinylangParserCURLY_BRACKET_OPEN  = 9
	TinylangParserCURLY_BRACKET_CLOSE = 10
	TinylangParserSEMICOLON           = 11
	TinylangParserEQUALS_TEST         = 12
	TinylangParserGREATER_THAN        = 13
	TinylangParserGREATER_OR_EQUALS   = 14
	TinylangParserLESS_THAN           = 15
	TinylangParserLESS_OR_EQUALS      = 16
	TinylangParserEQUALS              = 17
	TinylangParserCOMMA               = 18
	TinylangParserAND_TEST            = 19
	TinylangParserOR_TEST             = 20
	TinylangParserNUMBER              = 21
	TinylangParserTYPE_VOID           = 22
	TinylangParserTYPE_INT            = 23
	TinylangParserTYPE_STRING         = 24
	TinylangParserTYPE_BOOLEAN        = 25
	TinylangParserVALUE_TRUE          = 26
	TinylangParserVALUE_FALSE         = 27
	TinylangParserIDENT               = 28
	TinylangParserSTRING              = 29
)

// TinylangParser rules.
const (
	TinylangParserRULE_start            = 0
	TinylangParserRULE_function         = 1
	TinylangParserRULE_instruction_list = 2
	TinylangParserRULE_instruction      = 3
	TinylangParserRULE_type             = 4
	TinylangParserRULE_expression       = 5
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinylangParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinylangParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) CopyFrom(ctx *StartContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StartFunctionContext struct {
	*StartContext
}

func NewStartFunctionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StartFunctionContext {
	var p = new(StartFunctionContext)

	p.StartContext = NewEmptyStartContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StartContext))

	return p
}

func (s *StartFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartFunctionContext) Function() IFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionContext)
}

func (s *StartFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.EnterStartFunction(s)
	}
}

func (s *StartFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.ExitStartFunction(s)
	}
}

func (p *TinylangParser) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, TinylangParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewStartFunctionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(12)
		p.Function()
	}

	return localctx
}

// IFunctionContext is an interface to support dynamic dispatch.
type IFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionContext differentiates from other interfaces.
	IsFunctionContext()
}

type FunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionContext() *FunctionContext {
	var p = new(FunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinylangParserRULE_function
	return p
}

func (*FunctionContext) IsFunctionContext() {}

func NewFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionContext {
	var p = new(FunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinylangParserRULE_function

	return p
}

func (s *FunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionContext) CopyFrom(ctx *FunctionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type FunctContext struct {
	*FunctionContext
	t         ITypeContext
	name      antlr.Token
	instrlist IInstruction_listContext
}

func NewFunctContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctContext {
	var p = new(FunctContext)

	p.FunctionContext = NewEmptyFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionContext))

	return p
}

func (s *FunctContext) GetName() antlr.Token { return s.name }

func (s *FunctContext) SetName(v antlr.Token) { s.name = v }

func (s *FunctContext) GetT() ITypeContext { return s.t }

func (s *FunctContext) GetInstrlist() IInstruction_listContext { return s.instrlist }

func (s *FunctContext) SetT(v ITypeContext) { s.t = v }

func (s *FunctContext) SetInstrlist(v IInstruction_listContext) { s.instrlist = v }

func (s *FunctContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctContext) PARENTHESIS_OPEN() antlr.TerminalNode {
	return s.GetToken(TinylangParserPARENTHESIS_OPEN, 0)
}

func (s *FunctContext) PARENTHESIS_CLOSE() antlr.TerminalNode {
	return s.GetToken(TinylangParserPARENTHESIS_CLOSE, 0)
}

func (s *FunctContext) CURLY_BRACKET_OPEN() antlr.TerminalNode {
	return s.GetToken(TinylangParserCURLY_BRACKET_OPEN, 0)
}

func (s *FunctContext) CURLY_BRACKET_CLOSE() antlr.TerminalNode {
	return s.GetToken(TinylangParserCURLY_BRACKET_CLOSE, 0)
}

func (s *FunctContext) Type() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *FunctContext) IDENT() antlr.TerminalNode {
	return s.GetToken(TinylangParserIDENT, 0)
}

func (s *FunctContext) Instruction_list() IInstruction_listContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstruction_listContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstruction_listContext)
}

func (s *FunctContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.EnterFunct(s)
	}
}

func (s *FunctContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.ExitFunct(s)
	}
}

func (p *TinylangParser) Function() (localctx IFunctionContext) {
	this := p
	_ = this

	localctx = NewFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, TinylangParserRULE_function)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewFunctContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(14)

		var _x = p.Type()

		localctx.(*FunctContext).t = _x
	}
	{
		p.SetState(15)

		var _m = p.Match(TinylangParserIDENT)

		localctx.(*FunctContext).name = _m
	}
	{
		p.SetState(16)
		p.Match(TinylangParserPARENTHESIS_OPEN)
	}
	{
		p.SetState(17)
		p.Match(TinylangParserPARENTHESIS_CLOSE)
	}
	{
		p.SetState(18)
		p.Match(TinylangParserCURLY_BRACKET_OPEN)
	}
	{
		p.SetState(19)

		var _x = p.Instruction_list()

		localctx.(*FunctContext).instrlist = _x
	}
	{
		p.SetState(20)
		p.Match(TinylangParserCURLY_BRACKET_CLOSE)
	}

	return localctx
}

// IInstruction_listContext is an interface to support dynamic dispatch.
type IInstruction_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstruction_listContext differentiates from other interfaces.
	IsInstruction_listContext()
}

type Instruction_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstruction_listContext() *Instruction_listContext {
	var p = new(Instruction_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinylangParserRULE_instruction_list
	return p
}

func (*Instruction_listContext) IsInstruction_listContext() {}

func NewInstruction_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instruction_listContext {
	var p = new(Instruction_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinylangParserRULE_instruction_list

	return p
}

func (s *Instruction_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Instruction_listContext) CopyFrom(ctx *Instruction_listContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Instruction_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instruction_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ListInstructionContext struct {
	*Instruction_listContext
}

func NewListInstructionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ListInstructionContext {
	var p = new(ListInstructionContext)

	p.Instruction_listContext = NewEmptyInstruction_listContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Instruction_listContext))

	return p
}

func (s *ListInstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListInstructionContext) AllInstruction() []IInstructionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IInstructionContext); ok {
			len++
		}
	}

	tst := make([]IInstructionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IInstructionContext); ok {
			tst[i] = t.(IInstructionContext)
			i++
		}
	}

	return tst
}

func (s *ListInstructionContext) Instruction(i int) IInstructionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IInstructionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IInstructionContext)
}

func (s *ListInstructionContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(TinylangParserSEMICOLON)
}

func (s *ListInstructionContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(TinylangParserSEMICOLON, i)
}

func (s *ListInstructionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.EnterListInstruction(s)
	}
}

func (s *ListInstructionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.ExitListInstruction(s)
	}
}

func (p *TinylangParser) Instruction_list() (localctx IInstruction_listContext) {
	this := p
	_ = this

	localctx = NewInstruction_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, TinylangParserRULE_instruction_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewListInstructionContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TinylangParserTYPE_VOID)|(1<<TinylangParserTYPE_INT)|(1<<TinylangParserTYPE_STRING)|(1<<TinylangParserTYPE_BOOLEAN)|(1<<TinylangParserIDENT))) != 0 {
		{
			p.SetState(22)
			p.Instruction()
		}
		{
			p.SetState(23)
			p.Match(TinylangParserSEMICOLON)
		}

		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IInstructionContext is an interface to support dynamic dispatch.
type IInstructionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInstructionContext differentiates from other interfaces.
	IsInstructionContext()
}

type InstructionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInstructionContext() *InstructionContext {
	var p = new(InstructionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinylangParserRULE_instruction
	return p
}

func (*InstructionContext) IsInstructionContext() {}

func NewInstructionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstructionContext {
	var p = new(InstructionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinylangParserRULE_instruction

	return p
}

func (s *InstructionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstructionContext) CopyFrom(ctx *InstructionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *InstructionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstructionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type InstrDeclareContext struct {
	*InstructionContext
	t   ITypeContext
	id  antlr.Token
	exp IExpressionContext
}

func NewInstrDeclareContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InstrDeclareContext {
	var p = new(InstrDeclareContext)

	p.InstructionContext = NewEmptyInstructionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InstructionContext))

	return p
}

func (s *InstrDeclareContext) GetId() antlr.Token { return s.id }

func (s *InstrDeclareContext) SetId(v antlr.Token) { s.id = v }

func (s *InstrDeclareContext) GetT() ITypeContext { return s.t }

func (s *InstrDeclareContext) GetExp() IExpressionContext { return s.exp }

func (s *InstrDeclareContext) SetT(v ITypeContext) { s.t = v }

func (s *InstrDeclareContext) SetExp(v IExpressionContext) { s.exp = v }

func (s *InstrDeclareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstrDeclareContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(TinylangParserEQUALS, 0)
}

func (s *InstrDeclareContext) Type() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *InstrDeclareContext) IDENT() antlr.TerminalNode {
	return s.GetToken(TinylangParserIDENT, 0)
}

func (s *InstrDeclareContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InstrDeclareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.EnterInstrDeclare(s)
	}
}

func (s *InstrDeclareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.ExitInstrDeclare(s)
	}
}

type InstrAffectContext struct {
	*InstructionContext
	id  antlr.Token
	exp IExpressionContext
}

func NewInstrAffectContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InstrAffectContext {
	var p = new(InstrAffectContext)

	p.InstructionContext = NewEmptyInstructionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*InstructionContext))

	return p
}

func (s *InstrAffectContext) GetId() antlr.Token { return s.id }

func (s *InstrAffectContext) SetId(v antlr.Token) { s.id = v }

func (s *InstrAffectContext) GetExp() IExpressionContext { return s.exp }

func (s *InstrAffectContext) SetExp(v IExpressionContext) { s.exp = v }

func (s *InstrAffectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstrAffectContext) EQUALS() antlr.TerminalNode {
	return s.GetToken(TinylangParserEQUALS, 0)
}

func (s *InstrAffectContext) IDENT() antlr.TerminalNode {
	return s.GetToken(TinylangParserIDENT, 0)
}

func (s *InstrAffectContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InstrAffectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.EnterInstrAffect(s)
	}
}

func (s *InstrAffectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.ExitInstrAffect(s)
	}
}

func (p *TinylangParser) Instruction() (localctx IInstructionContext) {
	this := p
	_ = this

	localctx = NewInstructionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, TinylangParserRULE_instruction)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(38)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TinylangParserIDENT:
		localctx = NewInstrAffectContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(30)

			var _m = p.Match(TinylangParserIDENT)

			localctx.(*InstrAffectContext).id = _m
		}
		{
			p.SetState(31)
			p.Match(TinylangParserEQUALS)
		}
		{
			p.SetState(32)

			var _x = p.expression(0)

			localctx.(*InstrAffectContext).exp = _x
		}

	case TinylangParserTYPE_VOID, TinylangParserTYPE_INT, TinylangParserTYPE_STRING, TinylangParserTYPE_BOOLEAN:
		localctx = NewInstrDeclareContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(33)

			var _x = p.Type()

			localctx.(*InstrDeclareContext).t = _x
		}
		{
			p.SetState(34)

			var _m = p.Match(TinylangParserIDENT)

			localctx.(*InstrDeclareContext).id = _m
		}
		{
			p.SetState(35)
			p.Match(TinylangParserEQUALS)
		}
		{
			p.SetState(36)

			var _x = p.expression(0)

			localctx.(*InstrDeclareContext).exp = _x
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinylangParserRULE_type
	return p
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinylangParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) CopyFrom(ctx *TypeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TypeBooleanContext struct {
	*TypeContext
}

func NewTypeBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeBooleanContext {
	var p = new(TypeBooleanContext)

	p.TypeContext = NewEmptyTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeContext))

	return p
}

func (s *TypeBooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeBooleanContext) TYPE_BOOLEAN() antlr.TerminalNode {
	return s.GetToken(TinylangParserTYPE_BOOLEAN, 0)
}

func (s *TypeBooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.EnterTypeBoolean(s)
	}
}

func (s *TypeBooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.ExitTypeBoolean(s)
	}
}

type TypeVoidContext struct {
	*TypeContext
}

func NewTypeVoidContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeVoidContext {
	var p = new(TypeVoidContext)

	p.TypeContext = NewEmptyTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeContext))

	return p
}

func (s *TypeVoidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeVoidContext) TYPE_VOID() antlr.TerminalNode {
	return s.GetToken(TinylangParserTYPE_VOID, 0)
}

func (s *TypeVoidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.EnterTypeVoid(s)
	}
}

func (s *TypeVoidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.ExitTypeVoid(s)
	}
}

type TypeIntContext struct {
	*TypeContext
}

func NewTypeIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeIntContext {
	var p = new(TypeIntContext)

	p.TypeContext = NewEmptyTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeContext))

	return p
}

func (s *TypeIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeIntContext) TYPE_INT() antlr.TerminalNode {
	return s.GetToken(TinylangParserTYPE_INT, 0)
}

func (s *TypeIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.EnterTypeInt(s)
	}
}

func (s *TypeIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.ExitTypeInt(s)
	}
}

type TypeStringContext struct {
	*TypeContext
}

func NewTypeStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TypeStringContext {
	var p = new(TypeStringContext)

	p.TypeContext = NewEmptyTypeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeContext))

	return p
}

func (s *TypeStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeStringContext) TYPE_STRING() antlr.TerminalNode {
	return s.GetToken(TinylangParserTYPE_STRING, 0)
}

func (s *TypeStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.EnterTypeString(s)
	}
}

func (s *TypeStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.ExitTypeString(s)
	}
}

func (p *TinylangParser) Type() (localctx ITypeContext) {
	this := p
	_ = this

	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, TinylangParserRULE_type)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(44)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TinylangParserTYPE_VOID:
		localctx = NewTypeVoidContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(40)
			p.Match(TinylangParserTYPE_VOID)
		}

	case TinylangParserTYPE_INT:
		localctx = NewTypeIntContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(41)
			p.Match(TinylangParserTYPE_INT)
		}

	case TinylangParserTYPE_STRING:
		localctx = NewTypeStringContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(42)
			p.Match(TinylangParserTYPE_STRING)
		}

	case TinylangParserTYPE_BOOLEAN:
		localctx = NewTypeBooleanContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(43)
			p.Match(TinylangParserTYPE_BOOLEAN)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = TinylangParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = TinylangParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ParenthesisContext struct {
	*ExpressionContext
}

func NewParenthesisContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesisContext {
	var p = new(ParenthesisContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ParenthesisContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesisContext) PARENTHESIS_OPEN() antlr.TerminalNode {
	return s.GetToken(TinylangParserPARENTHESIS_OPEN, 0)
}

func (s *ParenthesisContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenthesisContext) PARENTHESIS_CLOSE() antlr.TerminalNode {
	return s.GetToken(TinylangParserPARENTHESIS_CLOSE, 0)
}

func (s *ParenthesisContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.EnterParenthesis(s)
	}
}

func (s *ParenthesisContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.ExitParenthesis(s)
	}
}

type NumberContext struct {
	*ExpressionContext
	e antlr.Token
}

func NewNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberContext {
	var p = new(NumberContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NumberContext) GetE() antlr.Token { return s.e }

func (s *NumberContext) SetE(v antlr.Token) { s.e = v }

func (s *NumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(TinylangParserNUMBER, 0)
}

func (s *NumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.EnterNumber(s)
	}
}

func (s *NumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.ExitNumber(s)
	}
}

type ExprStringContext struct {
	*ExpressionContext
	e antlr.Token
}

func NewExprStringContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprStringContext {
	var p = new(ExprStringContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExprStringContext) GetE() antlr.Token { return s.e }

func (s *ExprStringContext) SetE(v antlr.Token) { s.e = v }

func (s *ExprStringContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprStringContext) STRING() antlr.TerminalNode {
	return s.GetToken(TinylangParserSTRING, 0)
}

func (s *ExprStringContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.EnterExprString(s)
	}
}

func (s *ExprStringContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.ExitExprString(s)
	}
}

type MulDivContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MulDivContext) GetOp() antlr.Token { return s.op }

func (s *MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MulDivContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MulDivContext) MUL() antlr.TerminalNode {
	return s.GetToken(TinylangParserMUL, 0)
}

func (s *MulDivContext) DIV() antlr.TerminalNode {
	return s.GetToken(TinylangParserDIV, 0)
}

func (s *MulDivContext) MOD() antlr.TerminalNode {
	return s.GetToken(TinylangParserMOD, 0)
}

func (s *MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.EnterMulDiv(s)
	}
}

func (s *MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.ExitMulDiv(s)
	}
}

type AddSubContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AddSubContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AddSubContext) ADD() antlr.TerminalNode {
	return s.GetToken(TinylangParserADD, 0)
}

func (s *AddSubContext) SUB() antlr.TerminalNode {
	return s.GetToken(TinylangParserSUB, 0)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.ExitAddSub(s)
	}
}

type TrueContext struct {
	*ExpressionContext
	e antlr.Token
}

func NewTrueContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueContext {
	var p = new(TrueContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *TrueContext) GetE() antlr.Token { return s.e }

func (s *TrueContext) SetE(v antlr.Token) { s.e = v }

func (s *TrueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TrueContext) VALUE_TRUE() antlr.TerminalNode {
	return s.GetToken(TinylangParserVALUE_TRUE, 0)
}

func (s *TrueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.EnterTrue(s)
	}
}

func (s *TrueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.ExitTrue(s)
	}
}

type CompareContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewCompareContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CompareContext {
	var p = new(CompareContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *CompareContext) GetOp() antlr.Token { return s.op }

func (s *CompareContext) SetOp(v antlr.Token) { s.op = v }

func (s *CompareContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompareContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *CompareContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CompareContext) GREATER_THAN() antlr.TerminalNode {
	return s.GetToken(TinylangParserGREATER_THAN, 0)
}

func (s *CompareContext) GREATER_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(TinylangParserGREATER_OR_EQUALS, 0)
}

func (s *CompareContext) LESS_THAN() antlr.TerminalNode {
	return s.GetToken(TinylangParserLESS_THAN, 0)
}

func (s *CompareContext) LESS_OR_EQUALS() antlr.TerminalNode {
	return s.GetToken(TinylangParserLESS_OR_EQUALS, 0)
}

func (s *CompareContext) EQUALS_TEST() antlr.TerminalNode {
	return s.GetToken(TinylangParserEQUALS_TEST, 0)
}

func (s *CompareContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.EnterCompare(s)
	}
}

func (s *CompareContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.ExitCompare(s)
	}
}

type FalseContext struct {
	*ExpressionContext
	e antlr.Token
}

func NewFalseContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FalseContext {
	var p = new(FalseContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *FalseContext) GetE() antlr.Token { return s.e }

func (s *FalseContext) SetE(v antlr.Token) { s.e = v }

func (s *FalseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FalseContext) VALUE_FALSE() antlr.TerminalNode {
	return s.GetToken(TinylangParserVALUE_FALSE, 0)
}

func (s *FalseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.EnterFalse(s)
	}
}

func (s *FalseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.ExitFalse(s)
	}
}

type AndOrContext struct {
	*ExpressionContext
	op antlr.Token
}

func NewAndOrContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndOrContext {
	var p = new(AndOrContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *AndOrContext) GetOp() antlr.Token { return s.op }

func (s *AndOrContext) SetOp(v antlr.Token) { s.op = v }

func (s *AndOrContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndOrContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AndOrContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AndOrContext) AND_TEST() antlr.TerminalNode {
	return s.GetToken(TinylangParserAND_TEST, 0)
}

func (s *AndOrContext) OR_TEST() antlr.TerminalNode {
	return s.GetToken(TinylangParserOR_TEST, 0)
}

func (s *AndOrContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.EnterAndOr(s)
	}
}

func (s *AndOrContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.ExitAndOr(s)
	}
}

type ExprIdentContext struct {
	*ExpressionContext
	e antlr.Token
}

func NewExprIdentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprIdentContext {
	var p = new(ExprIdentContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExprIdentContext) GetE() antlr.Token { return s.e }

func (s *ExprIdentContext) SetE(v antlr.Token) { s.e = v }

func (s *ExprIdentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprIdentContext) IDENT() antlr.TerminalNode {
	return s.GetToken(TinylangParserIDENT, 0)
}

func (s *ExprIdentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.EnterExprIdent(s)
	}
}

func (s *ExprIdentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(TinylangListener); ok {
		listenerT.ExitExprIdent(s)
	}
}

func (p *TinylangParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *TinylangParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, TinylangParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(56)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case TinylangParserNUMBER:
		localctx = NewNumberContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(47)

			var _m = p.Match(TinylangParserNUMBER)

			localctx.(*NumberContext).e = _m
		}

	case TinylangParserVALUE_TRUE:
		localctx = NewTrueContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(48)

			var _m = p.Match(TinylangParserVALUE_TRUE)

			localctx.(*TrueContext).e = _m
		}

	case TinylangParserVALUE_FALSE:
		localctx = NewFalseContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(49)

			var _m = p.Match(TinylangParserVALUE_FALSE)

			localctx.(*FalseContext).e = _m
		}

	case TinylangParserIDENT:
		localctx = NewExprIdentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(50)

			var _m = p.Match(TinylangParserIDENT)

			localctx.(*ExprIdentContext).e = _m
		}

	case TinylangParserSTRING:
		localctx = NewExprStringContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(51)

			var _m = p.Match(TinylangParserSTRING)

			localctx.(*ExprStringContext).e = _m
		}

	case TinylangParserPARENTHESIS_OPEN:
		localctx = NewParenthesisContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(52)
			p.Match(TinylangParserPARENTHESIS_OPEN)
		}
		{
			p.SetState(53)
			p.expression(0)
		}
		{
			p.SetState(54)
			p.Match(TinylangParserPARENTHESIS_CLOSE)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(70)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TinylangParserRULE_expression)
				p.SetState(58)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(59)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TinylangParserMUL)|(1<<TinylangParserDIV)|(1<<TinylangParserMOD))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(60)
					p.expression(11)
				}

			case 2:
				localctx = NewAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TinylangParserRULE_expression)
				p.SetState(61)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(62)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TinylangParserADD || _la == TinylangParserSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(63)
					p.expression(10)
				}

			case 3:
				localctx = NewCompareContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TinylangParserRULE_expression)
				p.SetState(64)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(65)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*CompareContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<TinylangParserEQUALS_TEST)|(1<<TinylangParserGREATER_THAN)|(1<<TinylangParserGREATER_OR_EQUALS)|(1<<TinylangParserLESS_THAN)|(1<<TinylangParserLESS_OR_EQUALS))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*CompareContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(66)
					p.expression(9)
				}

			case 4:
				localctx = NewAndOrContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, TinylangParserRULE_expression)
				p.SetState(67)

				if !(p.Precpred(p.GetParserRuleContext(), 7)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 7)", ""))
				}
				{
					p.SetState(68)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AndOrContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == TinylangParserAND_TEST || _la == TinylangParserOR_TEST) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AndOrContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(69)
					p.expression(8)
				}

			}

		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
	}

	return localctx
}

func (p *TinylangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *TinylangParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 7)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
