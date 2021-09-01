package parser

type Parser interface {
	Expression(string) (string, error)
}

type ExpressionParser struct{}

func NewExpressionParser() *ExpressionParser {
	return &ExpressionParser{}
}

func (*ExpressionParser) Expression(expression string) (string, error) {
	return "", nil
}
