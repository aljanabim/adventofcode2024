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
		got := generateSecret(val.arg, 2000)
		if got != val.want {
			t.Errorf("generateSecret(%d) got %d want %d", val.arg, got, val.want)
		}
	}
}
