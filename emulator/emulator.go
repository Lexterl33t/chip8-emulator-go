package emulator

import "fmt"

func Exec(bytescode []byte) {

	var cpu *CPU = NewCPU()

	fmt.Println(cpu.Stack)
	fmt.Println("Execute bytecodes")
}
