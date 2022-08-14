package main

import (
	"fmt"
	"github.com/abarhub/goparser/tinylang/parser"
	"log"
	"os"
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
	EXPR_CODE_MOD
	EXPR_CODE_STR
	EXPR_CODE_LT
	EXPR_CODE_LTE
	EXPR_CODE_GT
	EXPR_CODE_GTE
	EXPR_CODE_EQU
	EXPR_CODE_NEQ
	EXPR_CODE_TRUE
	EXPR_CODE_FALSE
	EXPR_CODE_AND
	EXPR_CODE_OR
	EXPR_CODE_NOT
)

var conv_operator = map[int]ExprCode{parser.TinylangParserMUL: EXPR_CODE_MUL,
	parser.TinylangParserDIV:               EXPR_CODE_DIV,
	parser.TinylangParserADD:               EXPR_CODE_ADD,
	parser.TinylangParserSUB:               EXPR_CODE_SUB,
	parser.TinylangParserMOD:               EXPR_CODE_MOD,
	parser.TinylangParserEQUALS_TEST:       EXPR_CODE_EQU,
	parser.TinylangParserNOT_EQUALS_TEST:   EXPR_CODE_NEQ,
	parser.TinylangParserGREATER_THAN:      EXPR_CODE_GT,
	parser.TinylangParserGREATER_OR_EQUALS: EXPR_CODE_GTE,
	parser.TinylangParserLESS_THAN:         EXPR_CODE_LT,
	parser.TinylangParserLESS_OR_EQUALS:    EXPR_CODE_LTE,
	parser.TinylangParserAND_TEST:          EXPR_CODE_AND,
	parser.TinylangParserOR_TEST:           EXPR_CODE_OR,
	parser.TinylangParserNOT_TEST:          EXPR_CODE_NOT,
}

type Expression struct {
	code         ExprCode
	valeurInt    int
	variable     string
	valeurString string
	left         *Expression
	right        *Expression
	position     *Position
	returnType   *Type
}

func main() {

	functionList, err := Parser(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("nb funct", len(functionList))

	//functionList := functionList
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

			fmt.Printf("compile llvm ...\n")
			err = compileLLVM(functionList)
			fmt.Printf("compile llvm ok\n")
			if err != nil {
				fmt.Println("error compile", err)
			}
		}
	}
}
