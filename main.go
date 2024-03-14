package main

import (
	"log"
	"net/http"
	"github.com/ttamre/go.tv/api"
)


func main() {
	// movies, _ := api.GetMovies("The Godfather")
	// for _, movie := range movies {
	// 	fmt.Printf("(%v) %v (%v, %v)\n", movie.Year, movie.Title, movie.Genre, movie.IMDBRating)
	// }

	http.HandleFunc("/", api.RootHandler)
	http.HandleFunc("/user", api.UserHandler)
	http.HandleFunc("/review", api.ReviewHandler)
	log.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

// Database initialization
func init() {
	// create monboDB connection
}
