package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("add two numbers", func(t *testing.T) {
		sum := Add(2, 3)
		expected := 5
		if sum != expected {
			t.Errorf("expected %q but got %q", expected, sum)
		}
	})
}

func ExampleAdd() {
	sum := Add(4, 7)
	fmt.Println(sum)
	// Output: 11
}
