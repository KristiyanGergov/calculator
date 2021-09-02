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

func (app *App) Start() {
	expression, err := app.reader.ReadExpression()
	for err != nil {
		fmt.Println("couldn't read your input. please, try again")
		expression, err = app.reader.ReadExpression()
	}

	parsedExpression, err := app.parser.Expression(expression)
	if err != nil {
		fmt.Println("you have entered invalid input. please, try again. cause:", err)
	}

	result, err := app.calculator.Expression(parsedExpression)
	if err != nil {
		fmt.Println("expression could not be calculated. please, try again. cause:", err)
	}

	fmt.Printf("Expression: %s = %v", expression, result)
}
