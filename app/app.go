package app

import (
	"calculator_assignment/calculator"
	"calculator_assignment/parser"
	"calculator_assignment/reader"
	"fmt"
)

type App struct {
	reader     reader.Reader
	parser     parser.Parser
	calculator calculator.Calculator
}

func New(reader reader.Reader, parser parser.Parser, calculator calculator.Calculator) *App {
	return &App{
		reader:     reader,
		parser:     parser,
		calculator: calculator,
	}
}

func (app *App) Start() string {
	fmt.Printf("Enter the expression to be calculated: ")
	expression := app.reader.ReadExpression()

	parsedExpression, err := app.parser.Expression(expression)
	if err != nil {
		return fmt.Sprintf("you have entered invalid input. please, try again. cause: %s", err)
	}

	result, err := app.calculator.Expression(parsedExpression)
	if err != nil {
		return fmt.Sprintf("expression could not be calculated. please, try again. cause: %s", err)
	}

	return fmt.Sprintf("Expression: %s = %v", expression, result)
}
