package it

import "iter"

func Naturals() iter.Seq[int] {
	return func(yield func(int) bool) {
		var i int
		for {
			if !yield(i) {
				break
			}
			i++
		}
	}
}
