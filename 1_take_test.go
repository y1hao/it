package it

import (
	"slices"
	"testing"
)

func TestTakeCollect(t *testing.T) {
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := slices.Collect(Take(10, Naturals()))

	if !slices.Equal(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}
