// Code generated from Tinylang.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Tinylang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// TinylangListener is a complete listener for a parse tree produced by TinylangParser.
type TinylangListener interface {
	antlr.ParseTreeListener

	// EnterStartFunction is called when entering the StartFunction production.
	EnterStartFunction(c *StartFunctionContext)

	// EnterFunct is called when entering the Funct production.
	EnterFunct(c *FunctContext)

	// EnterListInstruction is called when entering the ListInstruction production.
	EnterListInstruction(c *ListInstructionContext)

	// EnterInstrAffect is called when entering the InstrAffect production.
	EnterInstrAffect(c *InstrAffectContext)

	// EnterInstrDeclare is called when entering the InstrDeclare production.
	EnterInstrDeclare(c *InstrDeclareContext)

	// EnterTypeVoid is called when entering the TypeVoid production.
	EnterTypeVoid(c *TypeVoidContext)

	// EnterTypeInt is called when entering the TypeInt production.
	EnterTypeInt(c *TypeIntContext)

	// EnterTypeString is called when entering the TypeString production.
	EnterTypeString(c *TypeStringContext)

	// EnterTypeBoolean is called when entering the TypeBoolean production.
	EnterTypeBoolean(c *TypeBooleanContext)

	// EnterParenthesis is called when entering the Parenthesis production.
	EnterParenthesis(c *ParenthesisContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterExprString is called when entering the ExprString production.
	EnterExprString(c *ExprStringContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterTrue is called when entering the True production.
	EnterTrue(c *TrueContext)

	// EnterCompare is called when entering the Compare production.
	EnterCompare(c *CompareContext)

	// EnterFalse is called when entering the False production.
	EnterFalse(c *FalseContext)

	// EnterAndOr is called when entering the AndOr production.
	EnterAndOr(c *AndOrContext)

	// EnterExprIdent is called when entering the ExprIdent production.
	EnterExprIdent(c *ExprIdentContext)

	// ExitStartFunction is called when exiting the StartFunction production.
	ExitStartFunction(c *StartFunctionContext)

	// ExitFunct is called when exiting the Funct production.
	ExitFunct(c *FunctContext)

	// ExitListInstruction is called when exiting the ListInstruction production.
	ExitListInstruction(c *ListInstructionContext)

	// ExitInstrAffect is called when exiting the InstrAffect production.
	ExitInstrAffect(c *InstrAffectContext)

	// ExitInstrDeclare is called when exiting the InstrDeclare production.
	ExitInstrDeclare(c *InstrDeclareContext)

	// ExitTypeVoid is called when exiting the TypeVoid production.
	ExitTypeVoid(c *TypeVoidContext)

	// ExitTypeInt is called when exiting the TypeInt production.
	ExitTypeInt(c *TypeIntContext)

	// ExitTypeString is called when exiting the TypeString production.
	ExitTypeString(c *TypeStringContext)

	// ExitTypeBoolean is called when exiting the TypeBoolean production.
	ExitTypeBoolean(c *TypeBooleanContext)

	// ExitParenthesis is called when exiting the Parenthesis production.
	ExitParenthesis(c *ParenthesisContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitExprString is called when exiting the ExprString production.
	ExitExprString(c *ExprStringContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitTrue is called when exiting the True production.
	ExitTrue(c *TrueContext)

	// ExitCompare is called when exiting the Compare production.
	ExitCompare(c *CompareContext)

	// ExitFalse is called when exiting the False production.
	ExitFalse(c *FalseContext)

	// ExitAndOr is called when exiting the AndOr production.
	ExitAndOr(c *AndOrContext)

	// ExitExprIdent is called when exiting the ExprIdent production.
	ExitExprIdent(c *ExprIdentContext)
}
