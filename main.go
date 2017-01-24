package main

import "flag"

func main() {
	var movieName string
	flag.StringVar(&movieName, "movie", "", "Movie for which meta data is to be fetched")
	flag.Parse()

	GetMovieDetails(movieName)
}
