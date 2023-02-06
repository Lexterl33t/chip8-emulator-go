package emulator

type Stack struct {
	Stack []byte
}

func NewStack() *Stack {
	return &Stack{
		Stack: make([]byte, 4096),
	}
}

func (s *Stack) Top() byte {
	if len(s.Stack) > 0 {
		return s.Stack[len(s.Stack)-1]
	}

	return 00
}
