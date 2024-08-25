package main

import (
	"fmt"
	"iter"
	"strconv"
	"strings"
)

type Rating struct {
	MovieID int
	Rating  float64
}

func Ratings(lines iter.Seq[string]) iter.Seq2[*Rating, error] {
	return func(yield func(*Rating, error) bool) {
		for l := range Drop(1, lines) {
			fields := strings.Split(l, ",")
			if len(fields) != 4 {
				if !yield(nil, fmt.Errorf("invalid entry %q", l)) {
					break
				}
			}
			movieId, err := strconv.Atoi(fields[1])
			if err != nil {
				if !yield(nil, fmt.Errorf("cannot parse movie ID: %w", err)) {
					break
				}
			}
			rating, err := strconv.ParseFloat(fields[2], 64)
			if err != nil {
				if !yield(nil, fmt.Errorf("cannot parse rating: %w", err)) {
					break
				}
			}
			if !yield(&Rating{
				MovieID: movieId,
				Rating:  rating,
			}, nil) {
				break
			}
		}
	}
}
