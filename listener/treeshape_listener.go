package listener

import (
	. "whilego/ast_type"
	"whilego/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

type TreeShapeListener struct {
	*parser.BaseWhilelangListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

var root Node = Node{}

type Node struct {
	NodeType WhileGoType
	children []Node
	content  string
}

// VisitTerminal is called when a terminal node is visited.
func (s *TreeShapeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *TreeShapeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *TreeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *TreeShapeListener) EnterProgram(ctx *parser.ProgramContext) {
	root.NodeType = TypeProgram
}

// ExitProgram is called when production program is exited.
func (s *TreeShapeListener) ExitProgram(ctx *parser.ProgramContext) {
	// Проверятор вызвать
}

// EnterSeqStatement is called when production seqStatement is entered.
func (s *TreeShapeListener) EnterSeqStatement(ctx *parser.SeqStatementContext) {
	//s.EnterEveryRule ctx.GetChildren()[0]
}

// ExitSeqStatement is called when production seqStatement is exited.
func (s *TreeShapeListener) ExitSeqStatement(ctx *parser.SeqStatementContext) {}

// EnterAttrib is called when production attrib is entered.
func (s *TreeShapeListener) EnterAttrib(ctx *parser.AttribContext) {}

// ExitAttrib is called when production attrib is exited.
func (s *TreeShapeListener) ExitAttrib(ctx *parser.AttribContext) {}

// EnterSkip is called when production skip is entered.
func (s *TreeShapeListener) EnterSkip(ctx *parser.SkipContext) {}

// ExitSkip is called when production skip is exited.
func (s *TreeShapeListener) ExitSkip(ctx *parser.SkipContext) {}

// EnterIf is called when production if is entered.
func (s *TreeShapeListener) EnterIf(ctx *parser.IfContext) {}

// ExitIf is called when production if is exited.
func (s *TreeShapeListener) ExitIf(ctx *parser.IfContext) {}

// EnterWhile is called when production while is entered.
func (s *TreeShapeListener) EnterWhile(ctx *parser.WhileContext) {}

// ExitWhile is called when production while is exited.
func (s *TreeShapeListener) ExitWhile(ctx *parser.WhileContext) {}

// EnterPrint is called when production print is entered.
func (s *TreeShapeListener) EnterPrint(ctx *parser.PrintContext) {}

// ExitPrint is called when production print is exited.
func (s *TreeShapeListener) ExitPrint(ctx *parser.PrintContext) {}

// EnterWrite is called when production write is entered.
func (s *TreeShapeListener) EnterWrite(ctx *parser.WriteContext) {}

// ExitWrite is called when production write is exited.
func (s *TreeShapeListener) ExitWrite(ctx *parser.WriteContext) {}

// EnterBlock is called when production block is entered.
func (s *TreeShapeListener) EnterBlock(ctx *parser.BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *TreeShapeListener) ExitBlock(ctx *parser.BlockContext) {}

// EnterRead is called when production read is entered.
func (s *TreeShapeListener) EnterRead(ctx *parser.ReadContext) {}

// ExitRead is called when production read is exited.
func (s *TreeShapeListener) ExitRead(ctx *parser.ReadContext) {}

// EnterId is called when production id is entered.
func (s *TreeShapeListener) EnterId(ctx *parser.IdContext) {}

// ExitId is called when production id is exited.
func (s *TreeShapeListener) ExitId(ctx *parser.IdContext) {}

// EnterExpParen is called when production expParen is entered.
func (s *TreeShapeListener) EnterExpParen(ctx *parser.ExpParenContext) {}

// ExitExpParen is called when production expParen is exited.
func (s *TreeShapeListener) ExitExpParen(ctx *parser.ExpParenContext) {}

// EnterInt is called when production int is entered.
func (s *TreeShapeListener) EnterInt(ctx *parser.IntContext) {}

// ExitInt is called when production int is exited.
func (s *TreeShapeListener) ExitInt(ctx *parser.IntContext) {}

// EnterBinOp is called when production binOp is entered.
func (s *TreeShapeListener) EnterBinOp(ctx *parser.BinOpContext) {}

// ExitBinOp is called when production binOp is exited.
func (s *TreeShapeListener) ExitBinOp(ctx *parser.BinOpContext) {}

// EnterNot is called when production not is entered.
func (s *TreeShapeListener) EnterNot(ctx *parser.NotContext) {}

// ExitNot is called when production not is exited.
func (s *TreeShapeListener) ExitNot(ctx *parser.NotContext) {}

// EnterBoolean is called when production boolean is entered.
func (s *TreeShapeListener) EnterBoolean(ctx *parser.BooleanContext) {}

// ExitBoolean is called when production boolean is exited.
func (s *TreeShapeListener) ExitBoolean(ctx *parser.BooleanContext) {}

// EnterAnd is called when production and is entered.
func (s *TreeShapeListener) EnterAnd(ctx *parser.AndContext) {}

// ExitAnd is called when production and is exited.
func (s *TreeShapeListener) ExitAnd(ctx *parser.AndContext) {}

// EnterBoolParen is called when production boolParen is entered.
func (s *TreeShapeListener) EnterBoolParen(ctx *parser.BoolParenContext) {}

// ExitBoolParen is called when production boolParen is exited.
func (s *TreeShapeListener) ExitBoolParen(ctx *parser.BoolParenContext) {}

// EnterRelOp is called when production relOp is entered.
func (s *TreeShapeListener) EnterRelOp(ctx *parser.RelOpContext) {}

// ExitRelOp is called when production relOp is exited.
func (s *TreeShapeListener) ExitRelOp(ctx *parser.RelOpContext) {}
