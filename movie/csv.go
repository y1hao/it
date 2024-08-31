package main

import (
	"encoding/csv"
	"errors"
	"io"
	"iter"
)

// Entries returns a sequence of entries from a CSV file.
// The first line is assumed to be header and is dropped.
func Entries(r io.Reader) iter.Seq2[[]string, error] {
	reader := csv.NewReader(r)
	_, _ = reader.Read() // discard the first line
	return func(yield func([]string, error) bool) {
		for {
			entry, err := reader.Read()
			if errors.Is(err, io.EOF) || !yield(entry, err) {
				break
			}
		}
	}
}
