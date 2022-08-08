package main

import (
	"fmt"
	"github.com/abarhub/goparser/tinylang/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"log"
	"os"
	"strconv"
)

type TypeCode int

const (
	TYPE_INT TypeCode = iota
	TYPE_VOID
	TYPE_STRING
	TYPE_BOOLEAN
)

type InstructionCode int

const (
	INSTRUCTION_AFFECTATION InstructionCode = iota
	INSTRUCTION_CALL
)

type Position struct {
	line   int
	column int
	pos    int
}

type Type struct {
	code     TypeCode
	position *Position
}

type Function struct {
	ReturnType  Type
	Name        string
	Instruction []*Instruction
	position    *Position
}

type Instruction struct {
	Code         InstructionCode
	FunctionName string
	Variable     string
	Valeur       *Expression
	Parameter    []*Expression
	position     *Position
}

type ExprCode int

const (
	EXPR_CODE_INT ExprCode = iota
	EXPR_CODE_VAR
	EXPR_CODE_ADD
	EXPR_CODE_SUB
	EXPR_CODE_MUL
	EXPR_CODE_DIV
	EXPR_CODE_STR
	EXPR_CODE_LT
	EXPR_CODE_LTE
	EXPR_CODE_GT
	EXPR_CODE_GTE
	EXPR_CODE_EQU
	EXPR_CODE_TRUE
	EXPR_CODE_FALSE
)

type Expression struct {
	code         ExprCode
	valeurInt    int
	variable     string
	valeurString string
	left         *Expression
	right        *Expression
	position     *Position
}

type calcListener struct {
	*parser.BaseTinylangListener

	stackExpr    []*Expression
	stackInstr   []*Instruction
	stackType    []*Type
	functionList []*Function
}

func (l *calcListener) pushExpr(e *Expression) {
	l.stackExpr = append(l.stackExpr, e)

}

func (l *calcListener) popExpr() *Expression {
	if len(l.stackExpr) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stackExpr[len(l.stackExpr)-1]

	// Remove the last element from the stack.
	l.stackExpr = l.stackExpr[:len(l.stackExpr)-1]

	return result
}

func (l *calcListener) pushInstr(e *Instruction) {
	l.stackInstr = append(l.stackInstr, e)
}

func (l *calcListener) pushType(i *Type) {
	l.stackType = append(l.stackType, i)
}

func (l *calcListener) popType() *Type {
	if len(l.stackType) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stackType[len(l.stackType)-1]

	// Remove the last element from the stack.
	l.stackType = l.stackType[:len(l.stackType)-1]

	return result
}

func (l *calcListener) addFunction(i *Function) {
	l.functionList = append(l.functionList, i)
}

