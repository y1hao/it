package it

import (
	"slices"
	"testing"
)

func TestNaturals(t *testing.T) {
	var got []int
	max, cur := 5, 0
	for n := range Naturals() {
		if cur == max {
			break
		}
		cur++
		got = append(got, n)
	}

	want := []int{0, 1, 2, 3, 4}
	if !slices.Equal(got, want) {
		t.Fatalf("want %v, got %v", want, got)
	}
}
