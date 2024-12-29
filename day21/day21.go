package day21

import (
	"fmt"
	"strconv"
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

func createKeyMap(grid [][]string, avoid [2]int) map[[2]string][]string {
	key2Seq := map[[2]string][]string{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			from := grid[i][j]
			for ii := 0; ii < len(grid); ii++ {
				for jj := 0; jj < len(grid[0]); jj++ {
					to := grid[ii][jj]
					if len(from) > 0 && len(to) > 0 {
						vertMoves := ii - i // + = move down	- = move up
						horzMoves := jj - j // + = move right	- = move left
						paths := []string{}
						// fmt.Println("Skip check horz")
						skipCheckHorzFirst := i == avoid[0] && jj == avoid[1]
						skipCheckVertFirst := j == avoid[1] && ii == avoid[0]
						checkDiag := utils.Abs(vertMoves) > 0 && utils.Abs(horzMoves) > 0

						if !skipCheckVertFirst {
							s := strings.Builder{}
							if vertMoves < 0 { // move up
								s.WriteString(strings.Repeat("^", utils.Abs(vertMoves)))
							} else { // move down
								s.WriteString(strings.Repeat("v", vertMoves))
							}
							if horzMoves > 0 { // move right
								s.WriteString(strings.Repeat(">", horzMoves))
							} else { // move left
								s.WriteString(strings.Repeat("<", utils.Abs(horzMoves)))
							}
							s.WriteRune('A')
							paths = append(paths, s.String())
						}
						if !skipCheckHorzFirst && checkDiag {
							s := strings.Builder{}
							if horzMoves > 0 { // move right
								s.WriteString(strings.Repeat(">", horzMoves))
							} else { // move left
								s.WriteString(strings.Repeat("<", utils.Abs(horzMoves)))
							}
							if vertMoves < 0 { // move up
								s.WriteString(strings.Repeat("^", utils.Abs(vertMoves)))
							} else { // move down
								s.WriteString(strings.Repeat("v", vertMoves))
							}
							s.WriteRune('A')
							paths = append(paths, s.String())
						}
						key2Seq[[2]string{from, to}] = paths
					}
				}
			}
		}
	}
	return key2Seq
}

var cache = map[string][]string{}

func computeKeySequence(code string, key2Seq map[[2]string][]string, seq string, seqs *[]string) {
	if len(code) < 2 {
		*seqs = append(*seqs, seq)
		return
	}
	var from string
	var to string
	var nextSeqStartIdx int
	if len(seq) == 0 {
		from = "A"
		to = string(code[0])
		nextSeqStartIdx = 0
	} else {
		from = string(code[0])
		to = string(code[1])
		nextSeqStartIdx = 1
	}
	paths := key2Seq[[2]string{from, to}]
	for _, path := range paths {
		computeKeySequence(code[nextSeqStartIdx:], key2Seq, seq+path, seqs)
	}
}
func computeDirBotSequence(seqs []string, key2Seq map[[2]string][]string) []string {
	dirBotSeqs := []string{}
	for _, seq := range seqs {
		computeKeySequence(seq, key2Seq, "", &dirBotSeqs)
	}
	return dirBotSeqs
}

func computeFinalSeq(numRobots int, code string, num2Seq map[[2]string][]string, dir2Seq map[[2]string][]string) string {
	defer utils.Duration(utils.Track(fmt.Sprintf("computeFinalSeq code: %s", code)))

	seqs := []string{}
	computeKeySequence(code, num2Seq, "", &seqs)
	// fmt.Println("Dir bot 1 Seqs", dirBot1Seqs)
	// for _, seq := range seqs {
	// 	fmt.Println("Code", code, "robot", 0, seq)
	// }

	for range numRobots {
		seqs = computeDirBotSequence(seqs, dir2Seq)
	}
	minLen := len(seqs[0])
	minIdx := 0
	for i, seq := range seqs {
		if len(seq) < minLen {
			minLen = len(seq)
			minIdx = i
		}
	}
	return seqs[minIdx]
}

var codes = []string{
	"789A",
	"540A",
	"285A",
	"140A",
	"189A",
}

func solvePart(numRobots int, codes []string) int {
	numpadGrid := [][]string{
		{"7", "8", "9"},
		{"4", "5", "6"},
		{"1", "2", "3"},
		{"", "0", "A"},
	}

	dirpadGrid := [][]string{
		{"", "^", "A"},
		{"<", "v", ">"},
	}

	num2Seq := createKeyMap(numpadGrid, [2]int{3, 0})

	// for nums, seq := range num2Seq {
	// 	if nums[0] == "A" {
	// 		fmt.Println(nums, seq)
	// 	}
	// }
	// fmt.Println()
	dir2Seq := createKeyMap(dirpadGrid, [2]int{0, 0})
	// for nums, seq := range dir2Seq {
	// 	if nums[0] == "A" {
	// 		fmt.Println(nums, seq)
	// 	}
	// }
	tot := 0
	for _, code := range codes {
		finalSeq := computeFinalSeq(numRobots, code, num2Seq, dir2Seq)

		numI64, err := strconv.ParseInt(code[:len(code)-1], 10, 64)
		num := int(numI64)
		if err != nil {
			panic(err)
		}
		// fmt.Println(num, len(finalSeq))
		tot += num * len(finalSeq)
	}

	return tot
}

func Solve() {
	res := solvePart(2, codes)
	utils.PrintSolution(21, 1, res)
	// res = solvePart(3, codes)
	// utils.PrintSolution(21, 2, res)
}
