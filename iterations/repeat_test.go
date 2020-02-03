package iterations

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("add two numbers", func(t *testing.T) {
		got := Repeat("a")
		expected := "aaaaa"

		if got != expected {
			t.Errorf("expected %q but got %q", expected, got)
		}
	})
}