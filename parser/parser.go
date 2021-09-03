package parser

import (
	"fmt"
	"strconv"
)

type Parser interface {
	Expression([]string) ([]Token, error)
}

type ExpressionParser struct{}

func NewExpressionParser() *ExpressionParser {
	return &ExpressionParser{}
}

var mapStringToOperatorType = map[string]OperatorType{
	"+": Add,
	"-": Subtract,
	"/": Divide,
	"*": Multiply,
}

var priorities = map[OperatorType]int{
	Add:      0,
	Subtract: 0,
	Multiply: 1,
	Divide:   1,
}

func (*ExpressionParser) Expression(expression []string) ([]Token, error) {
	var tokens []Token
	operators := NewStack()

	expressionSize := len(expression)

	if expressionSize%2 == 0 {
		return nil, fmt.Errorf("each two operands should have exactly one operator")
	}

	if expressionSize < 3 {
		return nil, fmt.Errorf("the expression must contain at least two operands and one operator")
	}

	// We should always start with an operand
	nextTokenShouldBeOperand := true

	for _, currentToken := range expression {
		operand, tokenIsAnOperand := parseTokenIntoOperand(currentToken)
		_, tokenIsAnOperator := mapStringToOperatorType[currentToken]

		if tokenIsAnOperand {
			if !nextTokenShouldBeOperand {
				return nil, fmt.Errorf("each expression should start and end with operand and each operand should be followed by an operator")
			}

			nextTokenShouldBeOperand = false
			tokens = append(tokens, Token{
				Type:  Operand,
				Value: operand,
			})
		} else if tokenIsAnOperator {
			if nextTokenShouldBeOperand {
				return nil, fmt.Errorf("each expression should start and end with operand and each operator should be followed by an operand")
			}

			nextTokenShouldBeOperand = true
			tokens, operators = handleOperatorToken(currentToken, tokens, operators)
		} else {
			return nil, fmt.Errorf("invalid input. the character %s is not valid", currentToken)
		}
	}

	for operators.Size() > 0 {
		tokens = append(tokens, Token{
			Type:  Operator,
			Value: operators.Pop(),
		})
	}

	return tokens, nil
}

func parseTokenIntoOperand(token string) (float64, bool) {
	value, err := strconv.ParseFloat(token, 64)
	if err != nil {
		return 0, false
	}
	return value, true
}

func handleOperatorToken(token string, tokens []Token, operators *Stack) ([]Token, *Stack) {
	operator := mapStringToOperatorType[token]

	if operators.Size() == 0 {
		operators.Push(operator)
		return tokens, operators
	}

	previousOperator := operators.Peek()
	priority := priorities[operator]
	previousPriority := priorities[previousOperator]

	for operators.Size() > 0 && priority <= previousPriority {
		operators.Pop()
		tokens = append(tokens, Token{
			Type:  Operator,
			Value: previousOperator,
		})
	}

	operators.Push(operator)

	return tokens, operators
}
