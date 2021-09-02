package parser

type tokenType int

const (
	Operand tokenType = iota
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
	Type  tokenType
	Value interface{}
}
