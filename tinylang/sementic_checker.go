package main

import "fmt"

var typeBoolean = Type{code: TYPE_BOOLEAN}
var typeInt = Type{code: TYPE_INT}
var typeString = Type{code: TYPE_STRING}
var typeVoid = Type{code: TYPE_VOID}

func showPosition(position *Position) string {
	if position == nil {
		return ""
	} else {
		return fmt.Sprintf("(line:%d, column:%d)", position.line, position.column)
	}
}

func checkExpression(expression *Expression, symbolTable map[string]*Type) (*Type, error) {
	var val = expression
	if val.code == EXPR_CODE_VAR {
		variable := val.variable
		if typeVar, ok := symbolTable[variable]; !ok {
			return nil, fmt.Errorf("error: variable %s not declared "+
				showPosition(expression.position), variable)
		} else {
			expression.returnType = typeVar
			return typeVar, nil
		}
	} else if val.code == EXPR_CODE_TRUE || val.code == EXPR_CODE_FALSE {
		expression.returnType = &typeBoolean
		return &typeBoolean, nil
	} else if val.code == EXPR_CODE_INT {
		expression.returnType = &typeInt
		return &typeInt, nil
	} else if val.code == EXPR_CODE_STR {
		expression.returnType = &typeString
		return &typeString, nil
	} else if expression.code == EXPR_CODE_ADD || expression.code == EXPR_CODE_SUB ||
		expression.code == EXPR_CODE_EQU || expression.code == EXPR_CODE_LT ||
		expression.code == EXPR_CODE_LTE ||
		expression.code == EXPR_CODE_GT || expression.code == EXPR_CODE_GTE ||
		expression.code == EXPR_CODE_MUL || expression.code == EXPR_CODE_DIV ||
		expression.code == EXPR_CODE_MOD || expression.code == EXPR_CODE_AND ||
		expression.code == EXPR_CODE_OR || expression.code == EXPR_CODE_NEQ {
		typeLeft, err2 := checkExpression(expression.left, symbolTable)
		if err2 != nil {
			return nil, err2
		} else if typeLeft == nil {
			return nil, fmt.Errorf("error: expression type left nil : %v "+
				showPosition(expression.position), expression)
		}
		typeRight, err2 := checkExpression(expression.right, symbolTable)
		if err2 != nil {
			return nil, err2
		} else if typeRight == nil {
			return nil, fmt.Errorf("error: expression type right nil : %v", expression)
		}
		if expression.code == EXPR_CODE_ADD || expression.code == EXPR_CODE_SUB ||
			expression.code == EXPR_CODE_MUL || expression.code == EXPR_CODE_DIV ||
			expression.code == EXPR_CODE_MOD {
			if typeLeft.code == TYPE_INT && typeRight.code == TYPE_INT {
				expression.returnType = &typeInt
				return &typeInt, nil
			} else {
				return nil, fmt.Errorf("error: operator '%v' need int parameters "+
					showPosition(expression.position), expression.code)
			}
		} else if expression.code == EXPR_CODE_LT ||
			expression.code == EXPR_CODE_LTE || expression.code == EXPR_CODE_GT ||
			expression.code == EXPR_CODE_GTE {
			if typeLeft.code == TYPE_INT && typeRight.code == TYPE_INT {
				expression.returnType = &typeBoolean
				return &typeBoolean, nil
			} else {
				return nil, fmt.Errorf("error: operator '%v' need int parameters "+
					showPosition(expression.position), expression.code)
			}
		} else if expression.code == EXPR_CODE_AND || expression.code == EXPR_CODE_OR {
			if typeLeft.code == TYPE_BOOLEAN && typeRight.code == TYPE_BOOLEAN {
				expression.returnType = &typeBoolean
				return &typeBoolean, nil
			} else {
				return nil, fmt.Errorf("error: operator '%v' need boolean parameters "+
					showPosition(expression.position), expression.code)
			}
		} else if expression.code == EXPR_CODE_EQU || expression.code == EXPR_CODE_NEQ {
			if typeLeft.code == TYPE_BOOLEAN && typeRight.code == TYPE_BOOLEAN {
				expression.returnType = &typeBoolean
				return &typeBoolean, nil
			} else if typeLeft.code == TYPE_INT && typeRight.code == TYPE_INT {
				expression.returnType = &typeBoolean
				return &typeBoolean, nil
			} else {
				return nil, fmt.Errorf("error: operator '%v' need boolean parameters "+
					showPosition(expression.position), expression.code)
			}
		} else {
			return nil, fmt.Errorf("error: invalid operator '%v' "+
				showPosition(expression.position), expression.code)
		}
	} else if expression.code == EXPR_CODE_NOT {
		typeLeft, err2 := checkExpression(expression.left, symbolTable)
		if err2 != nil {
			return nil, err2
		} else if typeLeft == nil {
			return nil, fmt.Errorf("error: expression type nil : %v "+
				showPosition(expression.position), expression)
		}
		if expression.code == EXPR_CODE_NOT {
			if typeLeft.code == TYPE_BOOLEAN {
				expression.returnType = &typeBoolean
				return &typeBoolean, nil
			} else {
				return nil, fmt.Errorf("error: operator '%v' need boolean parameter "+
					showPosition(expression.position), expression.code)
			}
		} else {
			return nil, fmt.Errorf("error: invalid operator '%v' "+
				showPosition(expression.position), expression.code)
		}
	} else {
		return nil, fmt.Errorf("error: expression unknow : %v "+
			showPosition(expression.position), val)
	}
}

func Checker(functionList []*Function) error {

	for _, function := range functionList {
		var varDeclared = make(map[string]*Type)

		for _, instr := range function.Instruction {
			if instr.Code == INSTRUCTION_AFFECTATION {

				var typeVar *Type
				var err error
				variable := instr.Variable
				if instr.Valeur != nil {
					typeVar, err = checkExpression(instr.Valeur, varDeclared)
					if err != nil {
						return err
					}
				} else {
					return fmt.Errorf("error: variable %s not affected "+
						showPosition(instr.position), variable)
				}

				if typeVar == nil {
					return fmt.Errorf("error: variable %s not affected "+showPosition(instr.position), variable)
				}
				if varTypeDeclared, ok := varDeclared[variable]; !ok {
					varDeclared[variable] = typeVar
				} else if varTypeDeclared.code != typeVar.code {
					return fmt.Errorf("error: invalide type for variable %v "+
						showPosition(instr.position), variable)
				}
			} else if instr.Valeur != nil {
				var val = instr.Valeur
				if val.code == EXPR_CODE_VAR {
					variable := val.variable
					if _, ok := varDeclared[variable]; !ok {
						return fmt.Errorf("error: variable %s not declared "+
							showPosition(instr.position), variable)
					}
				}
			}
		}
	}
	return nil
}
