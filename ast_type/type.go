package ast_type

type WhileGoType int

const (
	Undefined WhileGoType = iota
	TypeProgram
	TypeStatement
	TypeAttribution
	TypeSkip
	TypeIf
	TypeWhile
	TypePrint
	TypeWrite
	TypeBlock
	TypeExpression
	TypeConst
	TypeRead
	TypeId
	TypeOp
	TypeParen
)
