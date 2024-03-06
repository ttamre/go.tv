package api

import ("fmt")

type Review struct {
	ID     int      // primary key
	UserID int		// foreign key
	MovieID int		// foreign key

	Rating int
	Review string
}

type Movie struct {
    ID     int      // primary key
    Title  string
    Year   int
    Genre  string
    Director string
}

type User struct {
    ID       int        // primary key
    Username string
    Password string
    Reviews []Review
    ProfilePic file
}

/*
[GET] 	user searches for a movie in database
(func) SearchMovie(title string) (Movie, error)
	if found
		return it
	if not found
		[IMDB]	search for movie in IMDB
		[POST] 	if the user selects a movie, it is returned and saved in the database
		(func) AddMovie(movie Movie) error


[IMDB] user can submit a review for the given movie
[POST] save the review in the database
[POST] save the movie in the database if it doesn't already exist
(func) AddReview(movieID int, review Review) error

[GET] user can get a list of reviews on their profile from the database
(func) GetReviews(userID int) ([]Review, error)

[GET] user can get an individual review from database
(func) GetReview(reviewID int) (Review, error)

[PATCH] if a user wants to update a review, the review is updated in the database
(func) UpdateReview(review *Review) error

[DELETE] if a user wants to delete a review, the review is removed from the database
(func) DeleteReview(reviewID int) error

[AI]	user can view a list of recommended movies based on their reviews
(func)	GetRecommendations(userID int) ([]Movie, error)
*/
