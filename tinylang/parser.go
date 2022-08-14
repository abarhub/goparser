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
	fmt.Println("error:" + msg)
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

type calcListener struct {
	*parser.BaseTinylangListener

	stackExpr    []*Expression
	stackInstr   []*Instruction
	stackType    []*Type
	functionList []*Function
}

func (l calcListener) pushExpr(e *Expression) {
	l.stackExpr = append(l.stackExpr, e)
}

func (l calcListener) popExpr() *Expression {
	if len(l.stackExpr) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stackExpr[len(l.stackExpr)-1]

	// Remove the last element from the stack.
	l.stackExpr = l.stackExpr[:len(l.stackExpr)-1]

	return result
}

func (l calcListener) pushInstr(e *Instruction) {
	l.stackInstr = append(l.stackInstr, e)
}

func (l calcListener) pushType(i *Type) {
	l.stackType = append(l.stackType, i)
}

func (l calcListener) popType() *Type {
	if len(l.stackType) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stackType[len(l.stackType)-1]

	// Remove the last element from the stack.
	l.stackType = l.stackType[:len(l.stackType)-1]

	return result
}

func (l calcListener) addFunction(i *Function) {
	l.functionList = append(l.functionList, i)
}

func (l calcListener) ExitMulDiv(c *parser.MulDivContext) {
	expr, expr2 := l.popExpr(), l.popExpr()

	var code ExprCode
	switch c.GetOp().GetTokenType() {
	case parser.TinylangParserMUL:
	case parser.TinylangParserDIV:
	case parser.TinylangParserMOD:
		value, ok := conv_operator[c.GetOp().GetTokenType()]
		if ok {
			code = value
		} else {
			panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
		}
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
	expr3 := Expression{code: code, left: expr, right: expr2, position: getPosition(c.GetOp())}
	l.pushExpr(&expr3)
}

func (l calcListener) ExitTrue(c *parser.TrueContext) {
	expr := Expression{code: EXPR_CODE_TRUE, position: getPosition(c.GetStart())}
	l.pushExpr(&expr)
}

func (l calcListener) ExitFalse(c *parser.FalseContext) {
	expr := Expression{code: EXPR_CODE_FALSE, position: getPosition(c.GetStart())}
	l.pushExpr(&expr)
}

func (l calcListener) ExitAddSub(c *parser.AddSubContext) {
	expr, expr2 := l.popExpr(), l.popExpr()

	var code ExprCode
	switch c.GetOp().GetTokenType() {
	case parser.TinylangParserADD:
	case parser.TinylangParserSUB:
		value, ok := conv_operator[c.GetOp().GetTokenType()]
		if ok {
			code = value
		} else {
			panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
		}
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
	expr3 := Expression{code: code, left: expr, right: expr2, position: getPosition(c.GetOp())}
	l.pushExpr(&expr3)
}

func (l calcListener) ExitCompare(c *parser.CompareContext) {
	expr, expr2 := l.popExpr(), l.popExpr()

	var code ExprCode
	switch c.GetOp().GetTokenType() {
	case parser.TinylangParserEQUALS_TEST:
	case parser.TinylangParserNOT_EQUALS_TEST:
	case parser.TinylangParserGREATER_THAN:
	case parser.TinylangParserGREATER_OR_EQUALS:
	case parser.TinylangParserLESS_THAN:
	case parser.TinylangParserLESS_OR_EQUALS:
		value, ok := conv_operator[c.GetOp().GetTokenType()]
		if ok {
			code = value
		} else {
			panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
		}
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
	expr3 := Expression{code: code, left: expr, right: expr2, position: getPosition(c.GetOp())}
	l.pushExpr(&expr3)
}

func (l calcListener) ExitAndOr(c *parser.AndOrContext) {
	expr, expr2 := l.popExpr(), l.popExpr()

	var code ExprCode
	switch c.GetOp().GetTokenType() {
	case parser.TinylangParserAND_TEST:
	case parser.TinylangParserOR_TEST:
		value, ok := conv_operator[c.GetOp().GetTokenType()]
		if ok {
			code = value
		} else {
			panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
		}
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
	expr3 := Expression{code: code, left: expr, right: expr2, position: getPosition(c.GetOp())}
	l.pushExpr(&expr3)
}

func (l calcListener) ExitNot(c *parser.NotContext) {
	expr := l.popExpr()

	var code ExprCode
	switch c.GetOp().GetTokenType() {
	case parser.TinylangParserNOT_TEST:
		value, ok := conv_operator[c.GetOp().GetTokenType()]
		if ok {
			code = value
		} else {
			panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
		}
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
	expr3 := Expression{code: code, left: expr, position: getPosition(c.GetOp())}
	l.pushExpr(&expr3)
}

func (l calcListener) ExitNumber(c *parser.NumberContext) {
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err.Error())
	}

	expr := Expression{code: EXPR_CODE_INT, valeurInt: i, position: getPosition(c.GetStart())}
	l.pushExpr(&expr)
}

func (l calcListener) ExitExprIdent(c *parser.ExprIdentContext) {
	name := c.GetText()
	//fmt.Printf("ident: %v;%v;%v;%v;%v;%v;%v\n", name, c, c.GetStart(), c.GetStart().GetLine(), c.GetStart().GetColumn(), c.GetStart().GetStart(), c.GetStart().GetStop())
	//position:=Position{line: c.GetStart().GetLine(),column: c.GetStart().GetColumn(),pos: c.GetStart().GetStart()}
	expr := Expression{code: EXPR_CODE_VAR, variable: name, position: getPosition(c.GetStart())}
	l.pushExpr(&expr)
}

func (l calcListener) ExitExprString(c *parser.ExprStringContext) {
	str := c.GetText()
	str = str[1 : len(str)-1]
	expr := Expression{code: EXPR_CODE_STR, valeurString: str, position: getPosition(c.GetStart())}
	l.pushExpr(&expr)
}

func getPosition(token antlr.Token) (position *Position) {
	position = &Position{line: token.GetLine(), column: token.GetColumn(), pos: token.GetStart()}
	return
}

func (l calcListener) ExitInstrAffect(c *parser.InstrAffectContext) {
	res3 := l.popExpr()
	instr := Instruction{Code: INSTRUCTION_AFFECTATION, Variable: c.IDENT().GetText(),
		Valeur: res3, position: getPosition(c.GetStart())}
	l.pushInstr(&instr)
	fmt.Println("affect", instr)
}

func (l calcListener) EnterTypeVoid(c *parser.TypeVoidContext) {
	l.pushType(&Type{code: TYPE_VOID, position: getPosition(c.GetStart())})
}

func (l calcListener) EnterTypeInt(c *parser.TypeIntContext) {
	l.pushType(&Type{code: TYPE_INT, position: getPosition(c.GetStart())})
}

func (l calcListener) EnterTypeString(c *parser.TypeStringContext) {
	l.pushType(&Type{code: TYPE_STRING, position: getPosition(c.GetStart())})
}

func (l calcListener) EnterTypeBoolean(c *parser.TypeBooleanContext) {
	l.pushType(&Type{code: TYPE_BOOLEAN, position: getPosition(c.GetStart())})
}

func (l calcListener) ExitFunct(c *parser.FunctContext) {
	res := c.GetChild(1)
	typeReturn := l.popType()
	funct := Function{Name: c.GetName().GetText(), Instruction: l.stackInstr,
		ReturnType: *typeReturn, position: getPosition(c.GetStart())}
	l.stackInstr = []*Instruction{}
	l.addFunction(&funct)
	fmt.Println("funct", res.GetPayload(), res.GetPayload(), funct)
}

func (l calcListener) EnterStartFunction(c *parser.StartFunctionContext) {
	fmt.Printf("start2\n")
}

type TotoListener struct {
	*parser.BaseTinylangListener
	//parser.TinylangListener
}

func (l TotoListener) EnterStartFunction(c *parser.StartFunctionContext) {
	fmt.Printf("start\n")
}

func Parser(filename string) ([]*Function, error) {
	fs, err := antlr.NewFileStream(filename)
	if err != nil {
		return nil, err
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
	var listener calcListener = calcListener{}
	var tmp antlr.ParseTreeListener
	//var totoListener TotoListener = TotoListener{}
	tmp = listener
	//tmp = totoListener
	if listenerT, ok := tmp.(parser.TinylangListener); ok {
		fmt.Println("type=true", listenerT)
	} else {
		fmt.Println("type=false", listenerT)
	}
	antlr.ParseTreeWalkerDefault.Walk(&listener, tree)
	//antlr.ParseTreeWalkerDefault.Walk(tmp, tree)
	//antlr.ParseTreeWalkerDefault.Walk(&totoListener, tree)
	//fmt.Println("res error:", errorListener)
	if len(errorListener.errorMessage) == 0 {
		return listener.functionList, nil
	} else {
		return nil, &errorListener
	}
}
