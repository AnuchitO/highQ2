package sum

import (
	"testing"
)

func TestSum(t *testing.T) {
	got := Sum(1, 2)
	want := 3

	if got != want {
		t.Fatalf("got %d want %d", got, want)
	}
}
