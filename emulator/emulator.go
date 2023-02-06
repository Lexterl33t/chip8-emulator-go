package emulator

import "fmt"

// > Execute the bytecode
func Exec(bytescode []byte) {

	var runtime *Runtime = NewRuntime(bytescode)

	for runtime.PC < PC(len(bytescode)) {

		if bytecode, ok := runtime.Fetch(); ok {
			fmt.Println(bytecode)
		}

		runtime.NextOpcode()
	}

	fmt.Println("Execute bytecodes")
}
