package emulator

type CPU struct {
	Stack *Stack
}

func NewCPU() *CPU {
	return &CPU{
		Stack: NewStack(),
	}
}
