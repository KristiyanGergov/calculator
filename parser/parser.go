package parser

import (
	"fmt"
	"regexp"
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

var isTokenAnOperandPattern = "^[0-9]+$"          // Matches only numbers
var isTokenAnOperatorPattern = "^[\\*\\+\\-\\/]$" // Matches only the characters * / - +

func (*ExpressionParser) Expression(expression []string) ([]Token, error) {
	var tokens []Token
	operators := newStack()

	for _, currentToken := range expression {
		tokenIsAnOperand, _ := regexp.MatchString(isTokenAnOperandPattern, currentToken)
		tokenIsAnOperator, _ := regexp.MatchString(isTokenAnOperatorPattern, currentToken)

		if tokenIsAnOperand {
			value, _ := strconv.Atoi(currentToken)

			tokens = append(tokens, Token{
				Type:  Operand,
				Value: value,
			})
		} else if tokenIsAnOperator {
			operator, ok := mapStringToOperatorType[currentToken]
			if !ok {
				return nil, fmt.Errorf("unknown operator: %s", currentToken)
			}

			priority := priorities[operator]

			for operators.size() > 0 {
				previousOperator, err := operators.peek()
				if err != nil {
					return nil, fmt.Errorf("stack was corrupted")
				}

				previousPriority := priorities[previousOperator]

				if priority <= previousPriority {
					operators.pop()
					tokens = append(tokens, Token{
						Type:  Operator,
						Value: previousOperator,
					})
				} else {
					break
				}
			}

			operators.push(operator)
		} else {
			return nil, fmt.Errorf("invalid input")
		}
	}

	for operators.size() > 0 {
		operator, _ := operators.pop()
		tokens = append(tokens, Token{
			Type:  Operator,
			Value: operator,
		})
	}

	return tokens, nil
}

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
