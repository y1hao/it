package it

import "iter"

func Map[T1, T2 any](f func(T1) T2, seq iter.Seq[T1]) iter.Seq[T2] {
	return func(yield func(T2) bool) {
		for x := range seq {
			if !yield(f(x)) {
				break
			}
		}
	}
}

func Filter[T any](predicate func(T) bool, seq iter.Seq[T]) iter.Seq[T] {
	return func(yield func(T) bool) {
		for x := range seq {
			if predicate(x) {
				if !yield(x) {
					break
				}
			}
		}
	}
}

func Reduce[T1, T2 any](base T2, f func(cur T1, acc T2) T2, seq iter.Seq[T1]) T2 {
	for x := range seq {
		base = f(x, base)
	}
	return base
}
