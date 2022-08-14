// Code generated from C:/projet/goparser\Tinylang.g4 by ANTLR 4.10.1. DO NOT EDIT.

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

// EnterTypeInt is called when production TypeInt is entered.
func (s *BaseTinylangListener) EnterTypeInt(ctx *TypeIntContext) {}

// ExitTypeInt is called when production TypeInt is exited.
func (s *BaseTinylangListener) ExitTypeInt(ctx *TypeIntContext) {}

// EnterTypeString is called when production TypeString is entered.
func (s *BaseTinylangListener) EnterTypeString(ctx *TypeStringContext) {}

// ExitTypeString is called when production TypeString is exited.
func (s *BaseTinylangListener) ExitTypeString(ctx *TypeStringContext) {}

// EnterTypeBoolean is called when production TypeBoolean is entered.
func (s *BaseTinylangListener) EnterTypeBoolean(ctx *TypeBooleanContext) {}

// ExitTypeBoolean is called when production TypeBoolean is exited.
func (s *BaseTinylangListener) ExitTypeBoolean(ctx *TypeBooleanContext) {}

// EnterNot is called when production Not is entered.
func (s *BaseTinylangListener) EnterNot(ctx *NotContext) {}

// ExitNot is called when production Not is exited.
func (s *BaseTinylangListener) ExitNot(ctx *NotContext) {}

// EnterParenthesis is called when production Parenthesis is entered.
func (s *BaseTinylangListener) EnterParenthesis(ctx *ParenthesisContext) {}

// ExitParenthesis is called when production Parenthesis is exited.
func (s *BaseTinylangListener) ExitParenthesis(ctx *ParenthesisContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseTinylangListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseTinylangListener) ExitNumber(ctx *NumberContext) {}

// EnterExprString is called when production ExprString is entered.
func (s *BaseTinylangListener) EnterExprString(ctx *ExprStringContext) {}

// ExitExprString is called when production ExprString is exited.
func (s *BaseTinylangListener) ExitExprString(ctx *ExprStringContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseTinylangListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseTinylangListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseTinylangListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseTinylangListener) ExitAddSub(ctx *AddSubContext) {}

// EnterTrue is called when production True is entered.
func (s *BaseTinylangListener) EnterTrue(ctx *TrueContext) {}

// ExitTrue is called when production True is exited.
func (s *BaseTinylangListener) ExitTrue(ctx *TrueContext) {}

// EnterCompare is called when production Compare is entered.
func (s *BaseTinylangListener) EnterCompare(ctx *CompareContext) {}

// ExitCompare is called when production Compare is exited.
func (s *BaseTinylangListener) ExitCompare(ctx *CompareContext) {}

// EnterFalse is called when production False is entered.
func (s *BaseTinylangListener) EnterFalse(ctx *FalseContext) {}

// ExitFalse is called when production False is exited.
func (s *BaseTinylangListener) ExitFalse(ctx *FalseContext) {}

// EnterAndOr is called when production AndOr is entered.
func (s *BaseTinylangListener) EnterAndOr(ctx *AndOrContext) {}

// ExitAndOr is called when production AndOr is exited.
func (s *BaseTinylangListener) ExitAndOr(ctx *AndOrContext) {}

// EnterExprIdent is called when production ExprIdent is entered.
func (s *BaseTinylangListener) EnterExprIdent(ctx *ExprIdentContext) {}

// ExitExprIdent is called when production ExprIdent is exited.
func (s *BaseTinylangListener) ExitExprIdent(ctx *ExprIdentContext) {}
