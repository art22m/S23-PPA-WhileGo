// Code generated from Whilelang.g4 by ANTLR 4.12.0. DO NOT EDIT.

package parser // Whilelang

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type WhilelangParser struct {
	*antlr.BaseParser
}

var whilelangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	literalNames           []string
	symbolicNames          []string
	ruleNames              []string
	predictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func whilelangParserInit() {
	staticData := &whilelangParserStaticData
	staticData.literalNames = []string{
		"", "';'", "':='", "'skip'", "'if'", "'then'", "'else'", "'while'",
		"'do'", "'print'", "'{'", "'}'", "'read'", "'*'", "'+'", "'-'", "'('",
		"')'", "'true'", "'false'", "'='", "'<='", "'not'", "'and'",
	}
	staticData.symbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
		"", "", "", "", "", "", "", "INT", "ID", "Text", "Space",
	}
	staticData.ruleNames = []string{
		"program", "seqStatement", "statement", "expression", "bool",
	}
	staticData.predictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 27, 94, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 5, 1, 16, 8, 1, 10, 1, 12, 1, 19, 9, 1,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2,
		3, 2, 45, 8, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 55,
		8, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 5, 3, 63, 8, 3, 10, 3, 12, 3,
		66, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1,
		4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 3, 4, 84, 8, 4, 1, 4, 1, 4, 1, 4, 5, 4,
		89, 8, 4, 10, 4, 12, 4, 92, 9, 4, 1, 4, 0, 2, 6, 8, 5, 0, 2, 4, 6, 8, 0,
		2, 1, 0, 14, 15, 1, 0, 18, 19, 105, 0, 10, 1, 0, 0, 0, 2, 12, 1, 0, 0,
		0, 4, 44, 1, 0, 0, 0, 6, 54, 1, 0, 0, 0, 8, 83, 1, 0, 0, 0, 10, 11, 3,
		2, 1, 0, 11, 1, 1, 0, 0, 0, 12, 17, 3, 4, 2, 0, 13, 14, 5, 1, 0, 0, 14,
		16, 3, 4, 2, 0, 15, 13, 1, 0, 0, 0, 16, 19, 1, 0, 0, 0, 17, 15, 1, 0, 0,
		0, 17, 18, 1, 0, 0, 0, 18, 3, 1, 0, 0, 0, 19, 17, 1, 0, 0, 0, 20, 21, 5,
		25, 0, 0, 21, 22, 5, 2, 0, 0, 22, 45, 3, 6, 3, 0, 23, 45, 5, 3, 0, 0, 24,
		25, 5, 4, 0, 0, 25, 26, 3, 8, 4, 0, 26, 27, 5, 5, 0, 0, 27, 28, 3, 4, 2,
		0, 28, 29, 5, 6, 0, 0, 29, 30, 3, 4, 2, 0, 30, 45, 1, 0, 0, 0, 31, 32,
		5, 7, 0, 0, 32, 33, 3, 8, 4, 0, 33, 34, 5, 8, 0, 0, 34, 35, 3, 4, 2, 0,
		35, 45, 1, 0, 0, 0, 36, 37, 5, 9, 0, 0, 37, 45, 5, 26, 0, 0, 38, 39, 5,
		9, 0, 0, 39, 45, 3, 6, 3, 0, 40, 41, 5, 10, 0, 0, 41, 42, 3, 2, 1, 0, 42,
		43, 5, 11, 0, 0, 43, 45, 1, 0, 0, 0, 44, 20, 1, 0, 0, 0, 44, 23, 1, 0,
		0, 0, 44, 24, 1, 0, 0, 0, 44, 31, 1, 0, 0, 0, 44, 36, 1, 0, 0, 0, 44, 38,
		1, 0, 0, 0, 44, 40, 1, 0, 0, 0, 45, 5, 1, 0, 0, 0, 46, 47, 6, 3, -1, 0,
		47, 55, 5, 24, 0, 0, 48, 55, 5, 12, 0, 0, 49, 55, 5, 25, 0, 0, 50, 51,
		5, 16, 0, 0, 51, 52, 3, 6, 3, 0, 52, 53, 5, 17, 0, 0, 53, 55, 1, 0, 0,
		0, 54, 46, 1, 0, 0, 0, 54, 48, 1, 0, 0, 0, 54, 49, 1, 0, 0, 0, 54, 50,
		1, 0, 0, 0, 55, 64, 1, 0, 0, 0, 56, 57, 10, 3, 0, 0, 57, 58, 5, 13, 0,
		0, 58, 63, 3, 6, 3, 4, 59, 60, 10, 2, 0, 0, 60, 61, 7, 0, 0, 0, 61, 63,
		3, 6, 3, 3, 62, 56, 1, 0, 0, 0, 62, 59, 1, 0, 0, 0, 63, 66, 1, 0, 0, 0,
		64, 62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 7, 1, 0, 0, 0, 66, 64, 1, 0,
		0, 0, 67, 68, 6, 4, -1, 0, 68, 84, 7, 1, 0, 0, 69, 70, 3, 6, 3, 0, 70,
		71, 5, 20, 0, 0, 71, 72, 3, 6, 3, 0, 72, 84, 1, 0, 0, 0, 73, 74, 3, 6,
		3, 0, 74, 75, 5, 21, 0, 0, 75, 76, 3, 6, 3, 0, 76, 84, 1, 0, 0, 0, 77,
		78, 5, 22, 0, 0, 78, 84, 3, 8, 4, 3, 79, 80, 5, 16, 0, 0, 80, 81, 3, 8,
		4, 0, 81, 82, 5, 17, 0, 0, 82, 84, 1, 0, 0, 0, 83, 67, 1, 0, 0, 0, 83,
		69, 1, 0, 0, 0, 83, 73, 1, 0, 0, 0, 83, 77, 1, 0, 0, 0, 83, 79, 1, 0, 0,
		0, 84, 90, 1, 0, 0, 0, 85, 86, 10, 2, 0, 0, 86, 87, 5, 23, 0, 0, 87, 89,
		3, 8, 4, 3, 88, 85, 1, 0, 0, 0, 89, 92, 1, 0, 0, 0, 90, 88, 1, 0, 0, 0,
		90, 91, 1, 0, 0, 0, 91, 9, 1, 0, 0, 0, 92, 90, 1, 0, 0, 0, 7, 17, 44, 54,
		62, 64, 83, 90,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// WhilelangParserInit initializes any static state used to implement WhilelangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewWhilelangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func WhilelangParserInit() {
	staticData := &whilelangParserStaticData
	staticData.once.Do(whilelangParserInit)
}

// NewWhilelangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewWhilelangParser(input antlr.TokenStream) *WhilelangParser {
	WhilelangParserInit()
	this := new(WhilelangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &whilelangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.predictionContextCache)
	this.RuleNames = staticData.ruleNames
	this.LiteralNames = staticData.literalNames
	this.SymbolicNames = staticData.symbolicNames
	this.GrammarFileName = "Whilelang.g4"

	return this
}

// WhilelangParser tokens.
const (
	WhilelangParserEOF   = antlr.TokenEOF
	WhilelangParserT__0  = 1
	WhilelangParserT__1  = 2
	WhilelangParserT__2  = 3
	WhilelangParserT__3  = 4
	WhilelangParserT__4  = 5
	WhilelangParserT__5  = 6
	WhilelangParserT__6  = 7
	WhilelangParserT__7  = 8
	WhilelangParserT__8  = 9
	WhilelangParserT__9  = 10
	WhilelangParserT__10 = 11
	WhilelangParserT__11 = 12
	WhilelangParserT__12 = 13
	WhilelangParserT__13 = 14
	WhilelangParserT__14 = 15
	WhilelangParserT__15 = 16
	WhilelangParserT__16 = 17
	WhilelangParserT__17 = 18
	WhilelangParserT__18 = 19
	WhilelangParserT__19 = 20
	WhilelangParserT__20 = 21
	WhilelangParserT__21 = 22
	WhilelangParserT__22 = 23
	WhilelangParserINT   = 24
	WhilelangParserID    = 25
	WhilelangParserText  = 26
	WhilelangParserSpace = 27
)

// WhilelangParser rules.
const (
	WhilelangParserRULE_program      = 0
	WhilelangParserRULE_seqStatement = 1
	WhilelangParserRULE_statement    = 2
	WhilelangParserRULE_expression   = 3
	WhilelangParserRULE_bool         = 4
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SeqStatement() ISeqStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhilelangParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhilelangParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) SeqStatement() ISeqStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISeqStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISeqStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *WhilelangParser) Program() (localctx IProgramContext) {
	this := p
	_ = this

	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, WhilelangParserRULE_program)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(10)
		p.SeqStatement()
	}

	return localctx
}

