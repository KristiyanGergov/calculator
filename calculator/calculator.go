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

var operatorCalculation = map[parser.OperatorType]func(int, int) (int, error){
	parser.Add: func(a int, b int) (int, error) { return a + b, nil },
	parser.Divide: func(a int, b int) (int, error) {
		if b == 0 {
			return 0, fmt.Errorf("division with 0 is impossible")
		}
		return a / b, nil
	},
	parser.Subtract: func(a int, b int) (int, error) { return a - b, nil },
	parser.Multiply: func(a int, b int) (int, error) { return a * b, nil },
}

func (*ExpressionCalculator) Expression(expression []parser.Token) (float64, error) {
	var stack []int

	for _, token := range expression {
		if token.Type == parser.Operand {
			val := token.Value.(int)
			stack = append(stack, val)
		} else {
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]

			operatorType := token.Value.(parser.OperatorType)

			res, err := operatorCalculation[operatorType](a, b)
			if err != nil {
				return 0.0, err
			}

			stack = append(stack, res)
		}
	}

	return float64(stack[len(stack)-1]), nil
}
