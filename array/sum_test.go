package array

import "testing"

func TestSum(t *testing.T) {
	numbers := []int{3, 4, 5}
	got := Sum(numbers)
	want := 12
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
