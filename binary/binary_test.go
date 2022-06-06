package binary

import (
	"testing"
)

var testCases = []struct {
	binary string
	want   int
}{
	{"1", 1},
	{"10", 2},
	{"11", 3},
	{"100", 4},
	{"1001", 9},
	{"11010", 26},
	{"10001101000", 1128},
	{"0", 0},
}

func TestParseBinary(t *testing.T) {
	for _, tc := range testCases {
		got, _ := Parse(tc.binary)

		if got != tc.want {
			t.Errorf("got %d; but want %d\n", got, tc.want)
		}
	}
}

func TestParseBinaryError(t *testing.T) {
	want := "\"i\" is not a vaid digit"

	_, err := Parse("invalid binary")

	if err.Error() != want {
		t.Errorf("got %#v; but want %#v\n", err.Error(), want)
	}
}
