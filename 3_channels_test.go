package it

import "testing"

func TestChannels(t *testing.T) {
	// calculate the sum of squares of first 10 odd numbers
	isOdd := func(x int) bool { return x%2 != 0 }
	square := func(x int) int { return x * x }
	add := func(a, b int) int { return a + b }

	odds := FilterCh(isOdd, NaturalsCh())
	squares := MapCh(square, odds)
	got := ReduceCh(0, add, TakeCh(10, squares))

	want := 1 + 3*3 + 5*5 + 7*7 + 9*9 + 11*11 + 13*13 + 15*15 + 17*17 + 19*19
	if got != want {
		t.Fatalf("want %d, got %d", want, got)
	}
}
