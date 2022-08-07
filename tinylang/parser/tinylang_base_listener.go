// Code generated from Tinylang.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // Tinylang

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseTinylangListener is a complete listener for a parse tree produced by TinylangParser.
type BaseTinylangListener struct{}

var _ TinylangListener = &BaseTinylangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTinylangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTinylangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTinylangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTinylangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStartFunction is called when production StartFunction is entered.
func (s *BaseTinylangListener) EnterStartFunction(ctx *StartFunctionContext) {}

// ExitStartFunction is called when production StartFunction is exited.
func (s *BaseTinylangListener) ExitStartFunction(ctx *StartFunctionContext) {}

// EnterFunct is called when production Funct is entered.
func (s *BaseTinylangListener) EnterFunct(ctx *FunctContext) {}

// ExitFunct is called when production Funct is exited.
func (s *BaseTinylangListener) ExitFunct(ctx *FunctContext) {}

// EnterListInstruction is called when production ListInstruction is entered.
func (s *BaseTinylangListener) EnterListInstruction(ctx *ListInstructionContext) {}

// ExitListInstruction is called when production ListInstruction is exited.
func (s *BaseTinylangListener) ExitListInstruction(ctx *ListInstructionContext) {}

// EnterInstrAffect is called when production InstrAffect is entered.
func (s *BaseTinylangListener) EnterInstrAffect(ctx *InstrAffectContext) {}

// ExitInstrAffect is called when production InstrAffect is exited.
func (s *BaseTinylangListener) ExitInstrAffect(ctx *InstrAffectContext) {}

// EnterInstrDeclare is called when production InstrDeclare is entered.
func (s *BaseTinylangListener) EnterInstrDeclare(ctx *InstrDeclareContext) {}

// ExitInstrDeclare is called when production InstrDeclare is exited.
func (s *BaseTinylangListener) ExitInstrDeclare(ctx *InstrDeclareContext) {}

// EnterTypeVoid is called when production TypeVoid is entered.
func (s *BaseTinylangListener) EnterTypeVoid(ctx *TypeVoidContext) {}

// ExitTypeVoid is called when production TypeVoid is exited.
func (s *BaseTinylangListener) ExitTypeVoid(ctx *TypeVoidContext) {}

// EnterTYPEINT is called when production TYPEINT is entered.
func (s *BaseTinylangListener) EnterTYPEINT(ctx *TYPEINTContext) {}

// ExitTYPEINT is called when production TYPEINT is exited.
func (s *BaseTinylangListener) ExitTYPEINT(ctx *TYPEINTContext) {}

// EnterTypeString is called when production TypeString is entered.
func (s *BaseTinylangListener) EnterTypeString(ctx *TypeStringContext) {}

// ExitTypeString is called when production TypeString is exited.
func (s *BaseTinylangListener) ExitTypeString(ctx *TypeStringContext) {}

// EnterParenthesis is called when production Parenthesis is entered.
func (s *BaseTinylangListener) EnterParenthesis(ctx *ParenthesisContext) {}

// ExitParenthesis is called when production Parenthesis is exited.
func (s *BaseTinylangListener) ExitParenthesis(ctx *ParenthesisContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseTinylangListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseTinylangListener) ExitNumber(ctx *NumberContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseTinylangListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseTinylangListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseTinylangListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseTinylangListener) ExitAddSub(ctx *AddSubContext) {}
