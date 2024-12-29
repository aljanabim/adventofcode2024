package day21

import (
	"testing"
)

func TestComputeFinalSeq(t *testing.T) {
	var testCodes = []struct {
		code   string
		length int
	}{
		{"029A", 68},
		{"980A", 60},
		{"179A", 68},
		{"456A", 64},
		{"379A", 64},
	}

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
	dir2Seq := createKeyMap(dirpadGrid, [2]int{0, 0})

	for _, c := range testCodes {
		finalSeq := computeFinalSeq(2, c.code, num2Seq, dir2Seq)
		if c.length != len(finalSeq) {
			t.Errorf("computeFinalSeq(%q) got len %d want %d\nSeq: %s", c.code, len(finalSeq), c.length, finalSeq)
		}
	}
}

func TestSolvePart1(t *testing.T) {
	codes := []string{
		"029A",
		"980A",
		"179A",
		"456A",
		"379A",
	}
	got := solvePart(2, codes)
	want := 126384
	if got != want {
		t.Fatalf("got %d want %d", got, want)
	}
}
