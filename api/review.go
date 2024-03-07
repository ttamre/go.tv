package api

import (
	"strconv"
	"encoding/json"
	"net/http"
)

type Review struct {
	ID      int 	`json:_id`		// primary key
	UserID  int 	`json:user_id`	// foreign key
	MovieID int 	`json:movie_id`	// foreign key

	Rating int		`json:rating`
	Review string	`json:review`
}


// func CreateReview(user *User, movie *Movie, rating int, review string) (Review, error) {
func CreateReview(r *http.Response, newReview *Review, user *User) {
	// newReview := Review{UserID: user.ID, MovieID: movie.ID, Rating: rating, Review: review}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newReview)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// *newReview = Review{
	// 	UserID: r.FormValue("user_id"),
	// 	MovieID: r.FormValue("movie_id"),
	// 	Rating: r.FormValue("rating"),
	// 	Review: r.FormValue("review")}

	// add to database under user[reviews] (UPDATE users SET reviews = append(reviews, newReview) WHERE id = user.ID)
}

func GetReview(r *http.Response, review *Review) {
	// get review from database (SELECT * FROM reviews WHERE id = reviewID)
	review = Review{ID: r.FormValue("id")}

}

func UpdateRating(r *http.Response, review *Review) {
	// update database (UPDATE reviews SET rating = newRating WHERE id = review.ID)
	newRating := strconv.Atoi(r.FormValue("rating"))
	review.Rating = newRating
}

func UpdateReview(r *http.Response, review *Review) {
	// update database (UPDATE reviews SET review = newReview WHERE id = review.ID)
	review.Review = r.FormValue("review")
}

func DeleteReview(review *Review) {
	// remove from database (DELETE FROM reviews WHERE id = review.ID)
	review = nil
}
