package main

import (
	"fmt"
	"os"
	"reflect"

	"whilego/parser"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func getPrints(expr antlr.Tree) map[string]bool {
	res := map[string]bool{}

	tokens := expr.GetChildren()

	switch e := expr.(type) {
	case *parser.AttribContext:
		return res

	case *parser.WriteContext:
		return getPrints(tokens[1])

	case *parser.PrintContext:
		return res

	case *parser.IdContext:
		res[e.GetText()] = true

	case *parser.IfContext:
		return mergeMap(getPrints(tokens[3]), getPrints(tokens[5]))

	case *parser.WhileContext:
		return getPrints(tokens[3])

	case *parser.SkipContext:
		return res

	case *parser.BlockContext:
		return getPrints(tokens[1])

	case *parser.SeqStatementContext:
		for i := 0; i < len(tokens); i++ {
			res = mergeMap(res, getPrints(tokens[i]))
		}

	case *parser.ExpParenContext:
		return getPrints(tokens[1])

	case *antlr.TerminalNodeImpl:
		{
			switch e.GetText() {
			case ";":
				return res

			default:
				res[e.GetText()] = true
			}
		}

	default:
		fmt.Print("Not implemented get prints() -- ")
		fmt.Println(reflect.TypeOf(expr))
	}

	return res
}

func getOutputs(expr antlr.Tree) map[string]bool {
	res := map[string]bool{}

	tokens := expr.GetChildren()

	switch e := expr.(type) {
	case *parser.AttribContext:
		return getOutputs(tokens[0])

	case *parser.WriteContext:
		return res

	case *parser.PrintContext:
		return res

	case *parser.IdContext:
		res[e.GetText()] = true

	case *parser.IfContext:
		return mergeMap(getOutputs(tokens[3]), getOutputs(tokens[5]))

	case *parser.WhileContext:
		return getOutputs(tokens[3])

	case *parser.SkipContext:
		return res

	case *parser.BlockContext:
		return getOutputs(tokens[1])

	case *parser.SeqStatementContext:
		for i := 0; i < len(tokens); i++ {
			res = mergeMap(res, getOutputs(tokens[i]))
		}

	case *antlr.TerminalNodeImpl:
		{
			switch e.GetText() {
			case ";":
				return res
			default:
				res[e.GetText()] = true
			}
		}

	default:
		fmt.Print("Not implemented get outputs() -- ")
		fmt.Println(reflect.TypeOf(expr))

	}

	return res
}

func getInputs(expr antlr.Tree) map[string]bool {
	res := map[string]bool{}

	tokens := expr.GetChildren()

	switch e := expr.(type) {
	case *parser.AttribContext:
		return getInputs(tokens[2])

	case *parser.WriteContext:
		return getInputs(tokens[1])

	case *parser.PrintContext:
		return res

	case *parser.IdContext:
		res[e.GetText()] = true

	case *parser.ReadContext:
		return res

	case *parser.BinOpContext:
		return mergeMap(getInputs(tokens[0]), getInputs(tokens[2]))

	case *parser.ExpParenContext:
		return getInputs(tokens[1])

	case *parser.IntContext:
		return res

	case *parser.IfContext:
		return mergeMap(getInputs(tokens[1]), mergeMap(getInputs(tokens[3]), getInputs(tokens[5])))

	case *parser.WhileContext:
		return mergeMap(getInputs(tokens[1]), getInputs(tokens[3]))

	case *parser.SkipContext:
		return res

	case *parser.BlockContext:
		return getInputs(tokens[1])

	case *parser.SeqStatementContext:
		{
			for i := 0; i < len(tokens); i++ {
				res = mergeMap(res, getInputs(tokens[i]))
			}
		}

	case *parser.BoolParenContext:
		return getInputs(tokens[1])

	case *parser.RelOpContext:
		return mergeMap(getInputs(tokens[0]), getInputs(tokens[2]))

	case *antlr.TerminalNodeImpl:
		{
			switch e.GetText() {
			case ";":
				return res

			default:
				res[e.GetText()] = true
			}
		}

	default:
		fmt.Print("Not implemented get inputs() -- ")
		fmt.Println(reflect.TypeOf(expr))
	}

	return res
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])

	lexer := parser.NewWhilelangLexer(input)

	stream := antlr.NewCommonTokenStream(lexer, 0)

	p := parser.NewWhilelangParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	tree := p.Program()

	spiedVars := map[string]bool{}

	for i := tree.GetChild(0).GetChildCount() - 1; i >= 0; i-- {
		line := tree.GetChild(0).GetChild(i)
		inputs := getInputs(line)
		outputs := getOutputs(line)
		prints := getPrints(line)

		if (reflect.TypeOf(line) != reflect.TypeOf(&parser.PrintContext{})) &&
			(reflect.TypeOf(line) != reflect.TypeOf(&antlr.TerminalNodeImpl{})) {
			if hasCommon(spiedVars, outputs) || len(prints) > 0 {
				spiedVars = subMap(spiedVars, outputs)
				spiedVars = mergeMap(spiedVars, prints)
				spiedVars = mergeMap(spiedVars, inputs)
			} else {
				switch e := line.(type) {
				case *parser.AttribContext:
					fmt.Print("Dead attribution to \"")
					fmt.Print(e.GetChildren()[0])
					fmt.Println("\"")
				case *parser.IfContext:
					fmt.Println("Dead If Statement")
				case *parser.WhileContext:
					fmt.Println("Dead While Statement")
				}

			}
		}
	}
}

func copyMap(dst, pat map[string]bool) {
	for k, v := range pat {
		dst[k] = v
	}
}

func mergeMap(lhs, rhs map[string]bool) map[string]bool {
	res := map[string]bool{}

	copyMap(res, lhs)

	for k, v := range rhs {
		el, ok := res[k]
		if !ok || !el {
			res[k] = v
		}
	}

	return res
}

func subMap(lhs, rhs map[string]bool) map[string]bool {
	res := map[string]bool{}

	copyMap(res, lhs)

	for k, v := range rhs {
		el, ok := res[k]
		if ok && v == el {
			res[k] = false
		}
	}

	return res
}

func hasCommon(lhs, rhs map[string]bool) bool {
	for k, v := range lhs {
		el, ok := rhs[k]
		if ok && v == el && v {
			return true
		}
	}

	return false
}
