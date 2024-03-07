package api

import (
	"fmt"
	// "io"
	"os"
	"log"
	// "strconv"
	// "net/http"
	// "net/url"
	// "encoding/json"
	"database/sql"
	"github.com/ryanbradynd05/go-tmdb"
)

type APIResponse struct {
	Page    int `json:"page"`
	Results []struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		GenreIds         []int   `json:"genre_ids"`
		ID               int     `json:"id"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		Popularity       float64 `json:"popularity"`
		PosterPath       string  `json:"poster_path"`
		ReleaseDate      string  `json:"release_date"`
		Title            string  `json:"title"`
		Video            bool    `json:"video"`
		VoteAverage      float64 `json:"vote_average"`
		VoteCount        int     `json:"vote_count"`
	} `json:"results"`
	TotalPages   int `json:"total_pages"`
	TotalResults int `json:"total_results"`
}

type Movie struct {
	ID       	int		`json:id`		// primary key
	Title    	string	`json:title`
	Year     	int		`json:year`
	Genre    	string	`json:genre`
	IMDBRating 	float64	`json:imdb_rating`
}

var tmdbAPI *tmdb.TMDb
var genreMap map[int]string


func init() {
	genreMap = map[int]string{
		28: "Action",
	    12: "Adventure",
	    16: "Animation",
	    35: "Comedy",
	    80: "Crime",
	    99: "Documentary",
	    18: "Drama",
	    10751: "Family",
	    14: "Fantasy",
	    36: "History",
	    27: "Horror",
	    10402: "Music",
	    9648: "Mystery",
	    10749: "Romance",
	    878: "Science Fiction",
	    10770: "TV Movie",
	    53: "Thriller",
	    10752: "War",
	    37: "Western",
	    10759: "Action & Adventure",
	    10762: "Kids",
	    10763: "News",
	    10764: "Reality",
	    10765: "Sci-Fi & Fantasy",
	    10766: "Soap",
	    10767: "Talk",
	    10768: "War & Politics",
	}
}

func SaveMovie(db *sql.DB, movie Movie) error {
	// Insert data into the table
    insertQuery := "INSERT INTO movies (id, title, year, genre, imdb_rating) VALUES ($1, $2, $3, $4, $5)"
    _, err := db.Exec(insertQuery, movie.ID, movie.Title, movie.Year, movie.Genre, movie.IMDBRating)
    if err != nil {return err;}
    return nil
}


/* Search for a movie and return a list of results */
func GetMovies(db *sql.DB, title string) ([]Movie, error) {
	tmdbAPI = tmdb.Init(tmdb.Config{
    	APIKey:   os.Getenv("TMDB_API_KEY"),
  		Proxies:  nil,
       	UseProxy: false,
    })
    options := make(map[string]string)
    options["language"] = "en-US"

    response, err := tmdbAPI.SearchMovie(title, options)
    if err != nil {log.Println(err); return []Movie{}, err;}

    responseJSON, err := tmdb.ToJSON(response)
    if err != nil {log.Println(err); return []Movie{}, err;}

    fmt.Println(string(responseJSON))
}

// func GetMovies(db *sql.DB, title string) ([]Movie, error) {
	// Create properly formatted URL
	// target, err := url.Parse(endpoint)
	// if err != nil {log.Println(err); return []Movie{}, err;}

	// params := url.Values{}
	// params.Add("query", title)
	// params.Add("language", "en-US")
	// target.RawQuery = params.Encode() // Escape Query Parameters

	// // Create and send GET request
	// req, err := http.NewRequest("GET", target.String(), nil)
	// if err != nil {log.Println(err); return []Movie{}, err;}

	// req.Header.Add("accept", "application/json")
	// req.Header.Add("Authorization", os.Getenv("TMDB_API_TOKEN"))
	// res, err := http.DefaultClient.Do(req)
	// if err != nil {log.Println(err); return []Movie{}, err;}

	// defer res.Body.Close()

	// // Create a list of movies based on the response and return it
	// body, err := io.ReadAll(res.Body)
	// if err != nil {log.Println(err); return []Movie{}, err;}

	// var apiResponse APIResponse
    // if err := json.Unmarshal(body, &apiResponse); err != nil {
    //     log.Fatal("Error unmarshalling JSON:", err)
    // }

    // movies := make([]Movie, len(apiResponse.Results))

	// for i, movie := range apiResponse.Results {
	// 	// Convert genre ids to genre names
	// 	var genre string
	// 	for _, id := range movie.GenreIds {
	// 		genre += genreMap[id] + ", "
	// 	}

	// 	year, err := strconv.Atoi(movie.ReleaseDate[:4])
	// 	if err != nil {log.Println(err); return []Movie{}, err;}

	// 	movies[i] = Movie{
	// 		ID: movie.ID,
	// 		Title: movie.Title,
	// 		Year: year,
	// 		Genre: genre[:len(genre)-2],
	// 		IMDBRating: movie.VoteAverage,
	// 	}
	// }

	// return []Movie{}, nil
}
