package day21

import (
	"fmt"
	"slices"
	"strconv"
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
	}
	return result
}

func getAllSeqs(seq string, cache map[string][]string, key2Seq map[[2]string][]string) []string {
	if val, ok := cache[seq]; ok {
		return val
	} else {
		seqs := getKeySequence(fmt.Sprintf("A%s", seq), "", 0, key2Seq)
		cache[seq] = seqs
		return seqs
	}
}

type Record struct {
	string
	int
}

func getLengthOfShortestString(list []string) int {
	minLen := len(list[0])
	for _, seq := range list {
		if len(seq) < minLen {
			minLen = len(seq)
		}
	}
	return minLen
}

func getMinSequence(seqs []string, depth int, recordCache map[Record]int, allSeqCache map[string][]string, key2Seq map[[2]string][]string) int {
	if depth == 0 {
		return getLengthOfShortestString(seqs)
	}
	seqLengths := []int{}
	for _, seq := range seqs {
		seqLen := 0
		seqSplit := strings.SplitAfter(seq, "A")
		for _, subSeq := range seqSplit[:len(seqSplit)-1] {
			if val, ok := recordCache[Record{subSeq, depth}]; ok {
				seqLen += val
				continue
			}
			subSeqLen := getMinSequence(getAllSeqs(subSeq, allSeqCache, key2Seq), depth-1, recordCache, allSeqCache, key2Seq)
			recordCache[Record{subSeq, depth}] = subSeqLen
			seqLen += subSeqLen
		}
		seqLengths = append(seqLengths, seqLen)
	}
	return slices.Min(seqLengths)
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
	seq2Seqs := map[string][]string{}
	num2Seq := createKeyMap(numpadGrid, [2]int{3, 0})
	dir2Seq := createKeyMap(dirpadGrid, [2]int{0, 0})

	tot := 0
	for _, code := range codes {
		seqs := getAllSeqs(code, seq2Seqs, num2Seq)
		defer utils.Duration(utils.Track("GetAllSeqs"))
		minLen := getMinSequence(seqs, numRobots, map[Record]int{}, seq2Seqs, dir2Seq)

		numI64, err := strconv.ParseInt(code[:len(code)-1], 10, 64)
		num := int(numI64)
		if err != nil {
			panic(err)
		}
		tot += num * minLen
	}
	return tot
}
