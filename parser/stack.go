package parser

type stack struct {
	s []OperatorType
}

func newStack() *stack {
	return &stack{make([]OperatorType, 0)}
}

func (s *stack) push(v OperatorType) {
	s.s = append(s.s, v)
}

func (s *stack) pop() OperatorType {
	l := len(s.s)

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res
}

func (s *stack) peek() OperatorType {
	l := len(s.s)
	return s.s[l-1]
}

func (s *stack) size() int {
	return len(s.s)
}