// ISeqStatementContext is an interface to support dynamic dispatch.
type ISeqStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsSeqStatementContext differentiates from other interfaces.
	IsSeqStatementContext()
}

type SeqStatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySeqStatementContext() *SeqStatementContext {
	var p = new(SeqStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhilelangParserRULE_seqStatement
	return p
}

func (*SeqStatementContext) IsSeqStatementContext() {}

func NewSeqStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SeqStatementContext {
	var p = new(SeqStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhilelangParserRULE_seqStatement

	return p
}

func (s *SeqStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SeqStatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *SeqStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *SeqStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SeqStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SeqStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.EnterSeqStatement(s)
	}
}

func (s *SeqStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.ExitSeqStatement(s)
	}
}

func (p *WhilelangParser) SeqStatement() (localctx ISeqStatementContext) {
	this := p
	_ = this

	localctx = NewSeqStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, WhilelangParserRULE_seqStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(12)
		p.Statement()
	}
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == WhilelangParserT__0 {
		{
			p.SetState(13)
			p.Match(WhilelangParserT__0)
		}
		{
			p.SetState(14)
			p.Statement()
		}

		p.SetState(19)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhilelangParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhilelangParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyFrom(ctx *StatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrintContext struct {
	*StatementContext
}

func NewPrintContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintContext {
	var p = new(PrintContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) Text() antlr.TerminalNode {
	return s.GetToken(WhilelangParserText, 0)
}

func (s *PrintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.EnterPrint(s)
	}
}

func (s *PrintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.ExitPrint(s)
	}
}

type AttribContext struct {
	*StatementContext
}

func NewAttribContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AttribContext {
	var p = new(AttribContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *AttribContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttribContext) ID() antlr.TerminalNode {
	return s.GetToken(WhilelangParserID, 0)
}

func (s *AttribContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AttribContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.EnterAttrib(s)
	}
}

func (s *AttribContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.ExitAttrib(s)
	}
}

type SkipContext struct {
	*StatementContext
}

func NewSkipContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SkipContext {
	var p = new(SkipContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *SkipContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SkipContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.EnterSkip(s)
	}
}

func (s *SkipContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.ExitSkip(s)
	}
}

type BlockContext struct {
	*StatementContext
}

func NewBlockContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BlockContext {
	var p = new(BlockContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) SeqStatement() ISeqStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISeqStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISeqStatementContext)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.ExitBlock(s)
	}
}

type WhileContext struct {
	*StatementContext
}

func NewWhileContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileContext {
	var p = new(WhileContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *WhileContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileContext) Bool_() IBoolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolContext)
}

func (s *WhileContext) Statement() IStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *WhileContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.EnterWhile(s)
	}
}

func (s *WhileContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.ExitWhile(s)
	}
}

type IfContext struct {
	*StatementContext
}

func NewIfContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfContext {
	var p = new(IfContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *IfContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfContext) Bool_() IBoolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolContext)
}

func (s *IfContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *IfContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *IfContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.EnterIf(s)
	}
}

func (s *IfContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.ExitIf(s)
	}
}

type WriteContext struct {
	*StatementContext
}

func NewWriteContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WriteContext {
	var p = new(WriteContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *WriteContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WriteContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WriteContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.EnterWrite(s)
	}
}

func (s *WriteContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.ExitWrite(s)
	}
}

