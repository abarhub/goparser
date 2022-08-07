// for vistor pattern
package main

import (
	"fmt"
	"github.com/abarhub/goparser/tinylang/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
	"os"
	"strconv"
)

//type Visitor struct {
//	parser.BaseTinylangVisitor
//	stack []int
//}

//func NewVisitor() *Visitor {
//	return &Visitor{}
//}
//
//func (l *Visitor) push(i int) {
//	l.stack = append(l.stack, i)
//}
//
//func (l *Visitor) pop() int {
//	if len(l.stack) < 1 {
//		panic("stack is empty unable to pop")
//	}
//
//	// Get the last value from the stack.
//	result := l.stack[len(l.stack)-1]
//
//	// Remove the last element from the stack.
//	l.stack = l.stack[:len(l.stack)-1]
//
//	return result
//}
//
//func (l *Visitor) length() int {
//	return len(l.stack)
//}
//
//func (v *Visitor) visitRule(node antlr.RuleNode) interface{} {
//	node.Accept(v)
//	return nil
//}
//
//func (v *Visitor) VisitStart(ctx *parser.StartContext) interface{} {
//	v.visitRule(ctx.GetRuleContext())
//	return "abc"
//}
//
//func (v *Visitor) VisitNumber(ctx *parser.NumberContext) interface{} {
//	i, err := strconv.Atoi(ctx.NUMBER().GetText())
//	if err != nil {
//		panic(err.Error())
//	}
//
//	v.push(i)
//	return nil
//}
//
//func (v *Visitor) VisitMulDiv(ctx *parser.MulDivContext) interface{} {
//	//push expression result to stack
//	v.visitRule(ctx.Expression(0))
//	v.visitRule(ctx.Expression(1))
//
//	//push result to stack
//	var t antlr.Token = ctx.GetOp()
//	right := v.pop()
//	left := v.pop()
//	switch t.GetTokenType() {
//	case parser.TinylangParserMUL:
//		v.push(left * right)
//	case parser.TinylangParserDIV:
//		v.push(left / right)
//	default:
//		panic("should not happen")
//
//	}
//
//	return nil
//}
//
//func (v *Visitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
//	//push expression result to stack
//	v.visitRule(ctx.Expression(0))
//	v.visitRule(ctx.Expression(1))
//
//	//push result to stack
//	var t antlr.Token = ctx.GetOp()
//	right := v.pop()
//	left := v.pop()
//	switch t.GetTokenType() {
//	case parser.TinylangParserADD:
//		v.push(left + right)
//	case parser.TinylangParserSUB:
//		v.push(left - right)
//	default:
//		panic("should not happen")
//	}
//
//	return nil
//}
//
//func (v *Visitor) VisitParenthesis(ctx *parser.ParenthesisContext) interface{} {
//	v.visitRule(ctx.Expression())
//	return nil
//}

/*func calc(input string) int {

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
}*/

//func completer(in prompt.Document) []prompt.Suggest {
//	var ret []prompt.Suggest
//	return ret
//}

type calcListener struct {
	*parser.BaseTinylangListener

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
	case parser.TinylangParserMUL:
		l.push(left * right)
	case parser.TinylangParserDIV:
		l.push(left / right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *calcListener) ExitAddSub(c *parser.AddSubContext) {
	right, left := l.pop(), l.pop()

	switch c.GetOp().GetTokenType() {
	case parser.TinylangParserADD:
		l.push(left + right)
	case parser.TinylangParserSUB:
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

func (l *calcListener) ExitInstrAffect(c *parser.InstrAffectContext) {
	res := c.GetChild(0)
	variable := res.GetPayload()
	res2 := c.GetChild(2)
	fmt.Println("affect", variable, res2, c.IDENT().GetText())
}

func (l *calcListener) ExitFunct(c *parser.FunctContext) {
	res := c.GetChild(1)
	fmt.Println("funct", res.GetPayload(), res.GetPayload())
}

type TreeShapeListener struct {
	*parser.BaseTinylangListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println("text", ctx.GetText())

}

func main() {

	fs, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	lex := parser.NewTinylangLexer(fs)
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewTinylangParser(tokens)
	//v := NewVisitor()
	//visitor := parser.NewCalcBaseCalcVisitor{}
	//tree := p.Start().Accept(v)
	//parser.BuildParseTrees = true
	p.BuildParseTrees = true
	tree := p.Start()
	//var result = v.Visit(tree)
	//antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
	var listener calcListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, tree)
	fmt.Println("res", listener.pop())
	//p.Start().Accept(v)
	//fmt.Println("result", result)
	//fmt.Println("len", v.length())
	//for i := 0; i < v.length(); i++ {
	//	fmt.Println("res", v.pop())
	//}
	//fmt.Println(v.Visit(tree))

	//p := prompt.New(
	//	executor,
	//	completer,
	//	prompt.OptionPrefix(">>> "),
	//	prompt.OptionTitle("calc"),
	//)
	//p.Run()
}
