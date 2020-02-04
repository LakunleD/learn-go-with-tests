package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{3, 4, 5}
	got := Sum(numbers)
	want := 12
	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := sumAll([]int{3, 5}, []int{1, 0})
	want := []int{8, 1}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %v", got, want)
	}
}
