package emulator

type PC int
type Bytecode []byte

// A Runtime has a stack, a program counter, and a map of registers.
// @property Stack - The stack is a data structure that is used to store data. It is a LIFO (Last In
// First Out) data structure.
// @property {PC} PC - The program counter. This is a special register that keeps track of the current
// instruction.
// @property Registers - A map of the registers.
type Runtime struct {
	Stack           *Stack
	PC              PC
	CurrentBytecode Bytecode
	Registers       map[Register]int
	Program         []byte
}

// NewCPU returns a pointer to a new CPU struct with a new stack, a program counter set to 0x00, and a
// map of registers.
func NewRuntime(program []byte) *Runtime {
	return &Runtime{
		Stack:           NewStack(),
		PC:              -1,
		Registers:       make(map[Register]int),
		CurrentBytecode: make([]byte, 2),
		Program:         program,
	}
}

// Incrementing the program counter and then setting the current bytecode to the next two bytes in the
// program.
func (runtime *Runtime) NextOpcode() {

	runtime.PC++

	if runtime.PC >= PC(len(runtime.Program)) {
		runtime.CurrentBytecode = make([]byte, 2)
	} else {
		runtime.CurrentBytecode[0] = runtime.Program[runtime.PC]
		runtime.PC++
		runtime.CurrentBytecode[1] = runtime.Program[runtime.PC]
	}
}

func (runtime *Runtime) Fetch() (Bytecode, ok) {

	if runtime.PC >= 0 {
		return runtime.CurrentBytecode, true
	}

	return nil, false
}

func (runtime *Runtime) Decode(Bytecode) {

}
