package array

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{2, 4, 6, 1, 3}
	got := Sum(numbers)
	want := 16

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
