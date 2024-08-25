package main

import "iter"

func Drop[T any](n int, seq iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		var i int
		for x := range seq {
			if i < n {
				i++
			} else {
				if !yield(x) {
					break
				}
			}
		}
	}
}

func OmitError[T any](seq iter.Seq2[T, error]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for v, err := range seq {
			if err == nil {
				if !yield(v) {
					break
				}
			}
		}
	}
}
