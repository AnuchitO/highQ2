package fizzbuzz_test

import (
	"testing"

	"github.com/anuchito/highq2/fizzbuzz"
)

func TestFizzBuzz(t *testing.T) {
	got := fizzbuzz.Say(15)
	if got != "FizzBuzz" {
		t.Errorf("got %s; but want %s\n", got, "15")
	}
}
