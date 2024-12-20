package day15

import (
	"strings"
	"testing"
)

func TestComputeGPSSum(t *testing.T) {
	input := `##########
#.O.O.OOO#
#........#
#OO......#
#OO@.....#
#O#.....O#
#O.....OO#
#O.....OO#
#OO....OO#
##########`
	lines := strings.Split(input, "\n")
	warehouse := readMap(lines)
	want := 10092
	got := computeGPSSum(&warehouse)
	if got != want {
		t.Errorf("computeGPSSum got %d want %d", got, want)

	}

}
