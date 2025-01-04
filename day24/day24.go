package day24

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/aljanabim/adventofcode2024/utils"
)

type Gate struct {
	Input1 string
	Op     string
	Input2 string
	Target string
}

/*
=======================
Computation patterns
for each bit
x00 XOR y00 => z00 etc.
vvvvvvvvvvvvvvvvvvvvvvv
010  =[x00 XOR y00 z00]

030  &[y00 AND x00 kvj]
087  |[y01 XOR x01 bhq]
091  =[bhq XOR kvj z01]

056  &[y01 AND x01 rtg]
090  &[kvj AND bhq wkc]
092  |[rtg OR wkc htw]
085  ^[y02 XOR x02 mpj]
099  =[mpj XOR htw z02]

007  &[x02 AND y02 dmq]
094  &[htw AND mpj sdk]
095  |[dmq OR sdk gwf]
060  ^[y03 XOR x03 nbw]
098  =[nbw XOR gwf z03]

019  &[x03 AND y03 psr]
096  &[nbw AND gwf btm]
097  |[psr OR btm hct]
036  ^[x04 XOR y04 cnr]
101  =[cnr XOR hct z04]

&[y04 AND x04 ggt]
&[cnr AND hct nbp]
|[ggt OR nbp ffv]
^[y05 XOR x05 pmh]
=[pmh XOR ffv z05]

&[y05 AND x05 wnc]
&[ffv AND pmh brb]
|[brb OR wnc sfm]
^[x06 XOR y06 wsn]
=[sfm XOR wsn z06]

*/

// type Inputs map[string]bool

func checkSwapping(gateMappings map[string]Gate, maxZBit string) []string {
	swappedWires := []string{}

	for target, v := range gateMappings {
		// Check operation
		if target[0] == 'z' && target > "z01" && target < maxZBit {
			if v.Op != "XOR" {
				swappedWires = append(swappedWires, target)
				for _, v := range gateMappings {
					if v.Op == "XOR" && v.Input1[0] != 'x' && v.Input1[0] != 'y' && v.Target[0] != 'z' && v.Target != "z45" {
						if !slices.Contains(swappedWires, v.Target) {
							swappedWires = append(swappedWires, v.Target)
						}
					}
				}
			} else {
				targetInput1 := gateMappings[target].Input1
				targetInput2 := gateMappings[target].Input2
				// check XOR part due to correct operation but wrong input
				if strings.Contains(gateMappings[targetInput1].Op, "XOR") {
					// this check turned out not to be necessary for the given input
					// if !(strings.Contains(gateMappings[targetInput1].Input1, target[1:]) && strings.Contains(gateMappings[targetInput1].Input2, target[1:])) {
					// 	fmt.Println(">> XOR parts is broken due to wrong operation", targetInput1)
					// 	swappedWires = append(swappedWires, targetInput1)
					// }
				} else if strings.Contains(gateMappings[targetInput1].Op, "OR") {
				} else { // Input 1 has wrong operation
					swappedWires = append(swappedWires, targetInput1)
					if gateMappings[targetInput2].Op == "OR" { // We are looking for an XOR with wrong output
						for _, v := range gateMappings {
							if v.Input1[1:] == target[1:] && v.Input2[1:] == target[1:] && v.Op == "XOR" {
								swappedWires = append(swappedWires, v.Target)
								break
							}
						}
					} else if gateMappings[targetInput2].Op == "XOR" { // We are looking for an OR with wrong output
					}
				}

				if strings.Contains(gateMappings[targetInput2].Op, "XOR") {
					// this check turned out not to be necessary for the given input
					// if !(strings.Contains(gateMappings[targetInput2].Input1, target[1:]) && strings.Contains(gateMappings[targetInput2].Input2, target[1:])) {
					// 	fmt.Println(">> XOR parts is broken due to wrong operation", targetInput2)
					// 	swappedWires = append(swappedWires, targetInput2)
					// }
				} else if strings.Contains(gateMappings[targetInput2].Op, "OR") {
				} else { // Input 2 has wrong operation
					swappedWires = append(swappedWires, targetInput2)
					if gateMappings[targetInput1].Op == "OR" { // We are looking for an XOR with wrong output
						for _, v := range gateMappings {
							if v.Input1[1:] == target[1:] && v.Input2[1:] == target[1:] && v.Op == "XOR" {
								swappedWires = append(swappedWires, v.Target)
								break
							}
						}
					} else if gateMappings[targetInput1].Op == "XOR" { // We are looking for an OR with wrong output
					}
				}
			}
		}
	}

	return swappedWires
}

func solveParts(lines []string) (int, string) {
	defer utils.Duration(utils.Track("Both parts"))
	inputs := map[string]bool{}
	gates := [][]string{}
	gateMappings := map[string]Gate{}

	for _, line := range lines {
		if strings.Contains(line, ":") {
			row := strings.Split(strings.ReplaceAll(line, ":", ""), " ")
			val, err := strconv.ParseBool(row[1])
			if err != nil {
				panic(err)
			}
			inputs[row[0]] = val
		} else if strings.Contains(line, "->") {
			row := strings.Split(strings.ReplaceAll(line, "-> ", ""), " ")
			gates = append(gates, row)
		}
	}

	i := 0
	for len(gates) > 0 {
		row := gates[i]
		input1 := row[0]
		operation := row[1]
		input2 := row[2]
		target := row[3]

		gateMappings[target] = Gate{Target: target, Input1: input1, Input2: input2, Op: operation}

		_, ok1 := inputs[input1]
		_, ok2 := inputs[input2]
		if ok1 && ok2 {
			switch operation {
			case "AND":
				inputs[target] = inputs[input1] && inputs[input2]
			case "OR":
				inputs[target] = inputs[input1] || inputs[input2]
			case "XOR":
				inputs[target] = inputs[input1] != inputs[input2]
			}
			gates = slices.Delete(gates, i, i+1)
		}
		if i < len(gates)-1 {
			i++
		} else {
			i = 0
		}
	}
	num := 0
	res := 0
	for {
		val, ok := inputs[fmt.Sprintf("z%02d", num)]
		// fmt.Printf("z%02d=%t\n", num, val)
		if val {
			res |= 0b1 << num
		}
		if !ok {
			break
		}
		num++
	}
	// Check gate swapping
	swappedWires := checkSwapping(gateMappings, fmt.Sprintf("z%02d", num-1))
	slices.Sort(swappedWires)
	return res, strings.Join(swappedWires, ",")
}

func Solve() {
	lines, err := utils.ReadLines("day24/input.txt")
	if err != nil {
		panic(err)
	}
	num, swappedWires := solveParts(lines)
	utils.PrintSolution(24, 1, num)
	utils.PrintSolution(24, 2, swappedWires)
	fmt.Println("After fix")
	lines, err = utils.ReadLines("day24/input_fixed.txt")
	if err != nil {
		panic(err)
	}
	// Swap the following (deduced manually)
	// dvb <-> fsq
	// z10 <-> vcf
	// z17 <-> fhg
	// z39 <-> tnc
	num, swappedWires = solveParts(lines)
	utils.PrintSolution(24, 1, num)
	utils.PrintSolution(24, 2, swappedWires)
}
