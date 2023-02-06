package main

import (
	"github.com/Lexterl33t/chip8-emulator/emulator"
)

func main() {
	var bytescode []byte = []byte{
		0x10, 0x11,
	}

	emulator.Exec(bytescode)
}
