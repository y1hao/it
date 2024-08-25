package main

import (
	"fmt"
	"log"
)

func main() {
	lines, cleanUp, err := ReadLines("./movie/data/ml-32m/movies.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer cleanUp()

	movies := make(map[int]string)
	movieEntries := OmitError(Movies(lines))
	for m := range movieEntries {
		movies[m.ID] = m.Title
	}

	lines, cleanUp, err = ReadLines("./movie/data/ml-32m/ratings.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer cleanUp()

	ratings := make(map[string]float64)
	for r := range OmitError(Ratings(lines)) {
		title, ok := movies[r.MovieID]
		if ok {
			ratings[title] += r.Rating
		}
	}

	for t, r := range ratings {
		fmt.Printf("%s: %f\n", t, r)
	}
}
