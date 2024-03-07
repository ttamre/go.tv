package main

import (
	"fmt"
	"log"
	"database/sql"
	"github.com/ttamre/go.tv/api"
)

const (
	host 		= "localhost"
	port 		= 5432
	user	 	= "dev"
	password	= ""
	dbname 		= "go.tv"
)

var db sql.DB


func main() {
	movies, _ := api.GetMovies(&db, "The Matrix")
	fmt.Println(movies)
}

// Daabase initialization
func init() {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname))
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Initialize tables if they don't exist
	createUserTable := `CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, username TEXT, password TEXT, reviews UUID[]);`
	createMovieTable := `CREATE TABLE IF NOT EXISTS movies (id SERIAL PRIMARY KEY, name TEXT NOT NULL, age INT NOT NULL);`
	createReviewTable := `
	CREATE TABLE IF NOT EXISTS reviews (
		id SERIAL PRIMARY KEY,
		user_id SERIAL FOREIGN KEY,
		movie_id SERIAL FOREIGN KEY,
		rating INT NOT NULL,
		review TEXT
	);`

	for _, query := range []string{createUserTable, createMovieTable, createReviewTable} {
		if _, err := db.Exec(query); err != nil {
			log.Fatal("Error creating table:", err)
		}
	}
}
