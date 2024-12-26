package day19

import (
	"testing"
)

func _TestSolvePart1_simple(t *testing.T) {
	stripes := []int{0, 22, 3, 12, 1, 2}
	want := "312"
	// maxOrder := getOrder(want)
	got, _ := helper(stripes, want, 0, len(want), "", map[CacheKey]CacheVal{})
	if got != want {
		t.Fatalf("got %s want %s", got, want)
	}

}
func _TestSolvePart2_simple(t *testing.T) {
	stripes := []int{0, 31, 22, 3, 12, 1, 2}
	// 31 2
	// 3 12
	// 3 1 2
	wantFlag := "312"
	// maxOrder := getOrder(want)
	counter := 0
	helper2(stripes, wantFlag, 0, len(wantFlag), "", map[CacheKey]CacheVal{}, &counter)
	want := 3
	if counter != want {
		t.Fatalf("got %d want %d", counter, want)
	}

}

func _TestSolvePart2_mini(t *testing.T) {
	stripes := []int{0, 1, 21, 3, 4, 325, 13, 43, 31}
	cases := []struct {
		arg  string
		flag string
		want int
	}{
		// {"31211", "31211", 2},
		// {"3441", "3441", 1},
		// {"4331", "4331", 4},
		// 4 3 3 1 - check
		// 43 3 1 - check
		// 4 3 31 - check
		// 43 31 - check

		{"113431", "113431", 6},
		// 1 1 3 4 3 1 - check
		// 1 13 4 3 1 - check
		// 1 1 3 43 1 - skipped due to cache issue
		// 1 13 43 1 - check
		// 1 1 3 4 31
		// 1 13 4 31

		// {"5325", "", 0},
		// {"325114", "325114", 1},
		// {"3141", "3141", 2},
		// {"331423", "", 0},
	}
	for _, c := range cases {
		got := 0
		helper2(stripes, c.arg, 0, len(c.arg), "", map[CacheKey]CacheVal{}, &got)
		if got != c.want {
			t.Errorf("helper2 got %d want %d", got, c.want)
		}
	}
}

// func _TestSolvePart1(t *testing.T) {
// 	lines, err := utils.ReadLines("input_mini.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	stripesStr := lines[0]
// 	flags := lines[2:]
// 	stripes, color2Int := parseStripes(stripesStr)
// 	flagNums := parseFlags(flags, color2Int)

// 	got := solvePart1(stripes, flagNums)
// 	want := 6
// 	if got != want {
// 		t.Fatalf("got %d want %d", got, want)
// 	}
// }

// func _TestSolvePart2(t *testing.T) {
// 	lines, err := utils.ReadLines("input_mini.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	stripesStr := lines[0]
// 	flags := lines[2:]
// 	stripes, color2Int := parseStripes(stripesStr)
// 	flagNums := parseFlags(flags, color2Int)

// 	got := solvePart1(stripes, flagNums)
// 	want := 6
// 	if got != want {
// 		t.Fatalf("got %d want %d", got, want)
// 	}
// }
