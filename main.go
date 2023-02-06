package main

import (
	"encoding/binary"
	"fmt"
	"os"

	"github.com/Lexterl33t/chip8-emulator/emulator"
)

func main() {

	bytes, _ := os.ReadFile("./example/ibm-logo.ch8")

	bytescode := []byte{
		0x10, 0x20,
	}
	fmt.Println(bytes)
	emulator.Exec(bytes)
	a := binary.LittleEndian.Uint16(bytescode)
	fmt.Println(int16(a))

}
