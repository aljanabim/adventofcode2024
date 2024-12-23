package day17

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/aljanabim/adventofcode2024/utils"
)

type Register struct {
	PIdx int
	A    int
	B    int
	C    int
}

type Instruction struct {
	Opcode  int
	Operand int
}

func (inst *Instruction) Literal(register *Register) int {
	return inst.Operand
}
func (inst *Instruction) Combo(register *Register) int {
	switch inst.Operand {
	case 0, 1, 2, 3:
		return inst.Operand
	case 4:
		return register.A
	case 5:
		return register.B
	case 6:
		return register.C
	case 7:
		return -1
	default:
		return inst.Opcode
	}
}

// Opcode 0
func Adv(inst Instruction, register *Register) {
	register.A = int(float64(register.A) / math.Pow(2, float64(inst.Combo(register))))
}

// Opcode 1
func Bxl(inst Instruction, register *Register) {
	register.B = register.B ^ inst.Literal(register)
}

// Opcode 2
func Bst(inst Instruction, register *Register) {
	register.B = inst.Combo(register) % 8
}

// Opcode 3
func Jnz(inst Instruction, register *Register) bool {
	if register.A != 0 {
		register.PIdx = inst.Literal(register) / 2
		return true
	}
	return false
}

// Opcode 4
func Bxc(inst Instruction, register *Register) {
	register.B = register.B ^ register.C
}

// Opcode 5
func Out(inst Instruction, register *Register) int {
	return inst.Combo(register) % 8
}

// Opcode 6
func Bdv(inst Instruction, register *Register) {
	register.B = int(float64(register.A) / math.Pow(2, float64(inst.Combo(register))))
}

// Opcode 7
func Cdv(inst Instruction, register *Register) {
	register.C = int(float64(register.A) / math.Pow(2, float64(inst.Combo(register))))
}

func RunProgram(register *Register, instructions []Instruction) string {
	out := strings.Builder{}

	inc := true
	for register.PIdx < len(instructions) {
		inst := instructions[register.PIdx]
		inc = true
		switch inst.Opcode {
		case 0:
			Adv(inst, register)
		case 1:
			Bxl(inst, register)
		case 2:
			Bst(inst, register)
		case 3:
			jumped := Jnz(inst, register)
			if jumped {
				inc = false
			}
		case 4:
			Bxc(inst, register)
		case 5:
			out.WriteString(fmt.Sprintf("%d,", Out(inst, register)))
		case 6:
			Bdv(inst, register)
		case 7:
			Cdv(inst, register)
		}
		if inc {
			register.PIdx++
		}
	}
	return out.String()[:out.Len()-1]
}

func readInput(lines []string) (Register, []Instruction, error) {
	reg := Register{}
	program := []Instruction{}

	for _, line := range lines {
		if strings.Contains(line, "Register") {
			num, err := strconv.ParseInt(line[12:], 10, 64)
			if err != nil {
				return reg, program, err
			}
			switch line[9:10] {
			case "A":
				reg.A = int(num)
			case "B":
				reg.B = int(num)
			case "C":
				reg.C = int(num)
			}
		} else if strings.Contains(line, "Program") {
			numsStr := strings.Split(line[9:], ",")
			for i, numStr := range numsStr {
				if (i+1)%2 == 0 { // Odd number
					continue
				}
				opcode, err := strconv.ParseInt(numStr, 10, 64)
				if err != nil {
					return reg, program, err
				}
				operand, err := strconv.ParseInt(numsStr[i+1], 10, 64)
				if err != nil {
					return reg, program, err
				}
				program = append(program, Instruction{Opcode: int(opcode), Operand: int(operand)})
			}
		}
	}
	return reg, program, nil
}

func Solve() {
	lines, err := utils.ReadLines("day17/input.txt")
	if err != nil {
		panic(err)
	}
	register, instructions, err := readInput(lines)
	if err != nil {
		panic(err)
	}
	res := RunProgram(&register, instructions)
	utils.PrintSolution(17, 1, res)
}
