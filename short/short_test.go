package short

import (
	"fmt"
	"testing"
)

func TestShort(t *testing.T) {
	t.Log(testing.Short())
}

func TestShortExampleForTestThatIsLong(t *testing.T) {
	if testing.Short() {
		fmt.Println("skipping test")
		t.Skip()
	}

	fmt.Println("too long to run")
	// too long to run
}
