package emulator

type PC int

type CPU struct {
	Stack *Stack
	PC    PC
}

func NewCPU() *CPU {
	return &CPU{
		Stack: NewStack(),
	}
}
