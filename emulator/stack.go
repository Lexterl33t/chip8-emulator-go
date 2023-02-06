package emulator

import "errors"

// Stack is a type that contains a slice of int16s.
// @property {[]int16} Stack - This is the stack of the program.
type Stack struct {
	Stack []int16
}

type ok bool

// NewStack returns a pointer to a new Stack.
func NewStack() *Stack {
	return &Stack{}
}

// Checking if the stack is empty.
func (s *Stack) Empty() ok {
	if len(s.Stack) > 0 {
		return false
	}

	return true
}

// Returning the length of the stack.
func (s *Stack) Len() int {
	return len(s.Stack)
}

// Appending the data to the stack.
func (s *Stack) Push(data int16) {
	s.Stack = append(s.Stack, data)
}

// Removing the last element from the stack.
func (s *Stack) Pop() error {
	if !s.Empty() {
		s.Stack = s.Stack[:s.Len()-1]
		return nil
	}

	return errors.New("The stack was empty")
}

// Returning the top element of the stack.
func (s *Stack) Top() (int16, error) {

	if !s.Empty() {
		return s.Stack[s.Len()-1], nil
	}

	return -1, errors.New("Stack was empty")
}
