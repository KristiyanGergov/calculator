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
	var err error
	operators := newStack()

	for _, currentToken := range expression {
		operand, tokenIsAnOperand := parseTokenIntoOperand(currentToken)
		_, tokenIsAnOperator := mapStringToOperatorType[currentToken]

		if tokenIsAnOperand {
			tokens = append(tokens, Token{
				Type:  Operand,
				Value: operand,
			})
		} else if tokenIsAnOperator {
			tokens, operators, err = handleOperatorToken(currentToken, tokens, operators)

			if err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("invalid input")
		}
	}

	for operators.size() > 0 {
		tokens = append(tokens, Token{
			Type:  Operator,
			Value: operators.pop(),
		})
	}

	return tokens, nil
}

func parseTokenIntoOperand(token string) (int, bool) {
	value, err := strconv.Atoi(token)
	if err != nil {
		return 0, false
	}
	return value, true
}

func handleOperatorToken(token string, tokens []Token, operators *stack) ([]Token, *stack, error) {
	operator, ok := mapStringToOperatorType[token]
	if !ok {
		return nil, nil, fmt.Errorf("unknown operator: %s", token)
	}

	if operators.size() == 0 {
		operators.push(operator)
		return tokens, operators, nil
	}

	previousOperator := operators.peek()
	priority := priorities[operator]
	previousPriority := priorities[previousOperator]

	for operators.size() > 0 && priority <= previousPriority {
		operators.pop()
		tokens = append(tokens, Token{
			Type:  Operator,
			Value: previousOperator,
		})
	}

	operators.push(operator)

	return tokens, operators, nil
}
