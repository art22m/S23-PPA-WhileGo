// Code generated from Whilelang.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Whilelang

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// BaseWhilelangListener is a complete listener for a parse tree produced by WhilelangParser.
type BaseWhilelangListener struct{}

var _ WhilelangListener = &BaseWhilelangListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseWhilelangListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseWhilelangListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseWhilelangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseWhilelangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseWhilelangListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseWhilelangListener) ExitProgram(ctx *ProgramContext) {}

// EnterSeqStatement is called when production seqStatement is entered.
func (s *BaseWhilelangListener) EnterSeqStatement(ctx *SeqStatementContext) {}

// ExitSeqStatement is called when production seqStatement is exited.
func (s *BaseWhilelangListener) ExitSeqStatement(ctx *SeqStatementContext) {}

// EnterAttrib is called when production attrib is entered.
func (s *BaseWhilelangListener) EnterAttrib(ctx *AttribContext) {}

// ExitAttrib is called when production attrib is exited.
func (s *BaseWhilelangListener) ExitAttrib(ctx *AttribContext) {}

// EnterSkip is called when production skip is entered.
func (s *BaseWhilelangListener) EnterSkip(ctx *SkipContext) {}

// ExitSkip is called when production skip is exited.
func (s *BaseWhilelangListener) ExitSkip(ctx *SkipContext) {}

// EnterIf is called when production if is entered.
func (s *BaseWhilelangListener) EnterIf(ctx *IfContext) {}

// ExitIf is called when production if is exited.
func (s *BaseWhilelangListener) ExitIf(ctx *IfContext) {}

// EnterWhile is called when production while is entered.
func (s *BaseWhilelangListener) EnterWhile(ctx *WhileContext) {}

// ExitWhile is called when production while is exited.
func (s *BaseWhilelangListener) ExitWhile(ctx *WhileContext) {}

// EnterPrint is called when production print is entered.
func (s *BaseWhilelangListener) EnterPrint(ctx *PrintContext) {}

// ExitPrint is called when production print is exited.
func (s *BaseWhilelangListener) ExitPrint(ctx *PrintContext) {}

// EnterWrite is called when production write is entered.
func (s *BaseWhilelangListener) EnterWrite(ctx *WriteContext) {}

// ExitWrite is called when production write is exited.
func (s *BaseWhilelangListener) ExitWrite(ctx *WriteContext) {}

// EnterBlock is called when production block is entered.
func (s *BaseWhilelangListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BaseWhilelangListener) ExitBlock(ctx *BlockContext) {}

// EnterRead is called when production read is entered.
func (s *BaseWhilelangListener) EnterRead(ctx *ReadContext) {}

// ExitRead is called when production read is exited.
func (s *BaseWhilelangListener) ExitRead(ctx *ReadContext) {}

// EnterId is called when production id is entered.
func (s *BaseWhilelangListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseWhilelangListener) ExitId(ctx *IdContext) {}

// EnterExpParen is called when production expParen is entered.
func (s *BaseWhilelangListener) EnterExpParen(ctx *ExpParenContext) {}

// ExitExpParen is called when production expParen is exited.
func (s *BaseWhilelangListener) ExitExpParen(ctx *ExpParenContext) {}

// EnterInt is called when production int is entered.
func (s *BaseWhilelangListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production int is exited.
func (s *BaseWhilelangListener) ExitInt(ctx *IntContext) {}

// EnterBinOp is called when production binOp is entered.
func (s *BaseWhilelangListener) EnterBinOp(ctx *BinOpContext) {}

// ExitBinOp is called when production binOp is exited.
func (s *BaseWhilelangListener) ExitBinOp(ctx *BinOpContext) {}

// EnterNot is called when production not is entered.
func (s *BaseWhilelangListener) EnterNot(ctx *NotContext) {}

// ExitNot is called when production not is exited.
func (s *BaseWhilelangListener) ExitNot(ctx *NotContext) {}

// EnterBoolean is called when production boolean is entered.
func (s *BaseWhilelangListener) EnterBoolean(ctx *BooleanContext) {}

// ExitBoolean is called when production boolean is exited.
func (s *BaseWhilelangListener) ExitBoolean(ctx *BooleanContext) {}

// EnterAnd is called when production and is entered.
func (s *BaseWhilelangListener) EnterAnd(ctx *AndContext) {}

// ExitAnd is called when production and is exited.
func (s *BaseWhilelangListener) ExitAnd(ctx *AndContext) {}

// EnterBoolParen is called when production boolParen is entered.
func (s *BaseWhilelangListener) EnterBoolParen(ctx *BoolParenContext) {}

// ExitBoolParen is called when production boolParen is exited.
func (s *BaseWhilelangListener) ExitBoolParen(ctx *BoolParenContext) {}

// EnterRelOp is called when production relOp is entered.
func (s *BaseWhilelangListener) EnterRelOp(ctx *RelOpContext) {}

// ExitRelOp is called when production relOp is exited.
func (s *BaseWhilelangListener) ExitRelOp(ctx *RelOpContext) {}
