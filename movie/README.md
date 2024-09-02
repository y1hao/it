# What's the highest rated movie?

We have downloaded the MovieLens 32M dataset
from https://grouplens.org/datasets/movielens/.
It contains 32M entries of movie ratings.

All data are provided in the form of CSV files.
We care about the following 2 files:

ratings.csv:
(userId, movieId, rating, timestamp)

movies.csv:
(movieId, title, genres)

Ask:
Write a program to process the above 2 CSV files,
output the information of the highest rated movie.