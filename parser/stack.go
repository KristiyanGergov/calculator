package parser

type Stack struct {
	s []OperatorType
}

func NewStack() *Stack {
	return &Stack{make([]OperatorType, 0)}
}

func (s *Stack) Push(v OperatorType) {
	s.s = append(s.s, v)
}

func (s *Stack) Pop() OperatorType {
	l := len(s.s)

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res
}

func (s *Stack) Peek() OperatorType {
	l := len(s.s)
	return s.s[l-1]
}

func (s *Stack) Size() int {
	return len(s.s)
}
