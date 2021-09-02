package parser

type TokenType int

const (
	Operand TokenType = iota
	Operator
)

type OperatorType int

const (
	Multiply OperatorType = iota
	Subtract
	Divide
	Add
)

type Token struct {
	Type  TokenType
	Value interface{}
}
