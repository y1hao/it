package main

import "iter"

type Stats struct {
	MovieID      int
	TotalRatings float64
	Count        int
}

func (rs *Stats) Avg() float64 {
	return rs.TotalRatings / float64(rs.Count)
}

// GetStats aggregates a ratings sequence by the movie ID
func GetStats(ratings iter.Seq[*Rating]) iter.Seq[*Stats] {
	stats := make(map[int]*Stats)
	for r := range ratings {
		if stats[r.MovieID] == nil {
			stats[r.MovieID] = &Stats{MovieID: r.MovieID}
		}
		s := stats[r.MovieID]
		s.TotalRatings += r.Rating
		s.Count++
	}
	return func(yield func(*Stats) bool) {
		for _, stat := range stats {
			if !yield(stat) {
				break
			}
		}
	}
}
