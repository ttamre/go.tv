package api


type Review struct {
	ID      int 	`json:id`		// primary key
	UserID  int 	`json:user_id`	// foreign key
	MovieID int 	`json:movie_id`	// foreign key

	Rating int		`json:rating`
	Review string	`json:review`
}
