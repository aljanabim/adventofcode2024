package day19

import (
	"testing"

	"github.com/aljanabim/adventofcode2024/utils"
)

func TestSolvePart1_simple(t *testing.T) {
	stripes := []int{0, 22, 3, 12}
	want := "312"
	// maxOrder := getOrder(want)
	got, _ := helper(stripes, want, 0, len(want), "", map[CacheKey]CacheVal{})
	if got != want {
		t.Fatalf("got %s want %s", got, want)
	}

}

func TestSolvePart1_mini(t *testing.T) {
	stripes := []int{0, 1, 21, 3, 4, 325, 13, 43, 31}
	cases := []struct {
		arg  string
		want string
	}{
		{"31211", "31211"},
		{"3441", "3441"},
		{"4331", "4331"},
		{"113431", "113431"},
		{"5325", ""},
		{"325114", "325114"},
		{"3141", "3141"},
		{"331423", ""},
	}
	for _, c := range cases {
		got, ok := helper(stripes, c.arg, 0, len(c.arg), "", map[CacheKey]CacheVal{})
		if got != c.want {
			t.Errorf("solvePart1(%s)=%s want %s ok status %t", c.arg, got, c.want, ok)
		}
	}
}

func TestSolvePart1(t *testing.T) {
	lines, err := utils.ReadLines("input_mini.txt")
	if err != nil {
		panic(err)
	}
	stripesStr := lines[0]
	flags := lines[2:]
	stripes, color2Int := parseStripes(stripesStr)
	flagNums := parseFlags(flags, color2Int)

	got := solvePart1(stripes, flagNums)
	want := 6
	if got != want {
		t.Fatalf("got %d want %d", got, want)
	}
}
