package main

import (
	"fmt"
	"iter"
	"strconv"
)

type Rating struct {
	MovieID int
	Rating  float64
}

// Ratings creates a sequence of ratings data from a sequence of CSV file entries
func Ratings(entries iter.Seq2[[]string, error]) iter.Seq2[*Rating, error] {
	return func(yield func(*Rating, error) bool) {
		for fields, err := range entries {
			if err != nil {
				if !yield(nil, fmt.Errorf("error in ratings entry: %w", err)) {
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