func (l *calcListener) ExitMulDiv(c *parser.MulDivContext) {
	expr, expr2 := l.popExpr(), l.popExpr()

	var code ExprCode
	switch c.GetOp().GetTokenType() {
	case parser.TinylangParserMUL:
		code = EXPR_CODE_MUL
	case parser.TinylangParserDIV:
		code = EXPR_CODE_DIV
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
	expr3 := Expression{code: code, left: expr, right: expr2, position: getPosition(c.GetOp())}
	l.pushExpr(&expr3)
}

func (l *calcListener) ExitTrue(c *parser.TrueContext) {
	expr := Expression{code: EXPR_CODE_TRUE, position: getPosition(c.GetStart())}
	l.pushExpr(&expr)
}

func (l *calcListener) ExitFalse(c *parser.FalseContext) {
	expr := Expression{code: EXPR_CODE_FALSE, position: getPosition(c.GetStart())}
	l.pushExpr(&expr)
}

func (l *calcListener) ExitAddSub(c *parser.AddSubContext) {
	expr, expr2 := l.popExpr(), l.popExpr()

	var code ExprCode
	switch c.GetOp().GetTokenType() {
	case parser.TinylangParserADD:
		code = EXPR_CODE_ADD
	case parser.TinylangParserSUB:
		code = EXPR_CODE_SUB
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
	expr3 := Expression{code: code, left: expr, right: expr2, position: getPosition(c.GetOp())}
	l.pushExpr(&expr3)
}

func (l *calcListener) ExitCompare(c *parser.CompareContext) {
	expr, expr2 := l.popExpr(), l.popExpr()

	var code ExprCode
	switch c.GetOp().GetTokenType() {
	case parser.TinylangParserEQUALS_TEST:
		code = EXPR_CODE_EQU
	case parser.TinylangParserGREATER_THAN:
		code = EXPR_CODE_GT
	case parser.TinylangParserGREATER_OR_EQUALS:
		code = EXPR_CODE_GTE
	case parser.TinylangParserLESS_THAN:
		code = EXPR_CODE_LT
	case parser.TinylangParserLESS_OR_EQUALS:
		code = EXPR_CODE_LTE
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
	expr3 := Expression{code: code, left: expr, right: expr2, position: getPosition(c.GetOp())}
	l.pushExpr(&expr3)
}

func (l *calcListener) ExitNumber(c *parser.NumberContext) {
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err.Error())
	}

	expr := Expression{code: EXPR_CODE_INT, valeurInt: i, position: getPosition(c.GetStart())}
	l.pushExpr(&expr)
}

func (l *calcListener) ExitExprIdent(c *parser.ExprIdentContext) {
	name := c.GetText()
	//fmt.Printf("ident: %v;%v;%v;%v;%v;%v;%v\n", name, c, c.GetStart(), c.GetStart().GetLine(), c.GetStart().GetColumn(), c.GetStart().GetStart(), c.GetStart().GetStop())
	//position:=Position{line: c.GetStart().GetLine(),column: c.GetStart().GetColumn(),pos: c.GetStart().GetStart()}
	expr := Expression{code: EXPR_CODE_VAR, variable: name, position: getPosition(c.GetStart())}
	l.pushExpr(&expr)
}

func (l *calcListener) ExitExprString(c *parser.ExprStringContext) {
	str := c.GetText()
	str = str[1 : len(str)-1]
	expr := Expression{code: EXPR_CODE_STR, valeurString: str, position: getPosition(c.GetStart())}
	l.pushExpr(&expr)
}

func getPosition(token antlr.Token) (position *Position) {
	position = &Position{line: token.GetLine(), column: token.GetColumn(), pos: token.GetStart()}
	return
}

func (l *calcListener) ExitInstrAffect(c *parser.InstrAffectContext) {
	res3 := l.popExpr()
	instr := Instruction{Code: INSTRUCTION_AFFECTATION, Variable: c.IDENT().GetText(),
		Valeur: res3, position: getPosition(c.GetStart())}
	l.pushInstr(&instr)
	fmt.Println("affect", instr)
}

func (l *calcListener) EnterTypeVoid(c *parser.TypeVoidContext) {
	l.pushType(&Type{code: TYPE_VOID, position: getPosition(c.GetStart())})
}

func (l *calcListener) EnterTypeInt(c *parser.TypeIntContext) {
	l.pushType(&Type{code: TYPE_INT, position: getPosition(c.GetStart())})
}

func (l *calcListener) EnterTypeString(c *parser.TypeStringContext) {
	l.pushType(&Type{code: TYPE_STRING, position: getPosition(c.GetStart())})
}

func (l *calcListener) EnterTypeBoolean(c *parser.TypeBooleanContext) {
	l.pushType(&Type{code: TYPE_BOOLEAN, position: getPosition(c.GetStart())})
}

func (l *calcListener) ExitFunct(c *parser.FunctContext) {
	res := c.GetChild(1)
	typeReturn := l.popType()
	funct := Function{Name: c.GetName().GetText(), Instruction: l.stackInstr,
		ReturnType: *typeReturn, position: getPosition(c.GetStart())}
	l.stackInstr = []*Instruction{}
	l.addFunction(&funct)
	fmt.Println("funct", res.GetPayload(), res.GetPayload(), funct)
}

func main() {

	fs, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	lex := parser.NewTinylangLexer(fs)
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewTinylangParser(tokens)
	p.BuildParseTrees = true
	tree := p.Start()
	var listener calcListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, tree)
	//fmt.Println("res", listener.pop())

	fmt.Println("nb funct", len(listener.functionList))

	functionList := listener.functionList
	err = Checker(functionList)
	if err != nil {
		fmt.Println("error checker", err)
	} else {
		interpreter := NewInterpreter(functionList)
		var symbolTable []map[string]Valeur
		symbolTable, err = interpreter.interpreter()
		if err != nil {
			fmt.Println("error interpreter", err)
		} else {
			fmt.Printf("symbolTable: %v\n", symbolTable)
		}
	}
}
