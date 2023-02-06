package emulator_test

import (
	"testing"

	"github.com/Lexterl33t/chip8-emulator/emulator"
)

func TestEmpty(t *testing.T) {
	stack := emulator.NewStack()

	if stack.Empty() != true {
		t.Errorf("Stack was not empty")
	}
}

func TestNotEmpty(t *testing.T) {
	stack := emulator.NewStack()
	stack.Push(0x12)

	if stack.Empty() {
		t.Errorf("Stack was empty")
	}
}

func TestEmptyTop(t *testing.T) {
	stack := emulator.NewStack()

	_, err := stack.Top()
	if err == nil {
		t.Errorf("Stack Was not empty")
	}
}

func TestNotEmptyTop(t *testing.T) {
	stack := emulator.NewStack()
	stack.Push(0x10)

	_, err := stack.Top()
	if err != nil {
		t.Errorf("Stack was empty")
	}
}
