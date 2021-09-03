package main

import (
	"bufio"
	"calculator_assignment/app"
	"calculator_assignment/calculator"
	"calculator_assignment/parser"
	"calculator_assignment/reader"
	"os"
)

func main() {
	consoleReader := reader.NewConsoleReader(bufio.NewReader(os.Stdin))
	expressionParser := parser.NewExpressionParser()
	expressionCalculator := calculator.NewExpressionCalculator()

	app.New(consoleReader, expressionParser, expressionCalculator).Start()
}
