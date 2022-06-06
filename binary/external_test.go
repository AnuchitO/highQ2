package binary_test

import (
	"testing"

	"github.com/anuchito/highq2/binary"
)

func TestParseBinary(t *testing.T) {

	got, _ := binary.Parse("1")

	if got != 1 {
		t.Errorf("got %d; but want %d\n", got, 1)
	}
}
