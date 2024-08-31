package it

import (
	"slices"
	"testing"
)

func TestNaturals(t *testing.T) {
	var got []int
	max, i := 5, 0
	for n := range Naturals() {
		if i == max {
			break
		}
		i++
		got = append(got, n)
	}

	want := []int{0, 1, 2, 3, 4}
	if !slices.Equal(got, want) {
		t.Fatalf("want %v, got %v", want, got)
	}
}
