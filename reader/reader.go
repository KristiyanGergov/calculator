package reader

type Reader interface {
	ReadExpression() (string, error)
}

type ConsoleReader struct{}

func NewConsoleReader() *ConsoleReader {
	return &ConsoleReader{}
}

func (*ConsoleReader) ReadExpression() (string, error) {
	return "mocked response", nil
}