func (p *WhilelangParser) Statement() (localctx IStatementContext) {
	this := p
	_ = this

	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, WhilelangParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewAttribContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(20)
			p.Match(WhilelangParserID)
		}
		{
			p.SetState(21)
			p.Match(WhilelangParserT__1)
		}
		{
			p.SetState(22)
			p.expression(0)
		}

	case 2:
		localctx = NewSkipContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(23)
			p.Match(WhilelangParserT__2)
		}

	case 3:
		localctx = NewIfContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(24)
			p.Match(WhilelangParserT__3)
		}
		{
			p.SetState(25)
			p.bool_(0)
		}
		{
			p.SetState(26)
			p.Match(WhilelangParserT__4)
		}
		{
			p.SetState(27)
			p.Statement()
		}
		{
			p.SetState(28)
			p.Match(WhilelangParserT__5)
		}
		{
			p.SetState(29)
			p.Statement()
		}

	case 4:
		localctx = NewWhileContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(31)
			p.Match(WhilelangParserT__6)
		}
		{
			p.SetState(32)
			p.bool_(0)
		}
		{
			p.SetState(33)
			p.Match(WhilelangParserT__7)
		}
		{
			p.SetState(34)
			p.Statement()
		}

	case 5:
		localctx = NewPrintContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(36)
			p.Match(WhilelangParserT__8)
		}
		{
			p.SetState(37)
			p.Match(WhilelangParserText)
		}

	case 6:
		localctx = NewWriteContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(38)
			p.Match(WhilelangParserT__8)
		}
		{
			p.SetState(39)
			p.expression(0)
		}

	case 7:
		localctx = NewBlockContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(40)
			p.Match(WhilelangParserT__9)
		}
		{
			p.SetState(41)
			p.SeqStatement()
		}
		{
			p.SetState(42)
			p.Match(WhilelangParserT__10)
		}

	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhilelangParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhilelangParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ReadContext struct {
	*ExpressionContext
}

func NewReadContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReadContext {
	var p = new(ReadContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ReadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.EnterRead(s)
	}
}

func (s *ReadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.ExitRead(s)
	}
}

type IdContext struct {
	*ExpressionContext
}

func NewIdContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdContext {
	var p = new(IdContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IdContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdContext) ID() antlr.TerminalNode {
	return s.GetToken(WhilelangParserID, 0)
}

func (s *IdContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.EnterId(s)
	}
}

func (s *IdContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.ExitId(s)
	}
}

type ExpParenContext struct {
	*ExpressionContext
}

func NewExpParenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpParenContext {
	var p = new(ExpParenContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExpParenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpParenContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpParenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.EnterExpParen(s)
	}
}

func (s *ExpParenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.ExitExpParen(s)
	}
}

type IntContext struct {
	*ExpressionContext
}

func NewIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntContext {
	var p = new(IntContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntContext) INT() antlr.TerminalNode {
	return s.GetToken(WhilelangParserINT, 0)
}

func (s *IntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.EnterInt(s)
	}
}

func (s *IntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.ExitInt(s)
	}
}

type BinOpContext struct {
	*ExpressionContext
}

func NewBinOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinOpContext {
	var p = new(BinOpContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BinOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinOpContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BinOpContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.EnterBinOp(s)
	}
}

func (s *BinOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.ExitBinOp(s)
	}
}

func (p *WhilelangParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *WhilelangParser) expression(_p int) (localctx IExpressionContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, WhilelangParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(54)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case WhilelangParserINT:
		localctx = NewIntContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(47)
			p.Match(WhilelangParserINT)
		}

	case WhilelangParserT__11:
		localctx = NewReadContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(48)
			p.Match(WhilelangParserT__11)
		}

	case WhilelangParserID:
		localctx = NewIdContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(49)
			p.Match(WhilelangParserID)
		}

	case WhilelangParserT__15:
		localctx = NewExpParenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(50)
			p.Match(WhilelangParserT__15)
		}
		{
			p.SetState(51)
			p.expression(0)
		}
		{
			p.SetState(52)
			p.Match(WhilelangParserT__16)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(62)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
			case 1:
				localctx = NewBinOpContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, WhilelangParserRULE_expression)
				p.SetState(56)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(57)
					p.Match(WhilelangParserT__12)
				}
				{
					p.SetState(58)
					p.expression(4)
				}

			case 2:
				localctx = NewBinOpContext(p, NewExpressionContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, WhilelangParserRULE_expression)
				p.SetState(59)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(60)
					_la = p.GetTokenStream().LA(1)

					if !(_la == WhilelangParserT__13 || _la == WhilelangParserT__14) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(61)
					p.expression(3)
				}

			}

		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IBoolContext is an interface to support dynamic dispatch.
type IBoolContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsBoolContext differentiates from other interfaces.
	IsBoolContext()
}

type BoolContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoolContext() *BoolContext {
	var p = new(BoolContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = WhilelangParserRULE_bool
	return p
}

func (*BoolContext) IsBoolContext() {}

func NewBoolContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoolContext {
	var p = new(BoolContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = WhilelangParserRULE_bool

	return p
}

func (s *BoolContext) GetParser() antlr.Parser { return s.parser }

func (s *BoolContext) CopyFrom(ctx *BoolContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *BoolContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NotContext struct {
	*BoolContext
}

func NewNotContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotContext {
	var p = new(NotContext)

	p.BoolContext = NewEmptyBoolContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolContext))

	return p
}

func (s *NotContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotContext) Bool_() IBoolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolContext)
}

func (s *NotContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.EnterNot(s)
	}
}

func (s *NotContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.ExitNot(s)
	}
}

type BooleanContext struct {
	*BoolContext
}

func NewBooleanContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BooleanContext {
	var p = new(BooleanContext)

	p.BoolContext = NewEmptyBoolContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolContext))

	return p
}

func (s *BooleanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.EnterBoolean(s)
	}
}

func (s *BooleanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.ExitBoolean(s)
	}
}

type AndContext struct {
	*BoolContext
}

func NewAndContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AndContext {
	var p = new(AndContext)

	p.BoolContext = NewEmptyBoolContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolContext))

	return p
}

func (s *AndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AndContext) AllBool_() []IBoolContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBoolContext); ok {
			len++
		}
	}

	tst := make([]IBoolContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBoolContext); ok {
			tst[i] = t.(IBoolContext)
			i++
		}
	}

	return tst
}

func (s *AndContext) Bool_(i int) IBoolContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolContext)
}

func (s *AndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.EnterAnd(s)
	}
}

func (s *AndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.ExitAnd(s)
	}
}

type BoolParenContext struct {
	*BoolContext
}

func NewBoolParenContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolParenContext {
	var p = new(BoolParenContext)

	p.BoolContext = NewEmptyBoolContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolContext))

	return p
}

func (s *BoolParenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolParenContext) Bool_() IBoolContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBoolContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBoolContext)
}

func (s *BoolParenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.EnterBoolParen(s)
	}
}

func (s *BoolParenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.ExitBoolParen(s)
	}
}

type RelOpContext struct {
	*BoolContext
}

func NewRelOpContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RelOpContext {
	var p = new(RelOpContext)

	p.BoolContext = NewEmptyBoolContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BoolContext))

	return p
}

func (s *RelOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelOpContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *RelOpContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RelOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.EnterRelOp(s)
	}
}

func (s *RelOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(WhilelangListener); ok {
		listenerT.ExitRelOp(s)
	}
}

func (p *WhilelangParser) Bool_() (localctx IBoolContext) {
	return p.bool_(0)
}

func (p *WhilelangParser) bool_(_p int) (localctx IBoolContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewBoolContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IBoolContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, WhilelangParserRULE_bool, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(83)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewBooleanContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(68)
			_la = p.GetTokenStream().LA(1)

			if !(_la == WhilelangParserT__17 || _la == WhilelangParserT__18) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case 2:
		localctx = NewRelOpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(69)
			p.expression(0)
		}
		{
			p.SetState(70)
			p.Match(WhilelangParserT__19)
		}
		{
			p.SetState(71)
			p.expression(0)
		}

	case 3:
		localctx = NewRelOpContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(73)
			p.expression(0)
		}
		{
			p.SetState(74)
			p.Match(WhilelangParserT__20)
		}
		{
			p.SetState(75)
			p.expression(0)
		}

	case 4:
		localctx = NewNotContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(77)
			p.Match(WhilelangParserT__21)
		}
		{
			p.SetState(78)
			p.bool_(3)
		}

	case 5:
		localctx = NewBoolParenContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(79)
			p.Match(WhilelangParserT__15)
		}
		{
			p.SetState(80)
			p.bool_(0)
		}
		{
			p.SetState(81)
			p.Match(WhilelangParserT__16)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewAndContext(p, NewBoolContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, WhilelangParserRULE_bool)
			p.SetState(85)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(86)
				p.Match(WhilelangParserT__22)
			}
			{
				p.SetState(87)
				p.bool_(3)
			}

		}
		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}

	return localctx
}

func (p *WhilelangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 4:
		var t *BoolContext = nil
		if localctx != nil {
			t = localctx.(*BoolContext)
		}
		return p.Bool__Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *WhilelangParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *WhilelangParser) Bool__Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
