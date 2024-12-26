package day19

import (
	"testing"

	"github.com/aljanabim/adventofcode2024/utils"
)

func TestSolvePart1_mini(t *testing.T) {
	lines, err := utils.ReadLines("input_mini.txt")
	if err != nil {
		panic(err)
	}
	stripes, _ := readInput(lines)
	cases := []struct {
		flag string
		want int
	}{
		{"brwrr", 2},
		{"bggr", 1},
		{"gbbr", 4},
		{"rrbgbr", 6},
		{"ubwu", 0},
		{"bwurrg", 1},
		{"brgr", 2},
		{"bbrgwb", 0},
	}
	for _, c := range cases {
		got := findCombos(0, c.flag, stripes, map[key]int{})
		if got != c.want {
			t.Errorf("findCombos(%q) got %d want %d", c.flag, got, c.want)
		}
	}
}
