// Code generated from Whilelang.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Whilelang

import "github.com/antlr/antlr4/runtime/Go/antlr/v4"

// WhilelangListener is a complete listener for a parse tree produced by WhilelangParser.
type WhilelangListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterSeqStatement is called when entering the seqStatement production.
	EnterSeqStatement(c *SeqStatementContext)

	// EnterAttrib is called when entering the attrib production.
	EnterAttrib(c *AttribContext)

	// EnterSkip is called when entering the skip production.
	EnterSkip(c *SkipContext)

	// EnterIf is called when entering the if production.
	EnterIf(c *IfContext)

	// EnterWhile is called when entering the while production.
	EnterWhile(c *WhileContext)

	// EnterPrint is called when entering the print production.
	EnterPrint(c *PrintContext)

	// EnterWrite is called when entering the write production.
	EnterWrite(c *WriteContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterRead is called when entering the read production.
	EnterRead(c *ReadContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterExpParen is called when entering the expParen production.
	EnterExpParen(c *ExpParenContext)

	// EnterInt is called when entering the int production.
	EnterInt(c *IntContext)

	// EnterBinOp is called when entering the binOp production.
	EnterBinOp(c *BinOpContext)

	// EnterNot is called when entering the not production.
	EnterNot(c *NotContext)

	// EnterBoolean is called when entering the boolean production.
	EnterBoolean(c *BooleanContext)

	// EnterAnd is called when entering the and production.
	EnterAnd(c *AndContext)

	// EnterBoolParen is called when entering the boolParen production.
	EnterBoolParen(c *BoolParenContext)

	// EnterRelOp is called when entering the relOp production.
	EnterRelOp(c *RelOpContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitSeqStatement is called when exiting the seqStatement production.
	ExitSeqStatement(c *SeqStatementContext)

	// ExitAttrib is called when exiting the attrib production.
	ExitAttrib(c *AttribContext)

	// ExitSkip is called when exiting the skip production.
	ExitSkip(c *SkipContext)

	// ExitIf is called when exiting the if production.
	ExitIf(c *IfContext)

	// ExitWhile is called when exiting the while production.
	ExitWhile(c *WhileContext)

	// ExitPrint is called when exiting the print production.
	ExitPrint(c *PrintContext)

	// ExitWrite is called when exiting the write production.
	ExitWrite(c *WriteContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitRead is called when exiting the read production.
	ExitRead(c *ReadContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitExpParen is called when exiting the expParen production.
	ExitExpParen(c *ExpParenContext)

	// ExitInt is called when exiting the int production.
	ExitInt(c *IntContext)

	// ExitBinOp is called when exiting the binOp production.
	ExitBinOp(c *BinOpContext)

	// ExitNot is called when exiting the not production.
	ExitNot(c *NotContext)

	// ExitBoolean is called when exiting the boolean production.
	ExitBoolean(c *BooleanContext)

	// ExitAnd is called when exiting the and production.
	ExitAnd(c *AndContext)

	// ExitBoolParen is called when exiting the boolParen production.
	ExitBoolParen(c *BoolParenContext)

	// ExitRelOp is called when exiting the relOp production.
	ExitRelOp(c *RelOpContext)
}
