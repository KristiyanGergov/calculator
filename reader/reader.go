package reader

type Reader interface {
	ReadExpression() ([]string, error)
}

type ConsoleReader struct{}

func NewConsoleReader() *ConsoleReader {
	return &ConsoleReader{}
}

func (*ConsoleReader) ReadExpression() ([]string, error) {
	return []string{"5", "/", "2", "-", "3"}, nil
}
