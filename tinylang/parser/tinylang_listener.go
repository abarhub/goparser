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

	// EnterTYPEINT is called when entering the TYPEINT production.
	EnterTYPEINT(c *TYPEINTContext)

	// EnterTypeString is called when entering the TypeString production.
	EnterTypeString(c *TypeStringContext)

	// EnterParenthesis is called when entering the Parenthesis production.
	EnterParenthesis(c *ParenthesisContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

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

	// ExitTYPEINT is called when exiting the TYPEINT production.
	ExitTYPEINT(c *TYPEINTContext)

	// ExitTypeString is called when exiting the TypeString production.
	ExitTypeString(c *TypeStringContext)

	// ExitParenthesis is called when exiting the Parenthesis production.
	ExitParenthesis(c *ParenthesisContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)
}
