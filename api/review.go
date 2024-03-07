package api

import (
	"log"
	"database/sql"
)

// type Review struct {
// 	ID      int 	`json:id`		// primary key
// 	UserID  int 	`json:user_id`	// foreign key
// 	MovieID int 	`json:movie_id`	// foreign key

// 	Rating int		`json:rating`
// 	Review string	`json:review`
// }

func GetReview(db *sql.DB, reviewID int) error {
	getQuery := "SELECT * FROM reviews WHERE id = $1"
	_, err := db.Exec(getQuery, reviewID)
	if err != nil {log.Println(err); return err;}
	return nil
}

func SaveReview(db *sql.DB, userID string, movieID string, rating int, review string) error {
	// Insert data into the table
	insertQuery := "INSERT INTO reviews (user_id, movie_id, rating, review) VALUES ($1, $2, $3, $4)"
	_, err := db.Exec(insertQuery, userID, movieID, rating, review)
	if err != nil {log.Println(err); return err;}
	return nil
}
