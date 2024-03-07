package api


import (
	"os"
	"log"
	"strconv"
	"encoding/json"
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

sq
var tmdbAPI *tmdb.TMDb
var genreMap map[int]string


// func SaveMovie(db *sql.DB, movie Movie) error {}


/* Search for a movie and return a list of results */
func GetMovies(title string) ([]Movie, error) {
	tmdbAPI = tmdb.Init(tmdb.Config{
    	APIKey:   os.Getenv("TMDB_API_KEY"),
  		Proxies:  nil,
       	UseProxy: false,
    })
    options := make(map[string]string)
    options["language"] = "en-US"

    // Search for movie
    response, err := tmdbAPI.SearchMovie(title, options)
    if err != nil {log.Println(err); return []Movie{}, err;}

    // Convert response to JSON
    responseJSON, err := tmdb.ToJSON(response)
    if err != nil {log.Println(err); return []Movie{}, err;}

    // Convert JSON to APIResponse
    var apiResponse APIResponse
    if err := json.Unmarshal([]byte(responseJSON), &apiResponse); err != nil {
        log.Fatal("Error unmarshalling JSON:", err)
    }

    // Convert APIResponse to list of Movies
    var movies []Movie
    for _, movie := range apiResponse.Results {
    	// Parse genre ID into string
    	var genre string
		for _, id := range movie.GenreIds {genre += genreMap[id] + ", ";}

		// Parse release date string into year int
		year, err := strconv.Atoi(movie.ReleaseDate[:4])
		if err != nil {log.Println(err); return []Movie{}, err;}

		// Append movie to list
		movies = append(movies, Movie{
			ID: movie.ID,
			Title: movie.Title,
			Year: year,
			Genre: genre[:len(genre)-2],
			IMDBRating: movie.VoteAverage,
		})
	}

    return movies, nil
}


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
