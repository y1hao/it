package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"y1hao.github.com/it"
)

// data files are downloaded from MovieLens
// https://grouplens.org/datasets/movielens/

func main() {
	// process ratings file
	ratingsFile, err := os.Open("./movie/data/ml-32m/ratings.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer ratingsFile.Close()

	// get all entries
	ratings := Ratings(Entries(ratingsFile))

	// remove invalid entries
	validRatings := OmitError(ratings)

	// aggregate entries to total
	stats := GetStats(validRatings)

	// only consider popular movies
	effective := it.Filter(func(r *Stats) bool {
		return r.Count > 50
	}, stats)

	// find the max
	max := it.Reduce(nil, func(cur, acc *Stats) *Stats {
		if acc == nil {
			return cur
		}
		if cur.Avg() > acc.Avg() {
			return cur
		}
		return acc
	}, effective)

	// find the movie with maxID
	moviesFile, err := os.Open("./movie/data/ml-32m/movies.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer moviesFile.Close()

	for m := range OmitError(Movies(Entries(moviesFile))) {
		if m.ID == max.MovieID {
			fmt.Printf("The highest rated movie is %q\n", m.Title)
			fmt.Printf("Rating: %.2f\n", max.Avg())
			fmt.Printf("Rated by: %d\n", max.Count)
			fmt.Printf("Genres: %s\n", strings.Join(m.Genres, ", "))
			return
		}
	}

	log.Fatal("Could not find the highest rated movie is not found\n")
}
