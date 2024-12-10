package day9

import "testing"

func TestPart1(t *testing.T) {
	input := "2333133121414131402"
	got := solvePart1(input)
	want := 1928
	if got != want {
		t.Errorf("solvePart1 got %d want %d", got, want)

	}

}
