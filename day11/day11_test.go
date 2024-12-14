package day11

import (
	"reflect"
	"testing"
)

func aTestPart1(t *testing.T) {
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

func aTestBlink(t *testing.T) {
	got := []string{"0"}
	for range 4 {
		got = blink(got)
	}
	want := []string{"2", "2", "0", "4"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("blink got %v want %v", got, want)
	}
}

func TestPart2(t *testing.T) {
	input := []string{"0", "20"}
	solvePart2(input, 4)
	// input := []string{"125", "17"}
	// got := solvePart2(input, 6)
	// want := 22
	// if got != want {
	// 	t.Errorf("solvePart2 got %d want %d", got, want)
	// }

	// input = []string{"125", "17"}
	// got = solvePart2(input, 25)
	// want = 55312
	// if got != want {
	// 	t.Errorf("solvePart2 got %d want %d", got, want)
	// }

	// input = []string{"0", "1", "1001", "99", "999"} // 1 2024 1 0 9 9 2021976
	// got = solvePart2(input, 1)
	// want = 7
	// if got != want {
	// 	t.Errorf("solvePart2 got %d want %d", got, want)
	// }
}
