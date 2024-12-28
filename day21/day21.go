package day21

import (
	"fmt"
	"strings"

	"github.com/aljanabim/adventofcode2024/utils"
)

/*
<vA<AA>>^A	vAA<^A>A	<v<A>>^A	vA^A	<vA>^A		<v<A>^A>AA	vA^A	<v<A>A>^AAA	vA<^A>A
v<<A		>>^A		<A			>A		vA			<^AA		>A		<vAAA		>^A
< 			A 			^ 			A 		> 			^^			A 		vvv			A

<			A			^			A		^^			>			A		vvv			A
v<<A		^>>A		<A			>A		<AA			>vA			^A		v<AAA		^>A
v<A<AA^>>A	<A>vAA^A	v<<A^>>A	vA^A	v<<A^>>AA	vA<A^>A		<A>A	v<A<A^>>AAA	<A>vA^A


<v<A>>^A 	vA^A 			<vA<AA>>^AA vA<^A>AA 	vA^A 				<vA>^AA 	<A>A 	<v<A>A>^AAA vA<^A>A
<A 			>A 				v<<AA 		>^AA 		>A 					vAA 		^A 		<vAAA 		>^A
^			A				<< 			^^			A 					>>			A		vvv			A

^			A				^^			<<			A					>>			A		vvv			A
<A			>A				<AA			v<AA		^>>A				vAA			^A		v<AAA		^>A
v<<A^>>A	vA^A			v<<A^>>AA	v<A<A^>>AA	<A>vAA^A			v<A^>AA		<A>A	v<A<A^>>AAA	<A>vA^A

R0
R1
R2
*/

func createKeyMap(grid [][]string, avoid [2]int) map[[2]string]string {
	key2Seq := map[[2]string]string{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			from := grid[i][j]
			for ii := 0; ii < len(grid); ii++ {
				for jj := 0; jj < len(grid[0]); jj++ {
					to := grid[ii][jj]
					if len(from) > 0 && len(to) > 0 {
						s := strings.Builder{}
						vertMoves := ii - i // + = move down	- = move up
						horzMoves := jj - j // + = move right	- = move left
						// Flip priority between horz and vert check on
						// row where we have the missing button
						if avoid[0] == i {
							if vertMoves < 0 { // move up
								s.WriteString(strings.Repeat("^", utils.Abs(vertMoves)))
							}
							if horzMoves > 0 { // move right
								s.WriteString(strings.Repeat(">", horzMoves))
							}
							if horzMoves < 0 { // move left
								s.WriteString(strings.Repeat("<", utils.Abs(horzMoves)))
							}
							if vertMoves > 0 { // move down
								s.WriteString(strings.Repeat("v", vertMoves))
							}
						} else {
							if horzMoves < 0 { // move left
								s.WriteString(strings.Repeat("<", utils.Abs(horzMoves)))
							}
							if vertMoves < 0 { // move up
								s.WriteString(strings.Repeat("^", utils.Abs(vertMoves)))
							}
							if horzMoves > 0 { // move right
								s.WriteString(strings.Repeat(">", horzMoves))
							}
							if vertMoves > 0 { // move down
								s.WriteString(strings.Repeat("v", vertMoves))
							}
						}
						s.WriteRune('A')
						key2Seq[[2]string{from, to}] = s.String()
					}
				}
			}
		}
	}
	return key2Seq
}

func computeKeySequence(code string, key2Seq map[[2]string]string) string {
	from := "A"
	seq := strings.Builder{}
	for _, to := range code {
		seq.WriteString(key2Seq[[2]string{from, string(to)}])
		from = string(to)
	}
	return seq.String()
}

func computeFinalSeq(code string, num2Seq map[[2]string]string, dir2Seq map[[2]string]string) string {
	fmt.Println("Code", code)
	r0Seq := computeKeySequence(code, num2Seq)
	fmt.Println("R0", r0Seq)
	r1Seq := computeKeySequence(r0Seq, dir2Seq)
	fmt.Println("R1", r1Seq)
	r2Seq := computeKeySequence(r1Seq, dir2Seq)
	fmt.Println("R2", r2Seq)
	return r2Seq
}

var codes = []string{
	"789A",
	"540A",
	"285A",
	"140A",
	"189A",
}

func solvePart1(codes []string) int {
	numpadGrid := [][]string{
		{"7", "8", "9"},
		{"4", "5", "6"},
		{"1", "2", "3"},
		{"", "0", "A"},
	}

	// dirpadGrid := [][]string{
	// 	{"", "^", "A"},
	// 	{"<", "v", ">"},
	// }

	num2Seq := createKeyMap(numpadGrid, [2]int{3, 0})
	for nums, seq := range num2Seq {
		fmt.Println(nums, seq)
	}
	// dir2Seq := createKeyMap(dirpadGrid, [2]int{0, 0})
	tot := 0
	// for _, code := range codes {
	// 	numI64, err := strconv.ParseInt(code[:len(code)-1], 10, 64)
	// 	num := int(numI64)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	finalSeq := computeFinalSeq(code, num2Seq, dir2Seq)
	// 	tot += num * len(finalSeq)
	// }

	return tot
}

func Solve() {
	res := solvePart1(codes)
	// 138560 - too high
	utils.PrintSolution(21, 1, res)
}
