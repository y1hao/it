package main

import "iter"

// OmitError takes an [iter.Seq2] and returns an [iter.Seq].
// It skips elements that have a non-nil error
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
