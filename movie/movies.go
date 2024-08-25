package main

import (
	"fmt"
	"iter"
	"strconv"
	"strings"
)

type Movie struct {
	ID     int
	Title  string
	Genres []string
}

func Movies(lines iter.Seq[string]) iter.Seq2[*Movie, error] {
	return func(yield func(*Movie, error) bool) {
		for l := range Drop(1, lines) {
			fields := strings.Split(l, ",")
			if len(fields) != 3 {
				if !yield(nil, fmt.Errorf("invalid CSV entry: %q", l)) {
					break
				}
			}
			id, err := strconv.Atoi(fields[0])
			if err != nil {
				if !yield(nil, fmt.Errorf("cannot parse ID: %w", err)) {
					break
				}
			}
			if !yield(&Movie{
				ID:     id,
				Title:  fields[1],
				Genres: strings.Split(fields[2], "|"),
			}, nil) {
				break
			}
		}
	}
}
