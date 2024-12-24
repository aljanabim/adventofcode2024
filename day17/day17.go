package day17

import (
	"fmt"
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
	register.A >>= inst.Combo(register)
	// int(float64(register.A) / math.Pow(2, float64(inst.Combo(register))))
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
	register.B = register.A >> inst.Combo(register)
	// register.B = int(float64(register.A) / math.Pow(2, float64(inst.Combo(register))))
}

// Opcode 7
func Cdv(inst Instruction, register *Register) {
	register.C = register.A >> inst.Combo(register)
	// register.C = int(float64(register.A) / math.Pow(2, float64(inst.Combo(register))))
}

func RunProgram(register *Register, instructions []Instruction) []int {
	out := []int{}

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
			out = append(out, Out(inst, register))
		case 6:
			Bdv(inst, register)
		case 7:
			Cdv(inst, register)
		}
		if inc {
			register.PIdx++
		}
	}
	return out
}

func GetValidRegister(register *Register, instructions []Instruction, expected []int) int {
	validA := 1

	for {
		result := []int{}
		register.PIdx = 0
		register.A = validA
		register.B = 0
		register.C = 0
		inc := true
	progLoop:
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
				result = append(result, Out(inst, register))
				if result[len(result)-1] != expected[len(result)-1] {
					// fmt.Println("Oppps results not match expected", result, expected)
					break progLoop
				}
			case 6:
				Bdv(inst, register)
			case 7:
				Cdv(inst, register)
			}
			if inc {
				register.PIdx++
			}
		}
		if len(result) == len(expected) && result[len(result)-1] == expected[len(result)-1] {
			// fmt.Println("Found the fucker", validA, result)
			break
		}
		if validA%1_000_000 == 0 {
			fmt.Println("Testing", validA)
		}
		validA++
	}
	return validA
}

// The program boils down to the following
// while A not 0
//
//	B <- (A%8)^(A>>((A%8)^7))
//	out: B % 8
//	A >>= 3
//
// For 0 (last operand)     we need A=XXX B=000, thus, we need A=XXX such that B%8=000, so we search for all A \in [0,7].
// For 3 (last instruction) we need A=XXXYYY B=010, thus we need A=XXXYYYY such that B%8=010, then unshifting A<<3
// and so on until the first instruction
func SearchValidA(instructions []Instruction, expected []int, step int, ANext int) int {
	if step == -1 {
		return ANext
	}
	for i := range 8 {
		register := &Register{A: (ANext << 3) | i}
		out := RunProgram(register, instructions)
		if out[0] == expected[step] {
			// fmt.Printf("At step %d - Got with %d A=%d: %v\n", step, i, ANext<<3|i, out)
			res := SearchValidA(instructions, expected, step-1, ANext<<3|i)
			if res != -1 {
				return res
			}
		}
	}
	return -1
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

func parseOut(out []int) string {
	outStr := make([]string, len(out))
	for i, num := range out {
		outStr[i] = fmt.Sprint(num)
	}
	return strings.Join(outStr, ",")

}

func solvePart1() string {
	defer utils.Duration(utils.Track("Part 1"))
	lines, err := utils.ReadLines("day17/input.txt")
	if err != nil {
		panic(err)
	}
	register, instructions, err := readInput(lines)
	if err != nil {
		panic(err)
	}
	return parseOut(RunProgram(&register, instructions))
}

func solvePart2BruteForce() int {
	lines, err := utils.ReadLines("day17/input.txt")
	if err != nil {
		panic(err)
	}
	register, instructions, err := readInput(lines)
	if err != nil {
		panic(err)
	}
	expected := []int{}
	for _, inst := range instructions {
		expected = append(expected, inst.Opcode)
		expected = append(expected, inst.Operand)
	}
	return GetValidRegister(&register, instructions, expected)
}
func solvePart2() int {
	defer utils.Duration(utils.Track("Part 2"))
	lines, err := utils.ReadLines("day17/input.txt")
	if err != nil {
		panic(err)
	}
	_, instructions, err := readInput(lines)
	if err != nil {
		panic(err)
	}
	expected := []int{}
	for _, inst := range instructions {
		expected = append(expected, inst.Opcode)
		expected = append(expected, inst.Operand)
	}
	return SearchValidA(instructions, expected, len(expected)-1, 0)
}

func Solve() {
	res := solvePart1()
	utils.PrintSolution(17, 1, res)
	res2 := solvePart2()
	utils.PrintSolution(17, 2, res2)
}
