package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"whilego/parser"
)

func main() {
	// Setup the input
	is := antlr.NewInputStream("1 + 2 * 3")

	// Create the Lexer
	lexer := parser.NewWhilelangLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewWhilelangParser(stream)

	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&calcListener{}, p.Start())
}
