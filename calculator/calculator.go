package calculator

import (
	"calculator_assignment/parser"
	"fmt"
)

type Calculator interface {
	Expression([]parser.Token) (float64, error)
}

type ExpressionCalculator struct{}

func NewExpressionCalculator() *ExpressionCalculator {
	return &ExpressionCalculator{}
}

var operatorCalculation = map[parser.OperatorType]func(float64, float64) (float64, error){
	parser.Add: func(a float64, b float64) (float64, error) { return a + b, nil },
	parser.Divide: func(a float64, b float64) (float64, error) {
		if b == 0 {
			return 0, fmt.Errorf("division with 0 is impossible")
		}
		return a / b, nil
	},
	parser.Subtract: func(a float64, b float64) (float64, error) { return a - b, nil },
	parser.Multiply: func(a float64, b float64) (float64, error) { return a * b, nil },
}

func (*ExpressionCalculator) Expression(expression []parser.Token) (float64, error) {
	var calculations []float64

	for _, token := range expression {
		if token.Type == parser.Operand {
			val := token.Value.(float64)
			calculations = append(calculations, val)
		} else {
			a, b := calculations[len(calculations)-2], calculations[len(calculations)-1]
			calculations = calculations[:len(calculations)-2]

			operatorType := token.Value.(parser.OperatorType)

			res, err := operatorCalculation[operatorType](a, b)
			if err != nil {
				return 0.0, err
			}

			calculations = append(calculations, res)
		}
	}

	return calculations[len(calculations)-1], nil
}
