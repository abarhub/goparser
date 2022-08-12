package main

import (
	"fmt"
	"github.com/ichiban/prolog"
	"os"
	"strconv"
	"strings"
)

func test1() {
	p := prolog.New(os.Stdin, os.Stdout)

	if err := p.Exec(`
	human(socrates).       % This is a fact.
	mortal(X) :- human(X). % This is a rule.
`); err != nil {
		panic(err)
	} else {
		//return
	}

	sols, err := p.Query(`mortal(Who).`)
	if err != nil {
		panic(err)
	}
	defer sols.Close()

	// Iterates over solutions.
	for sols.Next() {
		// Prepare a struct with fields which name corresponds with a variable in the query.
		var s struct {
			Who string
		}
		if err := sols.Scan(&s); err != nil {
			panic(err)
		}
		fmt.Printf("Who = %s\n", s.Who) // ==> Who = socrates
	}

	// Check if an error occurred while querying.
	if err := sols.Err(); err != nil {
		panic(err)
	}
}

func Checker2(functionList []*Function) error {

	for _, function := range functionList {
		var varDeclared, varUsed map[string]bool
		varDeclared = make(map[string]bool)
		varUsed = make(map[string]bool)
		var kb []string = []string{}
		for no, instr := range function.Instruction {
			if instr.Code == INSTRUCTION_AFFECTATION {
				variable := instr.Variable
				if _, ok := varDeclared[variable]; !ok {
					varDeclared[variable] = true
					kb = append(kb, "declare("+variable+","+strconv.Itoa(no)+"). % declare "+variable)
				}
			}
			if instr.Valeur != nil {
				var val = instr.Valeur
				if val.code == EXPR_CODE_VAR {
					variable := val.variable
					if _, ok := varUsed[variable]; !ok {
						varUsed[variable] = true
						kb = append(kb, "used("+variable+","+strconv.Itoa(no)+"). % used "+variable)
					}
				}
			}
		}
		//fmt.Printf("prolog: %v\n", kb)
		//checkVar(kb)
	}
	//test1()
	return nil
}

func checkVar(kb []string) {
	p := prolog.New(os.Stdin, os.Stdout)

	kb2 := strings.Join(kb[:], "\n")
	kb2 = kb2 + "\nusedOk(X) :- declare(X,_),used(X,_)."
	println("kb2:", kb2)
	if err := p.Exec(kb2); err != nil {
		panic(err)
	} else {
		//return
	}

	sols, err := p.Query(`usedOk(Who).`)
	if err != nil {
		panic(err)
	}
	defer sols.Close()

	// Iterates over solutions.
	for sols.Next() {
		// Prepare a struct with fields which name corresponds with a variable in the query.
		var s struct {
			Who string
		}
		if err := sols.Scan(&s); err != nil {
			panic(err)
		}
		fmt.Printf("Who = %s\n", s.Who) // ==> Who = socrates
	}

	// Check if an error occurred while querying.
	if err := sols.Err(); err != nil {
		panic(err)
	}

}
