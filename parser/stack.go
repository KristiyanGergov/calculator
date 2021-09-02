package parser

import (
	"fmt"
)

type stack struct {
	s []OperatorType
}

func newStack() *stack {
	return &stack{make([]OperatorType, 0)}
}

func (s *stack) push(v OperatorType) {
	s.s = append(s.s, v)
}

func (s *stack) pop() (OperatorType, error) {
	l := len(s.s)
	if l == 0 {
		return -1, fmt.Errorf("empty stack")
	}

	res := s.s[l-1]
	s.s = s.s[:l-1]
	return res, nil
}

func (s *stack) peek() (OperatorType, error) {
	l := len(s.s)
	if l == 0 {
		return -1, fmt.Errorf("empty stack")
	}

	return s.s[l-1], nil
}

func (s *stack) size() int {
	return len(s.s)
}
