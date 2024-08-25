package it

import (
	"iter"
	"slices"
	"testing"
)

func Collect[T any](seq iter.Seq[T]) []T {
	var slc []T
	for x := range seq {
		slc = append(slc, x)
	}
	return slc
}

func TestTakeCollect(t *testing.T) {
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := Collect(Take(10, Naturals()))

	if !slices.Equal(want, got) {
		t.Fatalf("want %v, got %v", want, got)
	}
}
