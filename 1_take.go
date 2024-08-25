package it

import "iter"

func Take[T any](n int, seq iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		if n <= 0 {
			return
		}
		var i int
		for x := range seq {
			if i == n || !yield(x) {
				break
			}
			i++
		}
	}
}
