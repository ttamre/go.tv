package api


type Review struct {
	ID      int 	`json:_id`		// primary key
	UserID  int 	`json:user_id`	// foreign key
	MovieID int 	`json:movie_id`	// foreign key

	Rating int		`json:rating`
	Review string	`json:review`
}


func CreateReview(user *User, movie *Movie, rating int, review string) (Review, error) {
	newReview := Review{UserID: user.ID, MovieID: movie.ID, Rating: rating, Review: review}
	user.Reviews = append(user.Reviews, newReview)
	// add review to database under user[reviews]
	return newReview, nil
}

func ReadReview(reviewID int) (Review, error) {
	// get review from database
	return Review{}, nil
}

func UpdateRating(review *Review, newRating int) {
	// update database
	review.Rating = newRating
}

func UpdateReview(review *Review, newReview string) {
	// update database
	review.Review = newReview
}

func DeleteReview(review *Review) {
	// remove from database
	review = nil
}
