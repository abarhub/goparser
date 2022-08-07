// package listen
package main

import (
	"fmt"
	"github.com/abarhub/goparser/listen/parser"
	"os"
	"strconv"

	//"parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type calcListener struct {
	*parser.BaseCalcListener

	stack []int
}

func (l *calcListener) push(i int) {
	l.stack = append(l.stack, i)

}

func (l *calcListener) pop() int {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (l *calcListener) ExitMulDiv(c *parser.MulDivContext) {
	right, left := l.pop(), l.pop()

	switch c.GetOp().GetTokenType() {
	case parser.CalcParserMUL:
		l.push(left * right)
	case parser.CalcParserDIV:
		l.push(left / right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *calcListener) ExitAddSub(c *parser.AddSubContext) {
	right, left := l.pop(), l.pop()

	switch c.GetOp().GetTokenType() {
	case parser.CalcParserADD:
		l.push(left + right)
	case parser.CalcParserSUB:
		l.push(left - right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *calcListener) ExitNumber(c *parser.NumberContext) {
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err.Error())
	}

	l.push(i)
}

// calc takes a string expression and returns the evaluated result.
func calc(input string) int {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(stream)

	// Finally parse the expression (by walking the tree)
	var listener calcListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	return listener.pop()

}

func executor(in string) {
	fmt.Printf("Answer: %d\n", calc(in))
}

//func completer(in prompt.Document) []prompt.Suggest {
//	var ret []prompt.Suggest
//	return ret
//}

type TreeShapeListener struct {
	*parser.BaseCalcListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println("text", ctx.GetText())
}

func main() {
	//p := prompt.New(
	//	executor,
	//	completer,
	//	prompt.OptionPrefix(">>> "),
	//	prompt.OptionTitle("calc"),
	//)
	//p.Run()

	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parser.NewCalcLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCalcParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Start()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
