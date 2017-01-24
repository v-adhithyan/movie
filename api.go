package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// MovieDetails : Structure to store response from omdb api
type MovieDetails struct {
	Title      string
	Year       string
	Runtime    string
	Genre      string
	Plot       string
	imdbRating string
}

// GetMovieDetails : Fetches details of a movie from omdb, if movie name is given as argument
func GetMovieDetails(title string) (movieDetails MovieDetails) {
	title = strings.Replace(title, " ", "+", -1)
	url := "http://www.omdbapi.com/?t=" + title + "&y=&plot=full&r=json"

	response, e := http.Get(url)
	if e != nil {
		log.Fatal("Error while getting response ==>", e)
	}
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	err := json.Unmarshal(body, &movieDetails)
	if err != nil {
		log.Fatal("Error while decoding response", err)
	}
	fmt.Printf("%+v", movieDetails)
	return movieDetails
}
