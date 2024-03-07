package api


type User struct {
	ID       int 		`json:id`		// primary key
	Username string		`json:username`
	Password string		`json:password`
	// Reviews  []int	`json:reviews`	// review IDs
}






/*
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
