package main

import "fmt"

func Checker(functionList []*Function) error {

	for _, function := range functionList {
		var varDeclared map[string]bool
		varDeclared = make(map[string]bool)

		for _, instr := range function.Instruction {
			if instr.Code == INSTRUCTION_AFFECTATION {
				variable := instr.Variable
				if _, ok := varDeclared[variable]; !ok {
					varDeclared[variable] = true
				}
			}
			if instr.Valeur != nil {
				var val = instr.Valeur
				if val.code == EXPR_CODE_VAR {
					variable := val.variable
					if _, ok := varDeclared[variable]; !ok {
						return fmt.Errorf("error: variable %s not declared", variable)
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
