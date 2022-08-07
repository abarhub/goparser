// for vistor pattern
package main

import (
	"fmt"
	"github.com/abarhub/goparser/visitor/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
	"os"
	"strconv"
)

type Visitor struct {
	parser.BaseCalcVisitor
	stack []int
}

func NewVisitor() *Visitor {
	return &Visitor{}
}

func (l *Visitor) push(i int) {
	l.stack = append(l.stack, i)
}

func (l *Visitor) pop() int {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (l *Visitor) length() int {
	return len(l.stack)
}

func (v *Visitor) visitRule(node antlr.RuleNode) interface{} {
	node.Accept(v)
	return nil
}

func (v *Visitor) VisitStart(ctx *parser.StartContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *Visitor) VisitNumber(ctx *parser.NumberContext) interface{} {
	i, err := strconv.Atoi(ctx.NUMBER().GetText())
	if err != nil {
		panic(err.Error())
	}

	v.push(i)
	return nil
}

func (v *Visitor) VisitMulDiv(ctx *parser.MulDivContext) interface{} {
	//push expression result to stack
	v.visitRule(ctx.Expression(0))
	v.visitRule(ctx.Expression(1))

	//push result to stack
	var t antlr.Token = ctx.GetOp()
	right := v.pop()
	left := v.pop()
	switch t.GetTokenType() {
	case parser.CalcParserMUL:
		v.push(left * right)
	case parser.CalcParserDIV:
		v.push(left / right)
	default:
		panic("should not happen")

	}

	return nil
}

func (v *Visitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
	//push expression result to stack
	v.visitRule(ctx.Expression(0))
	v.visitRule(ctx.Expression(1))

	//push result to stack
	var t antlr.Token = ctx.GetOp()
	right := v.pop()
	left := v.pop()
	switch t.GetTokenType() {
	case parser.CalcParserADD:
		v.push(left + right)
	case parser.CalcParserSUB:
		v.push(left - right)
	default:
		panic("should not happen")
	}

	return nil
}

func (v *Visitor) VisitParenthesis(ctx *parser.ParenthesisContext) interface{} {
	v.visitRule(ctx.Expression())
	return nil
}

func calc(input string) int {

	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(tokens)

	v := NewVisitor()
	//Start is rule name of Calc.g4
	p.Start().Accept(v)
	return v.pop()
}

func executor(in string) {
	fmt.Printf("Answer: %d\n", calc(in))
}

//func completer(in prompt.Document) []prompt.Suggest {
//	var ret []prompt.Suggest
//	return ret
//}

func main() {

	fs, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	lex := parser.NewCalcLexer(fs)
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewCalcParser(tokens)
	v := NewVisitor()
	//visitor := parser.NewCalcBaseCalcVisitor{}
	//tree := p.Start().Accept(v)
	p.Start().Accept(v)
	for i := 0; i < v.length(); i++ {
		fmt.Println("res", v.pop())
	}
	//fmt.Println(v.Visit(tree))

	//p := prompt.New(
	//	executor,
	//	completer,
	//	prompt.OptionPrefix(">>> "),
	//	prompt.OptionTitle("calc"),
	//)
	//p.Run()
}
