package emulator

import "errors"

type Stack struct {
	Stack []int16
}

type ok bool

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Empty() ok {
	if len(s.Stack) > 0 {
		return false
	}

	return true
}

func (s *Stack) Len() int {
	return len(s.Stack)
}

func (s *Stack) Push(data int16) {
	s.Stack = append(s.Stack, data)
}

func (s *Stack) Pop() error {
	if !s.Empty() {
		s.Stack = s.Stack[:s.Len()-1]
		return nil
	}

	return errors.New("The stack was empty")
}

func (s *Stack) Top() (int16, error) {

	if !s.Empty() {
		return s.Stack[s.Len()-1], nil
	}

	return -1, errors.New("Stack was empty")
}
