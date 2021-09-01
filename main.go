package main

import (
	"calculator_assignment/app"
	"calculator_assignment/calculator"
	"calculator_assignment/parser"
	"calculator_assignment/reader"
)

func main() {
	consoleReader := reader.NewConsoleReader()
	expressionParser := parser.NewExpressionParser()
	expressionCalculator := calculator.NewExpressionCalculator()

	app.New(consoleReader, expressionParser, expressionCalculator).Start()
}
