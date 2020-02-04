package structs

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 3.0)
	want := 26.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(10.0, 3.0)
	want := 30.0
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
