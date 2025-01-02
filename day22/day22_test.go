package day22

import "testing"

func TestGenerateSecret(t *testing.T) {
	values := []struct {
		arg  int
		want int
	}{
		{1, 8685429},
		{10, 4700978},
		{100, 15273692},
		{2024, 8667524},
	}

	for _, val := range values {
		got := val.arg
		for range 2000 {
			got = generateSecret(got)
		}
		if got != val.want {
			t.Errorf("generateSecret(%d) got %d want %d", val.arg, got, val.want)
		}
	}
}

func TestSolvePart2(t *testing.T) {
	lines := []string{"1", "2", "3", "2024"}
	got := solvePart2(lines)
	want := 23
	if got != want {
		t.Fatalf("solvePart1(%v)=%d want %d", lines, got, want)
	}
}
