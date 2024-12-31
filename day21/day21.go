package day21

import (
	"fmt"
	"strings"

	"github.com/aljanabim/adventofcode2024/utils"
)

func getKeySequence(seq string, keys string, depth int, key2Seq map[[2]string][]string) []string {
	if depth == len(seq)-1 {
		return []string{keys}
	}
	from := string(seq[depth])
	to := string(seq[depth+1])

	paths := key2Seq[[2]string{from, to}]

	result := []string{}
	for _, path := range paths {
		result = append(result, getKeySequence(seq, fmt.Sprintf("%s%s", keys, path), depth+1, key2Seq)...)
		// getKeySequence(code[nextSeqStartIdx:], key2Seq, seq+path, seqs)
	}
	return result
}

func getAllSequences(seq string, path string, depth int, cache map[string][]string, key2Seq map[[2]string][]string) []string {
	seqSplit := strings.SplitAfter(seq, "A")
	subSeq := seqSplit[depth]
	if depth == len(seqSplit)-1 {
		return []string{path}
	}
	result := []string{}
	var subSeqseqs []string
	if val, ok := cache[subSeq]; ok {
		subSeqseqs = val
	} else {
		subSeqseqs = getKeySequence(fmt.Sprintf("A%s", subSeq), "", 0, key2Seq)
		cache[subSeq] = subSeqseqs
	}
	for _, subSeqSeq := range subSeqseqs {
		result = append(result, getAllSequences(seq, fmt.Sprintf("%s%s", path, subSeqSeq), depth+1, cache, key2Seq)...)
	}
	return result
}

func solvePart2(numRobots int, codes []string) int {
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
	seq2Seqs := map[string][]string{
		// "029A": {"^^", "vv"},
		// "980A": {">>", "<<"},
	}
	num2Seq := createKeyMap(numpadGrid, [2]int{3, 0})
	dir2Seq := createKeyMap(dirpadGrid, [2]int{0, 0})

	defer utils.Duration(utils.Track("GetAllSeqs"))

	code := codes[4]
	seqs := getAllSequences(code, "", 0, seq2Seqs, num2Seq)
	for range 2 {
		newSeqs := []string{}
		for _, seq := range seqs {
			newSeqs = append(newSeqs, getAllSequences(seq, "", 0, seq2Seqs, dir2Seq)...)
		}
		seqs = newSeqs
	}
	for _, seq := range seqs {
		fmt.Println("Final Seq for", code, ":", seq)
	}
	// fmt.Println("results", res)

	return numRobots
}

/*

5
^^<A, <^^A => min(3,3)
<AAv<A>>^A, v<<A>^AA>A => min(10,10)



*/
