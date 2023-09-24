package say

import (
	"testing"
)

func TestHello(t *testing.T) {
	Hello("Rob", "Muhlestein")
}

func ExampleHello() {

	Hello("Rob", "Muhlestein")

	// Output:
	// Hello, Rob Muhlestein.

}
