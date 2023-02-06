package main

import (
	"github.com/Lexterl33t/chip8-emulator/emulator"
)

func main() {
	var bytecodes []byte = []byte{
		0x10, 0x23, 0x23, 0x24,
	}

	emulator.Exec(bytecodes)
}
