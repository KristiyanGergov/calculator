package main

import (
	"calculator_assignment/calculate"
	"calculator_assignment/parser"
	"calculator_assignment/reader"
	"fmt"
)

func main() {
	expression, err := reader.ReadExpression()
	for err != nil {
		fmt.Println("couldn't read your input. please, try again")
		expression, err = reader.ReadExpression()
	}

	parsedExpression, err := parser.ParseExpression(expression)
	if err != nil {
		fmt.Println("you have entered invalid input. please, try again. cause:", err)
	}

	result, err := calculate.Expression(parsedExpression)
	if err != nil {
		fmt.Println("expression could not be calculated. please, try again. cause:", err)
	}

	fmt.Printf("Expression: %s = %v", expression, result)
}
