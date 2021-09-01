package calculator

type Calculator interface {
	Expression(string) (float64, error)
}

type ExpressionCalculator struct{}

func NewExpressionCalculator() *ExpressionCalculator {
	return &ExpressionCalculator{}
}

func (*ExpressionCalculator) Expression(expression string) (float64, error) {
	return 0.0, nil
}
