package calculator

import "calculator_assignment/parser"

type Calculator interface {
	Expression([]parser.Token) (float64, error)
}

type ExpressionCalculator struct{}

func NewExpressionCalculator() *ExpressionCalculator {
	return &ExpressionCalculator{}
}

func (*ExpressionCalculator) Expression(expression []parser.Token) (float64, error) {
	return 0.0, nil
}
