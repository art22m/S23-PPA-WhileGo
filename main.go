package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"os"
	"whilego/parser"
)

type TreeShapeListener struct {
	*parser.BaseWhilelangListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])

	lexer := parser.NewWhilelangLexer(input)

	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewWhilelangParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true

	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), p.Program())
}
