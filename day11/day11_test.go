package day11

import "testing"

func TestPart1(t *testing.T) {
	input := []string{"125", "17"}
	got := solvePart1(input, 6)
	want := 22
	if got != want {
		t.Errorf("solvePart1 got %d want %d", got, want)
	}

	input = []string{"125", "17"}
	got = solvePart1(input, 25)
	want = 55312
	if got != want {
		t.Errorf("solvePart1 got %d want %d", got, want)
	}

	input = []string{"0", "1", "1001", "99", "999"} // 1 2024 1 0 9 9 2021976
	got = solvePart1(input, 1)
	want = 7
	if got != want {
		t.Errorf("solvePart1 got %d want %d", got, want)
	}
}
