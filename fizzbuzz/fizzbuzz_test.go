package fizzbuzz

import "testing"

var fizzbuzzTests = []struct {
	n    int
	want string
}{
	{1, "1"},
	{2, "2"},
	{3, "Fizz"},
	{4, "4"},
	{5, "Buzz"},
	{6, "Fizz"},
	{7, "7"},
	{8, "8"},
	{9, "Fizz"},
	{10, "Buzz"},
	{11, "11"},
	{12, "Fizz"},
	{13, "13"},
	{14, "14"},
	{15, "FizzBuzz"},
}

func TestFizzBuzzShouldReturn1WhenInput1(t *testing.T) {
	for _, tc := range fizzbuzzTests {
		got := Say(tc.n)
		if got != tc.want {
			t.Errorf("got %s; but want %s\n", got, tc.want)
		}
	}
}
