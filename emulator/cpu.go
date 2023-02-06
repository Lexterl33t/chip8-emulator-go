package emulator

type PC int16

type CPU struct {
	Stack     *Stack
	PC        PC
	Registers map[Register]int
}

func NewCPU() *CPU {
	return &CPU{
		Stack: NewStack(),
	}
}
