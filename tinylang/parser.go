package main

import (
	"fmt"
	"github.com/abarhub/goparser/tinylang/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"strconv"
	"strings"
)

type customErrorListener struct {
	errorMessage []string
}

func (c *customErrorListener) Error() string {
	if len(c.errorMessage) == 0 {
		return ""
	}
	return strings.Join(c.errorMessage, ",")
}

func (c *customErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.errorMessage = append(c.errorMessage, fmt.Sprint("error: line "+strconv.Itoa(line)+":"+strconv.Itoa(column)+" "+msg))
	//fmt.Println("error:" + msg)
	//fmt.Printf("var:%p", &c)
}

func (c *customErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	c.errorMessage = append(c.errorMessage, "error parsing")
	fmt.Println("error:")
}

func (c *customErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	c.errorMessage = append(c.errorMessage, "error parsing")
	fmt.Println("error:")
}

func (c *customErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	c.errorMessage = append(c.errorMessage, "error parsing")
	fmt.Println("error:")
}

func Parser(filename string, listener antlr.ParseTreeListener) error {
	fs, err := antlr.NewFileStream(filename)
	if err != nil {
		return err
	}
	lex := parser.NewTinylangLexer(fs)
	var errorListener customErrorListener
	//fmt.Printf("var:%p", &errorListener)
	lex.AddErrorListener(&errorListener)
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewTinylangParser(tokens)
	p.BuildParseTrees = true
	p.AddErrorListener(&errorListener)
	tree := p.Start()
	antlr.ParseTreeWalkerDefault.Walk(listener, tree)
	fmt.Println("res error:", errorListener)
	if len(errorListener.errorMessage) == 0 {
		return nil
	} else {
		return &errorListener
	}
}
