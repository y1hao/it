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

func Movies(entries iter.Seq2[[]string, error]) iter.Seq2[*Movie, error] {
	return func(yield func(*Movie, error) bool) {
		for fields, err := range entries {
			if err != nil {
				if !yield(nil, fmt.Errorf("err in movie entry: %w", err)) {
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
