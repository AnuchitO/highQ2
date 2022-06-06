package teardown

import (
	"fmt"
	"testing"
)

func setup(t *testing.T) func() {
	t.Log("setup")

	return func() {
		t.Log("teardown")
	}
}

func TestTeardown(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	// test...
	t.Log("running test...")
}

func TestTeardownV2(t *testing.T) {
	t.Cleanup(setup(t))
	t.Cleanup(func() { fmt.Println("A") })
	t.Cleanup(func() { fmt.Println("B") })
	t.Cleanup(func() { fmt.Println("C") })
	t.Cleanup(func() { fmt.Println("D") })
	t.Cleanup(func() { fmt.Println("E") })

	// test...
	t.Log("running test...")
}
