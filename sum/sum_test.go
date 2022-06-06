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

func TestSumAll(t *testing.T) {
	// Arrange
	xs := []int{1, -2, 3}

	// Act
	got := Sum(xs...)

	// Assert
	if got != 2 {
		t.Error("Expected 2, got", got)
	}
}
