package day17

import (
	"testing"

	"github.com/aljanabim/adventofcode2024/utils"
)

func TestRunProgram(t *testing.T) {
	lines, err := utils.ReadLines("input_test.txt")
	if err != nil {
		panic(err)
	}
	register, instructions, err := readInput(lines)
	if err != nil {
		panic(err)
	}
	got := RunProgram(&register, instructions)
	want := "4,6,3,5,6,3,5,2,1,0"
	if got != want {
		t.Fatalf("got %s want %s", got, want)
	}

}

func TestBst(t *testing.T) {
	register := Register{C: 9}
	instruction := Instruction{Opcode: 2, Operand: 6}
	Bst(instruction, &register)
	want := 1
	got := register.B
	if got != want {
		t.Fatalf("got %d want %d", got, want)
	}
}
func TestBxl(t *testing.T) {
	register := Register{B: 29}
	instruction := Instruction{Opcode: 1, Operand: 7}
	Bxl(instruction, &register)
	want := 26
	got := register.B
	if got != want {
		t.Fatalf("got %d want %d", got, want)
	}
}

func TestBxc(t *testing.T) {
	register := Register{B: 2024, C: 43690}
	instruction := Instruction{Opcode: 4, Operand: 0}
	Bxc(instruction, &register)
	want := 44354
	got := register.B
	if got != want {
		t.Fatalf("got %d want %d", got, want)
	}
}

func TestMiniProgram(t *testing.T) {
	lines := []string{"Program: 5,0,5,1,5,4"}
	register, instructions, err := readInput(lines)
	if err != nil {
		panic(err)
	}

	register = Register{A: 10}
	got := RunProgram(&register, instructions)
	want := "0,1,2"
	if got != want {
		t.Fatalf("got %s want %s", got, want)
	}

}

func TestMiniProgram_registerA(t *testing.T) {
	lines := []string{"Program: 0,1,5,4,3,0"}
	register, instructions, err := readInput(lines)
	if err != nil {
		panic(err)
	}

	register = Register{A: 2024}
	got := RunProgram(&register, instructions)
	want := "4,2,5,6,7,7,7,7,3,1,0"
	if got != want && register.A != 0 {
		t.Fatalf("got %s want %s", got, want)
	}
}
