package main

import (
	"fmt"
	"github.com/ttamre/go.tv/api"
)


func main() {
	movies, _ := api.GetMovies("The Matrix")
	for _, movie := range movies {
		fmt.Printf("(%v) %v (%v, %v)\n", movie.Year, movie.Title, movie.Genre, movie.IMDBRating)
	}
}

// Daabase initialization
func init() {
}
