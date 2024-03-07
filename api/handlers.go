package api

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
)

var user User


func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to go.tv")
}


func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
			GetUser(r, &user)
		case "POST":
			CreateUser(r, &user)
		case "PUT":
			if r.FormValue("username") != "" {UpdateUsername(&user, r.FormValue("username"));}
			if r.FormValue("password") != "" {UpdatePassword(&user, r.FormValue("password"));}
		case "DELETE":
			DeleteUser(&user)
		}

		userJSON, _ := json.Marshal(user)
		log.Println(string(userJSON))
		fmt.Fprint(w,string(userJSON))
}

func ReviewHandler(w http.ResponseWriter, r *http.Request) {
	var review Review

	switch r.Method {
		case "GET":
			GetReview(r, &review)
		case "POST":
			CreateReview(r, &review, &user)
			user.Reviews = append(user.Reviews, review)
		case "PUT":
			if r.FormValue("rating") != "" {UpdateRating(r, &review);}
			if r.FormValue("review") != "" {UpdateReview(r, &review);}
		case "DELETE":
			DeleteReview(&review)
		}

		reviewJSON, _ := json.Marshal(review)
		log.Println(string(reviewJSON))
		fmt.Fprint(w,string(reviewJSON))
	}
