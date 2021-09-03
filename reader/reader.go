package reader

import (
	"strings"
	"text/scanner"
)

type Reader interface {
	ReadExpression() []string
}

type ConsoleReader interface {
	ReadString(byte) (string, error)
}

type ExpressionReader struct{
	consoleReader ConsoleReader
}

func NewConsoleReader(consoleReader ConsoleReader) *ExpressionReader {
	return &ExpressionReader{consoleReader: consoleReader}
}

func (c *ExpressionReader) ReadExpression() []string {
	input, _ := c.consoleReader.ReadString('\n')

	var s scanner.Scanner
	s.Init(strings.NewReader(input))

	var token rune
	var result = make([]string, 0)
	for token != scanner.EOF {
		token = s.Scan()
		value := s.TokenText()
		if len(value) > 0 {
			result = append(result, s.TokenText())
		}
	}

	return result
}
